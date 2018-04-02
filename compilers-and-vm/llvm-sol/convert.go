package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var replaceWith = map[string]string{
	"pragma":   "",
	"solidity": "",
	"^0.4.0;":  "",
	"contract": "class",
	"function": "func",
	"public":   "",
	"pure":     "",
	"returns":  "->",
	"(string)": "String",
	";":        "",
}

func main() {
	file, err := os.Open("main.sol")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	convertedLines := []string{}

	for scanner.Scan() {
		var line = scanner.Text()
		var lineAsStrings = strings.Split(line, " ")
		var convertedLine = ""
		for _, word := range lineAsStrings {
			var ok bool
			if _, ok = replaceWith[word]; ok {
				convertedLine = convertedLine + replaceWith[word]
			} else {
				convertedLine = convertedLine + " " + word
			}
		}

		convertedLines = append(convertedLines, strings.TrimRight(convertedLine, " "))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	writeLines(convertedLines, "main.swift")
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		if line != "" {
			fmt.Fprintln(w, line)
		}
	}
	return w.Flush()
}
