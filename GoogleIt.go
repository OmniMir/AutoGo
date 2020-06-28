package main

import (
	"github.com/go-vgo/robotgo/clipboard"
)

const (
	emptySearch    = ""
	chromeSettings = "--new-window"
	chromeSearch   = "www.google.ru/search?q="
)

func GoogleIt() {


	//Read clipboard
	wordsForSearch, _ := clipboard.ReadAll()
	//And to Chrome
	if wordsForSearch != emptySearch {
		chromeCommand := chromeSettings + Space + chromeSearch + wordsForSearch
		StartChrome(chromeCommand)
	}
}