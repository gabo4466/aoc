package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numberWords = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func ReplaceWordsWithDigits(s string) (newWord string) {
	newWord = s
	for number, word := range numberWords {
		index := strings.Index(newWord, word)
		for index != -1 {
			newWord = newWord[:index+1] + strconv.Itoa(number+1) + newWord[index+1:]
			index = strings.Index(newWord, word)
		}
	}
	fmt.Println(newWord)
	return newWord
}

func main() {
	file, err := os.ReadFile("01/input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	const chars = "1234567890"
	var result int
	for _, line := range strings.Fields(string(file)) {
		word := ReplaceWordsWithDigits(line)
		number, err := strconv.Atoi(string(word[strings.IndexAny(word, chars)]) + string(word[strings.LastIndexAny(word, chars)]))
		if err != nil {
			log.Fatal(err)
			return
		}
		result += number
	}
	fmt.Println(result)
}
