package main

import (
	"github.com/mike-seagull/wsl_tools/lib"
	"os"
	"fmt"
	"strings"
	"os/exec"
)
func open(file string) error{
	return exec.Command("/bin/sh", "-c", "cmd.exe /c start "+file).Run()
}
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Need a filename")
		os.Exit(1)
	}
	file := lib.GetFullPath(args[0])
	if strings.Contains(file, lib.UNIX_C) {
		fmt.Println("Cannot open")
		os.Exit(1)
	}
	file = lib.Convert(file, true)
	open(file)
	os.Exit(0)

}