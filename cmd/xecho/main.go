package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func xecho(w io.Writer, args []string, interpreted, noNewline bool) error {
	out := strings.Join(args, " ")

	if interpreted {
		var err error
		var b strings.Builder

		b.WriteString(`"`)
		b.WriteString(out)
		b.WriteString(`"`)

		out, err = strconv.Unquote(b.String())
		if err != nil {
			return err
		}
	}

	if noNewline {
		_, err := fmt.Fprint(w, out)
		return err
	}

	_, err := fmt.Fprintln(w, out)
	return err
}

func main() {
	noNewline := flag.Bool("n", false, "Ommit newline")
	interpreted := flag.Bool("e", false, "Interpreted args")
	flag.Parse()

	xecho(os.Stdout, flag.Args(), *interpreted, *noNewline)
}
