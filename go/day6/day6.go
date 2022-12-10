package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// open a file
	f, err := os.Open("day6/day6input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("Failed to close file")
		}
	}(f)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		index, sequence := FindIndexOfUniqueCharacterSequence(line, 4)
		fmt.Println("Index", index, "Sequence:", sequence)

		index, sequence = FindIndexOfUniqueCharacterSequence(line, 14)
		fmt.Println("Index", index, "Sequence:", sequence)
	}
}

func FindIndexOfUniqueCharacterSequence(text string, sequenceLength int) (int, string) {
	for i := 0; i < len(text)-sequenceLength; i++ {
		chunk := text[i : i+sequenceLength]
		if IsUniqueSequence(chunk) {
			return i + sequenceLength, chunk
		}
	}

	return -1, ""
}

func IsUniqueSequence(text string) bool {
	m := make(map[rune]bool)
	for _, i := range text {
		_, ok := m[i]
		if ok {
			return false
		}
		m[i] = true
	}
	return true
}
