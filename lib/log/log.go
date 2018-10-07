package log

import "github.com/sirupsen/logrus"

var log = logrus.NewEntry(logrus.New())

func Info(msg string) {
	log.Infoln(msg)
}

func Warn(msg string) {
	log.Warnln(msg)
}

func SetUser(user string) {
	log = log.WithFields(logrus.Fields{"user": user})
}

func init() {
	log.Logger.SetFormatter(&logrus.TextFormatter{})
}
