package main

import (
	. "github.com/OmniMir/AutoGo/libs"
)

const (
	premieres = "www.kinopoisk.ru/premiere/ru/"
)

func main() {
	StartChrome(premieres)
}
