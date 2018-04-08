package main

import "libs"

const (
	premieres = "http://www.kinopoisk.ru/premiere/ru/"
)

func main() {
	libs.StartChrome(premieres)
}
