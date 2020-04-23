package main

import (
    "fmt"
)

func main () {
    //var intf1 arith = addNum{}
    //var intf2 arith = subNum{}
    intf1 := addNum{}
    intf2 := subNum{}
    a := 20
    b := 10
    fmt.Println("Addition of num using interface", intf1.OpOnNum(a,b))
    fmt.Println("Subtraction of nums using interface", intf2.OpOnNum(a,b))
}

type arith interface {
    OpOnNum(int, int) int
}

type addNum struct {}

func (ad addNum) OpOnNum (a, b int) int {
    return a + b
}

type subNum struct {}

//func (sub subNum) OpOnNum (a, b int) int {
  //  return a - b
//}
