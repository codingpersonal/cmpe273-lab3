package main

import (
    "net/http"
    "fmt"
)

func main() {
	router_3000 := NewRouter_3000()
	router_3001 := NewRouter_3001()
	router_3002 := NewRouter_3002()
	
	fmt.Println("Hello")
	go func() {
		http.ListenAndServe("localhost:3000", router_3000)
	}()

	go func() {
			http.ListenAndServe("localhost:3002", router_3002)
		}()

	http.ListenAndServe("localhost:3001", router_3001)
	}

