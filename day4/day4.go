package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	plainText := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for idx, char := range line {
			row[idx] = rune(char)
		}
		plainText = append(plainText, row)
	}
	xmasLen := 4
	sum := 0

	for idx, row := range plainText {
		i := 0
		for i < len(row) {

			if row[i] == 'X' {
				// every check has to make sure the edge elements are included as well, we account for that with - 1
				// if we have enough elements to check in front of us
				if i+xmasLen <= len(row) {
					if row[i+1] == 'M' && row[i+2] == 'A' && row[i+3] == 'S' {
						sum++
					}

					// if we have enough elements to check the right diagonal upwards
					if idx >= xmasLen-1 {
						if plainText[idx-1][i+1] == 'M' && plainText[idx-2][i+2] == 'A' && plainText[idx-3][i+3] == 'S' {
							sum++
						}
					}
				}

				// if we have enough elements to check behind us
				if i >= xmasLen-1 {
					if row[i-1] == 'M' && row[i-2] == 'A' && row[i-3] == 'S' {
						sum++
					}

					// if we have enough elements to check the left diagonal upwards
					if idx >= xmasLen-1 {
						if plainText[idx-1][i-1] == 'M' && plainText[idx-2][i-2] == 'A' && plainText[idx-3][i-3] == 'S' {
							sum++
						}
					}
				}

				// if we have enough elements to check under us
				if idx+xmasLen <= len(plainText) {
					if plainText[idx+1][i] == 'M' && plainText[idx+2][i] == 'A' && plainText[idx+3][i] == 'S' {
						sum++
					}

					// if we have enough elements to check the right diagonal downwards
					if i+xmasLen <= len(row) {
						if plainText[idx+1][i+1] == 'M' && plainText[idx+2][i+2] == 'A' && plainText[idx+3][i+3] == 'S' {
							sum++
						}
					}

					// if we have enough elements to check the left diagonal downwards
					if i >= xmasLen-1 {
						if plainText[idx+1][i-1] == 'M' && plainText[idx+2][i-2] == 'A' && plainText[idx+3][i-3] == 'S' {
							sum++
						}
					}
				}

				// if we have enough elements to check above us
				if idx >= xmasLen-1 {
					if plainText[idx-1][i] == 'M' && plainText[idx-2][i] == 'A' && plainText[idx-3][i] == 'S' {
						sum++
					}
				}

			}
			i++
		}
	}
	fmt.Println(sum)
	partTwo(plainText)

}

func partTwo(plainText [][]rune) {
	sum := 0

	// we will target every A we find that is NOT in the first col, first row, last col, or last row as these dont have valid corners
	for idx := 1; idx < len(plainText)-1; idx++ {
		row := plainText[idx]
		for i := 1; i < len(row)-1; i++ {
			if row[i] == 'A' {

				// collect the corner characters
				var corners []rune
				var upperRightCorner rune = '0'
				var upperLeftCorner rune = '0'
				var lowerRightCorner rune = '0'
				var lowerLeftCorner rune = '0'

				// save us time if a corner is not an M or an S then it can never be MAS
				upperRightCorner = plainText[idx-1][i+1]
				if upperRightCorner != 'M' && upperRightCorner != 'S' {
					continue
				}
				corners = append(corners, upperRightCorner)

				upperLeftCorner = plainText[idx-1][i-1]
				if upperLeftCorner != 'M' && upperLeftCorner != 'S' {
					continue
				}
				corners = append(corners, upperLeftCorner)

				lowerRightCorner = plainText[idx+1][i+1]
				if lowerRightCorner != 'M' && lowerRightCorner != 'S' {
					continue
				}
				corners = append(corners, lowerRightCorner)

				lowerLeftCorner = plainText[idx+1][i-1]
				if upperLeftCorner != 'M' && upperLeftCorner != 'S' {
					continue
				}
				corners = append(corners, lowerLeftCorner)

				// we need to make sure we have two of each M and S, and a character cannot meet itself on the opposite side
				mCount := 0
				sCount := 0

				for _, el := range corners {
					c := rune(el)
					if c == 'M' {
						mCount++
					} else if c == 'S' {
						sCount++
					}
				}

				// this takes care of MMMS or similar
				if mCount != 2 || sCount != 2 {
					continue
				}

				// check if a character meets itself on the opppsite side
				if corners[0] != corners[3] || corners[1] != corners[2] {
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
}
