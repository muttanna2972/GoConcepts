package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
)

const (
    FILE="details.json"
)

type Addr struct {
    Hname   string  `json:"house-name"`
    Street  string  `json:"street"`
    Mod     bool    `json:"modified"`
}
type Details struct {
    Name     string   `json:"name"`
    Address  Addr     `json:"address"`
}
type ListDetails struct {
    AllDet []Details    `json:"personal_details"`
}
func read_write_json (pDet  *ListDetails, read bool) error {
    if read {
        data, err := ioutil.ReadFile(FILE);
        if err != nil {
            return err
        }
        fmt.Println("details read", string(data), "\n")
        if err = json.Unmarshal(data, pDet); err != nil {
            return err
        }
        fmt.Println("details read", *pDet, "\n")
        return nil
    }
    data, err := json.MarshalIndent(pDet, "", "  ")
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(FILE, data, 0644)
    return err
}

func main() {
    var pDet  ListDetails

    fmt.Println ("In main", read_write_json(&pDet, true), "\n")

    for i,_ := range pDet.AllDet {
        pDet.AllDet[i].Name = pDet.AllDet[i].Name + "1"
        pDet.AllDet[i].Address.Mod = true
    }
    fmt.Println ("In write", read_write_json(&pDet, false), "\n")
    fmt.Println(pDet)
}
