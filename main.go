package main

import (
	"fmt"
	"strings"
)

var alphabet = [26]int{1, 3, 3, 4, 1, 4, 4, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

var pl = fmt.Println

func main() {
	// testing 123 567

	text := ""
	fmt.Print("Enter your scrabble word: ")
	fmt.Scanf("%s", &text)
	text = strings.ToUpper(text)

	score := 0
	for _, ltr := range text {
		ltr := string(ltr)
		ascii := []byte(ltr)[0] - 65
		pl("Letter:", ltr, "\tValue:\t", alphabet[ascii])
		score += alphabet[ascii]
	}

	pl("Your score:", score)

}
