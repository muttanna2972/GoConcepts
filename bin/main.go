package main

import "fmt"

func test_callback(str string, cb func(string, bool)) {
    cb(str, true)
}

func main () {
    test_callback("muttu", func(str string, d bool)) {
        fmt.Println(str, d)
    }
}

