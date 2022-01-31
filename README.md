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

### Receiving a byte array
```go
package main

import (
    "io"
    "github.com/LegendaryB/go-native-messaging"
)

func main() {
    host := nativemessaging.NewNativeMessagingHost(nil)

    for {
        b, err := host.ReadBytes()
	
        if err != nil {
	    if err == io.EOF {
		break
	    }
	}
	
	// just echo the received message back to the extension
	host.Write(msg)
    }
}
```

**Note**: If you need to write a plain byte array by yourself you can use the following method:  
`func (host *NativeMessagingHost) WriteBytes(bytes []byte) error`

### Receiving a complex object
```go
package main

import (
    "io"
    "github.com/LegendaryB/go-native-messaging"
)

type Message struct {
    Value string `json:"value"`
}

func main() {
    host := nativemessaging.NewNativeMessagingHost(nil)
    msg := &Message{}

    for {
        if err := host.Read(msg); err != nil {
	    if err == io.EOF {
		break
	    }
	}

	// just echo the received message back to the extension
	host.Write(msg)
    }
}
```

### Sending a byte array
See last line in section [Receiving a byte array](#receiving-a-byte-array).

### Sending a complex object
See last line in section [Receiving a complex object](#receiving-a-complex-object).

## ⚙️ Debugging
In the samples above we were passing a `nil` to the `NewNativeMessagingHost` method. If you need to receive the error messages you can pass in a `io.Writer` instance. If the parameter is `nil` the error messages are written to `os.stderr`.

### Redirecting error messages to a file
```go
package main

import (
    "io"
    "os"
    "github.com/LegendaryB/go-native-messaging"
)

func main() {
    f, _ := os.Create("stderr.txt")
    host := nativemessaging.NewNativeMessagingHost(f)
}
```
