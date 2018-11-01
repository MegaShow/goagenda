package tty

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func printlnErrorUnix(msg string) {
	fmt.Println(aurora.Red(msg))
}
