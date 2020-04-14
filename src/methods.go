package main

import "fmt"

type student struct {
    name    string
    age     int
    place   string
}

func methodInvoke (name, place string, age int) {
    var std1 student = student {
                            name: name,
                            place: place,
                            age:  age,
                       }
    std1.printStudent()
}

func (s1 student) printStudent () {
    fmt.Println(s1.name, s1.age, s1.place)
}

func main () {
    var a, b int
    //func is can be a type
    addNum := func(a, b int) (sum int) {
        sum = a + b
        return
    }
    a = 10
    b = 20
    c := addNum(a,b)
    fmt.Println("the sum is", addNum(a,c))
    methodInvoke ("Priya", "Bengaluru", 29)
}
