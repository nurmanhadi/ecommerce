package logger

import "github.com/sirupsen/logrus"

var Logrus *logrus.Logger

func init() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	Logrus = log
}
