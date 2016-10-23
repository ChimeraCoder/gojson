package json2struct

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// Test example document
func TestExampleArray(t *testing.T) {
	i, err := os.Open(filepath.Join("examples", "example_array.json"))
	if err != nil {
		t.Fatal("error opening example.json: %s", err)
	}
	defer i.Close()

	expectedf, err := os.Open(filepath.Join("examples", "example_array.go"))
	if err != nil {
		t.Fatal("error opening example_array.go: %s", err)
	}
	defer expectedf.Close()

	expectedBts, err := ioutil.ReadAll(expectedf)
	if err != nil {
		t.Fatalf("error reading example_array.go: %s", err)
	}

	actual, err := Generate(i, ParseJson, "Users", "main", []string{"json"}, false)
	if err != nil {
		t.Fatal(err)
	}
	sactual, sexpected := string(actual), string(expectedBts)
	if sactual != sexpected {
		t.Fatalf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}
