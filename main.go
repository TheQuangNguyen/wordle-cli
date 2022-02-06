package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"

	c "wordle-cli/constants"
)

func main() {
	fmt.Printf(c.RULES)
	numGuessLeft := c.NUM_TRIES
	word := getWordOfDay()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n" + word)

	for numGuessLeft > 0 {
		guess, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again!")
			return
		}
	
		guess = strings.TrimSuffix(guess, "\n")
		if isWordValid(guess) {
			printGuessResult(guess, word)
		} else {
			fmt.Println(guess + " is not a valid word. Please try again.")
		}
	}
	
}

func getWordOfDay() string {
	for k, _ := range c.WORDS {
		return k
	}
	return ""
}

func isWordValid(word string) bool {
	if len(word) != c.WORD_LENGTH {
		return false
	}
	if _ , isValid := c.WORDS[word]; !isValid {
		return false
	}
	return true
}

func printGuessResult(guess string, word string) {
	bguess = []byte(guess)
	bword = []byte(word)

	
}
