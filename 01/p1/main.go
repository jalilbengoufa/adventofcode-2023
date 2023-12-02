package main


import (
    "bufio"
    "fmt"
    "os"
    "unicode"
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

    scanner := bufio.NewScanner(file)
    result := 0
    for scanner.Scan() {
        
        line := scanner.Text()
        var numbers string
        for _, char := range line {
            if unicode.IsDigit(char) {
                numbers += string(char)
            }
        }   
        
        lineNumber := 0

        numbersLen := len(numbers)

        runes := []rune(numbers)
        if numbersLen <= 1{

            i, _ := strconv.Atoi(string(runes[0]))
            lineNumber = i*11
        } else {
        
            i, _ := strconv.Atoi(string(runes[0])+string(runes[numbersLen-1])) 
            lineNumber = i

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
