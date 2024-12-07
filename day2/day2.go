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
	countGood := 0
	countGoodAdj := 0
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())

		var level []int
		for i := 0; i < len(row); i++ {
			number, _ := strconv.Atoi(row[i])
			level = append(level, number)
		}

		inc := CheckIncrease(level)
		err := CheckLevel(level, inc)
		if err == false {
			countGood++
		} else {
			for i := 0; i < len(level); i++ {
				var adjLevel []int
				adjLevel = append(adjLevel, level[:i]...)
				adjLevel = append(adjLevel, level[i+1:]...)
				errAdj := CheckLevel(adjLevel, inc)

				if errAdj == false {
					countGoodAdj++
					break
				}
			}
		}
	}
	fmt.Println(countGood)
	fmt.Println(countGood + countGoodAdj)
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

func CheckLevel(level []int, inc bool) bool {
	err := false
	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]

		if diff == 0 {
			err = true
			break
		}

		if inc == true {
			if diff > 3 {
				err = true
				break
			}
			if diff < 0 {
				err = true
				break
			}
		}

		if inc == false {
			if diff < -3 {
				err = true
				break
			}
			if diff > 0 {
				err = true
				break
			}
		}
	}
	return err
}
