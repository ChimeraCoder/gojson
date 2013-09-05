package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
)

var (
	name = flag.String("name", "Foo", "the name of the struct")
	pkg  = flag.String("pkg", "main", "the name of the package for the generated code")
)

// Given a JSON string representation of an object and a name structName,
// attemp to generate a struct definition
func generate(input io.Reader, structName, pkgName string) ([]byte, error) {
	result := map[string]interface{}{}
	if err := json.NewDecoder(input).Decode(&result); err != nil {
		return nil, err
	}

	src := fmt.Sprintf("package %s\ntype %s %s}",
		pkgName,
		structName,
		generateTypes(result, 0))
	return format.Source([]byte(src))
}

// Generate go struct entries for a map[string]interface{} structure
func generateTypes(obj map[string]interface{}, depth int) string {
	structure := "struct {"

	keys := make([]string, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := obj[key]
		valueType := typeForValue(value)

		//If a nested value, recurse
		if value, nested := value.(map[string]interface{}); nested {
			valueType = generateTypes(value, depth+1) + "}"
		}

		fieldName := fmtFieldName(key)
		structure += fmt.Sprintf("\n%s %s `json:\"%s\"`",
			fieldName,
			valueType,
			key)
	}
	return structure
}

var uppercaseFixups = map[string]bool{"id": true, "url": true}

// fmtFieldName formats a string as a struct key
//
// Example:
// 	fmtFieldName("foo_id")
// Output: FooID
func fmtFieldName(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	if len(parts) > 0 {
		last := parts[len(parts)-1]
		if uppercaseFixups[strings.ToLower(last)] {
			parts[len(parts)-1] = strings.ToUpper(last)
		}
	}
	return strings.Join(parts, "")
}

// generate an appropriate struct type entry
func typeForValue(value interface{}) string {
	//Check if this is an array

	if objects, ok := value.([]interface{}); ok {
		types := make(map[reflect.Type]bool, 0)
		for _, o := range objects {
			types[reflect.TypeOf(o)] = true
		}
		if len(types) == 1 {
			return "[]" + reflect.TypeOf(objects[0]).Name()
		}
		return "[]interface{}"
	} else if reflect.TypeOf(value) == nil {
		return "interface{}"
	}
	return reflect.TypeOf(value).Name()
}

// Return true if os.Stdin appears to be interactive
func isInteractive() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fileInfo.Mode()&(os.ModeCharDevice|os.ModeCharDevice) != 0
}

func main() {
	flag.Parse()

	if isInteractive() {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "Expects input on stdin")
		os.Exit(1)
	}

	if output, err := generate(os.Stdin, *name, *pkg); err != nil {
		fmt.Fprintln(os.Stderr, "error parsing", err)
		fmt.Fprintln(os.Stderr, os.Args[0], "expects a top-level object (not an array)")
		os.Exit(1)
	} else {
		fmt.Print(string(output))
	}
}
