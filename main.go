package main

import (
	"errors"
	"os"
	"strings"
	"time"

	fu "github.com/jqk/futool4go/fileutils"
	"github.com/jqk/futool4go/timeutils"
)

func main() {
	showVersion()

	var path string
	var caseSensitive bool
	var sortFunc func([]fu.FileExtension)

	argCount := len(os.Args)
	if argCount == 2 {
		if isCommand(os.Args[1]) {
			showHelp()
			return
		}

		caseSensitive = false
		sortFunc = fu.SortFileExtensionsByName
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
			sortFunc = fu.SortFileExtensionsByCount
		} else if ch == 'n' {
			sortFunc = fu.SortFileExtensionsByName
		} else if ch == 's' {
			sortFunc = fu.SortFileExtensionsBySize
		} else {
			showHelp()
			return
		}

		path = os.Args[2]
		exists, isDir, err := fu.FileExists(path)
		if err != nil {
			showError("Path error", err)
			return
		}
		if !exists || !isDir {
			showError("Path error", errors.New("given path does not exist or is a file"))
			return
		}
	}

	// 解析完参数，执行实际任务。
	extensions, elapsed, err := countFileExtensions(path, caseSensitive)
	if err != nil {
		showError("Get file extensions error", err)
		return
	}

	sortFunc(extensions)
	showExtentions(path, caseSensitive, extensions, elapsed)
}

func isCommand(arg string) bool {
	return strings.Index(arg, "-") == 0
}

func countFileExtensions(path string, caseSensitive bool) (
	[]fu.FileExtension, time.Duration, error) {
	showSearchingStart()

	var extensions []fu.FileExtension
	var err error

	done := make(chan struct{})
	dirCount := 0
	fileCount := 0
	extCount := 0
	elapsed := time.Duration(0)

	go func() {
		sw := timeutils.Stopwatch{}
		sw.Start()
		extensions, err = fu.GetFileExtensions(path, caseSensitive,
			func(path string, info os.FileInfo, ext *fu.FileExtension) error {
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

		sw.Stop()
		elapsed = sw.ElapsedTime()

		close(done)
	}()

	sleepTime := 200 * time.Millisecond
	for {
		time.Sleep(sleepTime)

		select {
		case <-done: // 等待扩展名扫描结束。
			if err != nil {
				return nil, time.Duration(0), err
			}

			showSearchingEnd()
			return extensions, elapsed, nil
		default: // 扩展名扫描中，打印进度。
			showSearchProgress(dirCount, fileCount, extCount)
		}
	}
}
