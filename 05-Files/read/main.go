package main

import (
	"fmt"
	"io/ioutil"
)
func main(){
	data,_ := ioutil.ReadFile("sample.text")
		fmt.Println(string(data))

}