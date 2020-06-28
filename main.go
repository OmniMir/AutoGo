package main

import (
	"flag"
	"fmt"
)

func main() {
	//Reading special flags
	comics := flag.Bool("comics", false, "Comics Increment")
	google := flag.Bool("google", false, "Call to Google")
	movies := flag.Bool("movies", false, "Comics Increment")
	opera := flag.Bool("opera", false, "From Chrome to Opera")
	paradigm := flag.Bool("paradigm", false, "From Chrome to Opera")
	update := flag.Bool("update", false, "Comics Increment")
	flag.Parse()
	//Choosing right mode
	if *comics {
		ComicsIncrement()
	} else if *google {
		GoogleIt()
	} else if *movies {
		ToMovies()
	} else if *opera {
		OperaIt()
	} else if *paradigm {
		ParadigmShift()
	} else if *update {
		UpdateFriday()
	} else {
		fmt.Println("You need right flag")
	}

}
