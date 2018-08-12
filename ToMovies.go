package main

import (
	. "./libs"
)

const (
	premieres = "http://www.kinopoisk.ru/premiere/ru/"
)

func main() {
	StartChrome(premieres)
}
