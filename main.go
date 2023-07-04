package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jqk/futool4go/fileutils"
)

func main() {
	showVersion()

	var path string
	var caseSensitive bool
	argCount := len(os.Args)

	if argCount == 2 {
		if strings.Index(os.Args[1], "-") == 0 {
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

	var extensions []fileutils.FileExtension
	var err error
	done := make(chan struct{})
	sleepTime := 200 * time.Millisecond
	dirCount := 0
	fileCount := 0
	extCount := 0

	go func() {
		extensions, err = fileutils.GetFileExtensions(path, caseSensitive,
			func(path string, info os.FileInfo, ext *fileutils.FileExtension) error {
				if ext == nil {
					dirCount++
				} else {
					fileCount++
					if ext.Count == 1 {
						extCount++
					}
				}

				return nil
			})

		close(done)
	}()

	stepPrinted := false

	for {
		time.Sleep(sleepTime)

		select {
		case <-done:
			if err != nil {
				showError("GetExtensions error", err)
				return
			}
			if stepPrinted {
				fmt.Println()
			}
			showExtentions(path, caseSensitive, extensions)
			return
		default:
			fmt.Printf("searching...   dir: %6d,  file: %7d,  ext: %5d\n", dirCount, fileCount, extCount)
			stepPrinted = true
		}
	}
}
