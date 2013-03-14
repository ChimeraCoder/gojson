package main

import (
	"io/ioutil"
	"testing"
)

//TestSimpleJson tests that a simple JSON string with a single key and a single (string) value returns no error
//It does not (yet) test for correctness of the end result
func TestSimpleJson(t *testing.T) {
	js := []byte(`{"foo" : "bar"}`)
	if _, err := generate(js, "TestStruct"); err != nil {
		panic(err)
	}
}

//TestNullableJson tests that a null JSON value is handled properly
func TestNullableJson(t *testing.T) {
	js := []byte(`{"foo" : "bar", "baz" : null}`)
	if _, err := generate(js, "TestStruct"); err != nil {
		panic(err)
	}
}

func TestExample(t *testing.T) {
	js, err := ioutil.ReadFile("json_example.json")
	if err != nil {
		panic(err)
	}
	expected := `package main

type TestStruct struct {
	_comment	string
	glossary	struct {
		GlossDiv	struct {
			GlossList	struct {
				GlossEntry struct {
					Abbrev		string
					Acronym		string
					GlossDef	struct {
						GlossSeeAlso	[]interface{}
						para		string
					}
					GlossSee	string
					GlossTerm	string
					ID		string
					SortAs		string
				}
			}
			title	string
		}
		title	string
	}
}
`
	actual, _ := generate(js, "TestStruct")
	if actual != expected {
		t.Errorf("'%s' (expected) != '%s' (actual)", expected, actual)
	}
}
