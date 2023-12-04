package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Invalid command - <COMMAND> 'INPUT'")
		return
	}

	startTime := time.Now()

	//ADVENT DAY ONE - START PART ONE
	filePath := os.Args[1]

	fileLines := fileToSlice(filePath)

	calibrationValues := make([]int, 0)

	for _, v := range fileLines {
		calibrationValues = append(calibrationValues, findCalibrationValueFromString(v))
	}

	fmt.Println(calibrationValues)

	sum := 0

	for _, i := range calibrationValues {
		sum += i
	}

	fmt.Printf("Calibration Sum: %d \n", sum)

	//ADVENT DAY ONE -- END PART ONE

	timeElapsed := time.Since(startTime)

	fmt.Printf("Took: %s", timeElapsed)

}

func findCalibrationValueFromString(line string) int {
	splitString := strings.Split(line, "")

	numOne := ""
	numTwo := ""

	//find first number
	for one := 0; one <= len(line); one++ {
		digit, isDigit := isValidDigit(splitString, one, "fwd")

		if isNum(splitString[one]) {
			numOne = splitString[one]
			break
		} else if isDigit {
			numOne = string(rune(digit))
		}
	}

	//find last number
	for two := len(line) - 1; two >= 0; two-- {
		digit, isDigit := isValidDigit(splitString, two, "bck")
		if isNum(splitString[two]) {
			numTwo = splitString[two]
			break
		} else if isDigit {
			numOne = string(rune(digit))
		}
	}

	//Combine and convert numbers
	finalNumber := numOne + numTwo
	converted, _ := strconv.Atoi(finalNumber)

	//return number
	return converted
}

func fileToSlice(filePath string) []string {
	lines := make([]string, 0)

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Invalid file - %s \n", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid file - %s \n", err)
	}

	return lines
}

func isNum(input string) bool {
	_, err := strconv.Atoi(input)

	return err == nil
}

func isValidDigit(input []string, startIndex int, dir string) (int, bool) {
	threeCharDigits := []string{"one", "two", "six"}
	// fourCharDigits := []string{"four", "five", "nine"}
	// fiveCharDigits := []string{"three", "seven", "eight"}

	//Three Character Digits
	for _, t := range threeCharDigits {
		var checkingValue string

		//Makes sure we don't go outside of the slice
		if (startIndex+2 >= len(input)) || (startIndex-2 <= 0) {
			return 0, false
		}

		//gets the value to check
		if dir == "fwd" {
			checkingValue = input[startIndex] + input[startIndex+1] + input[startIndex+2]
		} else if dir == "bck" {
			checkingValue = input[startIndex-2] + input[startIndex-1] + input[startIndex]
		}

		if checkingValue == t {
			return returnNumFromString(t), true
		}
	}

	return 0, false
}

func returnNumFromString(number string) int {
	switch number {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return -1
	}
}
