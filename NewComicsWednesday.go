package main

import "libs"

const (
	dc         = "C:\\Program Files\\ApexDC\\ApexDC-x64.exe"
	comixology = "--app-id=kkjoppilfbjlenhanlgakjknjdijblhl"
	new28oi = "https://msk.28oi.ru/item/by_article/74"
)

func main() {
	libs.StartApp(dc)
	libs.StartChrome(comixology)
	libs.StartOpera(new28oi)
}
