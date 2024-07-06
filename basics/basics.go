package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	low := 1
	high := 100
	fmt.Println("Guess a number between, ", low, "and ", high)
	scanner.Scan()

	trycount := 0
	for {
		guess := (low + high) / 2
		trycount++
		fmt.Println("I think the number is", guess)
		fmt.Println("Is that: ")
		fmt.Println("(a) too high")
		fmt.Println("(b) too low")
		fmt.Println("(c) correct")
		scanner.Scan()
		response := scanner.Text()

		if response == "a" {
			high = guess - 1
		} else if response == "b" {
			low = guess + 1
		} else if response == "c" {
			fmt.Println("I Win! It took me", trycount, "tries")
			break
		} else {
			fmt.Println("Invalid response, try again")
		}
	}
}
