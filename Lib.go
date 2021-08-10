package main

import "log"

const (
	Space = " "
)

func Check(error error, message ...string) {
	if error != nil {
		log.Print(error)
		log.Print(message)
	}

}
