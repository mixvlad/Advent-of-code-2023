package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	partOne()
	partTwo()
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

	for scanner.Scan() {
		str := scanner.Text()
		val1, val2 := 0, 0
		for i := 0; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				val1, _ = strconv.Atoi(string(str[i]))
				break
			}
		}

		for i := len(str) - 1; i >= 0; i-- {
			if str[i] >= '0' && str[i] <= '9' {
				val2, _ = strconv.Atoi(string(str[i]))
				break
			}
		}

		sum += val1*10 + val2
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part one: ", sum)
}

func MatchNumber(str string) int {
	if str[0] >= '0' && str[0] <= '9' {
		res, _ := strconv.Atoi(string(str[0]))
		return res
	}

	numbers := [...]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i := 0; i < len(numbers); i++ {
		num := numbers[i]
		for j := 0; j < len(num); j++ {
			if len(num) > len(str) || str[j] != num[j] {
				break
			}
			if j == len(num)-1 {
				return i + 1
			}
		}
	}

	return 0
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		str := scanner.Text()

		val1, val2 := 0, 0
		for i := 0; i < len(str); i++ {
			val1 = MatchNumber(str[i:])
			if val1 > 0 {
				break
			}
		}

		for i := len(str) - 1; i >= 0; i-- {
			val2 = MatchNumber(str[i:])
			if val2 > 0 {
				break
			}
		}

		res := val1*10 + val2
		fmt.Println("Part two: ", res)

		sum += res
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part two: ", sum)
}
