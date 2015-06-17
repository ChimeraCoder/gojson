package gojson_test

import (
	"fmt"
	"strings"

	"github.com/ChimeraCoder/gojson"
)

func ExampleGenerate() {
	structName := "test"
	pkgName := "main"
	input := strings.NewReader(`{"sample":"json"}`)
	goStruct, err := gojson.Generate(input, structName, pkgName)
	if err != nil {
		fmt.Printf("Error generating json: %s", err.Error())
	}
	fmt.Printf("%s", goStruct)
	// Output:
	// package main
	//
	// type test struct {
	//	Sample string `json:"sample"`
	// }
}
