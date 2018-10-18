package log

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.NewEntry(logrus.New())
var file *os.File
var verbose bool
var label = make(map[string]string)
var params string

// log verbose message, this won't store in log file
func Verbose(msg string) {
	if verbose {
		fmt.Println(label["Verbose"] + msg)
	}
}

// log info message, this won't store in log file
func Show(msg string) {
	fmt.Println(label["Show"] + msg)
}

// log info message
func Info(msg string) {
	fmt.Println(label["Info"] + msg)
	log.Infoln(msg)
}

// log error message and exit with status 1
func Error(msg string) {
	fmt.Println(aurora.Red(label["Error"] + msg))
	log.Errorln(msg)
	Release()
	os.Exit(1)
}

func Release() {
	file.Close()
}

func SetUser(user string) {
	log = log.WithFields(logrus.Fields{"user": user})
}

func AddParams(key, value string) {
	if value == "" {
		return
	}
	if params != "" {
		params += ","
	}
	params += key + ":" + value
	log = log.WithFields(logrus.Fields{"params": params})
}

func SetCommand(cmd string) {
	log = log.WithFields(logrus.Fields{"command": cmd})
}

func SetVerbose(isVerbose bool) {
	verbose = isVerbose
}

func SetLabel(isLabel bool) {
	if isLabel {
		label["Verbose"] = "[VERB] "
		label["Show"] = "[SHOW] "
		label["Info"] = "[INFO] "
		label["Error"] = "[ERRO] "
	}
}

func init() {
	log.Logger.SetFormatter(&logrus.TextFormatter{})
	f, err := os.OpenFile("agenda.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	file = f
	log.Logger.SetOutput(file)
}
