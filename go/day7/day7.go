package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name            string
	parentDirectory *Directory
	subDirectories  map[string]Directory
	files           map[string]File
	size            int
}

type File struct {
	name string
	size int
}

func main() {
	TOTAL_DISK_SPACE := 70000000
	REQUIRED_FREE_SPACE := 30000000
	// open a file
	f, err := os.Open("day7/input.txt")
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
	rootDirectory := Directory{name: "/", subDirectories: make(map[string]Directory, 0), files: make(map[string]File, 0)}
	currentDirectory := rootDirectory
	for scanner.Scan() {
		line := scanner.Text()
		lineTokens := strings.Split(line, " ")
		if string(lineTokens[0]) == "$" {
			command := lineTokens[1]
			if command == "cd" {
				newDirectoryName := lineTokens[2]
				if newDirectoryName == ".." {
					currentDirectory = *currentDirectory.parentDirectory
				} else if newDirectoryName == "/" {
					//ignore
				} else {
					dir := currentDirectory.subDirectories[newDirectoryName]
					currentDirectory = dir
				}
			} else if command == "ls" {
				// do nothing?
			}

		} else if string(lineTokens[0]) == "dir" {
			dirName := lineTokens[1]
			newDir, ok := CreateDirectoryIfNew(currentDirectory, dirName)
			if ok {
				currentDirectory.subDirectories[newDir.name] = newDir
			}
		} else if filesize, err := strconv.Atoi(lineTokens[0]); err == nil {
			// we have a file
			fileName := lineTokens[1]
			newFile, ok := CreateFileIfNew(currentDirectory, fileName, filesize)
			if ok {
				currentDirectory.files[fileName] = newFile
			}
		}
	}

	PrintFileStructure(&rootDirectory, 0)
	totalSpaceUsed := CalculateDirectorySize(&rootDirectory)

	smallDirectories := FindSmallDirectories(rootDirectory, 100000)
	totalSmallDirSize := 0
	for _, dir := range smallDirectories {
		totalSmallDirSize += dir.size
	}
	fmt.Println("TOTAL SMALL DIR SIZE:", totalSmallDirSize)

	// part 2
	currentFreeSpace := TOTAL_DISK_SPACE - totalSpaceUsed
	fmt.Println("TOTAL SPACE USED:", totalSpaceUsed)
	fmt.Println("TOTAL FREE SPACE:", currentFreeSpace)
	spaceNeeded := REQUIRED_FREE_SPACE - currentFreeSpace
	fmt.Println("\tNeed to free", spaceNeeded)

	smallestPossible := totalSpaceUsed
	var dirToDelete Directory
	for _, dir := range FindDirectoriesLargerThanValue(rootDirectory, spaceNeeded) {
		dirSize := CalculateDirectorySize(&dir)
		if dirSize < smallestPossible {
			smallestPossible = dirSize
			dirToDelete = dir
		}
	}

	fmt.Println("DELETE", dirToDelete.name, "WITH SIZE", smallestPossible)
}

func FindDirectoriesLargerThanValue(rootDir Directory, reqFileSize int) []Directory {
	largeDirs := make([]Directory, 0)
	for _, dir := range rootDir.subDirectories {
		if CalculateDirectorySize(&dir) > reqFileSize {
			largeDirs = append(largeDirs, dir)
		}

		largeDirs = append(largeDirs, FindDirectoriesLargerThanValue(dir, reqFileSize)...)
	}

	return largeDirs
}

func FindSmallDirectories(rootDir Directory, maxFileSize int) []Directory {
	smallDirs := make([]Directory, 0)

	for _, dir := range rootDir.subDirectories {
		dirSize := CalculateDirectorySize(&dir)
		if dirSize <= maxFileSize {
			smallDirs = append(smallDirs, dir)
		}

		smallDirs = append(smallDirs, FindSmallDirectories(dir, maxFileSize)...)
	}

	return smallDirs
}

func PrintFileStructure(directory *Directory, level int) {
	spacer := ""
	for i := 0; i < level; i++ {
		spacer = spacer + " - "
	}

	dirSize := CalculateDirectorySize(directory)
	directory.size = dirSize
	fmt.Println(spacer, "DIR", directory.name, "size:", dirSize)

	for _, dir := range directory.subDirectories {
		PrintFileStructure(&dir, level+1)
	}

	for _, file := range directory.files {
		fmt.Println(spacer, "-", file.name, ":", file.size)
	}
}

func CalculateDirectorySize(dir *Directory) int {
	dirSize := 0
	for _, subDir := range dir.subDirectories {
		dirSize += CalculateDirectorySize(&subDir)
	}

	for _, file := range dir.files {
		dirSize += file.size
	}

	dir.size = dirSize

	return dirSize
}

func CreateDirectoryIfNew(currentDirectory Directory, newSubdirectoryName string) (Directory, bool) {
	existingDirs := currentDirectory.subDirectories
	if x, found := existingDirs[newSubdirectoryName]; found {
		return x, false
	}
	newSubdirectory := Directory{name: newSubdirectoryName,
		parentDirectory: &currentDirectory,
		subDirectories:  make(map[string]Directory, 0),
		files:           make(map[string]File, 0)}
	return newSubdirectory, true
}

func CreateFileIfNew(currentDirectory Directory, newFilename string, fileSize int) (File, bool) {
	existingFiles := currentDirectory.files
	if x, found := existingFiles[newFilename]; found {
		return x, false
	}

	newFile := File{name: newFilename, size: fileSize}
	return newFile, true
}
