package main

import (
	. "github.com/OmniMir/AutoGo/libs"
)
const (
	dc         = "C:\\Program Files\\ApexDC\\ApexDC-x64.exe"
	comixology = "www.comixology.com/free-comics"
	new28oi    = "msk28oi.ru/collection/novinki"
)

func main() {
	StartApp(dc)
	StartChrome(comixology)
	StartOpera(new28oi)
}
