package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var ZeroChar = '0'

func main() {
	// open a file
	f, err := os.Open("day8/day8input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	grid := make([][]int, 0)
	scanner := bufio.NewScanner(f)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, ReadGridRow(line))
		row++
	}

	PrintGrid(grid)
	PartOne(grid)
	PartTwo(grid)
}

func PartTwo(grid [][]int) {
	scenicScore := make([][]int, len(grid))
	largestScore := 0
	for i := 0; i < len(grid); i++ {
		scenicScore[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			scenicScore[i][j] = CalculateScenicScoreForIndex(i, j, grid)
			if scenicScore[i][j] > largestScore {
				largestScore = scenicScore[i][j]
			}
		}
	}

	//PrintGrid(scenicScore)

	fmt.Println("Largest scenic score", largestScore)
}

func CalculateScenicScoreForIndex(row int, col int, grid [][]int) int {
	if row == 0 || col == 0 || row == len(grid)-1 || col == len(grid[row]) {
		return 0
	}

	heightAtIndex := grid[row][col]

	visibleUp := 0
	visibleRight := 0
	visibleDown := 0
	visibleLeft := 0

	// look up
	for j := row - 1; j >= 0; j-- {
		visibleUp++
		if grid[j][col] >= heightAtIndex {
			break
		}
	}

	// look down
	for j := row + 1; j < len(grid); j++ {
		visibleDown++
		if grid[j][col] >= heightAtIndex {
			break
		}
	}

	// look right
	for i := col + 1; i < len(grid[row]); i++ {
		visibleRight++
		if grid[row][i] >= heightAtIndex {
			break
		}
	}

	// look left
	for i := col - 1; i >= 0; i-- {
		visibleLeft++
		if grid[row][i] >= heightAtIndex {
			break
		}
	}

	return visibleUp * visibleRight * visibleDown * visibleLeft
}

func PartOne(grid [][]int) {
	visibleTreesMask := MarkVisibleTrees(grid)

	sum := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if visibleTreesMask[row][col] {
				sum++
			}
		}
	}

	fmt.Println("Total: ", sum)
}

func MarkVisibleTrees(grid [][]int) [][]bool {
	visible := make([][]bool, 0)

	// mark everything visible from the sides
	firstRow := GetFullVisibleRow(len(grid[0]))
	visible = append(visible, firstRow)

	for i := 1; i < len(grid)-1; i++ {
		visible = append(visible, RowVisible(grid[i]))
	}

	lastRow := GetFullVisibleRow(len(grid[0]))
	visible = append(visible, lastRow)

	for col := 0; col < len(grid)-1; col++ {
		tallestTree := -1
		// mark everything visible from the top
		for row := 0; row < len(grid); row++ {
			if grid[row][col] > tallestTree {
				visible[row][col] = true
				tallestTree = grid[row][col]
			}
		}
		// mark everything visible from the bottom
		tallestTree = -1
		for row := len(grid) - 1; row > 0; row-- {
			if grid[row][col] > tallestTree {
				visible[row][col] = true
				tallestTree = grid[row][col]
			}
		}
	}

	// mark everything in the last column
	for row := range visible {
		lastIndexInRow := len(visible[row]) - 1
		visible[row][lastIndexInRow] = true
	}

	return visible
}

func GetFullVisibleRow(rowLen int) []bool {
	visibleRow := make([]bool, rowLen)
	for i := range visibleRow {
		visibleRow[i] = true
	}

	return visibleRow
}

func PrintGrid(grid [][]int) {
	for _, v := range grid {
		fmt.Println(v)
	}
}

func RowVisible(row []int) []bool {
	visibleTreesInRow := make([]bool, len(row))
	tallestTree := -1
	// visible from left
	for i, v := range row {
		if v > tallestTree {
			visibleTreesInRow[i] = true
			tallestTree = v
		}
	}
	// visible from right
	tallestTree = -1
	for i := len(row) - 1; i > 0; i-- {
		if row[i] > tallestTree {
			visibleTreesInRow[i] = true
			tallestTree = row[i]
		}
	}

	return visibleTreesInRow
}

func ReadGridRow(row string) []int {
	gridRow := make([]int, len(row))
	for i, ch := range row {
		gridRow[i] = int(ch - ZeroChar)
	}

	return gridRow
}
