package main

import (
	. "sort"
	"strings"

	"github.com/go-vgo/robotgo/clipboard"
)

const newline = "\n"

func AlphabetSort() {
	//Get Text from Clipboard
	text, _ := clipboard.ReadAll()
	//Split Text to Array of Strings
	list := strings.Split(text, newline)
	//Sort Array of Strings
	Sort(Alphabetic(list))
	text = strings.Join(list, newline)
	//Get Text to Clipboard
	clipboard.WriteAll(text)
}

//Alphabet Sort Implementation
type Alphabetic []string

func (list Alphabetic) Len() int           { return len(list) }
func (list Alphabetic) Swap(i, j int)      { list[i], list[j] = list[j], list[i] }
func (list Alphabetic) Less(i, j int) bool { return list[i] < list[j] }
