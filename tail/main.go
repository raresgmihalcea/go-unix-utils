package main

import (
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFilePath() string {
	args := os.Args
	if len(args) == 1 {
		panic("tail: provide a file for reading")
	}
	return args[1]
}

func openFile(path string) *os.File {
	f, err := os.Open(path)
	check(err)
	return f
}

func main() {
	//linesFlag := flag.Int("n", 10, "Output  the  last  n  lines, instead of the last 10")

	followFlag := flag.Bool("follow", false, "Output appended data as file grows")

	//verboseFlag := flag.Bool("verbose", false, "Output headers giving file name")

	retryFlag := flag.Bool("retry", false, "Keep trying to open a file if it is inaccessible")

	//sleepFlag := flag.Float64("sleep", 1.0, "with -f, sleep for N seconds (default 1.0) between iterations")

	flag.Parse()

	filePath := readFilePath()

	if *retryFlag && !*followFlag {
		fmt.Println("tail: warning: --retry ignored; --retry is useful only when following")
		*retryFlag = false
	}

	f := openFile(filePath)
	defer f.Close()

}
