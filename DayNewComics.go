package main

const (
	dc         = "C:\\Program Files\\ApexDC\\ApexDC-x64.exe"
	comixology = "www.comixology.com/free-comics"
	chookandgeek    = "https://www.chookandgeek.ru/collection/new?order=descending_age"
	book24 = "https://book24.ru/catalog/komiksy-manga-artbuki--1710/?sort=date&by=desc"
)

func NewComicsDay() {
	StartApp(dc)
	StartChrome(comixology)
	StartOpera(chookandgeek)
	StartOpera(book24)
}
