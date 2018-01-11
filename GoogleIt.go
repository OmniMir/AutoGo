package main

import (
	"os/exec"
	//Libs for clipboard and
	"github.com/atotto/clipboard"
	"libs"
)

const (
	emptySearch    = ""
	chromeCommand  = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
	chromeSettings = "--new-window"
	chromeSearch   = "https://www.google.ru/search?q="
)

func main() {
	//Clean clipboard

	clipboard.WriteAll(emptySearch)
	//Copy selected text to clipboard
	libs.PressCopy()
	//Read clipboard
	wordsForSearch, _ := clipboard.ReadAll()
	chromeSearchFinal := chromeSearch + wordsForSearch
	//And to Chrome
	chrome := exec.Command(chromeCommand, chromeSettings, chromeSearchFinal)
	if (wordsForSearch != emptySearch) {
		chrome.Start()
	}
}