package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func openFile(path string) *os.File {
	f, err := os.Open(path)
	check(err)
	return f
}

func printHeader(path string) {
	li := strings.LastIndex(path, "/")
	var toPrint string
	if li == -1 {
		toPrint = path
	} else {
		toPrint = path[li+1:]
	}
	println("==>", toPrint, "<==")
}

// TODO: Follow & Sleep flags
func main() {
	linesFlag := flag.Int("n", 10, "Output  the  last  n  lines, instead of the last 10")

	verboseFlag := flag.Bool("verbose", false, "Output headers giving file name")

	flag.Parse()

	filePath := flag.Arg(0)

	f := openFile(filePath)
	defer f.Close()

	fi, err := f.Stat()
	check(err)

	scanner := NewScanner(f, 512, int(fi.Size()))

	if *verboseFlag {
		printHeader(filePath)
	}

	for range *linesFlag {
		l, _, err := scanner.LineBytes()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}

		fmt.Println(string(l))
	}
}
