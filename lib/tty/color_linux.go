package tty

import (
	"fmt"
	"github.com/logrusorgru/aurora"
)

func printlnErrorLinux(msg string) {
	fmt.Println(aurora.Red(msg))
}
