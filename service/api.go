package main

import (
	"github.com/pilu/traffic"
	"net/http"
)

func itemsHandler(w traffic.ResponseWriter, r *http.Request) {
}

func createItemHandler(w traffic.ResponseWriter, r *http.Request) {
}

func itemHandler(w traffic.ResponseWriter, r *http.Request) {
}

func updateItemHandler(w traffic.ResponseWriter, r *http.Request) {
}

func deleteItemHandler(w traffic.ResponseWriter, r *http.Request) {
}

func main() {
	router := traffic.New()

	router.Get("/items", itemsHandler)
	router.Post("/items", createItemHandler)
	router.Get("/items/:id", itemHandler)
	router.Put("/items/:id", updateItemHandler)
	router.Delete("/items/:id", deleteItemHandler)

	http.Handle("/", router)
	http.ListenAndServe(":7000", nil)
}
