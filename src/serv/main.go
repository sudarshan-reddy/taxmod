package main

import (
	"net/http"
	"os"
	"fmt"
	//"taxmain"
)
func test(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "testing server")
}

func main(){
	fs := http.FileServer(http.Dir("./serv/app/public/"))
	port := os.Getenv("PORT")
	fmt.Println("Port being listened to: " , port)	
	if port == "" {
		port = "5000"
	}
	http.Handle("/" , fs)
	http.HandleFunc("/test" , test)
	http.ListenAndServe(":" + port, nil)
}
