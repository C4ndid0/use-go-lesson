package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to the jungle!</h1>")
}

func main(){
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Stanting the server on :3000...")
	http.ListenAndServe(":3000",nil)
}
