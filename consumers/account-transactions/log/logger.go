package log

import "github.com/sirupsen/logrus"

var Logger = logrus.NewEntry(logrus.New())

func SetServiceName(name string) {
	Logger = Logger.WithField("service", name)
}
