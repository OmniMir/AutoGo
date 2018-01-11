package main

import (
	"libs"
	"os/exec"
	"fmt"
)

const (
	npmUpdate = "--help"
)

func main() {
	libs.StartNPM(npmUpdate)
	npm:= exec.Command("npm", npmUpdate)
	fmt.Printf(npm)
	//npm.Start()
}
