package libs

import (
	"os/exec"
	"time"
)

const (

	//Commandline apps
	cmdCommand      = "cmd"
	startCommand    = "/C start"
	apmCommand      = "apm"
	composerCommand = "composer"
	npmCommand      = "npm"
	VirtualDesktop  = "C:\\Program Files\\AutoGo\\VirtualDesktop.exe"

	//Graphical apps
	chromeCommand   = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	operaCommand    = "C:\\Program Files (x86)\\Opera\\launcher.exe"
	ExplorerCommand = "C:\\Windows\\explorer.exe"
)

//Generic appstart
func start(command string, settings string) {
	newRun := exec.Command(command, settings)
	newRun.Start()
}
func StartApp(command string) {
	newRun := exec.Command(command, "")
	newRun.Start()
}
func StartCmd(command string, settings string) {
	newRun := exec.Command(command, settings)
	newRun.Start()
}

//Graphical apps
func StartExplorer(settings string) {
	start(ExplorerCommand, settings)
}

func StartChrome(settings string) {
	start(chromeCommand, settings)
}
func StartOpera(settings string) {
	start(operaCommand, settings)
}
func StartUWP(settings string) {
	start(ExplorerCommand, settings)
}

//Commandline apps
func StartAPM(settings string) {
	settings = startCommand + Space + apmCommand + Space + settings
	start(cmdCommand, settings)
}
func StartComposer(settings string) {
	settings = startCommand + Space + composerCommand + Space + settings
	start(cmdCommand, settings)
}
func StartNPM(settings string) {
	settings = startCommand + Space + npmCommand + Space + settings
	start(cmdCommand, settings)
}

//Subsidiary apps
func StartSleep() {
	time.Sleep(500 * time.Millisecond)
}
