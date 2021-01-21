package main

const (
	wweRaw = "https://pwnews.net/blog/1-0-2"
	wweSmackdown = "https://pwnews.net/blog/1-0-1"
)

func WrestlingDay() {
	StartChrome(wweRaw)
	StartChrome(wweSmackdown)
}
