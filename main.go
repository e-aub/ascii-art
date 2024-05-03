package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	lettersMap = make(map[rune][]string)
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Invalid arguments, usage: go run main.go <TEXT>")
		os.Exit(1)
	}
	toWrite := os.Args[1]
	if toWrite == "" {
		return
	}
	fontFile, err := os.Open("standard.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fontFile.Close()

	reader := bufio.NewReader(fontFile)

	fileContent, err := reader.ReadString('\x1e')
	if err.Error() != "EOF" {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	//Split text by double newline
	letters := regexp.MustCompile(`\n{2}`).Split(fileContent, -1)

	for letter := 32; letter <= 126; letter++ {
		lettersMap[rune(letter)] = strings.Split(letters[letter-32], "\n")
	}
	text := split(toWrite)
	for _, part := range text {
		if part == "\\n" {
			fmt.Println()
			continue
		}
		count := 1
		for count <= 8 {
			for _, letter := range part {
				fmt.Print(lettersMap[letter][count-1])
			}
			fmt.Println()
			count++
		}
	}
}

func split(text string) []string {
	result := strings.Split(text, "\\n")
	if result[len(result)-1] == "" && result[len(result)-2] == "" {
		result = result[:len(result)-1]
	}
	for index, line := range result {
		if line == "" {
			result[index] = "\\n"
		}
	}
	return result
}
