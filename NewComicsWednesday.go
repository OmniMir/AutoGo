package main

import (
	. "./libs"
)
const (
	dc         = "C:\\Program Files\\ApexDC\\ApexDC-x64.exe"
	comixology = "https://www.comixology.com/free-comics"
	new28oi    = "https://msk.28oi.ru/item/by_article/74"
)

func main() {
	StartApp(dc)
	StartChrome(comixology)
	StartOpera(new28oi)
}
