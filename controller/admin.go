package controller

import (
	"bufio"
	"fmt"
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
	"os"
)

type AdminCtrl interface {
	GetStatus()
	Login()
	Logout()
	Register()
	Log()
}

func (c *Controller) Register() {
	user, _ := c.Ctx.GetString("user")
	password, _ := c.Ctx.GetSecretString("password")
	email, _ := c.Ctx.GetString("email")
	telephone, _ := c.Ctx.GetString("telephone")

	verifyNonNilUser(user)
	verifyNonNilPassword(password)
	verifyEmail(email)
	verifyTelephone(telephone)

	err := c.Srv.Admin().Register(user, password, email, telephone)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("register account successfully")
}

func (c *Controller) Login() {
	user, _ := c.Ctx.GetString("user")
	password, _ := c.Ctx.GetSecretString("password")

	verifyNonNilUser(user)
	verifyNonNilPassword(password)

	log.Verbose("check status")
	currentUser := c.Ctx.User.Get()
	if currentUser == user {
		log.Error("you are already logged in with this account")
	} else if currentUser != "" {
		log.Error("you are already logged in with user '" + currentUser + "', please logout first")
	}

	err := c.Srv.Admin().Login(user, password)
	if err != nil {
		log.Error(err.Error())
	}
	err = c.Ctx.User.Set(user)
	if err != nil {
		log.Error(err.Error())
	}
	log.Info("login successfully")
}

func (c *Controller) Logout() {
	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("not logged user")
		return
	}
	c.Ctx.User.Set("")
	log.Info("user '" + currentUser + "' logged out")
}

func (c *Controller) GetStatus() {
	currentUser := c.Ctx.User.Get()
	if currentUser == "" {
		fmt.Println("not logged user")
	} else {
		fmt.Println("user '" + currentUser + "' logged in")
	}
}

func (c *Controller) Log() {
	isOpen := viper.GetBool("Log.IsOpen")
	if isOpen {
		fmt.Println(aurora.Red("warning!!!"))
		fmt.Println("you are trying to access the log file")
		fmt.Println("if you don't want that everyone can access log file, please set 'false' in the config file")
		fmt.Println()
	} else {
		fmt.Println(aurora.Red("permission denied"))
		return
	}
	fmt.Println("print only 10 lines, more in file")
	fmt.Println()
	f, err := os.Open(viper.GetString("Log.Path") + string(os.PathSeparator) + viper.GetString("Log.File"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var output []string
	for {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			break
		}
		output = append(output, string(line))
		if len(output) > 10 {
			output = append(output[len(output)-10:])
		}
	}
	for _, item := range output {
		fmt.Print(item)
	}
	fmt.Println()
}

func GetAdminCtrl() AdminCtrl {
	return &ctrl
}
