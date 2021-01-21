package main

import (
	"github.com/go-vgo/robotgo/clipboard"
)

func ChromeIt() {
	url, _ := clipboard.ReadAll()
	StartChrome(url)
}

func DoubleCommanderIt() {
	dir, _ := clipboard.ReadAll()
	StartDoubleCommander(dir)
}

func EdgeIt() {
	url, _ := clipboard.ReadAll()
	StartEdge(url)
}

func OperaIt() {
	url, _ := clipboard.ReadAll()
	StartOpera(url)
}

func TagScannerIt() {
	dir, _ := clipboard.ReadAll()
	StartTagScanner(dir)
}

func TerminalIt() {
	dir, _ := clipboard.ReadAll()
	StartTerminalDir(dir)
}

func WinRarIt() {
	dir, _ := clipboard.ReadAll()
	StartWinRar(dir)
}
