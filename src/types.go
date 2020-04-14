package main

import "fmt"

const (
    _ = iota
    KB = 1 << (10 * iota)
    MB
    GB
)

func const_type () {
    fmt.Printf("%v, %T\n", KB, KB)
    fmt.Printf("%v, %T\n", MB, MB)
    fmt.Printf("%v, %T\n", GB, GB)
}

func arrays_sec() {
    var arr1 = [...]int {1,2,3}
    fmt.Printf("%v", arr1)
    //var arr2 [3]int
    arr2 := arr1
    arr2[1] = 100
    fmt.Printf("%v", arr1)
    fmt.Printf("%v", arr2)

}

func slices() {
    var sl1 = []int{1,2,3,4,5,6,7,8,9,10}
    //sl2 := sl1[_:] list all, no it thows an error "cannot use _ as value"
    //fmt.Println(sl2)
    sl3 := sl1[:] //all
    fmt.Println(sl3)
    sl4 := sl1[5:] //starting from 4th index which is number 5
    fmt.Println(sl4)
    sl5 := sl4[1:] //from 7-10
    fmt.Println(sl5)
    sl5[1] = 900
    fmt.Println(sl5) //7,900,9,10
    fmt.Println(sl1) //1,2,3,4,5,6,7,900,9,10
    sl6 := sl1[5:10]
    fmt.Println(sl6) //6
    fmt.Println(cap(sl6), len(sl6)) //6
    fmt.Println(cap(sl1), len(sl1)) //6
}
func slice_op () {
    var sop = make([]int, 10, 100) //make (type, length, capacity)
    //var s1 = []int{1,2,3,4,5}
    //can not use slice or array as 2nd argument in append
    //sop = append(sop, s1)
    /*
    for i, _ := range(sop) {
        sop[i] = i
    }*/
    //instead of loop  assign directly as below
    sop = []int{0,1,2,3,4,5,6,7,8,9}
    fmt.Println(sop, len(sop),cap(sop))
    //sop = append(sop, 1,2,3,4,5) or  sop = append(sop, s2...)
    //or
    s1 := []int{1,2,3,4,5}
    sop = append(sop, s1...) //using the spread operator ...(3dots)
    fmt.Println(sop, len(sop),cap(sop), sop[0])
}

func main () {
    //var str1 string = "priyanka muttanna"
    //var byteStr = []byte(str1)
    //var a rune = 'a'
    //fmt.Printf("%v, %T\n", a, a)
    //const_type();
    //arrays_sec()
    //slices()
    slice_op()
}
