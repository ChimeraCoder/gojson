package main

import (
	"testing"
)


//TestSimpleJson tests that a simple JSON string with a single key and a single (string) value returns no error
//It does not (yet) test for correctness of the end result
func TestSimpleJson(t *testing.T) {
	js := []byte(`{"foo" : "bar"}`)
	if err := generate(js, "TestStruct"); err != nil {
		panic(err)
	}
}

//TestNullableJson tests that a null JSON value is handled properly
func TestNullableJson(t *testing.T) {
	js := []byte(`{"foo" : "bar", "baz" : null}`)
	if err := generate(js, "TestStruct"); err != nil {
		panic(err)
	}
}
