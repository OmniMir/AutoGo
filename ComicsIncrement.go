package main

import (
	"github.com/go-vgo/robotgo/clipboard"
	"strconv"
	"strings"
)

const (
	hashtag = "#"
)

func ComicsIncrement() {

	//Getting existed information
	existedName, _ := clipboard.ReadAll()
	if !strings.Contains(existedName, hashtag) {
		return
	}
	existedTitle, existedNumber := stringInTwo(existedName, hashtag)

	//Only Number
	number := ""
	appendix := ""
	if strings.Contains(existedNumber, Space) {
		number, appendix = stringInTwo(existedNumber, Space)
		appendix = Space + appendix
	} else {
		number = existedNumber
	}

	//Begin with 0 and 09
	prefix := ""
	if strings.HasPrefix(number, "0") {
		prefix = "0"
	}
	if strings.HasPrefix(number, "09") {
		prefix = ""
	}

	//Ranging
	if strings.Contains(existedNumber, "*") {
		//TODO
	}

	//Incrementing
	numberInt, _ := strconv.Atoi(number)
	numberInt = numberInt + 1
	newNumber := strconv.Itoa(numberInt)

	//New numbering
	newTitleAndNumber := existedTitle + hashtag + prefix + newNumber + appendix
	clipboard.WriteAll(newTitleAndNumber)
}

func stringInTwo(full string, delimiter string) (first string, second string) {
	numberAndAppendix := strings.SplitN(full, delimiter, 2)
	return numberAndAppendix[0], numberAndAppendix[1]
}
