package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var positive, negative, zero int

	// Count positive, negative, and zero elements
	for _, num := range arr {
		switch {
		case num > 0:
			positive++
		case num < 0:
			negative++
		default:
			zero++
		}
	}

	n := len(arr)

	// Calculate fractions
	posFraction := float64(positive) / float64(n)
	negFraction := float64(negative) / float64(n)
	zeroFraction := float64(zero) / float64(n)

	// Print fractions
	fmt.Printf("%.6f\n", posFraction)
	fmt.Printf("%.6f\n", negFraction)
	fmt.Printf("%.6f\n", zeroFraction)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
