package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type FileReadResult struct {
	filename string
	content  []byte
	err      error
}

func (f FileReadResult) String() string {
	var b strings.Builder

	b.WriteString(strings.Repeat("-", 10))
	b.WriteString(f.filename)
	b.WriteString(strings.Repeat("-", 10))
	b.WriteByte('\n')
	b.Write(f.content)

	return b.String()
}

func (f *FileReadResult) Write(p []byte) (n int, err error) {
	f.content = p
	return len(p), nil
}

func main() {
	signChan := make(chan os.Signal, 1)

	signal.Notify(signChan, syscall.SIGINT)
	go handleSignal(signChan)

	flag.Parse()
	files := flag.Args()

	if len(files) == 0 {
		readFromStdin(os.Stdin, os.Stdout)
		return
	}

	// Eead from file
	fileNumbers := len(files)

	fileData := make(chan FileReadResult, fileNumbers)

	for _, file := range files {
		go func() {
			content, err := os.ReadFile(file)
			fileData <- FileReadResult{file, content, err}
		}()
	}

	for range fileNumbers {
		data := <-fileData

		if data.err != nil {
			fmt.Println(data.err)
			continue
		}

		fmt.Print(data)
	}
}

func readFromStdin(in io.Reader, out io.Writer) {
	input := bufio.NewScanner(in)
	for input.Scan() {
		fmt.Fprintln(out, input.Text())
	}
}

func handleSignal(signChan <-chan os.Signal) {
	signal := <-signChan
	switch signal {
	case syscall.SIGINT:
		os.Exit(0)
	default:
	}
}
