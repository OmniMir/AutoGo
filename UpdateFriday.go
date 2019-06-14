package main

import (
	. "github.com/OmniMir/AutoGo/libs"
)

const (
	//Windows Applications
	sumo      = "C:\\Program Files (x86)\\SUMo\\SUMo.exe" //winUpdate not work -> run application as document
	winUpdate = "ms-settings:windowsupdate"
	winStore  = "shell:Appsfolder\\Microsoft.WindowsStore_8wekyb3d8bbwe!App"
	kaspersky = "C:\\Program Files (x86)\\Kaspersky Lab\\Kaspersky Internet Security 19.0.0\\avpui.exe"
	//steam = "C:\\Program Files (x86)\\Steam\\Steam.exe"

	//Sites
	chromeSettings = "--new-window"
	siteChromium   = "chromium.woolyss.com/#windows-64-bit"

	//CLI Utilities
	apmUpdate      = "upgrade --confirm=false"
	composerUpdate = "global update"
	npmUpdate      = "update -g"
)

func main() {

	//Swith to New Desktop
	StartCmd(VirtualDesktop, "/Switch:1")

	//Windows Applications
	StartApp(sumo) //Need F5
	StartSleep()
	StartUWP(winUpdate)
	StartSleep()
	StartUWP(winStore)
	StartSleep()

	//Sites
	StartChrome("--new-window")
	StartSleep()
	StartChrome(siteChromium)
	StartSleep()

	//CLI Utilities
	StartAPM(apmUpdate)
	StartComposer(composerUpdate)
	StartNPM(npmUpdate)
	StartSleep()

	//Steam
	/*
		//Before need start service
		Start(steam, "")
	*/

	//Kaspersky Software Update
	StartApp(kaspersky)
	/*
		Some mouse magic))
	*/

}
