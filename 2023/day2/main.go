package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var mapColors = map[string]int{
	"blue":  14,
	"green": 13,
	"red":   12,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	_ = total
	for scanner.Scan() {
		string := scanner.Text()
		gameId, err := strconv.Atoi(string[5:strings.Index(string, ":")])
		if err != nil {
			log.Fatal(err)
		}
		stringSplit := strings.Split(string, ";")
		fmt.Println("The gameId:", gameId, "has this many rounds:", len(stringSplit))
		for i, value := range stringSplit {
			if i == 0 {
				value = value[strings.Index(value, ":")+1:]
			}
			fmt.Println("This is the substring:", value)

		}
		total += int(gameId)

	}
	fmt.Println("The total is:", total)
}

// Lemme think about what I need:
// For example, the record of a few games might look like this:

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
// Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
// Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
// Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
// Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
//
// 12 red cubes, 13 green cubes, and 14 blue cubes
//
// Games 1, 2 and 5 are valid, therefore result is the sum which is: 8
// Figure out if all the subsets are valid.
//
// Each line is a game, I need a way to parse the numbers the names, the comma, and the semicolon.
