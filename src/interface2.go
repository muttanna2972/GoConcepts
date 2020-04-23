package main

import (
    "fmt"
)

func main () {
   var man1 details = details {
                        name:"muttu",
                        age: 30,
                      }
    var inval mint = 0
    var structiface BasicInterface = &man1
    var intiface  BasicInterface = &inval

    structiface.addone()
    intiface.addone()

    /*
     *  Instances passed to methods
     *  are passed by value or reference?
     */
    structiface.addtwo()
    intiface.addtwo()
}

//Interface
type BasicInterface interface {
    addone()
    addtwo()
}

type mint int

type details struct {
    name    string
    age     int
}

func (n mint) addone() {
    n = n + 1
    fmt.Println(n)
}

func (n details) addone() {
    n.age = n.age + 1
    fmt.Println("Age", n.age)
}

func (n mint) addtwo() {
    n = n + 2
    fmt.Println(n)
}

func (n details) addtwo() {
    n.age = n.age + 2
    fmt.Println("Age", n.age)
}
