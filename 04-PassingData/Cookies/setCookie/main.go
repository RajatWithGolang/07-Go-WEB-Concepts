package main

import "net/http"

func set(res http.ResponseWriter,req *http.Request){
	c1:= http.Cookie{
		Name : "first cookie",
		Value : "First Value",
		HttpOnly : true,
	}
	c2 := http.Cookie{
		Name : "second  cookie",
		Value : "Second Value",
		HttpOnly : true,
	}
	http.SetCookie(res,&c1)
	http.SetCookie(res,&c2)
}
func main(){
	http.HandleFunc("/",set)
	http.ListenAndServe(":8080",nil)
}