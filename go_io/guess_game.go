package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // creates a new *bufio.Reader

	var input string
	var won bool = false
	var num, guessCount int = 41, 0

	fmt.Println("Guess the number between 1-100")
	for !won {

		input, _ = reader.ReadString('\n')
		input = strings.Split(input, "\n")[0]
		if input == "" {
			continue
		}
		var n, err = strconv.Atoi(input)

		guessCount++
		if err != nil {
			fmt.Println("Invalid integer input. Please try again!")
		} else {
			if n == num {
				fmt.Println("\n > You guessed it correct in", guessCount, "guesses < \n")
				won = true
			} else if n < num {
				fmt.Println("Try again you guess is too low.")
			} else {
				fmt.Println("Try again you guess is too high.")
			}
		}
	}
}
