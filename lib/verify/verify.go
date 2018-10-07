package verify

import (
	"github.com/MegaShow/goagenda/lib/log"
	"regexp"
)

func AssertReg(pattern, s, msg string) {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil || !matched {
		log.Error(msg)
	}
}

func AssertNil(s, msg string) {
	if s == "" {
		log.Error(msg)
	}
}

func AssertLength(minLen, maxLen int, s, msg string) {
	if len(s) < minLen || len(s) > maxLen {
		log.Error(msg)
	}
}
