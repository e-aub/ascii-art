package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Taking arguments
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run . [STRING]\n\nEX: go run . \"something\"")
		return
	}
	//Declaring variables
	var (
		characters = make(map[rune][]string)
		counter    = 1
		toWrite    []string
	)
	//newline fix
	toWrite = NewLineFix(os.Args[1])
	//Read file
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	//Creating map
	for i := ' '; i <= '~'; i++ {
		for scanner.Scan() {
			if counter == 1 {
				counter++
				continue
			}
			if counter%9 != 1 {
				characters[i] = append(characters[i], scanner.Text())
				counter++
			} else {
				counter++
				break
			}
		}
	}
	for _, word := range toWrite {
		if word == "\\n" || word == "\n" {
			fmt.Println()
			continue
		} else {
			Printer(&characters, word)
		}
	}
}

func Printer(characters *map[rune][]string, word string) {
	line := 0
	for line < 8 {
		for _, char := range word {
			fmt.Print((*characters)[char][line])
		}
		if line != 7 {
			fmt.Println()
		}
		line++
	}
}

func NewLineFix(s string) []string {
	var result []string
	var temp string
	if len(s) <= 1 {
		result = append(result, s)
		return result
	}
	for i := 0; i < len(s); i++ {

		if s[i] == '\\' && s[i+1] == 'n' {
			if temp != "" {
				result = append(result, temp)
			}
			result = append(result, "\n")
			temp = ""
			i += 1
		} else {
			temp += string(s[i])
		}
		if i == len(s)-1 && temp != "" {
			result = append(result, temp)
		}
	}
	return result
}
