package controller

import (
	"github.com/MegaShow/goagenda/lib/log"
	"github.com/MegaShow/goagenda/lib/verify"
	"math"
	"time"
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

// Username's length should between 1 and 32.
// Username includes only letters, digits and underline, and it must start with letter.
func verifyNonNilUser(user string) {
	log.Verbose("check if parameter user matches rules")
	verify.AssertNonNil(user, `user name can't be empty`)
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

// Password's length should between 6 and 64.
func verifyNonNilPassword(password string) {
	log.Verbose("check if parameter password matches rules")
	verify.AssertNonNil(password, `password can't be empty`)
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

// Title shouldn't be empty.
func verifyNonNilTitle(title string) {
	log.Verbose("check if parameter title matches rules")
	verify.AssertNonNil(title, `title can't be empty`)
}

// Start time shouldn't be empty.
// The format of time is YYYY-MM-DD/hh:mm or YYYY-M-D/h:m.
// Time is of 24-hour.
func verifyNonNilStartTime(t time.Time) {
	log.Verbose("check if parameter start time matches rules")
	verify.AssertTimeNonEqual(t, time.Unix(0, 0), `start time can't be empty`)
	verify.AssertTimeNonEqual(t, time.Unix(0, 1), `start time invalid
the format of time is YYYY-MM-DD/hh:mm or YYYY-M-D/h:m, and time is of 24-hour`)
}

// End time shouldn't be empty.
// The format of time is YYYY-MM-DD/hh:mm or YYYY-M-D/h:m.
// Time is of 24-hour.
func verifyNonNilEndTime(t time.Time) {
	log.Verbose("check if parameter end time matches rules")
	verify.AssertTimeNonEqual(t, time.Unix(0, 0), `end time can't be empty`)
	verify.AssertTimeNonEqual(t, time.Unix(0, 1), `end time invalid
the format of time is YYYY-MM-DD/hh:mm or YYYY-M-D/h:m, and time is of 24-hour`)
}

func verifyNonNilParticipator(arr []string) {
	log.Verbose("check if parameter participator matches rules")
	verify.AssertArrayLength(1, math.MaxInt32, arr, `participator can't be empty`)
}
