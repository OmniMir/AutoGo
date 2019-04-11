package main

import (
	. "github.com/OmniMir/AutoGo/libs"
)

const (
	premieres = "http://www.kinopoisk.ru/premiere/ru/"
)

func main() {
	StartChrome(premieres)
}
