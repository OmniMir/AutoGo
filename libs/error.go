package libs

import "log"

func Сheck(error error, message... string) {
	if error != nil {
		log.Print(error)
		log.Print(message)
	}

}