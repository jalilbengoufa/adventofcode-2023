package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	for scanner.Scan() {
		line := scanner.Text()

		gameSplit := strings.Split(line, ": ")

		maxMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		gameRecords := strings.Split(gameSplit[1], "; ")
		for _, s := range gameRecords {

			rounds := strings.Split(s, ", ")
			for _, r := range rounds {

				roundSplit := strings.Split(r, " ")
				colorCount, _ := strconv.Atoi(roundSplit[0])

				if maxMap[roundSplit[1]] < colorCount {

					maxMap[roundSplit[1]] = colorCount

				}

			}

		}

		power := 1
		for _, colorV := range maxMap {

			power *= colorV
		}
		result += power
	}

	fmt.Println("result", result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
