package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"reflect"
)

func generateTypes(obj map[string]interface{}, layers int) string {
	structure := "struct {"

	for key, _ := range obj {
		curType := reflect.TypeOf(obj[key])
		indentation := "\t"
		for i := 0; i < layers; {
			indentation += "\t"
			i++
		}

		nested, isNested := obj[key].(map[string]interface{})
		if isNested {
			//This is a nested object
			structure += "\n" + indentation + key + " " + generateTypes(nested, layers+1) + "}"
		} else {
			//Check if this is an array
			_, isArray := obj[key].([]interface{})
			if isArray {
				//Currently defaults to interface{} because JSON allows for heterogeneous arrays
				//TODO Run type inference on array to see if it is an array of a single type
				structure += "\n" + indentation + key + " []interface{}"
			} else if curType == nil{
				structure += "\n" + indentation + key + " " + "*interface{}"
			} else {
				structure += "\n" + indentation + key + " " + curType.Name()
            }
		}
	}
	return structure
}

//Given a JSON string representation of an object and a name structName,
//generate the struct definition of the struct, and give it the specified name
func generate(jsn []byte, structName string) (err error) {
	result := map[string]interface{}{}
	json.Unmarshal(jsn, &result)

	typeString := generateTypes(result, 0)

	typeString = "package main \n type " + structName + " " + typeString + "}"

	fset := token.NewFileSet() // positions are relative to fset

	formatted, err := parser.ParseFile(fset, "", typeString, parser.ParseComments)
	if err != nil {
		return err
	}

	printer.Fprint(os.Stdout, fset, formatted)
	return
}

var input_file *string = flag.String("file", "", "the name of the file that contains the json")

func main() {
	flag.Parse()


	if *input_file == "" {
        //If '-file' was not provided, use the first command-line argument, if one exists
        //If no command-line arguments were provided, panic
        if len(os.Args) > 0{
            *input_file = os.Args[1]
        } else{
            panic(fmt.Errorf("No input file specified"))
        }
	}

	//Demontrate example
	//using http://json.org/example.html

	js, _ := ioutil.ReadFile(*input_file)

	if err := generate(js, "TestStruct"); err != nil {
		panic(err)
	}
}
