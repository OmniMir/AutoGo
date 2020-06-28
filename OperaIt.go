package main

import (
	"github.com/go-vgo/robotgo/clipboard"
)

func OperaIt() {
	url, _ := clipboard.ReadAll()
	StartOpera(url)
}