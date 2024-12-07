package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("data2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countErrors := 0
	countLevels := 0
	countErrorsAdj := 0
	for scanner.Scan() {
		countLevels += 1
		row := strings.Fields(scanner.Text())

		var level []int
		for i := 0; i < len(row); i++ {
			number, _ := strconv.Atoi(row[i])
			level = append(level, number)
		}

		inc := CheckIncrease(level)
		err, idx := CheckLevel(level, inc)
		if err == true {
			countErrors++

			adjLevel := append(level[:idx], level[idx+1:]...)
			errAdj, _ := CheckLevel(adjLevel, inc)
			if errAdj == true {
				countErrorsAdj++
			}
		}
	}

	fmt.Println(countLevels - countErrors)
	fmt.Println(countLevels - countErrorsAdj)
}

func CheckIncrease(level []int) bool {
	var increase bool
	increaseCount := 0
	decreaseCount := 0
	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]
		if diff > 0 {
			increaseCount++
		}
		if diff < 0 {
			decreaseCount++
		}
	}
	if increaseCount > decreaseCount {
		increase = true
	} else {
		increase = false
	}
	return increase
}

func CheckLevel(level []int, inc bool) (bool, int) {
	err := false
	var idx int
	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]

		if diff == 0 {
			err = true
			idx = i
			break
		}

		if inc == true {
			if diff > 3 {
				err = true
				idx = i
				break
			}
			if diff < 0 {
				err = true
				idx = i
				break
			}
		}

		if inc == false {
			if diff < -3 {
				err = true
				idx = i
				break
			}
			if diff > 0 {
				err = true
				idx = i
				break
			}
		}
	}
	return err, idx
}
