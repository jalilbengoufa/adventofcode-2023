package main


import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "sort"
    "strconv"
)

func main(){

    filename := "data.txt"
    file, err := os.Open(filename)
    if err != nil {
    fmt.Println("Error opening file:", err)
        os.Exit(1)
    }
    defer file.Close()
    

    numMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
        "four": 4,
        "five": 5,
        "six":6,
        "seven":7,
        "eight":8,
        "nine":9,
        "1":1,
        "2":2,
        "3":3,
        "4":4,
        "5":5,
        "6":6,
        "7":7,
        "8":8,
        "9":9,
	}


    scanner := bufio.NewScanner(file)
    result := 0
    for scanner.Scan() {
        
        line := scanner.Text()
        

        var indexes []int
        var indexNumMap = map[int]int{}
        for word, wordNumber := range(numMap){
            
            startIndex := 0
            for {
                index := strings.Index(line[startIndex:], word)
                    if index == -1 {
                        break
                    }

                actualIndex := startIndex + index

                indexes = append(indexes,actualIndex)
                
                indexNumMap[actualIndex] = wordNumber
                startIndex = actualIndex + len(word)
            }

                

        }


        sort.Ints(indexes)

        arrayLen := len(indexes)
        
        lineNumber := 0
        if arrayLen <= 1 {

            lineNumber = indexNumMap[indexes[0]]*11

        } else {

            num1 := strconv.Itoa(indexNumMap[indexes[0]])
            num2 := strconv.Itoa(indexNumMap[indexes[arrayLen-1]])
            number := num1+num2

            convNum, _ := strconv.Atoi(number)

            lineNumber = convNum
        }

        
        fmt.Println(line)
        fmt.Println(lineNumber)

        result+=lineNumber
        

    }


    fmt.Println("result",result)
    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }

    
}
