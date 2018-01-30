package gojson

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestSimpleJson tests that a simple JSON string with a single key and a single (string) value returns no error
// It does not (yet) test for correctness of the end result
func TestSimpleJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar"}`)
	if _, err := Generate(i, ParseJson, "TestStruct", "gojson", []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestNullableJson tests that a null JSON value is handled properly
func TestNullableJson(t *testing.T) {
	i := strings.NewReader(`{"foo" : "bar", "baz" : null}`)
	if _, err := Generate(i, ParseJson, "TestStruct", "gojson", []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestSimpleArray tests that an array without conflicting types is handled correctly
func TestSimpleArray(t *testing.T) {
	i := strings.NewReader(`{"foo" : [{"bar": 24}, {"bar" : 42}]}`)
	if _, err := Generate(i, ParseJson, "TestStruct", "gojson", []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestInvalidFieldChars tests that a document with invalid field chars is handled correctly
func TestInvalidFieldChars(t *testing.T) {
	i := strings.NewReader(`{"f.o-o" : 42}`)
	if _, err := Generate(i, ParseJson, "TestStruct", "gojson", []string{"json"}, false, true); err != nil {
		t.Error("Generate() error:", err)
	}
}

// TestDisambiguateFloatInt tests that disambiguateFloatInt correctly
// converts JSON numbers to the desired types.
func TestDisambiguateFloatInt(t *testing.T) {
	examples := []struct {
		FloatsOnly bool
		In         interface{}
		Out        string
	}{
		{FloatsOnly: false, In: 2.2, Out: "float64"},
		{FloatsOnly: false, In: 2.0, Out: "int64"},
		{FloatsOnly: false, In: float64(2), Out: "int64"},
		{FloatsOnly: true, In: 2.2, Out: "float64"},
		{FloatsOnly: true, In: 2.0, Out: "float64"},
		{FloatsOnly: true, In: float64(2), Out: "float64"},
	}

	for i, ex := range examples {
		ForceFloats = ex.FloatsOnly
		if actual := disambiguateFloatInt(ex.In); actual != ex.Out {
			t.Errorf("[Example %d] got %q, but expected %q", i+1, actual, ex.Out)
		}
	}
	ForceFloats = false
}

// TestInferFloatInt tests that we can correctly infer a float or an int from a
// JSON number when no command-line flag is provided.
func TestInferFloatInt(t *testing.T) {
	f, err := os.Open(filepath.Join("examples", "floats.json"))
	if err != nil {
		t.Fatalf("error opening examples/floats.json: %s", err)
	}
	defer f.Close()

	expected, err := ioutil.ReadFile(filepath.Join("examples", "expected_floats.go.out"))
	if err != nil {
		t.Fatalf("error reading expected_floats.go.out: %s", err)
	}

	actual, err := Generate(f, ParseJson, "Stats", "gojson", []string{"json"}, false, true)
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}

}

// TestYamlNumbers tests that we handle Yaml's number system that has both floats and ints correctly
func TestYamlNumbers(t *testing.T) {
	f, err := os.Open(filepath.Join("examples", "numbers.yaml"))
	if err != nil {
		t.Fatalf("error opening examples/numbers.yaml: %s", err)
	}
	defer f.Close()

	expected, err := ioutil.ReadFile(filepath.Join("examples", "expected_numbers.go.out"))
	if err != nil {
		t.Fatalf("error reading expected_numbers.go.out: %s", err)
	}

	actual, err := Generate(f, ParseYaml, "Stats", "gojson", []string{"yaml"}, false, false)
	if err != nil {
		t.Error(err)
	}
	sactual, sexpected := string(actual), string(expected)
	if sactual != sexpected {
		t.Errorf("'%s' (expected) != '%s' (actual)", sexpected, sactual)
	}
}

// Test example document
func TestExample(t *testing.T) {
	i, err := os.Open(filepath.Join("examples", "example.json"))
	if err != nil {
		t.Error("error opening example.json", err)
	}

	expected, err := ioutil.ReadFile(filepath.Join("examples", "expected_output_test.go.out"))
	if err != nil {
		t.Error("error reading expected_output_test.go", err)
	}

	actual, err := Generate(i, ParseJson, "User", "gojson", []string{"json"}, false, true)
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
		{in: "foo_id", out: "FooID"},
		{in: "fooId", out: "FooID"},
		{in: "foo_url", out: "FooURL"},
		{in: "foobar", out: "Foobar"},
		{in: "url_sample", out: "URLSample"},
		{in: "_id", out: "ID"},
		{in: "__id", out: "ID"},
	}

	for _, testCase := range testCases {
		lintField := FmtFieldName(testCase.in)
		if lintField != testCase.out {
			t.Errorf("error fmtFiledName %s != %s (%s)", testCase.in, testCase.out, lintField)
		}
	}
}
