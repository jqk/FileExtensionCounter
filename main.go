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
	var err error

	argCount := len(os.Args)
	if argCount == 1 {
		showHelp()
		return
	} else if argCount == 2 {
		// 2 个参数时，第 2 个参数必须是待查看的路径，而不是命令。
		if isCommand(os.Args[1]) {
			showError("Argument error", errors.New("the only one argument must be the path to search"), true)
			return
		}

		caseSensitive = false
		sortFunc = fu.SortFileExtensionsByName
		path = os.Args[1]
	} else if argCount == 3 {
		// 3 个参数时，第 2 个是命令，第 3 个是待查看的路径。
		// 命令字符串只能是 3 个字符。
		command := strings.ToLower(os.Args[1])
		if len(command) != 3 || !isCommand(command) {
			showError("Command error", errors.New(os.Args[1]+" is not a valid command"), true)
			return
		}
		if caseSensitive, err = getCaseSensitive(command[1]); err != nil {
			showError("Command error", err, true)
			return
		}
		if sortFunc, err = getSortFunc(command[2]); err != nil {
			showError("Command error", err, true)
			return
		}

		path = os.Args[2]
	} else {
		showError("Argument error", errors.New("wrong number of argument"), true)
		return
	}

	if err = validatePath(path); err != nil {
		showError("Path error", err, false)
		return
	}

	// 解析完参数，执行实际任务。
	extensions, elapsed, err := countFileExtensions(path, caseSensitive)
	if err != nil {
		showError("Get file extensions error", err, false)
		return
	}

	sortFunc(extensions)
	showExtentions(path, caseSensitive, extensions, elapsed)
}

func countFileExtensions(path string, caseSensitive bool) (
	[]fu.FileExtension, time.Duration, error) {
	showSearchingStart()

	var extensions []fu.FileExtension
	var err error

	// 用于协程同步的通道。
	done := make(chan struct{})

	// 各个扩展名的统计信息。
	dirCount := 0
	fileCount := 0
	extCount := 0
	elapsed := time.Duration(0)

	// 启动单独的协程处理扩展名。
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

	// 等待扩展名扫描结束。并显示扩展名扫描进度。
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

func isCommand(arg string) bool {
	return strings.Index(arg, "-") == 0
}

func getCaseSensitive(ch byte) (bool, error) {
	if ch == 't' {
		return true, nil
	} else if ch == 'f' {
		return false, nil
	} else {
		return false, errors.New("not 't' or 'f' for case sensitive")
	}
}

func getSortFunc(ch byte) (func([]fu.FileExtension), error) {
	if ch == 'c' {
		return fu.SortFileExtensionsByCount, nil
	} else if ch == 'n' {
		return fu.SortFileExtensionsByName, nil
	} else if ch == 's' {
		return fu.SortFileExtensionsBySize, nil
	} else {
		return nil, errors.New("not 'c', 'n' or 's' for sort method")
	}
}

func validatePath(path string) error {
	exists, isDir, err := fu.FileExists(path)
	if err != nil {
		return err
	}
	if !exists || !isDir {
		return errors.New("given path does not exist or is a file")
	}

	return nil
}
