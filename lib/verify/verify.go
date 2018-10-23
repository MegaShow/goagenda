package verify

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
	"regexp"
)

func AssertReg(pattern, s, msg string) {
	matched, err := regexp.MatchString(pattern, s)
	if err != nil || !matched {
		fmt.Println(aurora.Red(msg))
		os.Exit(2)
	}
}

func AssertNil(s, msg string) {
	if s == "" {
		fmt.Println(aurora.Red(msg))
		os.Exit(2)
	}
}

func AssertLength(minLen, maxLen int, s, msg string) {
	if len(s) < minLen || len(s) > maxLen {
		fmt.Println(aurora.Red(msg))
		os.Exit(2)
	}
}

func AssertArrayLength(minLen, maxLen int, arr []string, msg string) {
	if len(arr) < minLen || len(arr) > maxLen {
		fmt.Println(aurora.Red(msg))
		os.Exit(2)
	}
}
