package main

import (
	. "github.com/OmniMir/AutoGo/libs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"net/http"
)

//Windows GUI application
func main() {
	var mainWindow *walk.MainWindow

	MainWindow{
		AssignTo: &mainWindow,
		Title:    "Парадигмы",
		MinSize:  Size{600, 400},
		Font:     Font{PointSize: 20},
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text: "⏬ Загрузки",
				OnClicked: func() {
					downloads()
					mainWindow.Close()
				},
			},
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

	mainWindow.Run()
}

func comics() {
	//Comics reading
	StartApp("C:\\Program Files\\CDisplayEx\\CDisplayEx.exe")
	StartCmd("notepad", "K:\\Комиксы\\Список.txt")
	StartChrome("--app-id=cbbipihhaanmdjaclfmpjfnnecifpjdn")
	//My comics folders
	StartExplorer("G:\\Comics")
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
	StartChrome("http://baibako.tv/browse.php")
	StartChrome("http://newstudio.tv/")
	StartChrome("http://www.lostfilm.tv/")
	StartChrome("https://vk.com/sunshine.studio")
	StartChrome("https://comicsdb.ru/")
	StartChrome("http://animanga.ru/")
	StartChrome("http://rikudou.ru/")
	StartChrome("http://mangachan.me/user/kapsilon/newest")
}

func manga() {
	//My manga folders
	StartExplorer("K:\\Манга")
	StartExplorer("G:\\Download")
	//Tool for renaming
	StartApp("C:\\Program Files (x86)\\Total Commander\\TOTALCMD64.EXE")
}

func series() {
	//Player and remote control
	StartApp("C:\\Program Files\\MPC-HC\\mpc-hc64.exe")
	StartApp("C:\\Program Files (x86)\\Unified Remote\\RemoteServerWin.exe")
	//Remote cotrol on smartphone via webhook
	http.Get("https://vash.omnimir.ru/macrodroid/serials.php")
	//My series folders
	StartExplorer("G:\\Torrent")
	StartSleep()
	StartExplorer("D:\\ОмниСериалы")
	StartSleep()
	StartExplorer("G:\\Сериалы")

}
