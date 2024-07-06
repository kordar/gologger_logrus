package gologger_logrus_test

import (
	"bufio"
	logger "github.com/kordar/gologger"
	"github.com/kordar/gologger_logrus"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestT1(t *testing.T) {
	log := gologger_logrus.NewLogrusAdapt(logrus.New())
	log.Info("aaaaaaaaaaaaaaa")
	log.Warnf("bbbbbbbbbbbbbbb")
	log.Error("eeeeeeeeeeee")
}

func TestInit(t *testing.T) {
	logrusLog := logrus.New()
	logrusLog.SetLevel(logrus.InfoLevel)
	src, err := os.OpenFile(os.DevNull, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(src)
	logrusLog.SetOutput(writer)
	adapt := gologger_logrus.NewLogrusAdapt(logrusLog)
	logger.InitGlobal(adapt)
}
