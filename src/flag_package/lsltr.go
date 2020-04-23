package main

import (
    "flag"
    "fmt"
    "os"
)

func main () {
    lscmd := flag.NewFlagSet("ls", flag.ExitOnError)
    options := lscmd.String("options", "-l", "list-files")
    if len(os.Args) < 2 {
        fmt.Println("Error in arguments")
        os.Exit(1)
    }
    _ = lscmd
    _ = options
    fmt.Println("the arguments", *options)
}
