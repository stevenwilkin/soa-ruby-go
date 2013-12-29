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

func getId(r *traffic.Request) uint {
	idString := r.URL.Query().Get("id")
	id, _ := strconv.ParseUint(idString, 10, 0)
	return uint(id)
}

func getItem(r *traffic.Request) (Item, bool) {
	id := getId(r)
	text, present := items[id]

	if present {
		return Item{id, text}, true
	} else {
		return Item{}, false
	}
}

func checkItemExists(w traffic.ResponseWriter, r *traffic.Request) {
	if _, present := getItem(r); !present {
		w.WriteHeader(http.StatusNotFound)
		w.WriteText("Item Not Found")
	}
}

func itemsHandler(w traffic.ResponseWriter, r *traffic.Request) {
	allItems := []Item{}
	for id, text := range items {
		allItems = append(allItems, Item{id, text})
	}

	b, _ := json.Marshal(allItems)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func createItemHandler(w traffic.ResponseWriter, r *traffic.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	id := nextId
	items[id] = string(body)
	nextId++

	returnItemAsJson(w, Item{id, string(body)})
}

func itemHandler(w traffic.ResponseWriter, r *traffic.Request) {
	item, _ := getItem(r)
	returnItemAsJson(w, item)
}

func updateItemHandler(w traffic.ResponseWriter, r *traffic.Request) {
	id := getId(r)
	body, _ := ioutil.ReadAll(r.Body)

	items[id] = string(body)

	returnItemAsJson(w, Item{id, string(body)})
}

func deleteItemHandler(w traffic.ResponseWriter, r *traffic.Request) {
	id := getId(r)
	delete(items, id)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := traffic.New()

	router.Get("/items", itemsHandler)
	router.Post("/items", createItemHandler)
	router.Get("/items/:id", itemHandler).
		AddBeforeFilter(checkItemExists)
	router.Put("/items/:id", updateItemHandler).
		AddBeforeFilter(checkItemExists)
	router.Delete("/items/:id", deleteItemHandler).
		AddBeforeFilter(checkItemExists)

	http.Handle("/", router)
	http.ListenAndServe(":7000", nil)
}
