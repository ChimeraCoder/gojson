package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/ChimeraCoder/gojson"
)

var (
	name       = flag.String("name", "Foo", "the name of the struct")
	pkg        = flag.String("pkg", "main", "the name of the package for the generated code")
	inputName  = flag.String("input", "", "the name of the input file containing JSON (if input not provided via STDIN)")
	outputName = flag.String("o", "", "the name of the file to write the output to (outputs to STDOUT by default)")
)

func main() {
	flag.Parse()

	if isInteractive() && *inputName == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "Expects input on stdin")
		os.Exit(1)
	}

	var input io.Reader
	input = os.Stdin
	if *inputName != "" {
		f, err := os.Open(*inputName)
		if err != nil {
			log.Fatalf("reading input file: %s", err)
		}
		defer f.Close()
		input = f
	}

	if output, err := gojson.Generate(input, *name, *pkg); err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		os.Exit(1)
	} else {
		if *outputName != "" {
			err := ioutil.WriteFile(*outputName, output, 0644)
			if err != nil {
				log.Fatalf("writing output: %s", err)
			}
		} else {
			fmt.Print(string(output))
		}
	}

}

// Return true if os.Stdin appears to be interactive
func isInteractive() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Mode()&(os.ModeCharDevice|os.ModeCharDevice) != 0
}
