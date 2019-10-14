package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	secretNum := generateRandInt(1, 100)
	fmt.Println("Guess a number between 1 and 100")

	attempts := 0
	for {
		attempts++

		fmt.Println("Input your guess now: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Provide an integer")
			continue
		}

		if guess > secretNum {
			fmt.Println("Your guess is larger than secret number. You have used", attempts, "attempts")
		} else if guess < secretNum {
			fmt.Println("Your guess is lower than secret number. You have used", attempts, "attempts")
		} else {
			fmt.Println("You got it right! You have used total", attempts, "attempts")
			break
		}

	}
}



func generateRandInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	fmt.
	return rand.Intn(max-min) + min
}
