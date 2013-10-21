package main

import (
	"encoding/json"
	"github.com/pilu/traffic"
	"io/ioutil"
	"net/http"
)

type Item struct {
	Id   uint
	Text string
}

var (
	items       = map[uint]string{}
	nextId uint = 1
)

func itemsHandler(w traffic.ResponseWriter, r *http.Request) {
}

func createItemHandler(w traffic.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	id := nextId
	items[id] = string(body)
	nextId++

	item := Item{id, string(body)}
	b, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
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
