package main

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
	//Graphical Apps
	doublecmdCommand  = "C:\\Program Files\\Double Commander\\doublecmd.exe"
	edgeCommand       = "C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe"
	ExplorerCommand   = "C:\\Windows\\explorer.exe"
	chromeCommand     = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	operaCommand      = "C:\\Program Files\\Opera\\launcher.exe"
	tagscannerCommand = "C:\\Program Files\\TagScanner\\Tagscan.exe"
	winrarCommand     = "C:\\Program Files\\WinRAR\\WinRAR.exe"
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

//Graphical Apps
func StartDoubleCommander(settings string) {
	start(doublecmdCommand, settings)
}
func StartEdge(settings string) {
	start(edgeCommand, settings)
}
func StartExplorer(settings string) {
	start(ExplorerCommand, settings)
}

func StartChrome(settings string) {
	start(chromeCommand, settings)
}
func StartOpera(settings string) {
	start(operaCommand, settings)
}
func StartTagScanner(settings string) {
	start(tagscannerCommand, settings)
}
func StartUWP(settings string) {
	start(ExplorerCommand, settings)
}
func StartWinRar(settings string) {
	start(winrarCommand, settings)
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
