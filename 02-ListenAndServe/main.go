package main

import (
	"fmt"
	"net/http"

)
type ownType int
func (o ownType) ServeHTTP(res http.ResponseWriter,req *http.Request){
	fmt.Fprintf(res,"Hello World")
}
func main(){
	var o ownType
	 http.ListenAndServe(":8080",o)
	
}