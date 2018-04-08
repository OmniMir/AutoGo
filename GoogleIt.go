package main

import (
	//Libs for clipboard and
	"github.com/atotto/clipboard"
	"libs"
)

const (
	emptySearch    = ""
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

	//And to Chrome
	if wordsForSearch != emptySearch {
		chromeCommand := chromeSettings + libs.Space + chromeSearch + wordsForSearch
		libs.StartChrome(chromeCommand)
	}
}
