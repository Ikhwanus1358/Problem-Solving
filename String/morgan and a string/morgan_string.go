package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'morganAndString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func morganAndString(a string, b string) string {
	var result strings.Builder

	// Convert strings to slices of runes for easier manipulation
	runeA := []rune(a)
	runeB := []rune(b)

	// Initialize indices for both strings
	indexA, indexB := 0, 0

	for indexA < len(runeA) && indexB < len(runeB) {
		// Compare substrings starting from the current indices
		substrA := runeA[indexA:]
		substrB := runeB[indexB:]

		// Compare substrings lexicographically
		if lexicographicCompare(substrA, substrB) <= 0 {
			result.WriteRune(runeA[indexA])
			indexA++
		} else {
			result.WriteRune(runeB[indexB])
			indexB++
		}
	}

	// Append the remaining characters from both strings
	result.WriteString(string(runeA[indexA:]))
	result.WriteString(string(runeB[indexB:]))

	return result.String()
}

// lexicographicCompare compares two substrings lexicographically.
// Returns a negative value if a < b, 0 if a == b, and a positive value if a > b.
func lexicographicCompare(a, b []rune) int {
	minLen := min(len(a), len(b))

	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}

	return len(a) - len(b)
}

// min returns the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		a := readLine(reader)

		b := readLine(reader)

		result := morganAndString(a, b)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
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
