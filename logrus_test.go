package gologger_logrus_test

import (
	"github.com/kordar/gologger_logrus"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestT1(t *testing.T) {
	logger := gologger_logrus.NewLogrusAdapt(logrus.New())
	logger.Info("xxxxxxxxxxxxxxxxxxxxxxxxxx")
	logger.Warnf("eeeeeeeeeee")
}
