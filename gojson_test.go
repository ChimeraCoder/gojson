package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// TestSimpleJson tests that a simple JSON string with a single key and a single (string) value returns no error
// It does not (yet) test for correctness of the end result
func TestSimpleJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar"}`)
	if _, err := generate(i, "TestStruct", "gojson_test"); err != nil {
		t.Error("generate() error:", err)
	}
}

// TestNullableJson tests that a null JSON value is handled properly
func TestNullableJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar", "baz" : null}`)
	if _, err := generate(i, "TestStruct", "gojson_test"); err != nil {
		t.Error("generate() error:", err)
	}
}

// Test example document
func TestExample(t *testing.T) {
	i, err := os.Open("example.json")
	if err != nil {
		t.Error("error opening example.json", err)
	}

	expected, err := ioutil.ReadFile("expected_output_test.go")
	if err != nil {
		t.Error("error reading expected_output_test.go", err)
	}

	actual, err := generate(i, "User", "main")
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}
