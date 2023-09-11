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
	var sortFunc func([]fu.FileExtension)
	var err error

	option := fu.NewWalkExtensionOption()
	argCount := len(os.Args)

	if argCount == 1 {
		showHelp()
		return
	} else if argCount == 2 {
		sortFunc = fu.SortFileExtensionsByName
		path = os.Args[1]
	} else if argCount == 6 {
		// os.Args[0] 是程序自己的名字。
		// os.Args[1] 到 os.Args[3] 是 3 个 bool 类型的参数。正好对应 option 的 3 个成员。
		if err = setOption(option, os.Args); err != nil {
			showError("Option error", err, true)
			os.Exit(1)
		}

		// os.Args[4] 是排序方法。
		if sortFunc, err = getSortFunc(os.Args[4]); err != nil {
			showError("Option error", err, true)
			os.Exit(1)
		}

		// os.Args[5] 是路径。
		path = os.Args[5]
	} else {
		showError("Argument error", errors.New("wrong number of argument"), true)
		os.Exit(1)
	}

	if err = validatePath(path); err != nil {
		showError("Path error", err, false)
		os.Exit(2)
	}

	// 解析完参数，执行实际任务。
	extensions, elapsed, err := countFileExtensions(path, option)
	if err != nil {
		showError("Get file extensions error", err, false)
		os.Exit(3)
	}

	sortFunc(extensions)
	showExtentions(path, option, extensions, elapsed)
}

func countFileExtensions(path string, option *fu.WalkExtensionOption) (
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
		extensions, err = fu.GetFileExtensions(path, option,
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

func setOption(option *fu.WalkExtensionOption, args []string) error {
	boolValues := make([]bool, 3)

	for i := 1; i < 4; i++ {
		arg := strings.ToLower(args[i])
		if arg == "-f" {
			boolValues[i-1] = false
		} else if arg == "-t" {
			boolValues[i-1] = true
		} else {
			return errors.New(args[i] + " is not a valid option")
		}
	}

	if !boolValues[0] {
		option.PathErrorHandler = nil
	}
	option.CaseSensitive = boolValues[1]
	option.Recursive = boolValues[2]

	return nil
}

func getSortFunc(arg string) (func([]fu.FileExtension), error) {
	arg = strings.ToLower(arg)

	if arg == "-c" {
		return fu.SortFileExtensionsByCount, nil
	} else if arg == "-e" {
		return fu.SortFileExtensionsByName, nil
	} else if arg == "-s" {
		return fu.SortFileExtensionsBySize, nil
	} else {
		return nil, errors.New("not 'c', 'e' or 's' for sort method")
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
