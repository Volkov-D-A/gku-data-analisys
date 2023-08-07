package logger

import (
	"github.com/sirupsen/logrus"
)

func Get() *logrus.Logger {
	ll := logrus.New()
	return ll
}
