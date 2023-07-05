package main

import (
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

	showSearchingStart()

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

	for {
		time.Sleep(sleepTime)

		select {
		case <-done:
			if err != nil {
				showError("GetExtensions error", err)
				return
			}

			showSearchingEnd()
			showExtentions(path, caseSensitive, extensions)
			return
		default:
			showSearchStep(dirCount, fileCount, extCount)
		}
	}
}
