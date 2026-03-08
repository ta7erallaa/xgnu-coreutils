package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

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

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		xcat(f, os.Stdout, flgOpt)
	}
}

type FlagOpts struct {
	showNumbers, showNumbersNonblank bool
}

func xcat(r io.Reader, w io.Writer, flgOpt FlagOpts) {
	fileReader := bufio.NewReader(r)
	lineNum := 1
	for {
		line, err := fileReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
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

	return fmt.Sprint(line), lineNum + 1

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
