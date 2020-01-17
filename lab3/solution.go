package lab3

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	wordOccurrences := make(map[string]int)

	for i := range words {
		currentWord := words[i]

		v, ok := wordOccurrences[currentWord]

		if !ok {
			wordOccurrences[currentWord] = 1
			continue
		}

		wordOccurrences[currentWord] = v + 1
	}

	return wordOccurrences
}

func Run() {
	wc.Test(WordCount)
}
