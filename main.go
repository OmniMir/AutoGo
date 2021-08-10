package main

import (
	"flag"
	"fmt"
)

func main() {
	//Reading special flags
	alphabet := flag.Bool("alphabet", false, "Sort Clipboard in Alphabetical Order")
	anime := flag.Bool("anime", false, "View All New Anime")
	chrome := flag.Bool("chrome", false, "From Chromium to Chrome")
	comics := flag.Bool("comics", false, "Comics Increment")
	doublecmd := flag.Bool("doublecmd", false, "From Explorer to Double Commander")
	edge := flag.Bool("edge", false, "From Chromium to Edge")
	freegames := flag.Bool("freegames", false, "Free Games from Digital Stores")
	google := flag.Bool("google", false, "Call to Google")
	movies := flag.Bool("movies", false, "Comics Increment")
	newcomics := flag.Bool("newcomics", false, "New Comics at Wednesday")
	opera := flag.Bool("opera", false, "From Chromium to Opera")
	paradigm := flag.Bool("paradigm", false, "Paradigm Shift")
	tagscanner := flag.Bool("tagscanner", false, "From Explorer to TagScanner")
	terminal := flag.Bool("terminal", false, "Windows Terminal")
	update := flag.Bool("update", false, "Update Your System at Friday")
	winrar := flag.Bool("winrar", false, "From Explorer to WinRAR")
	wrestling1 := flag.Bool("wrestling1", false, "View All Raw Wrestling")
	wrestling2 := flag.Bool("wrestling2", false, "View All SmackDown Wrestling")
	flag.Parse()

	//Choosing right mode
	if *alphabet {
		AlphabetSort()
	} else if *anime {
		ComicsIncrement()
	} else if *comics {
		ComicsIncrement()
	} else if *chrome {
		ChromeIt()
	} else if *doublecmd {
		DoubleCommanderIt()
	} else if *edge {
		EdgeIt()
	} else if *freegames {
		FreeGamesThursday()
	} else if *google {
		GoogleIt()
	} else if *movies {
		ToMovies()
	} else if *newcomics {
		NewComicsDay()
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
	} else if *wrestling1 {
		WrestlingDay1()
	} else if *wrestling2 {
		WrestlingDay2()
	} else {
		fmt.Println("You need right flag")
	}
}
