package main

import (
	"encoding/json"
	"github.com/pilu/traffic"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Item struct {
	Id   uint
	Text string
}

var (
	items       = map[uint]string{}
	nextId uint = 1
)

func returnItemAsJson(w traffic.ResponseWriter, item Item) {
	b, _ := json.Marshal(item)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func itemsHandler(w traffic.ResponseWriter, r *http.Request) {
	allItems := []Item{}
	for id, text := range items {
		println(id)
		allItems = append(allItems, Item{id, text})
	}

	b, _ := json.Marshal(allItems)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func createItemHandler(w traffic.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	id := nextId
	items[id] = string(body)
	nextId++

	returnItemAsJson(w, Item{id, string(body)})
}

func itemHandler(w traffic.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, _ := strconv.ParseUint(idString, 10, 0)

	if text, present := items[uint(id)]; present {
		returnItemAsJson(w, Item{uint(id), text})
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
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
