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
	if _, err := generate(i, "TestStruct", "main"); err != nil {
		t.Error("generate() error:", err)
	}
}

// TestNullableJson tests that a null JSON value is handled properly
func TestNullableJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar", "baz" : null}`)
	if _, err := generate(i, "TestStruct", "main"); err != nil {
		t.Error("generate() error:", err)
	}
}

// TestSimpleArray tests that an array without conflicting types is handled correctly
func TestSimpleArray(t *testing.T) {
	i := strings.NewReader(`{"foo" : [{"bar": 24}, {"bar" : 42}]}`)
	if _, err := generate(i, "TestStruct", "main"); err != nil {
		t.Error("generate() error:", err)
	}
}

// TestInvalidFieldChars tests that a document with invalid field chars is handled correctly
func TestInvalidFieldChars(t *testing.T) {
	i := strings.NewReader(`{"f.o-o" : 42}`)
	if _, err := generate(i, "TestStruct", "main"); err != nil {
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

func TestFmtFieldName(t *testing.T) {
	type TestCase struct {
		in  string
		out string
	}

	testCases := []TestCase{
		TestCase{in: "foo_id", out: "FooID"},
		TestCase{in: "fooId", out: "FooID"},
		TestCase{in: "foo_url", out: "FooURL"},
		TestCase{in: "foobar", out: "Foobar"},
		TestCase{in: "url_sample", out: "URLSample"},
	}

	for _, testCase := range testCases {
		lintField := fmtFieldName(testCase.in)
		if lintField != testCase.out {
			t.Errorf("error fmtFiledName %s != %s (%s)", testCase.in, testCase.out, lintField)
		}
	}
}
