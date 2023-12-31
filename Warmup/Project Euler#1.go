package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int(tTemp)

    for tItr := 0; tItr < t; tItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int(nTemp)

        result := sumMultiplesOfThreeOrFive(n)
        fmt.Println(result)
    }
}

func sumMultiplesOfThreeOrFive(limit int) int {
    sum := 0
    for i := 1; i < limit; i++ {
        if i%3 == 0 || i%5 == 0 {
            sum += i
        }
    }
    return sum
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}