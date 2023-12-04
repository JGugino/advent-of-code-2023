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
		if isNum(splitString[one]) {
			numOne = splitString[one]
			break
		}
	}

	//find last number
	for two := len(line) - 1; two >= 0; two-- {
		if isNum(splitString[two]) {
			numTwo = splitString[two]
			break
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
