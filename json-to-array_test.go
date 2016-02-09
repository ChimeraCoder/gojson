package json2struct

import (
	"os"
	"testing"
)

// Test example document
func TestExampleArray(t *testing.T) {
	i, err := os.Open("example_array.json")
	if err != nil {
		t.Error("error opening example.json", err)
	}

	// TODO we can do better than []interface{} for homogenous structs
	expected := `package main

type Users []interface{}
`

	actual, err := Generate(i, ParseJson, "Users", "main", []string{"json"})
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}
