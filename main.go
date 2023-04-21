package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secretNumber := rand.Intn(100) + 1
	fmt.Println("The secret number is ", secretNumber)

	var guess int
	for i := 0; ; i++ {
		fmt.Print("Please input your guess: ")
		if _, err := fmt.Scan(&guess); err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			fmt.Scanln()
			continue
		}
		if guess == secretNumber {
			fmt.Printf("Congratulations! You guessed the number in %d tries.\n", i+1)
			break
		} else if guess < secretNumber {
			fmt.Println("low,again")
		} else {
			fmt.Println("high,again")
		}
	}
}
