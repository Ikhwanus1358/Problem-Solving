package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the 'getWays' function below.
// The function is expected to return an INTEGER.
// The function accepts following parameters:
//  1. INTEGER n
//  2. INTEGER_ARRAY c

func getWays(n int, c []int) int {
	// Initialize an array to store the number of ways to make change for each amount from 0 to n
	ways := make([]int, n+1)
	ways[0] = 1 // There is one way to make change for 0

	// Iterate through each coin denomination
	for _, coin := range c {
		// Update the ways array for each amount from coin to n
		for i := coin; i <= n; i++ {
			ways[i] += ways[i-coin]
		}
	}

	// The final result is stored in ways[n]
	return ways[n]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// Read the amount to change and the number of coin types
	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	n, _ := strconv.Atoi(firstMultipleInput[0])
	m, _ := strconv.Atoi(firstMultipleInput[1])

	// Read the coin denominations
	cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	var c []int
	for i := 0; i < m; i++ {
		coin, _ := strconv.Atoi(cTemp[i])
		c = append(c, coin)
	}

	// Print the number of ways to make change for 'n' units using coins having the values given by 'c'
	ways := getWays(n, c)
	fmt.Println(ways)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
