package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"net/http"
)

//Windows GUI application
func ParadigmShift() {
	var mainWindow *walk.MainWindow

	err := MainWindow{
		AssignTo: &mainWindow,
		Title:    "Парадигмы",
		MinSize:  Size{600, 400},
		Font:     Font{PointSize: 20},
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text: "🎦 Сериалы",
				OnClicked: func() {
					series()
					mainWindow.Close()
				},
			},
			PushButton{
				Text: "📘 Комиксы",
				OnClicked: func() {
					comics()
					mainWindow.Close()
				},
			},
			PushButton{
				Text: "⏬ Загрузки",
				OnClicked: func() {
					downloads()
					mainWindow.Close()
				},
			},
			PushButton{
				Text: "📙 Манга",
				OnClicked: func() {
					manga()
					mainWindow.Close()
				},
			},
			PushButton{
				Text: "💑 Знакомства",
				OnClicked: func() {
					dating()
					mainWindow.Close()
				},
			},
		},
	}.Create()
	Check(err)

	mainWindow.Run()
}

func comics() {
	//Comics reading
	StartApp("C:\\Program Files\\CDisplayEx\\CDisplayEx.exe")
	StartCmd("C:\\Windows\\System32\\notepad.exe", "K:\\Комиксы\\Список.txt")
	//StartApp("K:\\Комиксы\\Список.txt")
	StartChrome("--app-id=cbbipihhaanmdjaclfmpjfnnecifpjdn")
	//My comics folders
	StartExplorer("G:\\Comics")
	StartSleep()
	StartExplorer("K:\\Комиксы")
	StartSleep()
	StartExplorer("G:\\Comics")
	StartSleep()
	StartExplorer("K:\\Комиксы")
}

func dating() {
	//How to start Edge))
	StartApp("%windir%\\explorer.exe shell:Appsfolder\\Microsoft.MicrosoftEdge_8wekyb3d8bbwe!MicrosoftEdge")
	//DatingDialog

}

func downloads() {
	//My Download folder
	StartExplorer("G:\\Download")
	//My sites
	StartChrome("--new-window")
	StartSleep()
	StartChrome("baibako.tv/browse.php?search=&incldead=0&cat=0&videoformat=4")
	StartSleep()
	StartChrome("newstudio.tv/")
	StartSleep()
	StartChrome("www.lostfilm.tv/")
	StartSleep()
	StartChrome("vk.com/koshara.serials")
	StartSleep()
	StartChrome("comicsdb.ru/")
	StartSleep()
	StartChrome("animanga.ru/")
	StartSleep()
	StartChrome("rikudou.ru/")
	StartSleep()
	StartChrome("manga-chan.me/user/kapsilon/newest")
}

func manga() {
	//My mangg folders
	StartExplorer("K:\\Манга")
	StartExplorer("G:\\Download")
	//Tool for renaming
	StartApp("C:\\Program Files (x86)\\Double Commander\\doublecmd.exe")
}

func series() {
	//Player and remote control
	StartApp("C:\\Program Files\\MPC-HC\\mpc-hc64.exe")
	StartApp("C:\\Program Files (x86)\\Unified Remote\\RemoteServerWin.exe")
	//Remote control on smartphone via webhook
	_ , err := http.Get("https://vash.omnimir.ru/macrodroid/serials.php")
	Check(err)
	//My series folders
	StartExplorer("G:\\Torrent")
	StartSleep()
	StartExplorer("D:\\ОмниСериалы")
	StartSleep()
	StartExplorer("G:\\Сериалы")

}
