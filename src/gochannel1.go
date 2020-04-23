package main

import (
        "fmt"
        "sync"
)
const (
    _ = iota
    petrol
    diesel
)
type mybike struct {
    etype       uint8
    power       float32
    capacity    float32
    name        string
    comp        string
}

var wg = sync.WaitGroup{}

func main () {
     //var mbike mybike
    var chv = make(chan mybike)

    wg.Add(2)
    go func () {
        //fmt.Printf("%V %T", wg, wg)
        chv <-  mybike {
                    etype:petrol,
                    power: 44.5,
                    capacity: 394.4,
                    name: "duke390",
                    comp: "KTM",
                }
        wg.Done()
    }()
    go func () {
        fmt.Println(<-chv)
        wg.Done()
    }()
    wg.Wait()
}
