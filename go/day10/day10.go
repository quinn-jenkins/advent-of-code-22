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
	partOne()

	// open a file
	f, err := os.Open("day10/day10input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	cycle := 0
	spritePosition := 1
	nextThreshold := 40
	scanner := bufio.NewScanner(f)
	crtRow := ""
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]
		if command == "addx" {
			input, _ := strconv.Atoi(line[1])
			endCycle := cycle + 2
			for i := cycle; i < endCycle; i++ {
				//fmt.Println("cycle", cycle, "sprite pos", spritePosition)
				if Abs(spritePosition-cycle) <= 1 {
					crtRow += "#"
				} else {
					crtRow += "."
				}

				cycle++

				if cycle >= nextThreshold {
					fmt.Println(crtRow)
					cycle = 0
					crtRow = ""
				}
			}
			spritePosition += input
		} else {
			//fmt.Println("--", command)
			cycle++
			if Abs(spritePosition-cycle) <= 1 {
				crtRow += "#"
			} else {
				crtRow += "."
			}

			if cycle >= nextThreshold {
				fmt.Println(crtRow)
				cycle = 0
				crtRow = ""
			}
		}
	}
}

func partOne() {
	// open a file
	f, err := os.Open("day10/day10input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	cycle := 0
	xreg := 1
	nextThreshold := 20
	signalStr := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]
		if command == "addx" {
			input, _ := strconv.Atoi(line[1])
			fmt.Println("--", command, input)
			fmt.Println("\tStarting cycle", cycle)
			cycle += 2
			if cycle >= nextThreshold {
				strToAdd := nextThreshold * xreg
				signalStr += strToAdd
				fmt.Println("***SignalSTR is now", signalStr, "after adding", strToAdd)
				nextThreshold += 40
			}
			xreg += input
			fmt.Println("Ending cycle", cycle, "XREG is", xreg)
		} else {
			fmt.Println("--", command)
			cycle++
			if cycle >= nextThreshold {
				strToAdd := nextThreshold * xreg
				signalStr += strToAdd
				fmt.Println("***SignalSTR is now", signalStr, "after adding", strToAdd)
				nextThreshold += 40
			}
		}
	}

	fmt.Println("SIGNAL STR:", signalStr)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
