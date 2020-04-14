package main

import (
    "fmt"
    //"reflect"
)

type Details struct {
    name         string
    age          int
    family      []string
    pin         string
}
const (
        _ = iota
    petrol
    diesel
)
type Engine struct {
    e_type      uint8   `required`
    power       float32
    capacity    float32
}
type Bike struct {
    Engine
    name    string
    comp    string
}

func struct_basic() {
    var s1 Details
    s1 = Details{
        name: "Priya",
        age: 29,
        family:[]string{
            "Muttanna",
            "New one",
        },
    }
    s1.pin = "560093"
    fmt.Println(s1, len(s1.family))
    s2 := s1 //unlike slices and maps when one struct var is assigned to another its copy of it and not the reference
    s2.pin = "587118"
    fmt.Println(s2)
}
//embed
func embed_ex() {
    es := Bike{}
    es.capacity = 394.4
//    es.e_type = diesel
    es.power = 44.5
    es.name = "duke_391"
    es.comp = "ktm"
    fmt.Println(es)
}

func main () {
    //struct_basic()
    embed_ex()
}
