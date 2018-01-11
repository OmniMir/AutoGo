package libs

import "os/exec"

const (
	chromeCommand   = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	uwpCommand      = "C:\\Windows\\explorer.exe"
	apmCommand      = "apm"
	composerCommand = "composer"
	npmCommand      = "npm"
)

func Start(command string, settings string) {
	newRun := exec.Command(command, settings)
	newRun.Start()
}

func StartChrome(settings string) {
	Start(chromeCommand, settings)
}

func StartUWP(settings string) {
	Start(uwpCommand, settings)
}

func StartAPM(settings string) {
	Start(apmCommand, settings)
}

func StartComposer(settings string) {
	Start(composerCommand, settings)
}

func StartNPM(settings string) {
	Start(npmCommand, settings)
}
