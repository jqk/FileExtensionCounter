package main

import (
	"fmt"
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
	white.Println(") 1.2.0, 2023-09-11")
	white.Println()
}

func showHelp() {
	yellow.Println("Usage:")
	yellow.Println("  fec [<byPassPermissionError> <caseSensitive> <recursive> <sortMethod>] <path/to/counting/extensions>")
	white.Println("       counting extensions in specified path")
	yellow.Println("\nOptions:")
	white.Print("  All options must be either specified in order or omitted. Default values are used when omitted: ")
	yellow.Println("-t -f -t -e\n")

	yellow.Print("  byPassPermissionError: ")
	white.Print("-t is true, skip the permission error ")
	yellow.Print("(default)")
	white.Println("; -f is false, throw an error")
	yellow.Print("  caseSensitive        : ")
	white.Print("-t is true, the extension is case sensitive; -f is false, the extension is case insensitive ")
	yellow.Println("(default)")
	yellow.Print("  recursive            : ")
	white.Print("-t is true, sub directories are included ")
	yellow.Print("(default)")
	white.Println("; -f is false, sub directories are excluded")
	yellow.Print("  sortMethod           : ")
	white.Println("how to sort the result")
	white.Println("                         -c means sort by count")
	white.Print("                         -e means sort by extension ")
	yellow.Println("(default)")
	white.Println("                         -s means sort by size\n")

	white.Println("Otherwise: show this help")
	white.Print("See <")
	yellow.Print("https://github.com/jqk/FileExtensionCounter")
	white.Println("> for more information\n")
}

func showError(header string, err error, includingHelp bool) {
	color.Errorf("%s: %s", header, err)
	white.Println("\n")

	if includingHelp {
		showHelp()
	}
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

func showExtentions(path string, option *fileutils.WalkExtensionOption,
	extensions []fileutils.FileExtension, elapsed time.Duration) {

	size := int64(0)
	extNameLength := 0
	seqNo := 0
	for _, ext := range extensions {
		size += ext.Size
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

	green.Print("Searching path          : ")
	yellow.Println(path)
	green.Print("Bypass permission error : ")
	yellow.Println(option.PathErrorHandler != nil)
	green.Print("Case sensitive          : ")
	yellow.Println(option.CaseSensitive)
	green.Print("Recursive               : ")
	yellow.Println(option.Recursive)
	green.Print("Found file              : ")
	yellow.Println(seqNo)
	green.Print("Found extension         : ")
	yellow.Println(len(extensions))
	green.Print("Total size              : ")
	yellow.Println(common.ToSizeString(size))
	green.Print("Elapsed time            : ")
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
