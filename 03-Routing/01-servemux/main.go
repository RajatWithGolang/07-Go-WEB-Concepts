package main

import (

   "net/http"
   "io"

)

type myType int
func (m myType) ServeHTTP(res http.ResponseWriter, req *http.Request){
       io.WriteString(res, "your text goes here")

}
func main(){
	var t myType
	mux := http.NewServeMux() 
    mux.Handle("/",t)
	http.ListenAndServe("8080",mux)
}