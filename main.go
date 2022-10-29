package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "day"+os.Args[1], ".")
	ret, err := cmd.Output()

	if err != nil {
		println(err.Error())
	}

	println(string(ret))
}
