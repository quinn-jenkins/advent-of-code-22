package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	startZone int
	endZone   int
}

func main() {
	// open a file
	f, err := os.Open("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	numOverlappingAssignments := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		assign1, assign2 := ParseLine(scanner.Text())
		isOverlapping := IsOverlapping(assign1, assign2)
		isFullyContained := IsFullyContained(assign1, assign2)

		if isOverlapping || isFullyContained {
			numOverlappingAssignments++
		}
	}
	fmt.Println("Total of overlapping assignments", numOverlappingAssignments)
}

func ParseLine(line string) (Assignment, Assignment) {
	assignments := strings.Split(line, ",")

	if len(assignments) != 2 {
		fmt.Println("ERROR!!!", line, "does not have exactly 2 assignments")
	}

	var assignSet [2]Assignment
	for idx, assignmentString := range assignments {
		assignment := ConvertNumRangeToAssignment(assignmentString)
		assignSet[idx] = assignment
	}

	return assignSet[0], assignSet[1]
}

func ConvertNumRangeToAssignment(numberRange string) Assignment {
	numbers := strings.Split(numberRange, "-")
	var assign Assignment

	start, _ := strconv.Atoi(numbers[0])
	end, _ := strconv.Atoi(numbers[1])
	assign.startZone = start
	assign.endZone = end
	return assign
}

func IsFullyContained(first Assignment, second Assignment) bool {
	if first.startZone <= second.startZone && first.endZone >= second.endZone {
		return true
	} else if second.startZone <= first.startZone && second.endZone >= first.endZone {
		return true
	}

	return false
}

func IsOverlapping(first Assignment, second Assignment) bool {
	if first.startZone <= second.endZone && second.startZone <= first.endZone {
		return true
	}
	return false
}
