package main

import (
	. "github.com/OmniMir/AutoGo/libs"
)

const (
	//Windows Applications
	sumo      = "C:\\Program Files (x86)\\SUMo\\SUMo.exe" //winUpdate not work -> run application as document
	winUpdate = "C:\\Users\\Sergey\\AppData\\Local\\Packages\\windows.immersivecontrolpanel_cw5n1h2txyewy\\LocalState\\Indexed\\Settings\\ru-RU\\AAA_SystemSettings_MusUpdate_UpdateActionButton.settingcontent-ms"
	winStore  = "shell:Appsfolder\\Microsoft.WindowsStore_8wekyb3d8bbwe!App"
	kaspersky = "C:\\Program Files (x86)\\Kaspersky Lab\\Kaspersky Internet Security 19.0.0\\avpui.exe"
	//steam = "C:\\Program Files (x86)\\Steam\\Steam.exe"

	//Sites
	chromeSettings = "--new-window"
	siteChromium   = "http://chromium.woolyss.com/#windows-64-bit"
	siteSamlab     = "http://samlab.ws/category/samsoft/"

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
	StartApp(winUpdate)
	StartSleep()
	StartUWP(winStore)
	StartSleep()

	//Sites
	sitesToChrome := chromeSettings + Space + siteChromium + Space + siteSamlab
	StartChrome(sitesToChrome)
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
