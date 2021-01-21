package main

const (
	//Sites
	epicGames = "https://www.epicgames.com/store/ru/free-games"
	steamGames = "https://steamdb.info/upcoming/free/"
	itchIO = "https://www.steamgifts.com/discussion/zCtNz/free-itchio-games-and-everything-on-100-sale-sorted"
	indieGala = "https://freebies.indiegala.com/"
	pikabu = "https://pikabu.ru/community/free"
)

func FreeGamesThursday() {
	//Sites
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
