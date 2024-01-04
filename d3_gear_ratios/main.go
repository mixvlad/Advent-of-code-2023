package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	partOne()
}

func regexNum(s string) int {
	re := regexp.MustCompile("[0-9]+")
	res := 0
	res, _ = strconv.Atoi(re.FindAllString(s, 1)[0])
	return res
}

func regexNums(s string) []string {
	re := regexp.MustCompile("[0-9]+")

	return re.FindAllString(s, -1)
}

func regexNumsPos(s string) [][]int {
	re := regexp.MustCompile("[0-9]+")

	return re.FindAllStringIndex(s, -1)
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

	row := 0
	for scanner.Scan() {
		curRow := scanner.Text()

		nums := regexNums(curRow)
		// numsPos := regexNumsPos(curRow)

		for _, num := range nums {
			fmt.Println(num)
		}

		board := make([][]string, len(curRow))
		for i := range board {
			board[i] = make([]string, len(curRow))
		}

		for k, v := range curRow {
			board[row][k] = string(v)
		}
		row++

		fmt.Println(curRow)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part one: ", sum)
	fmt.Println("Part two: ", sum2)
}
