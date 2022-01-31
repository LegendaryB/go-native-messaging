package nativemessaging

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type NativeMessagingHost struct {
	input  io.Reader
	output io.Writer
	error  *log.Logger
}

func NewNativeMessagingHost(stderr io.Writer) *NativeMessagingHost {
	if stderr == nil {
		stderr = os.Stderr
	}

	return &NativeMessagingHost{
		input:  os.Stdin,
		output: os.Stdout,
		error:  log.New(stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (host *NativeMessagingHost) Write(v interface{}) error {
	if v == nil {
		return errors.New("can't marshal nil interface")
	}

	bytes, err := json.Marshal(v)

	if err != nil {
		err = fmt.Errorf("failed to marshal interface to JSON: %v", err)
		host.error.Print(err)
		return err
	}

	if err = host.WriteBytes(bytes); err != nil {
		host.error.Print(err)
		return err
	}

	return nil
}

func (host *NativeMessagingHost) WriteBytes(bytes []byte) error {
	if bytes == nil {
		return errors.New("can't write nil buffer")
	}

	if err := host.write(bytes); err != nil {
		return err
	}

	return nil
}

func (host *NativeMessagingHost) write(bytes []byte) error {
	size := len(bytes)

	if err := host.writeSize(size); err != nil {
		err = fmt.Errorf("failed to write payload size to standard output: %v", err)
		return err
	}

	n, err := host.output.Write(bytes)

	if err != nil {
		err = fmt.Errorf("failed to write payload to standard output: %v", err)
		return err
	}

	if n != size {
		err = fmt.Errorf("count of written bytes differs from calculated size %v/%v", n, size)
		return err
	}

	return nil
}

func (host *NativeMessagingHost) writeSize(size int) error {
	const SIZE_OF_HEADER = 4

	buffer := make([]byte, SIZE_OF_HEADER)

	binary.LittleEndian.PutUint32(buffer, (uint32)(size))

	bytesWritten, err := host.output.Write(buffer)

	if err != nil {
		err = fmt.Errorf("failed to write payload size to standard output: %v", err)
		return err
	}

	if bytesWritten != SIZE_OF_HEADER {
		err = fmt.Errorf("count of written bytes differs from calculated size %v/%v", bytesWritten, size)
		return err
	}

	return nil
}

func (host *NativeMessagingHost) Read(v interface{}) error {
	bytes, err := host.ReadBytes()

	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, v); err != nil {
		err = fmt.Errorf("failed to unmarshal payload: %v", err)
		host.error.Print(err)
		return err
	}

	return nil
}

func (host *NativeMessagingHost) ReadBytes() ([]byte, error) {
	bytes, err := host.read()

	if err != nil {
		host.error.Print(err)
		return nil, err
	}

	return bytes, nil
}

func (host *NativeMessagingHost) read() ([]byte, error) {
	size, err := host.readSize()

	if err != nil {
		return nil, err
	}

	reader := io.LimitReader(host.input, int64(size))
	buffer := make([]byte, size)

	bytesRead, err := reader.Read(buffer)

	if err != nil {
		err = fmt.Errorf("failed to read payload from standard input: %v", err)
		return nil, err
	}

	if bytesRead != int(size) {
		err = fmt.Errorf("count of read bytes differs from calculated size %v/%v", bytesRead, size)
		return nil, err
	}

	return buffer, nil
}

func (host *NativeMessagingHost) readSize() (uint32, error) {
	var size uint32 = 0

	if err := binary.Read(host.input, binary.LittleEndian, &size); err != nil {
		if err == io.EOF {
			return size, err
		}

		err = fmt.Errorf("failed to read payload size from standard input: %v", err)
		return size, err
	}

	return size, nil
}
