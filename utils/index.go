package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = map[string]func(){
		"linux": func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		},
		"darwin": func() {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		},
		"windows": func() {
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		},
	}
}

func ClearConsole() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
		// Note: Clear the terminal bufer
		fmt.Print("\033[3J\033[;H")
	} else {
		panic("Your platform (" + runtime.GOOS + ") is unsupported! I can't clear terminal screen :(")
	}
}
