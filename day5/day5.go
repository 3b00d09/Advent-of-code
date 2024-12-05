// 5:47

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[int][]int)
	updates := make([][]int, 0)
	// flag to tell us what to expect. the updates will come after a line break
	finishedRules := false

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			finishedRules = true
			continue
		}
		if !finishedRules {
			numbers := strings.Split(line, "|")
			firstNum, err := strconv.Atoi(numbers[0])
			if err != nil {
				log.Fatal(err.Error())
			}
			secondNum, err := strconv.Atoi(numbers[len(numbers)-1])
			if err != nil {
				log.Fatal(err.Error())
			}
			rules[firstNum] = append(rules[firstNum], secondNum)
		} else {
			numbers := strings.Split(line, ",")
			row := make([]int, 0)
			for _, number := range numbers {
				tmp, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err.Error())
				}
				row = append(row, tmp)
			}
			updates = append(updates, row)
		}
	}

	sum := 0

outerLoop:
	for _, row := range updates {

		// loop over every number in the row
		for i, num := range row {
			// if the number exists as a key, then we know it has a set of rules
			if _, exists := rules[num]; exists {
				// loop over the row again and see if there are elements that exists in our set of rules
				for idx, num2 := range row {
					exists := existsInRules(num2, rules[num])
					// if a number exists in our set of rules, we need to make sure we exist before it
					if exists {
						if i > idx {
							continue outerLoop
						}
					}
				}
			}
		}
		midpointNumber := row[len(row)/2]
		sum += midpointNumber
	}

	fmt.Println(sum)
}

func existsInRules(targetNum int, rules []int) bool {
	for _, num := range rules {
		if num == targetNum {
			return true
		}
	}
	return false
}
