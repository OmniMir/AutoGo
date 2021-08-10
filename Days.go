package main

//ANIME
const (
	newAnime = "https://kg-portal.ru/news/anime/preview/"
)

func NewAnimeSeason() {
	StartChrome(newAnime)
}

//FREE GAMES
const (
	epicGames  = "https://www.epicgames.com/store/ru/free-games"
	steamGames = "https://steamdb.info/upcoming/free/"
	itchIO     = "https://www.steamgifts.com/discussion/zCtNz/free-itchio-games-and-everything-on-100-sale-sorted"
	indieGala  = "https://freebies.indiegala.com/"
	pikabu     = "https://pikabu.ru/community/free"
)

func FreeGamesThursday() {
	StartChrome(epicGames)
	StartSleep()
	StartChrome(steamGames)
	StartSleep()
	StartChrome(itchIO)
	StartSleep()
	StartChrome(indieGala)
	StartSleep()
	StartChrome(pikabu)
	StartSleep()
}

//MOVIES
const (
	premieres = "www.kinopoisk.ru/premiere/ru/"
)

func ToMovies() {
	StartChrome(premieres)
}

//NEW COMICS
const (
	dcpp         = "C:\\Program Files\\ApexDC\\ApexDC-x64.exe"
	comixology   = "www.comixology.com/free-comics"
	chookandgeek = "https://www.chookandgeek.ru/collection/new?order=descending_age"
	book24       = "https://book24.ru/catalog/komiksy-manga-artbuki--1710/?sort=date&by=desk"
)

func NewComicsDay() {
	StartApp(dcpp)
	StartChrome(comixology)
	StartOpera(chookandgeek)
	StartOpera(book24)
}

//UPDATE
const (
	//Windows Applications
	sumo      = "C:\\Program Files (x86)\\SUMo\\SUMo.exe" //winUpdate not work -> run application as document
	winUpdate = "ms-settings:windowsupdate"
	kaspersky = "C:\\Program Files (x86)\\Kaspersky Lab\\Kaspersky Internet Security 21.2\\avpui.exe"
	//steam = "C:\\Program Files (x86)\\Steam\\Steam.exe"
	//Sites
	chromium = "chromium.woolyss.com/#windows-64-bit"
	repacks  = "https://repack.me/"
	newApps  = "https://betanews.com/?s=Best+Windows+10+apps+this+week"
	//CLI Utilities
)

func UpdateFriday() {
	//Swith to New Desktop
	StartCmd(VirtualDesktop, "/Switch:1")
	//Windows Applications
	StartApp(sumo)
	StartSleep()
	StartUWP(winUpdate)
	StartSleep()
	StartStore()
	StartSleep()
	//Sites
	StartChrome(chromium)
	StartSleep()
	StartChrome(repacks)
	StartSleep()
	StartChrome(newApps)
	StartSleep()
	//CLI Utilities
	StartTerminal()
	//Steam
	/*
		//Before need start service
		Start(steam, "")
	*/
	//Kaspersky Software Update
	StartApp(kaspersky)
	/*
		Some mouse magic))
	*/
}

//WRESTLING
const (
	wweRaw       = "https://pwnews.net/blog/1-0-2"
	wweSmackdown = "https://pwnews.net/blog/1-0-1"
)

func WrestlingDay1() {
	StartChrome(wweRaw)
}
func WrestlingDay2() {
	StartChrome(wweSmackdown)
}
