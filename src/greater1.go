package main

import "fmt"

func arg_fun_cb_caller (numbers [] int, cb func(int) bool) []int {
    //g1 := [] int{}
    var g1 =  []int {100, 102}

    defer fmt.Println("Defering the print")
    for _, n := range numbers {
        if cb(n) {
            g1 = append(g1, n)
        }
    }
    fmt.Println("Returning")
    return g1;
}

func main () {
    number := []int{1,2,3,4}
    pg1 := []int{}
    pg1 = arg_fun_cb_caller(number, func (n int) bool {
        return n > 2
    })
    fmt.Print(pg1)
}
