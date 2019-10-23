package main

import (
	"html/template"
	"strconv"
	"net/http"
)
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type person struct{
	Name string
	Age int64
}

func foo(res http.ResponseWriter,req *http.Request){
	  name := req.FormValue("name")
	  age,_  := strconv.ParseInt(req.FormValue("age"),10,0)
	  tpl.ExecuteTemplate(res,"index.html",person{name,age})
}

func main(){
    http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}