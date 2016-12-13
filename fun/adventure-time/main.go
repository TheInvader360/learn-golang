package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type location struct {
	Name        string
	Exits       map[string]int
	Description string
}

func main() {
	//Build a map of all locations, set a starting location, and print a welcome message
	locations := make(map[int]location)
	locations[0] = location{"Development Office", map[string]int{"east": 1}, "The room is nothing special - it's a pretty standard windowless IT office, but it's full to the brim with awesome developers that brighten the place up a bit..."}
	locations[1] = location{"Foyer", map[string]int{"north": 4, "east": 2, "west": 0}, "This is the main visitor entrance vestibule of a large warehousing unit. There's a door to main reception to the north, a door with a biohazard warning sign on it to the west, and a beautiful glass panelled door to the east affording a view of a most magnificent car park... You can't help but notice a horrific buzzing sound coming from the fire alarm control unit mounted on the wall."}
	locations[2] = location{"Car Park", map[string]int{"north": 3, "west": 1}, "A sea of grey asphalt topped with grit, home to about a hundred cars during office hours. It's dull, grey, and miserable, but as soon as you lift your head up a little you see the natural beauty of the magnificent snow capped mountains in the distance... You're not in Slough now baby, this is Snowdonia! There's a large warehouse directly west of you, and some dangerous looking badlands to the north."}
	locations[3] = location{"Wilderness", map[string]int{"south": 2}, "You've ventured out beyond the workplace perimiter. Rumour has it that secret societies regularly meet in the area, keep an eye out for any strange phenomena!"}
	locations[4] = location{"Main Reception", map[string]int{"south": 1}, "There doesn't seem to be anyone around..."}
	var currentLocation = locations[1]
	fmt.Print("\nSUPER AWESOME TBP ADVENTURE\n===========================\n")
	//Main game loop
	for {
		fmt.Println("\nYou are in the " + strings.ToUpper(currentLocation.Name) + "\n\n" + currentLocation.Description)
		fmt.Print("\nYou can go:\n")
		for key := range currentLocation.Exits {
			fmt.Printf("  %s\n", key)
		}
		fmt.Println("\nWhich direction will you go?")
		//Read player input, trim whitespace, and handle quit command
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "quit" {
			os.Exit(0)
		}
		//If player input matches a valid exit in the current location, update current location - set it to be the destination/value of the specified exit
		if value, exists := currentLocation.Exits[strings.ToLower(text)]; exists {
			currentLocation = locations[value]
		} else {
			fmt.Println("\nYou can't go that way!")
		}
	}
}
