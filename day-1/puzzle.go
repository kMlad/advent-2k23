package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kMlad/advent-2k23/utils"
)

func Day1Test(testLine string) {
	reversedTestLine := utils.Reverse(testLine)

	firstDigit := utils.FindFirstDigit(testLine)
	lastDigit := utils.FindFirstDigit(reversedTestLine)

	fmt.Println(firstDigit, lastDigit)
}

func Day1Part1() {
	sum := 0
	inputLines := strings.Split(PuzzleInput, "\n")

	for _, line := range inputLines {
		reversedLine := utils.Reverse(line)

		firstDigit := utils.FindFirstDigit(line)
		lastDigit := utils.FindFirstDigit(reversedLine)

		finalNumberStringified := strconv.Itoa(firstDigit) + strconv.Itoa(lastDigit)
		finalNumber, err := strconv.Atoi(finalNumberStringified)
		if err == nil {
			sum += finalNumber
		}

	}

	fmt.Println(sum)
}
