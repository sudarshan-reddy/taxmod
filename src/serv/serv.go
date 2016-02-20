package main

import (
	"net/http"
	//"taxmain"
)


func main(){
	fs := http.FileServer(http.Dir("app/public"))
	http.Handle("/" , fs)
	http.ListenAndServe(":8001", nil)
}
