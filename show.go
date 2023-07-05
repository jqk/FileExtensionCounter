package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/jqk/futool4go/common"
	"github.com/jqk/futool4go/fileutils"
)

var blue color.Style = color.New(color.LightBlue)
var green color.Style = color.New(color.LightGreen)
var white color.Style = color.New(color.White)
var yellow color.Style = color.New(color.Yellow)

func showVersion() {
	white.Println()
	white.Println("Copyright (c) 1999-2023 Not a dream Co., Ltd.")
	white.Print("file extension counter (")
	blue.Print("fec")
	white.Println(") 1.0.2, 2023-07-05")
	white.Println()
}

func showHelp() {
	yellow.Println("Usage:")
	yellow.Println("  fec [command] <path/to/counting/extensions>")
	white.Println("          counting extensions in specified path")
	yellow.Println("\nCommand:")
	white.Println("  the first char of the command defines if the extension is case sensitive.")
	white.Println("      't' is true, 'f' is false.")
	white.Println("  the second one defines how to sort the result.")
	white.Println("      'c' means sort by count.")
	white.Println("      'n' means sort by extension.")
	white.Println("      's' means sort by size.\n")

	yellow.Print("  -fn: default command, can be omitted. ")
	white.Println("case insensitive and sort the result by extension.")
	yellow.Print("  -fc: ")
	white.Println("case insensitive and sort the result by count.")
	yellow.Print("  -fs: ")
	white.Println("case insensitive and sort the result by size.")

	yellow.Print("  -tn: ")
	white.Println("case sensitive and sort the result by extension.")
	yellow.Print("  -tc: ")
	white.Println("case sensitive and sort the result by count.")
	yellow.Print("  -ts: ")
	white.Println("case sensitive and sort the result by size.")

	yellow.Println()
	yellow.Println("  otherwise: show this help.")
	yellow.Println()
}

func showError(header string, err error) {
	color.Errorf("%s: %s\n", header, err)
	os.Exit(1)
}

func showSearchingStart() {
	yellow.Println("Searching...")
}

func showSearchingEnd() {
	yellow.Println("Searching done.\n")
}

func showSearchProgress(dirCount, fileCount, extCount int) {
	white.Print("found dir: ")
	yellow.Printf("%6d", dirCount)
	white.Print(",  file: ")
	yellow.Printf("%7d", fileCount)
	white.Print(",  ext: ")
	yellow.Printf("%5d\n", extCount)
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

	green.Print("Searching path : ")
	yellow.Println(path)
	green.Print("Case sensitive : ")
	yellow.Println(caseSensitive)
	green.Print("Found file     : ")
	yellow.Println(fileCount)
	green.Print("Found extension: ")
	yellow.Println(len(extensions))
	green.Println()

	if fileCount == 0 {
		return
	}

	t := fmt.Sprintf("  No.  %%-%ds  %%5s   %%11s\n", extNameLength)
	s := fmt.Sprintf(t, "Extension", "Count", "Size")
	white.Println(s)

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
