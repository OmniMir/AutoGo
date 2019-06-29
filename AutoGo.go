package main

import (
	. "github.com/OmniMir/AutoGo/libs"
	"github.com/go-vgo/robotgo/clipboard"
	"strconv"
	"strings"
)

func main() {
		googleIt()
		operaIt()
		comicsIncrement()
	}

}

func googleIt() {
	const (
		emptySearch    = ""
		chromeSettings = "--new-window"
		chromeSearch   = "www.google.ru/search?q="
	)

	//Read clipboard
	wordsForSearch, _ := clipboard.ReadAll()
	//And to Chrome
	if wordsForSearch != emptySearch {
		chromeCommand := chromeSettings + Space + chromeSearch + wordsForSearch
		StartChrome(chromeCommand)
	}
}

func operaIt() {
	url, _ := clipboard.ReadAll()
	StartOpera(url)
}

func comicsIncrement() {
	const (
		hashtag = "#"
	)
	//Getting existed information
	existedName, _ := clipboard.ReadAll()
	if !strings.Contains(existedName,hashtag){
		return
	}
	existedTitle, existedNumber := stringInTwo(existedName, hashtag)

	//Only Number
	number := ""
	appendix := ""
	if strings.Contains(existedNumber,Space){
		number, appendix = stringInTwo(existedNumber, Space)
		appendix = Space + appendix
	} else {
		number = existedNumber
	}

	//Begin with Zero
	prefix := ""
	if strings.HasPrefix(number, "0") {
		prefix = "0"
	}

	//Ranging
	if strings.Contains(existedNumber,"*"){
	//TODO
	}

	//Incrementing
	numberInt, _ := strconv.Atoi(number)
	numberInt = numberInt + 1
	newNumber := strconv.Itoa(numberInt)

	//New numbering
	newTitleAndNumber := existedTitle + hashtag + prefix + newNumber + appendix
	fmt.Println(newTitleAndNumber)
	clipboard.WriteAll(newTitleAndNumber)
}

func stringInTwo (full string, delimiter string) (first string, second string) {
	numberAndAppendix := strings.SplitN(full, delimiter, 2)
	return numberAndAppendix[0], numberAndAppendix[1]
}
