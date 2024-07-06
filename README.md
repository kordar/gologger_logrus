# gologger_logrus

包装`logrus`对象，实现[日志门面](https://github.com/kordar/gologger)接口

## 安装

```go
go get github.com/kordar/gologger_logrus v1.0.2
```

## 初始化

```go
import (
	"bufio"
	logger "github.com/kordar/gologger"
	"github.com/kordar/gologger_logrus"
	"github.com/sirupsen/logrus"
	"os"
)

func TestInit(t *testing.T) {
	logrusLog := logrus.New()
	logrusLog.SetLevel(logrus.InfoLevel)
	src, err := os.OpenFile(os.DevNull, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(src)
	logrusLog.SetOutput(writer)
	// 将logrusLog对象包装，设置即可
	adapt := gologger_logrus.NewLogrusAdapt(logrusLog)
	logger.InitGlobal(adapt)
}
```