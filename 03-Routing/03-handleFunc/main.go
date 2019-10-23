package main

import (

   "net/http"
   "io"

)


func custHandler(res http.ResponseWriter, req *http.Request){
       io.WriteString(res, "your text goes here")

}
func main(){	
   http.HandleFunc("/",custHandler)
	http.ListenAndServe("8080",nil)
}