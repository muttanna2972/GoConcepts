package main

import (
    "fmt"
)

func panic_func() {
    var i int =10
    m := 0
    //Anonymous func
    defer func(){
        if m * 5 == 25 {
            if  err := recover(); err != nil {
               fmt.Println("panic recovered")
            }
        }
    }()
    for ;m < i; m++ {
        if j := m * 5; j == 25 {
            panic("Err value of multiplication is 25")
        }
    }
    fmt.Println("last statement of panic func")
}
func main() {
    panic_func()
    fmt.Println("in main")
}
