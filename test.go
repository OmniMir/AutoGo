package main

import (
	"github.com/micmonay/keybd_event"
)

const (
	emptySearch = ""
)

var (
	wordsForSearch string
)

func main() {

	//Copy selected text to clipboard
	copy, _ := keybd_event.NewKeyBonding()
	copy.SetKeys(keybd_event.VK_C)
	copy.HasCTRL(true)
	copy.Launching()
}
