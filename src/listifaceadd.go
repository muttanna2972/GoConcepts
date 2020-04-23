package main

import (
    "fmt"
    "net"
)
//interface state monitor application
func main () {
    var addr []net.Addr
    var err error
    addr, err = net.InterfaceAddrs()
    if err == nil {
        fmt.Printf("%v %T,", addr, addr)
    }else {
        fmt.Println("Error reading", err)
    }
    iface,err1 := net.InterfaceByName("eth0")
    if err1 == nil {
        fmt.Printf("%v %T,", iface, iface)
    } else {
        fmt.Println("Error reading", err)
    }

}
