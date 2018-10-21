package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/lib/verify"
)

func verifyUser(user string) {
	if user == "" {
		return
	}
	log.Verbose("check if parameter user matches rules")
	verify.AssertLength(1, 32, user, "user name too long")
	verify.AssertReg(`^[a-zA-Z][a-zA-Z0-9_]{0,31}$`, user, "user name invalid")
}

func verifyPassword(password string) {
	/*
	if password == "" {
		return
	}
	*/
	log.Verbose("check if parameter password matches rules")
	verify.AssertLength(1, 64, password, "password too long")
}

func verifyEmail(email string) {
	if email == "" {
		return
	}
	log.Verbose("check if parameter email matches rules")
	verify.AssertReg(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, email, "email invalid")
}

func verifyTelephone(telephone string) {
	if telephone == "" {
		return
	}
	log.Verbose("check if parameter telephone matches rules")
	verify.AssertReg(`^((1[3-8][0-9])+\d{8})$`, telephone, "telephone invalid")
}
