package main

import "libs"

const (
	space = " "

	//Windows Applications
	sumo           = "C:\\Program Files (x86)\\SUMo\\SUMo.exe"	//winUpdate not work -> run application as document
	winUpdate      = "C:\\Users\\Sergey\\AppData\\Local\\Packages\\windows.immersivecontrolpanel_cw5n1h2txyewy\\LocalState\\Indexed\\Settings\\ru-RU\\AAA_SystemSettings_MusUpdate_UpdateActionButton.settingcontent-ms"
	winStore       = "shell:Appsfolder\\Microsoft.WindowsStore_8wekyb3d8bbwe!App"
	kaspersky = "C:\\Program Files (x86)\\Kaspersky Lab\\Kaspersky Internet Security 18.0.0\\avpui.exe"
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

	//Windows Applications
	libs.StartSimple(sumo) //Need F5
	libs.StartSleep()
	libs.StartSimple(winUpdate)
	libs.StartSleep()
	libs.StartUWP(winStore)
	libs.StartSleep()

	//Sites
	sitesToChrome := chromeSettings + space + siteChromium + space + siteSamlab
	libs.StartChrome(sitesToChrome)
	libs.StartSleep()

	//CLI Utilities
	libs.StartAPM(apmUpdate)
	libs.StartComposer(composerUpdate)
	libs.StartNPM(npmUpdate)
	libs.StartSleep()

	//Steam
	/*
	//Before need start service
	libs.Start(steam, "")
	 */

	//Kaspersky Software Update
	libs.Start(kaspersky, "")
	/*
	Some mouse magic))
	 */

}
