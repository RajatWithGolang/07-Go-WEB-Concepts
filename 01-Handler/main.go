package main

import (
   "fmt"
   "net/http"
)
//Any Type That implementes Handler interface in a Handler
//Handler Interface contains ServeHTTP Method which holds below two Paramenters
//ResponseWriter and Pointer to Request
type ownType int

func (o ownType) ServeHTTP(res http.ResponseWriter, req *http.Request){
   fmt.Println(res,"Hello World")
}
func main(){

}