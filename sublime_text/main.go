package main

import (
	"fmt"
	"os"
	"os/exec"
	"github.com/mike-seagull/wsl_tools/lib"
)

var SUBLIME_TEXT_UNIX string = os.Getenv("SUBLIME_TEXT_UNIX")

func open_in_sublime_text(file string) error {
	return exec.Command("/bin/sh", "-c", SUBLIME_TEXT_UNIX + " " + file).Run()
}
func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Need a filename")
		os.Exit(1)
	}
	for _, file := range args {
		file = lib.GetFullPath(file)
		file = lib.Convert(file, true)
		open_in_sublime_text(file)
	}
}