package constants

import (
	"fmt"
)

const (
	RULES = `Guess the WORDLE in 6 tries
Each guess must be a valid 5 letter word. Hit the enter button to submit.
After each guess, the color of the letters will change to show how close your guess was to the word.
`
	NOT_ENOUGH_LETTERS = "Not enough letters. Please try again."
	WORD_NOT_FOUND = "Word not found. Please try again."
	WIN_MESSAGE = "Congratulations! You figured out today WORDLE."
	LOSE_MESSAGE = "You can do it next time!"
)

func CorrectWordMsg(solution string) string {
	return fmt.Sprintf("The word was %s.\n", solution)
}