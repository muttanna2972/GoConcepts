package main

import (
    "fmt"
    "sync"
)

var wg = sync.WaitGroup{}

var ch1 = make(chan int)

func extread() {
    fmt.Println("by exthread", <-ch1)
    //ch1<-99
    wg.Done()
}

func  main () {
    wg.Add(3)
    go func () {
        i := <-ch1
        fmt.Println("channel read", i)
        ch1 <- 111
        //fmt.Println("func1", <-ch1)
        wg.Done()
    }()
   go extread()
   go func () {
        ch1 <- 33
        fmt.Println("func2", <-ch1)
        ch1 <- 222
        wg.Done()
    }()
    wg.Wait()
}
