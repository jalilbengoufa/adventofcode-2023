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
		valid := true
		line := scanner.Text()

		gameSplit := strings.Split(line, ": ")
		gameId, _ := strconv.Atoi(strings.Split(gameSplit[0], " ")[1])

		gameRecords := strings.Split(gameSplit[1], "; ")
		for _, s := range gameRecords {

			boundriesMap := map[string]int{
				"red":   12,
				"green": 13,
				"blue":  14,
			}
			rounds := strings.Split(s, ", ")
			for _, r := range rounds {

				roundSplit := strings.Split(r, " ")
				colorCount, _ := strconv.Atoi(roundSplit[0])

				if boundriesMap[roundSplit[1]]-colorCount < 0 {

					valid = false
				} else {

					boundriesMap[roundSplit[1]] -= colorCount
				}

				if !valid {
					break
				}

			}

			if !valid {
				break
			}
		}

		if valid {
			result += gameId
		}
	}

	fmt.Println("result", result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
