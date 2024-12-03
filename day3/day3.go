package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var plainText string = ""

	for scanner.Scan() {
		line := scanner.Text()
		plainText += line
	}

	// will use sliding window to check if we have a match
	// we are looking for keypharse mul( then we can check the next n characters until we reach a )
	startIdx := 0
	jumpIdx := 4
	endIdx := (len(plainText) - 1) - jumpIdx
	var expressions []string
	var validExpressions []string
	var sum int = 0

	for startIdx < endIdx {
		keypharse := plainText[startIdx : startIdx+jumpIdx]
		if keypharse == "mul(" {
			pointer := startIdx + jumpIdx
			// look for a closing paranthesis
			for pointer < endIdx {
				char := rune(plainText[pointer])
				if char == ')' {
					break
				} else {
					pointer++
				}
			}
			// gathering all potential expressions
			expression := plainText[startIdx:startIdx+jumpIdx] + plainText[startIdx+jumpIdx:pointer+1]
			expressions = append(expressions, expression)
		}
		startIdx++
	}

outerLoop:
	for _, e := range expressions {
		start := e[4:]
		for _, el := range start {
			char := rune(el)
			if isInvalidChar(char) {
				continue outerLoop
			}
		}
		validExpressions = append(validExpressions, e)
	}

	for _, el := range validExpressions {
		// remove mul( and the closing paranthesis
		el = el[4 : len(el)-1]
		commaIdx := indexOf(el, ',')
		// we ignore any elements that were parsed but dont have a , (spare me my algorithm isnt perfect)
		if commaIdx == -1 {
			continue
		}
		_num1 := el[0:commaIdx]
		_num2 := el[commaIdx+1:]

		num1, err := strconv.Atoi(_num1)
		if err != nil {
			log.Fatal(err.Error())
		}

		num2, err := strconv.Atoi(_num2)
		if err != nil {
			log.Fatal(err.Error())
		}

		sum += num1 * num2
	}

	fmt.Println(sum)
	partTwo()
}

func partTwo() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var plainText string = ""

	for scanner.Scan() {
		line := scanner.Text()
		plainText += line
	}

	// will use sliding window to check if we have a match
	// we are looking for keypharse mul( then we can check the next n characters until we reach a )
	startIdx := 0
	jumpIdx := 4
	endIdx := len(plainText) - 1
	var expressions []string
	var validExpressions []string
	var sum int = 0
	// these are similar to how jump index work but instead of mul( its do() and don't()
	doLength := 4
	dontLength := 7

	for startIdx < endIdx-dontLength {
		// if we hit a dont
		potentialDont := plainText[startIdx : startIdx+dontLength]
		if potentialDont == "don't()" {
			// look for the next available do so we can skip our startIdx to it as theres no point checking for mul( between a dont() and a do()
			for startIdx < endIdx-doLength {
				potentialDo := plainText[startIdx : startIdx+doLength]
				if potentialDo == "do()" {
					// once we find it we can break out of this loop
					break
				}
				startIdx++
			}
		}

		// rest is literally the same
		keypharse := plainText[startIdx : startIdx+jumpIdx]
		if keypharse == "mul(" {
			pointer := startIdx + jumpIdx
			// look for a closing paranthesis
			for pointer < endIdx-jumpIdx-1 {
				char := rune(plainText[pointer])
				if char == ')' {
					break
				} else {
					pointer++
				}
			}
			// gathering all potential expressions
			expression := plainText[startIdx:startIdx+jumpIdx] + plainText[startIdx+jumpIdx:pointer+1]
			expressions = append(expressions, expression)
		}
		startIdx++
	}

outerLoop:
	for _, e := range expressions {
		start := e[4:]
		for _, el := range start {
			char := rune(el)
			if isInvalidChar(char) {
				continue outerLoop
			}
		}
		validExpressions = append(validExpressions, e)
	}

	for _, el := range validExpressions {
		// remove mul( and the closing paranthesis
		el = el[4 : len(el)-1]
		commaIdx := indexOf(el, ',')
		// we ignore any elements that were parsed but dont have a , (spare me my algorithm isnt perfect)
		if commaIdx == -1 {
			continue
		}
		_num1 := el[0:commaIdx]
		_num2 := el[commaIdx+1:]

		num1, err := strconv.Atoi(_num1)
		if err != nil {
			log.Fatal(err.Error())
		}

		num2, err := strconv.Atoi(_num2)
		if err != nil {
			log.Fatal(err.Error())
		}

		sum += num1 * num2
	}

	fmt.Println(sum)
}

func isInvalidChar(c rune) bool {
	return c != ',' && c != ')' && !unicode.IsDigit(c)
}

func indexOf(s string, target rune) int {
	for i, r := range s {
		if r == target {
			return i
		}
	}
	return -1
}
