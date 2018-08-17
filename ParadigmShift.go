package main

import (
	. "./libs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
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
	//My comics folders
	StartExplorer("G:\\Comics")
	StartExplorer("G:\\Comics")
	StartExplorer("K:\\Комиксы")
	//Comics reading
	StartApp("C:\\Program Files\\CDisplayEx\\CDisplayEx.exe")
	StartApp("K:\\Комиксы\\Список.txt")
	StartChrome("--app-id=cbbipihhaanmdjaclfmpjfnnecifpjdn")
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
	//My series folders
	StartExplorer("D:\\ОмниСериалы")
	StartExplorer("G:\\Torrent")
	StartExplorer("G:\\Сериалы")
	//Player and remote control
	StartApp("C:\\Program Files\\MPC-HC\\mpc-hc64.exe")
	StartApp("C:\\Program Files (x86)\\Unified Remote\\RemoteServerWin.exe")
}
