package main

import (
	"fmt"
	"os"
	"syscall"
)

const defaultUser = "root"
const defaultShell = "/bin/bash"
const credits = "please-go - please don't take me too seriously (but in Go).\nMade by Cynthia Revström <me@cynthia.re> - https://cynthia.re\n"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please what?")
		return
	}

	if os.Args[1] != "root" {
		fmt.Print(credits)
		fmt.Println("type root as the first argument to please")
		return
	}

	fmt.Println("No worries.")

	pw := getPwNam(defaultUser)
	syscall.Setuid(pw.uid)
	syscall.Setgid(pw.gid)

	argv := []string{defaultShell}
	syscall.Exec(argv[0], argv, os.Environ())
}
