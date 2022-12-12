package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	//partOne()
	partTwo()
}

func partTwo() {
	// open a file
	f, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

		if len(lines) == 3 {
			// process block of 3 lines
			val := processGroup(lines)
			total += val
			// reset lines
			lines = []string{}
		}
	}

	fmt.Println("TOTAL:", total)
}

func processGroup(strings []string) int {
	intersection := HashGeneric([]rune(strings[0]), []rune(strings[1]))
	intersection = HashGeneric(intersection, []rune(strings[2]))
	if len(intersection) < 1 {
		fmt.Println("ERROR!!! No Intersection")
	}

	//for i := 0; i < len(intersection); i++ {
	//	fmt.Println("Repeat", string(intersection[i]))
	//}
	return ValueForChar(intersection[0])
}

func partOne() {
	// open a file
	f, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	total := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		val := processLine(scanner.Text())
		total += val
		fmt.Println("New Total", total)
	}

	fmt.Println("TOTAL:", total)
}

func processLine(text string) int {
	length := len(text)
	halfLength := length / 2

	fmt.Println("Full String", text)
	str1 := text[0:halfLength]
	//fmt.Println("\tFirstHalf ", str1, "Length", len(str1))
	str2 := text[halfLength:length]
	//fmt.Println("\tSecondHalf", str2, "Length", len(str2))

	intersection := HashGeneric([]rune(str1), []rune(str2))
	if len(intersection) < 1 {
		fmt.Println("ERROR!!!")
	}

	repeat := intersection[0]
	val := ValueForChar(repeat)
	fmt.Println("Repeat Char", string(repeat), "Value", val)

	return val

}

func ValueForChar(ch rune) int {
	if unicode.IsUpper(ch) {
		return int(ch) - 38
	} else {
		return int(ch) - 96
	}
}

// Hash has complexity: O(n * x) where x is a factor of hash function efficiency (between 1 and 2)
func HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}
