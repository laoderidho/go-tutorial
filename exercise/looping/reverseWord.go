package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := "Aku Sayang Ibu"
	reversed := ReverseWord(input)
	fmt.Println(reversed)
}

func ReverseWord(str string) string {
	words := strings.Fields(str) // Memisahkan kata-kata dalam string
	reversedWords := make([]string, len(words))

	for i, word := range words {
		reversedWords[i] = Reverse(word)
	}

	return strings.Join(reversedWords, " ")
}

func Reverse(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isCapitalized(s string) bool {
	r := []rune(s)
	return unicode.IsUpper(r[0])
}
