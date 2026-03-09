package main

import (
	"bytes"
	"io"
	"testing"
)

// TODO: test following condition:
// stdin
// file
// flags
var xcatTests = []struct {
	name                             string
	input                            *bytes.Buffer
	output                           *bytes.Buffer
	showNumbers, showNumbersNonblank bool

	want string
}{
	{
		name:                "Default test",
		input:               bytes.NewBufferString("welcome\n"),
		output:              new(bytes.Buffer),
		want:                "welcome\n",
		showNumbers:         false,
		showNumbersNonblank: false,
	},
	{
		name:                "Show line number",
		input:               bytes.NewBufferString("welcome\nhi\nfoo\n"),
		output:              new(bytes.Buffer),
		want:                "1 welcome\n2 hi\n3 foo\n",
		showNumbers:         true,
		showNumbersNonblank: false,
	},
	{
		name:                "Show line number non-blank",
		input:               bytes.NewBufferString("welcome\n\nfoo\n"),
		output:              new(bytes.Buffer),
		want:                "1 welcome\n\n2 foo\n",
		showNumbers:         false,
		showNumbersNonblank: true,
	},
}

func TestXcat(t *testing.T) {
	for _, test := range xcatTests {
		t.Run(test.name, func(t *testing.T) {
			flgOpt := FlagOpts{showNumbers: test.showNumbers, showNumbersNonblank: test.showNumbersNonblank}

			err := xcat(test.input, test.output, flgOpt)
			if err != nil {
				if err != io.EOF {
					t.Error(err)
				}
			}

			if test.output.String() != test.want {
				t.Error("output from xcat not equal what i'm want")
			}
		})
	}
}
