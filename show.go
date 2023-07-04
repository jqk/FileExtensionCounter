package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/jqk/futool4go/common"
	"github.com/jqk/futool4go/fileutils"
)

func showVersion() {
	w := color.New(color.White)
	b := color.New(color.LightBlue)
	w.Println()
	w.Println("Copyright (c) 1999-2023 Not a dream Co., Ltd.")
	w.Print("file extension counter (")
	b.Print("fec")
	w.Println(") 1.0.1, 2023-07-04")
	w.Println()
}

func showHelp() {
	w := color.New(color.LightWhite)
	y := color.New(color.LightYellow)
	g := color.New(color.LightGreen)

	y.Println("Usage:")
	y.Println("  fec [command] <path/to/counter/extensions>")
	y.Println("\nCommand:")
	y.Print("  -f, --false, or omit command: ")
	w.Print("extensions in specified path, extension is case ")
	g.Println("insensitive.")
	y.Print("  -t, --true                  : ")
	w.Print("extensions in specified path, extension is case ")
	g.Println("sensitive.")
	y.Println()
	y.Println("  otherwise: show this help.")
	y.Println()
}

func showError(header string, err error) {
	color.FgLightRed.Println("%s: %s\n", header, err)
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

	g := color.New(color.Green)
	y := color.New(color.Yellow)

	g.Print("Searching path : ")
	y.Println(path)
	g.Print("Case sensitive : ")
	y.Println(caseSensitive)
	g.Print("Found file     : ")
	y.Println(fileCount)
	g.Print("Found extension: ")
	y.Println(len(extensions))
	g.Println()

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
