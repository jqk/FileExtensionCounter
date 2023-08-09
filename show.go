package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

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
	white.Println(") 1.0.3, 2023-08-09")
	white.Println()
}

func showHelp() {
	yellow.Println("Usage:")
	yellow.Println("  fec [command] <path/to/counting/extensions>")
	white.Println("       counting extensions in specified path")
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

func showError(header string, err error, includingHelp bool) {
	color.Errorf("%s: %s\n\n", header, err)

	if includingHelp {
		showHelp()
	}

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

func showExtentions(path string, caseSensitive bool,
	extensions []fileutils.FileExtension, elapsed time.Duration) {

	extNameLength := 0
	seqNo := 0
	for _, ext := range extensions {
		seqNo += ext.Count

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
	yellow.Println(seqNo)
	green.Print("Found extension: ")
	yellow.Println(len(extensions))
	green.Print("Elapsed time   : ")
	yellow.Println(elapsed)
	green.Println()

	if seqNo == 0 {
		return
	}

	t := fmt.Sprintf("  No.  %%-%ds  %%5s   %%11s\n", extNameLength)
	s := fmt.Sprintf(t, "Extension", "Count", "Size")
	white.Println(s)

	// 1. sequence number, right aligned.
	// 2. extension, left aligned.
	// 3. count, right aligned.
	format := fmt.Sprintf("%%5d  %%-%ds  %%5d", extNameLength)

	var sizeSmall int64 = 1024 * 1024 * 10
	var sizeMiddle int64 = 1024 * 1024 * 100
	var sizeLarge int64 = 1024 * 1024 * 1024

	seqNo = 1
	for _, ext := range extensions {
		var c color.Style

		if ext.Size < sizeSmall {
			c = blue
		} else if ext.Size < sizeMiddle {
			c = green
		} else if ext.Size < sizeLarge {
			c = yellow
		} else {
			c = color.New(color.LightRed)
		}

		blue.Printf(format, seqNo, ext.Name, ext.Count)
		// 4. size, right aliged.
		c.Printf("   %11s\n", common.ToSizeString(ext.Size))
		seqNo++
	}

	fmt.Println()
}
