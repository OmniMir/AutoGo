package main

import "github.com/micmonay/keybd_event"

func PressCopy() {
	copyKey, _ := keybd_event.NewKeyBonding()
	copyKey.SetKeys(keybd_event.VK_INSERT)
	copyKey.HasCTRL(true)
	copyKey.Launching()
}
