package nativemessaging

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type GenericNativeMessagingHost[T any] struct {
	*NativeMessagingHost
}

func NewGenericNativeMessagingHost[T any](stderr io.Writer) *GenericNativeMessagingHost[T] {
	if stderr == nil {
		stderr = os.Stderr
	}

	return &GenericNativeMessagingHost[T]{
		NewNativeMessagingHost(stderr),
	}
}

func (host *GenericNativeMessagingHost[T]) Write(data T) error {
	bytes, err := json.Marshal(data)

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

func (host *GenericNativeMessagingHost[T]) ReadT() (*T, error) {
	data := host.createInstanceOf()

	bytes, err := host.ReadBytes()

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, data); err != nil {
		err = fmt.Errorf("failed to unmarshal payload: %v", err)
		host.error.Print(err)
		return nil, err
	}

	return data, nil
}

func (host GenericNativeMessagingHost[T]) createInstanceOf() *T {
	return new(T)
}
