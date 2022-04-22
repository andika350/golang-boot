package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	endpoint := http.HandlerFunc(home)
	router.Handle("/", middleware1(endpoint))

	http.ListenAndServe(":8800", router)
}




//middleware
func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware 1")
		next.ServeHTTP(rw, r)
	})
}


//controller
func home(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("oke"))
}