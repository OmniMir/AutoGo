package main

import (
	"flag"
	"fmt"
)

func main() {
	//Reading special flags
	comics := flag.Bool("comics", false, "Comics Increment")
	doublecmd := flag.Bool("doublecmd", false, "From Explorer to Double Commander")
	edge := flag.Bool("edge", false, "From Chrome to Edge")
	google := flag.Bool("google", false, "Call to Google")
	movies := flag.Bool("movies", false, "Comics Increment")
	opera := flag.Bool("opera", false, "From Chrome to Opera")
	paradigm := flag.Bool("paradigm", false, "Paradigm Shift")
	tagscanner := flag.Bool("tagscanner", false, "From Explorer to TagScanner")
	terminal := flag.Bool("terminal", false, "Windows Terminal")
	update := flag.Bool("update", false, "Update Your System")
	winrar := flag.Bool("winrar", false, "From Explorer to WinRAR")
	flag.Parse()

	//Choosing right mode
	if *comics {
		ComicsIncrement()
	} else if *doublecmd {
		DoubleCommanderIt()
	} else if *edge {
		EdgeIt()
	} else if *google {
		GoogleIt()
	} else if *movies {
		ToMovies()
	} else if *opera {
		OperaIt()
	} else if *paradigm {
		ParadigmShift()
	} else if *tagscanner {
		TagScannerIt()
	} else if *terminal {
		TerminalIt()
	} else if *update {
		UpdateFriday()
	} else if *winrar {
		WinRarIt()
	} else {
		fmt.Println("You need right flag")
	}
}
