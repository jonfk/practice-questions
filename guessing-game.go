package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const toGuess = 42

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Guess number : ")
		guessStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guessNum, err := strconv.Atoi(guessStr[:len(guessStr)-1])
		if err != nil {
			fmt.Printf("%v\n", err)
			fmt.Println("Please enter number")
			continue
		}
		if guessNum > toGuess {
			fmt.Println("Guess Lower")
		} else if guessNum < toGuess {
			fmt.Println("Guess Higher")
		} else {
			fmt.Println("Correct Guess")
			break
		}
	}
	fmt.Println("end")
}
