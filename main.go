package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

// The help menu...
func help() {
	helpMessage := `gorand a random file picker

Usage:
  gorand [path] [count]

Arguments:
  path: relative or full path to a directory containing files
  count (optional): the amount of random files to retrieve (defaults to 1)

Use "gorand --help" for more information 

Exit Codes:
  0	OK
  1	Failed at runtime
  2	Invalid user input provided
`
	fmt.Fprint(os.Stderr, helpMessage)
}

// Runs the cli and returns the exit code for the OS
// We keep it separate from main to facilitate test evaluation while maintaining it all in the same process as the tests
func run(args []string) int {
	if len(args) < 2 || args[1] == "--help" {
		help()
		return 2
	}

	directory := args[1]
	count := 1 //default

	if len(args) > 2 {
		var err error
		count, err = strconv.Atoi(args[2])
		if err != nil || count < 1 {
			fmt.Fprintln(os.Stderr, "Error: Invalid count. Must be a positive integer.")
			return 2
		}
	}

	randomFiles, err := getRandomFiles(directory, count)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}

	for _, file := range randomFiles {
		fmt.Println(file)
	}

	return 0
}

func main() {
	os.Exit(run(os.Args))
}

// Retrieves a random list of strings representing files in the target diretory
func getRandomFiles(directory string, count int) ([]string, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	// See how many files we have and append them
	var filePaths []string
	for _, file := range files {
		if !file.IsDir() {
			filePaths = append(filePaths, filepath.Join(directory, file.Name()))
		}
	}

	// Treat detected file count
	filesCount := len(filePaths)
	if filesCount == 0 {
		return nil, fmt.Errorf("no files found in the provided directory")
	}
	if count > filesCount {
		count = filesCount // we cap in case you want more files than are present in the directory
	}

	// mix things up
	result := make([]string, 0, count)
	// Use a map to track selected indices to avoid duplicates
	selected := make(map[int]struct{})

	for len(result) < count {
		index := rand.Intn(len(filePaths))
		if _, exists := selected[index]; !exists {
			selected[index] = struct{}{}
			result = append(result, filePaths[index])
		}
	}
	return result, nil
}
