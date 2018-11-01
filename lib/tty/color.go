package tty

import "runtime"

func ColorfulError(msg string) {
	if runtime.GOOS == "windows" {
		printlnErrorWin(msg)
	} else {
		printlnErrorUnix(msg)
	}
}
