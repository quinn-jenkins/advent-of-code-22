package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

func main() {
	// open a file
	f, err := os.Open("day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// number of points, not including the head of the rope. Minimum is 1
	numKnots := 9

	headLocation := Coord{}
	scanner := bufio.NewScanner(f)

	tailLocations := make(map[int][]Coord, 10)
	currentTailLocations := make([]Coord, 0)
	for i := 0; i < numKnots; i++ {
		tailLocations[i] = make([]Coord, 0)
		currentTailLocations = append(currentTailLocations, Coord{})
	}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		direction := split[0]
		distance, _ := strconv.Atoi(split[1])
		fmt.Println("-- MOVING", direction, "DISTANCE", distance)
		for i := 0; i < distance; i++ {
			headLocation = MoveHead(direction, headLocation)
			fmt.Println("Head", headLocation)

			for knot := 0; knot < len(tailLocations); knot++ {
				drivingLocation := headLocation
				if knot > 0 {
					drivingLocation = currentTailLocations[knot-1]
				}
				knotLocation := UpdateTail(currentTailLocations[knot], drivingLocation)
				if IsNewLocation(tailLocations[knot], knotLocation) {
					//fmt.Println("*Tail", knot, "Location is now", knotLocation.x, knotLocation.y)
					tailLocations[knot] = append(tailLocations[knot], knotLocation)
				}
				currentTailLocations[knot] = knotLocation
				fmt.Println("Knot", knot+1, "is at", knotLocation)
			}
		}
	}

	//fmt.Println("VISITED LOCATIONS")
	//for _, val := range tailLocations {
	//	fmt.Println("Visited", val)
	//}
	fmt.Println("TOTAL NUM", len(tailLocations[0]))
	fmt.Println("Number of knots", len(tailLocations))
	fmt.Println("TOTAL NUM LAST KNOT", len(tailLocations[len(tailLocations)-1]))
	//for _, v := range tailLocations[len(tailLocations)-1] {
	//	fmt.Println(v)
	//}
}

func MoveHead(direction string, currentLocation Coord) Coord {
	switch direction {
	case "U":
		return Coord{x: currentLocation.x, y: currentLocation.y + 1}
	case "D":
		return Coord{x: currentLocation.x, y: currentLocation.y - 1}
	case "L":
		return Coord{x: currentLocation.x - 1, y: currentLocation.y}
	case "R":
		return Coord{x: currentLocation.x + 1, y: currentLocation.y}
	}

	return currentLocation
}

func UpdateTail(currentTailCoord Coord, currentHeadCoord Coord) Coord {
	verticalSep := currentHeadCoord.y - currentTailCoord.y
	horizSep := currentHeadCoord.x - currentTailCoord.x
	newTailCoord := Coord{}
	if Abs(horizSep) > 1 {
		// need to move horizontally
		if horizSep > 0 {
			newTailCoord.x = currentTailCoord.x + 1
		} else {
			newTailCoord.x = currentTailCoord.x - 1
		}

		if Abs(verticalSep) > 0 {
			// need a diagonal move
			if verticalSep > 0 {
				newTailCoord.y = currentTailCoord.y + 1
			} else {
				newTailCoord.y = currentTailCoord.y - 1
			}
		} else {
			newTailCoord.y = currentTailCoord.y
		}
	} else if Abs(verticalSep) > 1 {
		// need to move vertically
		if verticalSep > 0 {
			newTailCoord.y = currentTailCoord.y + 1
		} else {
			newTailCoord.y = currentTailCoord.y - 1
		}
		if Abs(horizSep) > 0 {
			// need a diagonal move
			if horizSep > 0 {
				newTailCoord.x = currentTailCoord.x + 1
			} else {
				newTailCoord.x = currentTailCoord.x - 1
			}
		} else {
			newTailCoord.x = currentTailCoord.x
		}
	} else {
		return currentTailCoord
	}

	return newTailCoord
}

func IsNewLocation(alreadyVisited []Coord, location Coord) bool {
	for _, val := range alreadyVisited {
		if val.x == location.x && val.y == location.y {
			return false
		}
	}

	return true
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
