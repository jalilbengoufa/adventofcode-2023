package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type Point struct {
	y int
	x int
}

func main() {

	filename := "data.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0

	numbersMap := map[Point]string{}
	symMap := []Point{}
	index := 0
	for scanner.Scan() {
		line := scanner.Text()
		currentN := "?"
		for x, char := range line {
			if unicode.IsDigit(char) {
				if currentN == "?" {
					currentN = string(char)
				} else {

					currentN += string(char)

				}

				if x == len(line)-1 && currentN != "?" {

					numbersMap[Point{y: index, x: x}] = string(currentN)

				}
			} else if char != '.' {
				symMap = append(symMap, Point{y: index, x: x})

				if currentN != "?" {
					numbersMap[Point{y: index, x: x - 1}] = string(currentN)
					currentN = "?"
				}
			} else {

				if currentN != "?" {
					numbersMap[Point{y: index, x: x - 1}] = string(currentN)
					currentN = "?"
				}
			}
		}

		index++
	}

	gearsCount := make(map[Point][]int)
	for point, number := range numbersMap {

		nL := len(number)

		intNb, _ := strconv.Atoi(number)
		for _, sPoint := range symMap {

			if sPoint.y == point.y {
				if sPoint.x == point.x+1 || sPoint.x == point.x-nL {

					gearsCount[sPoint] = append(gearsCount[sPoint], intNb)

				}

			}

			if sPoint.y == point.y-1 {
				if sPoint.x <= point.x+1 && sPoint.x >= point.x-nL {
					gearsCount[sPoint] = append(gearsCount[sPoint], intNb)
				}

			}
			if sPoint.y == point.y+1 {
				if sPoint.x <= point.x+1 && sPoint.x >= point.x-nL {
					gearsCount[sPoint] = append(gearsCount[sPoint], intNb)
				}
			}
		}

	}

	for _, gears := range gearsCount {

		if len(gears) == 2 {

			subR := 1
			for _, gear := range gears {

				subR *= gear

			}

			result += subR
		}
	}

	fmt.Println("result", result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
