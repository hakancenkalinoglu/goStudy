package main

import (
	"fmt"
	"math/rand"
)

func PlayGuessGame() {
	targetNumber := rand.Intn(100) + 1
	var guess int
	attempts := 0

	fmt.Println("\n=== Welcome to the Number Guessing Game ===")
	fmt.Println("I have chosen a number between 1 and 100.")
	fmt.Println("Can you guess what it is?")
	fmt.Println("-------------------------------------------")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		attempts++

		if guess < targetNumber {
			fmt.Println("Too low! Try a bigger number.")
		} else if guess > targetNumber {
			fmt.Println("Too high! Try a smaller number.")
		} else {
			fmt.Printf("\n🎉 Congratulations! You found the number %v in %v attempts!\n", targetNumber, attempts)
			break
		}

		if attempts == 5 {
			fmt.Printf("\nGame Over! You've used all 5 attempts. The number was %v.\n", targetNumber)
			break
		}
	}
}
