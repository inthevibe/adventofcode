package main

import (
	"slices"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	// Slice of calories each elf is carrying
	var caloriesElf []int
	
	// Opening the input file
	file, err := os.Open("input.txt")
	if err != nil{
		fmt.Println("Error openning file: ", err)
		return
	}
	defer file.Close()
	
	// Start the scanner
	scanner := bufio.NewScanner(file) 
	
	// Calories the current elf is carrying
	currentElfCalories := 0
	
	// For scanning each line
	for scanner.Scan(){
	line := scanner.Text()
	
		//Check if the line is empty
		if strings.TrimSpace(line) == ""{
			caloriesElf = append(caloriesElf, currentElfCalories)
			currentElfCalories = 0
		} else {
		//Convert the line string into a int
		number, _ := strconv.Atoi(line)
		//Add to the number of calories the elf is carrying
		currentElfCalories = currentElfCalories + number			
		}
	}
	//Print the final result
	fmt.Println(slices.Max(caloriesElf))
	//Check for errors in the scanner
	if err := scanner.Err(); err != nil{
		fmt.Println("Error scanning file:", err)
	}
}

/*
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
*/
