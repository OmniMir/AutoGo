package libs

import (
	"fmt"
	"os/exec"
	"time"
)

const (

	//Commandline apps
	cmdCommand      = "cmd"
	startCommand    = "/C start "
	apmCommand      = "apm "
	composerCommand = "composer "
	npmCommand      = "npm "

	//Graphical apps
	chromeCommand = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	uwpCommand    = "C:\\Windows\\explorer.exe"
)

//Generic appstart
func Start(command string, settings string) {
	newRun := exec.Command(command, settings)
	newRun.Start()
}
func StartSimple(command string) {
	newRun := exec.Command(command, "")
	newRun.Start()
}
func StartInCmd(command string, settings string) {
	newRun := exec.Command(command, settings)
	stdout, _ := newRun.CombinedOutput()
	fmt.Printf(string(stdout))
}

//Graphical apps
func StartChrome(settings string) {
	Start(chromeCommand, settings)
}
func StartUWP(settings string) {
	Start(uwpCommand, settings)
}

//Commandline apps
func StartAPM(settings string) {
	settings = startCommand + apmCommand + settings
	Start(cmdCommand, settings)
}
func StartComposer(settings string) {
	settings = startCommand + composerCommand + settings
	Start(cmdCommand, settings)
}
func StartNPM(settings string) {
	settings = startCommand + npmCommand + settings
	Start(cmdCommand, settings)
}
//Subsidiary apps
func StartSleep() {
	time.Sleep(3 * time.Second)
}
