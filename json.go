package main

import (
    "encoding/json"
    "log"
    "reflect"
)


func generateTypes(obj map[string]interface{}, layers int) string {
    //structure := make(map[string]string)
    structure := "struct {"

    for key, _ := range obj{
        curType:= reflect.TypeOf(obj[key])
        log.Printf("Type is %v", curType)
        indentation := "\t"
        for i := 0; i < layers;{
            indentation += "\t"
            i++
        }

        //Check if this is a nested json object
        log.Printf("Object %v", obj[key])
        nested, isNested:= obj[key].(map[string]interface{})
        if isNested {
            //This is a nested object
            //structure[key] = generateTypes(nested)
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

            //structure[key] = curType.Name()
        }
    }
    //return prettyPrint(structure)
    return structure
}

func generate(jsn string, structName string) string{
    bts := []byte(jsn)
    result := map[string]interface{}{}
    json.Unmarshal(bts, &result)

    typeString := generateTypes(result, 0)

    return "type " + structName + " " + typeString + "}"
}



func main(){

    json := `{
    "glossary": {
        "title": "example glossary",
        "GlossDiv": {
            "title": "S",
            "GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
                    "SortAs": "SGML",
                    "GlossTerm": "Standard Generalized Markup Language",
                    "Acronym": "SGML",
                    "Abbrev": "ISO 8879:1986",
                    "GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
                        "GlossSeeAlso": ["GML", "XML"]
                    },
                    "GlossSee": "markup"
                }
            }
        }
    }
}`
    log.Printf("Result: \n%v", generate(json, "TestStruct"))
}
