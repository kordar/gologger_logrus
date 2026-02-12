# gologger-logrus

A [logrus](https://github.com/sirupsen/logrus) adapter for [gologger](https://github.com/kordar/gologger).

## Installation

```bash
go get github.com/kordar/gologger_logrus
```

## Usage

```go
package main

import (
	"github.com/kordar/gologger"
	"github.com/kordar/gologger_logrus"
	"github.com/sirupsen/logrus"
)

func main() {
	// Create a logrus logger instance
	l := logrus.New()

	// Create the adapter
	adapter := gologger_logrus.NewLogrusAdapt(l)

	// Initialize gologger with the adapter
	logger.InitGlobal(adapter)

	// Use gologger
	logger.Info("This message is logged via logrus")
	logger.WithField("key", "value").Info("Message with field")
}
```
