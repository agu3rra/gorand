package gorand

import (
	"fmt"
	"os"
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
	if len(args) < 1 {
		help()
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Args))
}
