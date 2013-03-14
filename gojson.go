package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
)

var input_file = flag.String("file", "", "the name of the file that contains the json")
var struct_name = flag.String("struct", "JsonStruct", "the desired name of the struct")
var export_fields = flag.Bool("export_fields", true, "should field names be automatically capitalized?")

func generateTypes(obj map[string]interface{}, layers int) string {
	structure := "struct {"

	keys := make([]string, 0, len(obj))
	for key, _ := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		curType := reflect.TypeOf(obj[key])
		indentation := "\t"
		for i := 0; i < layers; {
			indentation += "\t"
			i++
		}

		var typeForKey string
		if nested, isNested := obj[key].(map[string]interface{}); isNested {
			//This is a nested object
			typeForKey = generateTypes(nested, layers+1) + "}"
		} else {
			//Check if this is an array
			if objects, ok := obj[key].([]interface{}); ok {
				types := make(map[reflect.Type]bool, 0)
				for _, o := range objects {
					types[reflect.TypeOf(o)] = true
				}

				if len(types) == 1 {
					typeForKey = "[]" + reflect.TypeOf(objects[0]).Name()
				} else {
					typeForKey = "[]interface{}"
				}
			} else if curType == nil {
				typeForKey = "*interface{}"
			} else {
				typeForKey = curType.Name()
			}
		}
		if *export_fields {
			key = strings.Title(key)
		}
		structure += fmt.Sprintf("\n%s%s %s", indentation, key, typeForKey)
	}
	return structure
}

//Given a JSON string representation of an object and a name structName,
//generate the struct definition of the struct, and give it the specified name
func generate(input io.Reader, structName string) (js_s string, err error) {
	result := map[string]interface{}{}
	if err = json.NewDecoder(input).Decode(&result); err != nil {
		return
	}

	typeString := generateTypes(result, 0)
	typeString = "package main \n type " + structName + " " + typeString + "}"
	return fmtGo(typeString)
}

// pretty prints a piece of go code
func fmtGo(input string) (string, error) {
	fset := token.NewFileSet()

	formatted, err := parser.ParseFile(fset, "", input, parser.ParseComments)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	printer.Fprint(&buf, fset, formatted)
	return buf.String(), nil
}
}

func main() {
	flag.Parse()

	//If '-file' was not provided, use the first command-line argument (if one exists)
	if *input_file == "" && len(flag.Args()) > 0 {
		*input_file = flag.Args()[0]
	}

	var input io.Reader

	if *input_file == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "No input file specified")
		os.Exit(1)
	} else {
		var err error
		input, err = os.Open(*input_file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error reading", *input_file)
			os.Exit(1)
		}
	}

	if output, err := generate(input, *struct_name); err != nil {
		fmt.Fprintln(os.Stderr, "error parsing json", err)
		os.Exit(1)
	} else {
		fmt.Print(output)
	}
}
