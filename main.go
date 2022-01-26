package main

import (
	"io"
	"os"

	"github.com/LegendaryB/go-native-messaging/nativemessaging"
)

type Test struct {
	Text string `json:"text"`
}

func main() {
	f, _ := os.Create("log.txt")

	host := nativemessaging.NewNativeMessagingHost(f)

	value := &Test{}

	for {
		if err := host.Read(value); err != nil {
			if err == io.EOF {
				f.Write([]byte("Received EOF error, Chrome probably has closed"))
				break
			}
		}

		host.Write(value)
	}
}
