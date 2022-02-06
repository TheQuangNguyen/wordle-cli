package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	c "wordle-cli/constants"
)

var mword map[rune]int = make(map[rune]int)

func main() {
	fmt.Printf(c.RULES)
	numGuessLeft := c.NUM_TRIES
	word := getWordOfDay()
	for index, char := range word {
		mword[char] = index
	}

	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("\n" + word)

	for numGuessLeft > 0 {
		guess, err := reader.ReadString('\n');
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again!")
			return
		}
	
		guess = strings.TrimSuffix(guess, "\n")
		if isWordValid(guess) {
			if guess == word {
				fmt.Println(string(c.GREEN), strings.ToUpper(guess))
				fmt.Println(c.WIN_MESSAGE)
				return 
			}
			printGuessResult(guess, word)
			numGuessLeft--
			fmt.Printf("You have %d guesses left\n", numGuessLeft)
		} else {
			fmt.Println(guess + " is not a valid word. Please try again.")
		}
	}
	fmt.Println("Today WORDLE was " + strings.ToUpper(word))
	fmt.Println(c.LOSE_MESSAGE)
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
	for gIndex, char := range guess {
		if wIndex, isPresent := mword[char]; isPresent {
			if gIndex == wIndex {
				fmt.Printf(c.GREEN + "%s", strings.ToUpper(string(char)))
			} else {
				fmt.Printf(c.BLUE + "%s", strings.ToUpper(string(char)))
			}
		} else {
			fmt.Printf(c.RESET + "%s", strings.ToUpper(string(char)))
		}
	}
	fmt.Println(c.RESET)
}