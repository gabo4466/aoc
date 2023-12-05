package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	green int
	blue  int
	red   int
}

var config = Round{
	red:   12,
	green: 13,
	blue:  14,
}

func main() {
	file, _ := os.ReadFile("02/input.txt")
	var gameIdsSum int
	var minSum int
	for i, line := range strings.Split(string(file), "\n") {
		isPlayable := true
		var minViable Round
		for _, roundsInput := range strings.Split(strings.Split(line, ":")[1], ";") {
			var round Round
			for _, cubesInput := range strings.Split(roundsInput, ",") {
				trimmed := strings.Split(strings.TrimSpace(cubesInput), " ")
				amount, _ := strconv.Atoi(trimmed[0])
				switch trimmed[1] {
				case "blue":
					round.blue = amount
					minViable.blue = int(math.Max(float64(minViable.blue), float64(amount)))
				case "green":
					round.green = amount
					minViable.green = int(math.Max(float64(minViable.green), float64(amount)))
				case "red":
					round.red = amount
					minViable.red = int(math.Max(float64(minViable.red), float64(amount)))
				}
			}
			if !CanPlayWithConfig(round, config) {
				isPlayable = false
			}
		}
		if isPlayable {
			gameIdsSum += i + 1
		}
		minSum += minViable.red * minViable.green * minViable.blue
	}
	fmt.Println(gameIdsSum)
	fmt.Println(minSum)
}

func CanPlayWithConfig(round Round, config Round) (canPlay bool) {
	canPlay = true
	if round.red > config.red {
		canPlay = false
	}
	if round.blue > config.blue {
		canPlay = false
	}
	if round.green > config.green {
		canPlay = false
	}
	return
}
