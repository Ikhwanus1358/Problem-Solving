package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	// Convert int32 array to int array for sorting
	var intArr []int
	for _, num := range arr {
		intArr = append(intArr, int(num))
	}

	// Sort the array in ascending order
	sort.Ints(intArr)

	// Calculate the minimum sum
	minSum := int64(intArr[0] + intArr[1] + intArr[2] + intArr[3])

	// Calculate the maximum sum
	maxSum := int64(intArr[1] + intArr[2] + intArr[3] + intArr[4])

	// Print the results
	fmt.Printf("%d %d\n", minSum, maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
