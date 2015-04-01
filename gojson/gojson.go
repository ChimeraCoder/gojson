package main

import (
    "strings"

    "github.com/gopherjs/gopherjs/js"
    gojson "github.com/ChimeraCoder/gojson"
)

type Pet struct {
    name string
}

func New(name string) *js.Object {
    return js.MakeWrapper(&Pet{name})
}

func (p *Pet) Name() string {
    return p.name
}

func (p *Pet) SetName(name string) {
    p.name = name
}

func main() {
    js.Global.Set("pet", map[string]interface{}{
        "New": New,
    })  
    input := strings.NewReader("asdf")

    if output, err := gojson.Generate(input, "Test", "FooPkg"); err != nil {
        panic(err)
    } else {
        print(output)
    }   

}
