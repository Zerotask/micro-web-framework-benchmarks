package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id    int
	Name  string
	Price float32
}

func add(value1, value2 int) int {
	return value1 + value2
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		products := []Product{
			{1, "Fridge1", 200.99},
			{2, "Fridge2", 300.99},
			{3, "Fridge3", 400.99},
			{4, "Fridge4", 500.99},
			{5, "Fridge5", 600.99},
		}
		jsonResponse, err := json.Marshal(products)
		if err != nil {
			log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
		}
		fmt.Fprintf(w, string(jsonResponse))
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse, err := json.Marshal(add(3, 8))
		if err != nil {
			log.Fatalf("Error occured during marshaling. Error: %s", err.Error())
		}
		fmt.Fprintf(w, string(jsonResponse))
	})

	// http://127.0.0.1:13003
	log.Fatal(http.ListenAndServe(":13003", nil))
}
