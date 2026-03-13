package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
)

type FlagOpts struct {
	showNumbers, showNumbersNonblank bool
}

func main() {
	showNumbers := flag.Bool("n", false, "Show line numbers")
	showNumbersNonblank := flag.Bool("b", false, "Number nonempty")

	flag.Parse()
	files := flag.Args()

	// read from stdin
	if len(files) == 0 {
		if err := readFromSdtin(); err != nil {
			log.Fatal(err)
		}
		return
	}

	// read from files
	flgOpt := FlagOpts{showNumbers: *showNumbers, showNumbersNonblank: *showNumbersNonblank}
	var failedExitCode bool

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			if pathErr, ok := errors.AsType[*fs.PathError](err); ok {
				fmt.Fprintf(os.Stderr, "Error: file '%s' does not exist.\n\n", pathErr.Path)
			} else {
				fmt.Fprintf(os.Stderr, "%v\n\n", err)
			}

			failedExitCode = true
			continue
		}
		defer f.Close()

		err = xcat(f, os.Stdout, flgOpt)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			continue
		}
	}

	if failedExitCode {
		os.Exit(1)
	}
}

func xcat(r io.Reader, w io.Writer, flgOpt FlagOpts) error {
	fileReader := bufio.NewReader(r)
	lineNum := 1
	for {

		line, err := fileReader.ReadString('\n')
		if err != nil {
			return err
		}

		var formattedLine string
		formattedLine, lineNum = formatLine(line, lineNum, flgOpt.showNumbers, flgOpt.showNumbersNonblank)
		fmt.Fprint(w, formattedLine)
	}
}

func formatLine(line string, lineNum int, showNumbers, showNumbersNonblank bool) (string, int) {

	if showNumbersNonblank {
		if line == "\n" {
			return line, lineNum
		}

		return fmt.Sprintf("%d %s", lineNum, line), lineNum + 1
	}

	if showNumbers {
		return fmt.Sprintf("%d %s", lineNum, line), lineNum + 1
	}

	return fmt.Sprint(line), lineNum
}

func readFromSdtin() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				return err
			}
		}
		fmt.Fprint(os.Stdout, str)
	}
}
