package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open a file
	f, err := os.Open("day5/day5input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var containerLines []string
	var moveLines []string

	readingContainer := true
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if readingContainer && strings.Contains(line, "1") {
			readingContainer = false
			// skip the next empty line
			scanner.Scan()
		} else {
			if readingContainer {
				containerLines = append(containerLines, line)
			} else {
				moveLines = append(moveLines, line)
			}
		}
	}

	containerLayout := ParseContainerLayout(containerLines)

	//for _, move := range moveLines {
	//	PerformMovePart1(containerLayout, move)
	//}

	for _, move := range moveLines {
		PerformMovePart2(containerLayout, move)
	}

}

func PerformMovePart2(containerLayout [][]string, move string) {
	splits := strings.Split(move, " ")
	if len(splits) != 6 {
		fmt.Println("ERROR!!! Invalid Move", move)
	}

	numToMove, _ := strconv.Atoi(splits[1])
	fromStackIndex, _ := strconv.Atoi(splits[3])
	toStackIndex, _ := strconv.Atoi(splits[5])
	fromStackIndex--
	toStackIndex--

	fmt.Println("move", numToMove, "from", fromStackIndex, "to", toStackIndex)
	fromStack := containerLayout[fromStackIndex]
	toStack := containerLayout[toStackIndex]

	fromStackTopIndex := GetTopContainerIndex(fromStack)

	chunkToMove := fromStack[(fromStackTopIndex - numToMove + 1):]
	fromStack = fromStack[:(fromStackTopIndex - numToMove + 1)]
	containerLayout[fromStackIndex] = fromStack

	for _, e := range chunkToMove {
		toStack = append(toStack, e)
	}

	containerLayout[toStackIndex] = toStack

	PrintContainerLayout(containerLayout)
}

func GetTopContainerIndex(stack []string) int {
	topIndex := 0
	for i, e := range stack {
		if e != "" {
			topIndex = i
		} else {
			return topIndex
		}
	}

	return topIndex
}

func PerformMovePart1(containerLayout [][]string, move string) {
	splits := strings.Split(move, " ")
	if len(splits) != 6 {
		fmt.Println("ERROR!!! Invalid Move", move)
	}

	numToMove, _ := strconv.Atoi(splits[1])
	fromStackIndex, _ := strconv.Atoi(splits[3])
	toStackIndex, _ := strconv.Atoi(splits[5])
	fromStackIndex--
	toStackIndex--

	fmt.Println("move", numToMove, "from", fromStackIndex, "to", toStackIndex)
	fromStack := containerLayout[fromStackIndex]
	toStack := containerLayout[toStackIndex]

	for numToMove > 0 {
		for i := len(fromStack) - 1; i >= 0; i-- {
			if !(fromStack[i] == "") {
				moved := false
				for j := 0; j < len(toStack); j++ {
					isOpenSpot := toStack[j] == ""
					if isOpenSpot {
						toStack[j] = fromStack[i]
						fromStack = fromStack[:len(fromStack)-1]
						numToMove--
						moved = true
						break
					}
				}

				if !moved {
					fmt.Println("appending", fromStack[i], "to", toStack)
					toStack = append(toStack, fromStack[i])
					containerLayout[toStackIndex] = toStack
					fmt.Println(fromStack)
					fromStack = fromStack[:len(fromStack)-1]
					containerLayout[fromStackIndex] = fromStack
					fmt.Println(containerLayout[fromStackIndex])
					numToMove--
					moved = true
					break
				}
			}
		}

		fmt.Println("Need", numToMove, "more")
	}

	PrintContainerLayout(containerLayout)
}

// ParseContainerLayout takes the set of strings that makes up the container layout and parses them into a 2D array where the
// first index is the stack, and the second index is the position in that stack
func ParseContainerLayout(containerStrings []string) [][]string {
	height := len(containerStrings)

	// the number of elements in the bottom row will give us the number of columns
	// there are 4 characters per element in a row
	bottomRow := strings.Split(containerStrings[height-1], " ")
	numStacks := len(bottomRow)

	containerLayout := make([][]string, numStacks)
	for i := range containerLayout {
		containerLayout[i] = make([]string, 0)
	}

	for row := height - 1; row >= 0; row-- {
		rowElements := GetContainersInRow(containerStrings[row], numStacks)
		for stack, e := range rowElements {
			if !(e == "") {
				containerLayout[stack] = append(containerLayout[stack], e)
			}
		}
	}

	PrintContainerLayout(containerLayout)

	return containerLayout
}

func PrintContainerLayout(containers [][]string) {
	numStacks := len(containers)
	for stack := 0; stack < numStacks; stack++ {
		containersInStack := containers[stack]
		fmt.Println(containersInStack)
	}
}

func GetContainersInRow(row string, size int) []string {
	containers := make([]string, size)
	// split the row into 4 character chunks
	for i := 0; i < size; i++ {
		if i*4+3 <= len(row) {
			container := row[(i * 4):((i+1)*4 - 1)]
			if strings.Contains(container, "[") {
				containers[i] = container[1:2]
			} else {
				containers[i] = ""
			}
		}
	}

	return containers
}
