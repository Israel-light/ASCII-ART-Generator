package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintln(os.Stderr, "Invalid Input!\nCorrect usage: go run . [STRING] [BANNER]")
		os.Exit(1)
	}

	bannerName := "standard"
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	data, err := os.ReadFile(bannerName + ".txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading banner file", err)
	}

	block := strings.Split(string(data), "\n\n")
	mapChar := make(map[rune][]string)

	asciiCode := rune(32) // starting from SPACE char

	for i := 0; i < len(block); i++ {
		line := strings.Split(block[i], "\n")

		for j := 0; j < len(line); j++ {
			line[j] = strings.TrimRight(line[j], "\r")
		}

		mapChar[asciiCode] = line
		asciiCode++
	}

	// handling the input
	input := os.Args[1]
	// check for unprintable ascii characters
	printable := ""
	for _, char := range input {
		if char < 32 || char > 126 {
			printable += " "
		} else {
			printable += string(char)
		}
	}

	processedInput := strings.ReplaceAll(printable, "\\n", "\n")
	printLines := strings.Split(processedInput, "\n")

	for i := 0; i < len(printLines); i++ {
		if printLines[i] == "" {
			fmt.Println("")
		} else {
			for row := 0; row <= 7; row++ {
				var cumRow strings.Builder
				for _, char := range printLines[i] {
					rows, ok := mapChar[char]
					if !ok {
						rows = mapChar[' ']
					}

					cumRow.WriteString(rows[row])
				}
				fmt.Println(cumRow.String())
			}
		}
	}
}
