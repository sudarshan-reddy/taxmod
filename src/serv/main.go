package main

import (
	"net/http"
	"os"
	//"taxmain"
)


func main(){
	fs := http.FileServer(http.Dir("app/public"))
	port := os.Getenv("PORT")
	http.Handle("/" , fs)
	http.ListenAndServe(":" + port, nil)
}
