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
		log.Fatal(err.Error())
	}

	defer file.Close()

	var leftCol []int
	var rightCol []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// split into array by space
		// because the input has several spaces between the two numbers then we will have empty elements in our array
		// but we can guarantee that the first and last elements of the array are valid numbers, everything in-between is a space
		splitLine := strings.Split(line, " ")

		// convert the read text from string to int
		leftValue, err := strconv.Atoi(splitLine[0])

		if err != nil {
			log.Fatal(err.Error())
		}

		rightValue, err := strconv.Atoi(splitLine[len(splitLine)-1])

		if err != nil {
			log.Fatal(err.Error())
		}

		leftCol = append(leftCol, leftValue)
		rightCol = append(rightCol, rightValue)
	}

	sort.Ints(leftCol)
	sort.Ints(rightCol)

	totalDistance := 0

	for i := range len(leftCol) {
		distanceBetween := leftCol[i] - rightCol[i]

		// math.abs only works for floats for whatever reason
		if distanceBetween < 0 {
			distanceBetween = -1 * distanceBetween
		}

		totalDistance += distanceBetween
	}

	fmt.Println(totalDistance)

	//================================================//
	// part two

	similarityMap := make(map[int]int)

	// fill the map with keys from left col and values 0
	for _, num := range leftCol {
		similarityMap[num] = 0
	}

	// iterate over right col and check if the element in right col exists as a key in our map
	for _, num := range rightCol {
		if _, exists := similarityMap[num]; exists {
			similarityMap[num]++
		}
	}

	similarityScore := 0
	for key, value := range similarityMap {
		if value != 0 {
			similarityScore += key * value
		}
	}

	fmt.Println(similarityScore)

}
