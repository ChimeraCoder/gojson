package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io/ioutil"
    "reflect"
    "os"
)


func generateTypes(obj map[string]interface{}, layers int) string {
    structure := "struct {"

    for key, _ := range obj{
        curType:= reflect.TypeOf(obj[key])
        indentation := "\t"
        for i := 0; i < layers;{
            indentation += "\t"
            i++
        }

        nested, isNested:= obj[key].(map[string]interface{})
        if isNested {
            //This is a nested object
            structure += "\n" + indentation + key + " " + generateTypes(nested, layers + 1) + "}"
        } else{
            //Check if this is an array
            _, isArray := obj[key].([]interface{})
            if isArray {
                //Currently defaults to interface{} because JSON allows for heterogeneous arrays
                //TODO Run type inference on array to see if it is an array of a single type
                structure += "\n" + indentation + key + " []interface{}"
            } else {
                structure += "\n" + indentation + key + " " + curType.Name()
            }
        }
    }
    return structure
}


//Given a JSON string representation of an object and a name structName,
//generate the struct definition of the struct, and give it the specified name
func generate(jsn []byte, structName string) string{
    result := map[string]interface{}{}
    json.Unmarshal(jsn, &result)

    typeString := generateTypes(result, 0)

    return "type " + structName + " " + typeString + "}"
}


var filename *string =  flag.String("file", os.Args[1], "the name of the file that contains the json")

func main(){
    flag.Parse()
    if *filename == "" {
        panic(fmt.Errorf("No filename specified"))
    }

    //Demontrate example
    //using http://json.org/example.html

    js, _  := ioutil.ReadFile(*filename)

    fmt.Printf("package main\n%v\n", generate(js, "TestStruct"))
}
