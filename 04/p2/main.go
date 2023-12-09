package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	cardCopyCountMap := make(map[int]int)
	cardId := 1
	for scanner.Scan() {
		line := scanner.Text()

		splitedLine := strings.Split(line, " | ")
		cardLine := strings.Split(splitedLine[0], ": ")[1]

		currentN := ""
		var cardNumbers []string
		for _, carL := range cardLine {

			if string(carL) != " " {

				if currentN == "" {

					currentN += string(carL)
				} else {

					currentN += string(carL)
					cardNumbers = append(cardNumbers, currentN)
					currentN = ""
				}

			} else {

				if currentN != "" {
					cardNumbers = append(cardNumbers, currentN)
				}
				currentN = ""
			}
		}
		if currentN != "" {
			cardNumbers = append(cardNumbers, currentN)
		}

		var winNumbers []string
		currentN = ""
		for _, winN := range splitedLine[1] {

			if string(winN) != " " {

				if currentN == "" {

					currentN += string(winN)
				} else {

					currentN += string(winN)
					winNumbers = append(winNumbers, currentN)
					currentN = ""
				}

			} else {

				if currentN != "" {
					winNumbers = append(winNumbers, currentN)
				}
				currentN = ""
			}
		}
		if currentN != "" {
			winNumbers = append(winNumbers, currentN)
		}

		count := 0
		for _, cn := range cardNumbers {

			for _, wn := range winNumbers {
				if cn == wn {

					count++
				}
			}
		}

		copyNb, hasCopies := cardCopyCountMap[cardId]
		if hasCopies {

			for i := 0; i < copyNb; i++ {

				countCopies(count, cardCopyCountMap, cardId)
			}
			cardCopyCountMap[cardId] = cardCopyCountMap[cardId] + 1

		} else {

			cardCopyCountMap[cardId] = 1
		}

		countCopies(count, cardCopyCountMap, cardId)
		cardId++

	}

	result := 0
	for _, cardIndex := range cardCopyCountMap {

		result += cardIndex
	}

	fmt.Println("result", result)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}

func countCopies(count int, cardCopyCountMap map[int]int, cardId int) {

	for i := 1; i <= count; i++ {

		_, ok := cardCopyCountMap[cardId+i]
		if ok {

			cardCopyCountMap[cardId+i] = cardCopyCountMap[cardId+i] + 1
		} else {

			cardCopyCountMap[cardId+i] = 1
		}
	}
}
