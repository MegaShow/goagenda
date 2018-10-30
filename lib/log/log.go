package log

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var log = logrus.NewEntry(logrus.New())
var file *os.File
var verbose bool
var username string

// log verbose message, this won't store in log file
func Verbose(msg string) {
	if verbose {
		fmt.Println(msg)
	}
}

// log info message
func Info(msg string) {
	fmt.Println(msg)
	log.Infoln(msg)
}

// log error message and exit with status 1
func Error(msg string) {
	fmt.Println(aurora.Red(msg))
	log.Errorln(msg)
	Release()
	os.Exit(1)
}

func Release() {
	file.Close()
}

func SetUser(user string) {
	if username == user {
		return
	}
	username = user
	log = log.WithFields(logrus.Fields{"user": user})
}

func AddParams(key string, value interface{}) {
	if key == "" {
		return
	}
	key = "param_" + key
	log = log.WithFields(logrus.Fields{key: value})
}

func SetCommand(cmd string) {
	log = log.WithFields(logrus.Fields{"command": cmd})
}

func SetVerbose(isVerbose bool) {
	verbose = isVerbose
}

func Init() {
	log.Logger.SetFormatter(&logrus.TextFormatter{})
	f, err := os.OpenFile(viper.GetString("Log.Path")+string(os.PathSeparator)+viper.GetString("Log.File"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	file = f
	log.Logger.SetOutput(file)
}
