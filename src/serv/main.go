package main

import (
	"net/http"
	"os"
	"fmt"
	//"taxmain"
)


func main(){
	fs := http.FileServer(http.Dir("app/public"))
	port := os.Getenv("PORT")
	fmt.Println("Port being listened to: " , port)	
	if port == "" {
		port = "5000"
	}
	http.Handle("/" , fs)
	http.ListenAndServe(":" + port, nil)
}
