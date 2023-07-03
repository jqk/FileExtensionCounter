package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jqk/futool4go/common"
	fileutils "github.com/jqk/futool4go/fileutils"
)

func showVersion() {
	fmt.Println()
	fmt.Println("Copyright (c) 1999-2023 Not a dream Co., Ltd.")
	fmt.Println("file extension counter (fec) 1.0.0, 2023-07-03")
	fmt.Println()
}

func showHelp() {
	fmt.Println("Usage:")
	fmt.Println("  fec [command] <path/to/counter/extensions>")
	fmt.Println("\nCommand:")
	fmt.Println("  -f, --false, or omit command: Count extensions in specified path, extension is case insensitive.")
	fmt.Println("  -t, --true                  : Count extensions in specified path, extension is case sensitive.")
	fmt.Println()
	fmt.Println("  otherwise: show this help.")
	fmt.Println()
}

func showError(header string, err error) {
	fmt.Printf("%s: %s\n", header, err)
	os.Exit(1)
}

func showExtentions(path string, caseSensitive bool, extensions []fileutils.FileExtension) {
	extNameLength := 0
	fileCount := 0
	for _, ext := range extensions {
		fileCount += ext.Count

		n := len(ext.Name)
		if n > extNameLength {
			extNameLength = n
		}
	}

	// extNameLength 用于计算输出格式中扩展名部分的长度。
	// 此处的 9，是正面要打印的 "Extension" 的长度。
	if extNameLength < 9 {
		extNameLength = 9
	}

	path, _ = filepath.Abs(path)

	fmt.Println("Searching path :", path)
	fmt.Println("Case sensitive :", caseSensitive)
	fmt.Println("Found file     :", fileCount)
	fmt.Println("Found extension:", len(extensions))
	fmt.Println()

	if fileCount == 0 {
		return
	}

	t := fmt.Sprintf("  No.  %%-%ds  %%5s   %%11s\n", extNameLength)
	s := fmt.Sprintf(t, "Extension", "Count", "Size")
	fmt.Println(s)

	// 1. sequence number, right aligned.
	// 2. extension, left aligned.
	// 3. count, right aligned.
	// 4. size, right aliged.
	format := fmt.Sprintf("%%5d  %%-%ds  %%5d   %%11s\n", extNameLength)

	fileCount = 1
	for _, ext := range extensions {
		fmt.Printf(format, fileCount, ext.Name, ext.Count, common.ToSizeString(ext.Size))
		fileCount++
	}

	fmt.Println()
}
