package main

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/jqk/futool4go/fileutils"
)

func main() {
	showVersion()

	var path string
	var caseSensitive bool
	var sortFunc func([]fileutils.FileExtension)

	argCount := len(os.Args)
	if argCount == 2 {
		if isCommand(os.Args[1]) {
			showHelp()
			return
		}

		caseSensitive = false
		sortFunc = fileutils.SortFileExtensionsByName
		path = os.Args[1]
	} else if argCount != 3 {
		showHelp()
		return
	} else { // argCount == 3
		command := strings.ToLower(os.Args[1])
		if len(command) != 3 || !isCommand(command) {
			showHelp()
			return
		}

		ch := command[1]
		if ch == 't' {
			caseSensitive = true
		} else if ch == 'f' {
			caseSensitive = false
		} else {
			showHelp()
			return
		}

		ch = command[2]
		if ch == 'c' {
			sortFunc = fileutils.SortFileExtensionsByCount
		} else if ch == 'n' {
			sortFunc = fileutils.SortFileExtensionsByName
		} else if ch == 's' {
			sortFunc = fileutils.SortFileExtensionsBySize
		} else {
			showHelp()
			return
		}

		path = os.Args[2]
		exists, isDir, err := fileutils.FileExists(path)
		if err != nil {
			showError("Path error", err)
			return
		}
		if !exists || isDir {
			showError("Path error", errors.New("given path does not exist or is a file"))
			return
		}
	}

	// 解析完参数，执行实际任务。
	extensions, err := countFileExtensions(path, caseSensitive)
	if err != nil {
		showError("Get file extensions error", err)
		return
	}

	sortFunc(extensions)
	showExtentions(path, caseSensitive, extensions)
}

func isCommand(arg string) bool {
	return strings.Index(arg, "-") == 0
}

func countFileExtensions(path string, caseSensitive bool) ([]fileutils.FileExtension, error) {
	showSearchingStart()

	var extensions []fileutils.FileExtension
	var err error

	done := make(chan struct{})
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

	sleepTime := 200 * time.Millisecond
	for {
		time.Sleep(sleepTime)

		select {
		case <-done: // 等待扩展名扫描结束。
			if err != nil {
				return nil, err
			}

			showSearchingEnd()
			return extensions, nil
		default: // 扩展名扫描中，打印进度。
			showSearchProgress(dirCount, fileCount, extCount)
		}
	}
}
