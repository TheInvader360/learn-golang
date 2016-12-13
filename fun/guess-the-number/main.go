package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var active = true
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
var minValid = 10
var maxValid = 15

func main() {
	for active {
		var randomNumber = rng.Intn(maxValid-minValid+1) + minValid
		fmt.Println("Guess the number I'm thinking of (" + strconv.Itoa(minValid) + "-" + strconv.Itoa(maxValid) + ")")
		var playerNumber int
		_, err := fmt.Scanf("%d", &playerNumber)
		if err != nil {
			log.Fatal(err)
		}
		if playerNumber >= minValid && playerNumber <= maxValid {
			if playerNumber == randomNumber {
				fmt.Println("CORRECT!")
			} else {
				fmt.Println("You guessed " + strconv.Itoa(playerNumber) + " but the correct number was " + strconv.Itoa(randomNumber))
			}
		} else {
			fmt.Println("Invalid choice, the number is between " + strconv.Itoa(minValid) + " and " + strconv.Itoa(maxValid) + " inclusive...")
		}
		active = playAgain()
	}
}

func playAgain() bool {
	fmt.Println("Play again? (y/n)")
	var playAgain string
	fmt.Scanf("%s", &playAgain)
	playAgain = strings.ToLower(playAgain)
	if playAgain == "y" {
		return true
	}
	return false
}
