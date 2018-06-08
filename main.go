package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	port := 3000
	http.HandleFunc("/customer", customerHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func customerHandler(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case "GET":
			listALlCustomers(response, request)
		case "POST":
			createCustomer(response, request)
		case "DELETE":
			deleteCustomer(response, request)
		default:
			response.Write([]byte("Method not Allowed"))
	}
}

func listALlCustomers(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("{}"))
}

func createCustomer(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(201)
}

func deleteCustomer(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
}

