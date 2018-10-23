package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/lib/verify"
)

// Username's length should between 1 and 32.
// Username includes only letters, digits and underline, and it must start with letter.
//
// If username is empty, it will pass the check.
func verifyUser(user string) {
	if user == "" {
		return
	}
	log.Verbose("check if parameter user matches rules")
	verify.AssertLength(1, 32, user, `user name too long
the length of user name can't be larger than 32`)
	verify.AssertReg(`^[a-zA-Z][a-zA-Z0-9_]{0,31}$`, user, `user name invalid
username includes only letters, digits and underline, and it must start with letter`)
}

// Password's length should between 6 and 64.
//
// If password is empty, it will pass the check.
func verifyPassword(password string) {
	if password == "" {
		return
	}
	log.Verbose("check if parameter password matches rules")
	verify.AssertLength(6, 64, password, `password too long or too short
the length of password can't be larger than 64 and shorter than 6`)
}

// Email should match the follow regexp.
//
// If email is empty, it will pass the check.
func verifyEmail(email string) {
	if email == "" {
		return
	}
	log.Verbose("check if parameter email matches rules")
	verify.AssertReg(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, email, "email invalid")
}

// Telephone should match the follow regexp, and doesn't support country/area code.
//
// If telephone is empty, it will pass the check.
func verifyTelephone(telephone string) {
	if telephone == "" {
		return
	}
	log.Verbose("check if parameter telephone matches rules")
	verify.AssertReg(`^((1[3-8][0-9])+\d{8})$`, telephone, "telephone invalid")
}

// Args should be empty.
func verifyEmptyArgs(args []string) {
	log.Verbose("check if args is empty")
	verify.AssertArrayLength(0, 0, args, "parameters shouldn't needed")
}
