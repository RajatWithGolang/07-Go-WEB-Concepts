package main

import (
    "fmt"
    "io/ioutil"
)

func main() {

    mydata := []byte("Here goes my sample text")

    // the WriteFile method returns an error if unsuccessful
    err := ioutil.WriteFile("myfile.text", mydata, 0777)
    // handle this error
    if err != nil {
        // print it out
        fmt.Println(err)
    }

    data, err := ioutil.ReadFile("myfile.text")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Print(string(data))

}