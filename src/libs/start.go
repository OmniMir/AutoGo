package libs

import (
	"os/exec"
	"fmt"
	"time"
)

const (
	cmdCommand = "cmd"
	startCommand = "/C start "
	apmCommand      = "apm "
	composerCommand = "composer "
	npmCommand      = "npm "
	chromeCommand   = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	uwpCommand      = "C:\\Windows\\explorer.exe"
)

func Start(command string, settings string) {
	newRun := exec.Command(command, settings)
	newRun.Start()
}

func StartInCmd(command string, settings string) {
	newRun := exec.Command(command, settings)
	stdout, _ := newRun.CombinedOutput()
	fmt.Printf(string(stdout))
}

func StartChrome(settings string) {
	Start(chromeCommand, settings)
}

func StartUWP(settings string) {
	Start(uwpCommand, settings)
}

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

func StartSleep(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}
