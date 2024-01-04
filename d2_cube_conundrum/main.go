package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func main() {
	partOne()
}

func findMax(nums []string) int {
	max := 0
	for _, numColor := range nums {
		num := regexNum(numColor)
		if num > max {
			max = num
		}
	}
	return max
}

func regexNumColor(s string, color string) []string {
	re := regexp.MustCompile("[0-9]+ " + color)
	return re.FindAllString(s, -1)
}

func regexGame(s string) string {
	re := regexp.MustCompile("Game [0-9]+")
	return re.FindAllString(s, 1)[0]
}

func regexNum(s string) int {
	re := regexp.MustCompile("[0-9]+")
	res := 0
	res, _ = strconv.Atoi(re.FindAllString(s, 1)[0])
	return res
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	sum := 0
	sum2 := 0

	for scanner.Scan() {
		str := scanner.Text()
		reds := regexNumColor(str, "red")
		greens := regexNumColor(str, "green")
		blues := regexNumColor(str, "blue")

		redsNeeded := findMax(reds)
		greensNeeded := findMax(greens)
		bluesNeeded := findMax(blues)

		sum2 += redsNeeded * greensNeeded * bluesNeeded
		if MAX_RED < redsNeeded || MAX_GREEN < greensNeeded || MAX_BLUE < bluesNeeded {
			continue
		}

		game := regexGame(str)
		gameNum := regexNum(game)
		sum += gameNum
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part one: ", sum)
	fmt.Println("Part two: ", sum2)
}
