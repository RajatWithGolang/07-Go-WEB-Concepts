package main

import (
	"net/http"
	"fmt"
)
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
	fmt.Fprintln(res, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(res, "in chrome go to: dev tools / application / cookies")
}
func read(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("first-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "YOUR COOKIE:", c)
}

func main(){
	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.ListenAndServe(":8080",nil)
}