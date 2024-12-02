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

	var reports [][]int
	safeReportCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, " ")
		var tmp []int

		for _, num := range row {
			_num, err := strconv.Atoi(num)

			if err != nil {
				log.Fatal(err.Error())
			}

			tmp = append(tmp, _num)
		}

		reports = append(reports, tmp)

	}

	// loop over every row in the 2d array
	// label to help us skip reports if a violation is found
outerLoop:
	for _, report := range reports {
		// if our first element is less than our last element, we expect a gradually increasing report
		if report[0] < report[len(report)-1] {
			// loop over every element in the row we are currently in
			for i := range len(report) - 1 {
				difference := AbsSubtract(report[i], report[i+1])
				if report[i] > report[i+1] || difference < 1 || difference > 3 {
					// if violation skip iteration
					continue outerLoop
				}
			}
			safeReportCount++
		} else {
			// gradually decreasing
			for i := range len(report) - 1 {
				difference := AbsSubtract(report[i], report[i+1])
				if report[i] < report[i+1] || difference < 1 || difference > 3 {
					continue outerLoop
				}
			}
			safeReportCount++
		}
	}

	fmt.Println(safeReportCount)

	// =====================================================//
	// part two

	safeReportCountPartTwo := 0

	// loop over every row in the 2d array
outerLoop2:
	for _, report := range reports {

		// if our first element is less than our last element, we expect a gradually increasing report
		if report[0] < report[len(report)-1] {
			i := 0
			// loop over the elements in the report and look for violations
			for i < len(report)-1 {
				difference := AbsSubtract(report[i], report[i+1])
				if report[i] > report[i+1] || difference < 1 || difference > 3 {
					// if a violation is found, we need to look into every combination of that report by taking away 1 element at a time and validating that report
					for idx := 0; idx < len(report); idx++ {
						// copy so it doesnt affect the report itself
						reportCopy := make([]int, len(report))
						copy(reportCopy, report)
						tmp := append(reportCopy[:idx], reportCopy[idx+1:]...)
						// validate copy
						isValid := isValidIncreasingReport(tmp)
						if isValid {
							// valid combination exists, add and skip current iteration
							safeReportCountPartTwo++
							continue outerLoop2
						}
					}
					// exit the loop if there is a violation no valid combination is found
					break
				}
				i++
				// if we reach the end with no problems
				if i == len(report)-1 {
					safeReportCountPartTwo++
				}
			}

		} else {
			// gradually decrease
			i := 0
			for i < len(report)-1 {
				difference := AbsSubtract(report[i], report[i+1])
				if report[i] < report[i+1] || difference < 1 || difference > 3 {
					for idx := 0; idx < len(report); idx++ {
						reportCopy := make([]int, len(report))
						copy(reportCopy, report)
						tmp := append(reportCopy[:idx], reportCopy[idx+1:]...)
						isValid := isValidDecreasingReport(tmp)
						if isValid {
							safeReportCountPartTwo++
							continue outerLoop2
						}
					}
					break
				}
				i++
				if i == len(report)-1 {
					safeReportCountPartTwo++
				}
			}
		}
	}
	fmt.Println(safeReportCountPartTwo)
}

func AbsSubtract(x int, y int) int {
	res := x - y
	if res < 0 {
		res = res * -1
	}
	return res
}

func isValidDecreasingReport(report []int) bool {
	i := 0
	for i < len(report)-1 {
		difference := AbsSubtract(report[i], report[i+1])
		if report[i] < report[i+1] || difference < 1 || difference > 3 {
			return false
		}
		i++
	}
	return true
}

func isValidIncreasingReport(report []int) bool {
	i := 0
	for i < len(report)-1 {
		difference := AbsSubtract(report[i], report[i+1])
		if report[i] > report[i+1] || difference < 1 || difference > 3 {
			return false
		}
		i++
	}
	return true
}
