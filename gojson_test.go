package main

import (
	"os"
	"strings"
	"testing"
)

//TestSimpleJson tests that a simple JSON string with a single key and a single (string) value returns no error
//It does not (yet) test for correctness of the end result
func TestSimpleJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar"}`)
	if _, err := generate(i, "TestStruct"); err != nil {
		panic(err)
	}
}

//TestNullableJson tests that a null JSON value is handled properly
func TestNullableJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar", "baz" : null}`)
	if _, err := generate(i, "TestStruct"); err != nil {
		panic(err)
	}
}

func TestExample(t *testing.T) {
	i, err := os.Open("json_example.json")
	if err != nil {
		panic(err)
	}
	expected := `package main

type TestStruct struct {
	_comment	string
	Glossary	struct {
		GlossDiv	struct {
			GlossList	struct {
				GlossEntry struct {
					Abbrev		string
					Acronym		string
					GlossDef	struct {
						GlossSeeAlso	[]string
						Para		string
					}
					GlossSee	string
					GlossTerm	string
					ID		string
					SortAs		string
				}
			}
			Title	string
		}
		Title	string
	}
}
`
	actual, _ := generate(i, "TestStruct")
	if actual != expected {
		t.Errorf("'%s' (expected) != '%s' (actual)", expected, actual)
	}
}
