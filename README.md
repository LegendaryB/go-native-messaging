<h1 align="center">go-native-messaging</h1><div align="center">

[![forthebadge](https://forthebadge.com/images/badges/fuck-it-ship-it.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

[![GitHub license](https://img.shields.io/github/license/LegendaryB/go-native-messaging.svg?longCache=true&style=flat-square)](https://github.com/LegendaryB/go-native-messaging/blob/main/LICENSE)

<sub>Built with ❤︎ by LegendaryB</sub>

[Native Messaging](https://developer.chrome.com/docs/apps/nativeMessaging/) module powered by Go.
</div><br>

## 📝 Requirements
* Read and understand [this](https://developer.chrome.com/docs/apps/nativeMessaging/).

## 🚀 How to use it?

### Retrieving complex objects
```go
package main

import (
	"github.com/LegendaryB/go-native-messaging"
)

type Message struct {
    Value string `json:"data"`
}

func main() {
    host := nativemessaging.NewNativeMessagingHost(nil)

    msg := &Message{}

    for {
		if err := host.Read(msg); err != nil {
			if err == io.EOF {
				f.Write([]byte("Received EOF error, Browser or tab was probably closed"))
				break
			}
		}

		host.Write("pong: " + msg.Value)
	}
}
```
