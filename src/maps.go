package main

import "fmt"

func maps() {
    map1 := map[string]int{
        "clang": 1,
        "golang": 2,
        "python":3,
    }
    map1["java"] = 4 //add a new entry
    fmt.Println(map1)
    //initialize a map with make
    map2 := make(map[string]int)
    map2 = map[string]int{
        "linux": 1,
        "mac": 2,
        "windows":3,
    }
    //map3 := append(map1, map2) append is only for slices
    map3 := map1   //like slices maps are also passed by reference, so modifying assigned map will reflect in actual map
    map3["new"] = 10  //map1 ==> map[clang:1 golang:2 java:4 new:10 python:3]
    fmt.Println(map1)
    fmt.Println(map2)
    //check if the map has a key
    //val, ret := map1["ruby"]
    if val, ret := map1["ruby"]; ret {
    //if ret == true {
        fmt.Println("the key exists, value", val)
    } else {
        fmt.Println("the key does not exist, value", val)
    }
    fmt.Println("Length of map2", len(map2))
    //delete a member from the map 
    key := "java"
    //_,exists := map3[key]
    if _, exists := map1[key];exists {
        delete(map1, "java")
        fmt.Println("map1 after deleting a memeber", map1)
    } else {
        fmt.Println("can not deletea key from map1", key, map1)
    }
    fmt.Println(map1)
}
func main () {
   maps()
}
