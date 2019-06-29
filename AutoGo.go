package main

import (
	. "github.com/OmniMir/AutoGo/libs"
	"github.com/go-vgo/robotgo/clipboard"
)

func main() {
		googleIt()
	}

}

func googleIt() {
	const (
		emptySearch    = ""
		chromeSettings = "--new-window"
		chromeSearch   = "www.google.ru/search?q="
	)

	//Read clipboard
	wordsForSearch, _ := clipboard.ReadAll()
	//And to Chrome
	if wordsForSearch != emptySearch {
		chromeCommand := chromeSettings + Space + chromeSearch + wordsForSearch
		StartChrome(chromeCommand)
	}
}
