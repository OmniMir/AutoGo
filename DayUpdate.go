package main

const (
	//Windows Applications
	sumo      = "C:\\Program Files (x86)\\SUMo\\SUMo.exe" //winUpdate not work -> run application as document
	winUpdate = "ms-settings:windowsupdate"
	kaspersky = "C:\\Program Files (x86)\\Kaspersky Lab\\Kaspersky Internet Security 21.2\\avpui.exe"
	//steam = "C:\\Program Files (x86)\\Steam\\Steam.exe"

	//Sites
	chromium = "chromium.woolyss.com/#windows-64-bit"
	repacks = "https://repack.me/"
	newApps = "https://betanews.com/?s=Best+Windows+10+apps+this+week"

	//CLI Utilities

)

func UpdateFriday() {

	//Swith to New Desktop
	StartCmd(VirtualDesktop, "/Switch:1")
	//Windows Applications
	StartApp(sumo)
	StartSleep()
	StartUWP(winUpdate)
	StartSleep()
	StartStore()
	StartSleep()

	//Sites
	StartChrome(chromium)
	StartSleep()
	StartChrome(repacks)
	StartSleep()
	StartChrome(newApps)
	StartSleep()

	//CLI Utilities
	StartTerminal()

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
