package logger

import "github.com/sirupsen/logrus"

var mainLogger = logrus.New()

func GetLogger(app string) *logrus.Entry {
	return mainLogger.WithField("app", app)
}
