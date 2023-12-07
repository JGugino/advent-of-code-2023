package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	RedCubes   int = 12
	GreenCubes int = 13
	BlueCubes  int = 14
)

type Game struct {
	Id   int
	Sets []Cubes
}

type Cubes struct {
	Id    int
	Blue  int
	Red   int
	Green int
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Invalid command - <COMMAND> 'INPUT'")
		return
	}

	startTime := time.Now()

	//ADVENT DAY TWO -- START

	filePath := os.Args[1]

	fileLines := FileToStringSlice(filePath)

	deconstructedGames := make([]Game, 0)

	for _, g := range fileLines {
		deconstructedGames = append(deconstructedGames, mustDeconstructGameString(g))
	}

	validGames := determineValidGames(deconstructedGames)

	for _, v := range validGames {
		fmt.Println(v)
	}
	//ADVENT DAY TWO -- END

	timeElapsed := time.Since(startTime)

	fmt.Printf("Took: %s", timeElapsed)
}

func mustDeconstructGameString(game string) Game {
	var createdGame Game

	splitGameString := strings.Split(game, ":")

	if len(splitGameString) < 2 {
		panic(fmt.Sprintf("Invalid Game: %s", game))
	}

	gameID, err := strconv.Atoi(strings.Split(splitGameString[0], " ")[1])

	if err != nil {
		panic("Unable to convert game Id. Is it a number?")
	}

	gameSets := strings.Split(splitGameString[1], ";")

	sets := make([]Cubes, 0)

	for i, s := range gameSets {
		splitSets := strings.Split(s, ",")
		var r, g, b int
		for _, v := range splitSets {
			cubeSplit := strings.Split(v, " ")

			var err error

			if cubeSplit[2] == "red" {
				r, err = strconv.Atoi(cubeSplit[1])
			} else if cubeSplit[2] == "blue" {
				b, err = strconv.Atoi(cubeSplit[1])
			} else if cubeSplit[2] == "green" {
				g, err = strconv.Atoi(cubeSplit[1])
			}

			if err != nil {
				panic("Invalid Game Set")
			}
		}

		sets = append(sets, Cubes{
			Id:    i,
			Red:   r,
			Blue:  b,
			Green: g,
		})
	}

	createdGame = Game{
		Id:   gameID,
		Sets: sets,
	}

	return createdGame
}

func determineValidGames(games []Game) []Game {
	validGames := make([]Game, 0)

	for _, g := range games {
		gameValid := false

		for _, s := range g.Sets {
			if s.Red <= RedCubes {
				if s.Blue <= BlueCubes {
					if s.Green <= GreenCubes {
						gameValid = true
					}
				}
			}
		}

		if gameValid {
			validGames = append(validGames, g)
		}
	}

	return validGames
}

func FileToStringSlice(filePath string) []string {
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
