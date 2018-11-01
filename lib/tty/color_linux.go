package tty

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func printlnError(msg string) {
	fmt.Println(aurora.Red(msg))
}
