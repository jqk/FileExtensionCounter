package main

import (
	"os"
	"strings"

	"github.com/jqk/futool4go/fileutils"
)

func main() {
	showVersion()

	var path string
	var caseSensitive bool
	argCount := len(os.Args)

	if argCount == 2 {
		if isCommand(os.Args[1]) {
			showHelp()
			return
		}

		caseSensitive = false
		path = os.Args[1]
	} else if argCount != 3 {
		showHelp()
		return
	} else {
		command := strings.ToLower(os.Args[1])

		if command == "-f" || command == "--false" {
			caseSensitive = false
		} else if command == "-t" || command == "--true" {
			caseSensitive = true
		} else {
			showHelp()
			return
		}

		path = os.Args[2]
	}

	extensions, err := fileutils.GetFileExtensions(path, caseSensitive)
	if err != nil {
		showError("GetExtensions error", err)
		return
	}

	showExtentions(path, caseSensitive, extensions)
}

func isCommand(arg string) bool {
	return strings.Index(arg, "-") == 0
}
