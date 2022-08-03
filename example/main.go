package main

import (
	"fmt"
	"github.com/WGrape/apimock/example/server"
	"net/http"
)

func main() {
	http.HandleFunc("/api/get_user", server.GetUserHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("ListenAndServe error : ", err)
		return
	}
}
