package main

import (
	. "github.com/lxn/walk/declarative"
	. "libs"
)

func main() {
	MainWindow{
		Title:   "Смена Парадигмы",
		MinSize: Size{600, 400},
		Font:Font{PointSize: 20},
		Layout:  VBox{},
		Children: []Widget{
			PushButton{
				Text: "Загрузки",
				OnClicked: func() {
					downloads()
				},
			},
			PushButton{
				Text: "Сериалы",
				OnClicked: func() {
					series()
				},
			},
			PushButton{
				Text: "Комиксы",
				OnClicked: func() {
					comics()
				},
			},
			PushButton{
				Text: "Манга",
				OnClicked: func() {
					manga()
				},
			},
			PushButton{
				Text: "Знакомства",
				OnClicked: func() {
					dating()
				},
			},
		},
	}.Run()
}

func comics() {
	//My comics folders
	StartApp("G:\\Comics")
	StartApp("K:\\Комиксы")
	StartApp("K:\\Комиксы\\Список.txt")
	//Comics reading
	StartApp("C:\\Program Files\\CDisplayEx\\CDisplayEx.exe")
	StartChrome("--app-id=cbbipihhaanmdjaclfmpjfnnecifpjdn")
}

func dating() {
	//How to start Edge))
	StartApp("%windir%\\explorer.exe shell:Appsfolder\\Microsoft.MicrosoftEdge_8wekyb3d8bbwe!MicrosoftEdge")
	//DatingDialog
}

func downloads() {
	//My Download folder
	StartApp("G:\\Download")
	//My sites
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
	StartApp("K:\\Манга")
	StartApp("G:\\Download")
	//Tool for renaming
	StartApp("C:\\Program Files (x86)\\Total Commander\\TOTALCMD64.EXE")
}

func series() {
	//My series folders
	StartApp("D:\\ОмниСериалы")
	StartApp("G:\\Torrent")
	//Player and remote control
	StartApp("C:\\Program Files\\MPC-HC\\mpc-hc64.exe")
	StartApp("C:\\Program Files (x86)\\Unified Remote\\RemoteServerWin.exe")
}
