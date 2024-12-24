package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var LL []int
	var RR []int
	RC := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		left, right := 0, 0
		parts := strings.SplitN(line, "   ", 2)
		left, err = strconv.Atoi(parts[0])
		right, err = strconv.Atoi(parts[1])
		LL = append(LL, left)
		RR = append(RR, right)
		RC[right]++
	}

	sort.Ints(LL)
	sort.Ints(RR)
	total := 0
	for i := 0; i < len(LL); i++ {
		distance := LL[i] - RR[i]
		if distance < 0 {
			distance = -distance
		}
		total += distance
	}

	total2 := 0
	for i := 0; i < len(LL); i++ {
		total2 += LL[i] * RC[LL[i]]
	}

	fmt.Println("The total distance is:", total)
	fmt.Println("The total distance 2 is:", total2)
}
