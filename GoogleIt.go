package main

import (
	"os/exec"
	//Libs for clipboard and
	"github.com/atotto/clipboard"
	"github.com/micmonay/keybd_event"
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
	copyKey, _ := keybd_event.NewKeyBonding()
	copyKey.SetKeys(keybd_event.VK_INSERT)
	copyKey.HasCTRL(true)
	copyKey.Launching()
	//Read clipboard
	wordsForSearch, _ := clipboard.ReadAll()
	chromeSearchFinal := chromeSearch + wordsForSearch
	//And to Chrome
	chrome := exec.Command(chromeCommand, chromeSettings, chromeSearchFinal)
	if (wordsForSearch != emptySearch) {
		chrome.Start()
	}
}