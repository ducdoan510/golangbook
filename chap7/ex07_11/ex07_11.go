package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s: $%s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item := query.Get("item")
	price, err := strconv.ParseFloat(query.Get("price"), 32)

	if _, ok := db[item]; ok {
		if err != nil {
			msg := "Price is invalid"
			http.Error(w, msg, http.StatusBadRequest)
		} else {
			db[item] = dollars(price)
			fmt.Fprintf(w, "updated:\n%s: $%.2f\n", item, price)
		}
	} else {
		http.Error(w, fmt.Sprintf("Item %q not found", item), http.StatusNotFound)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item := query.Get("item")
	price, err := strconv.ParseFloat(query.Get("price"), 32)
	if _, ok := db[item]; ok {
		http.Error(w, fmt.Sprintf("Item %q already exists", item), http.StatusBadRequest)
	} else {
		if err != nil {
			http.Error(w, "Price is invalid", http.StatusBadRequest)
		} else {
			db[item] = dollars(price)
			fmt.Fprintf(w, "created:\n%s: $%.2f\n", item, price)
		}
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		delete(db, item)
		fmt.Fprintf(w, "deleted:\n%s: $%.2f\n", item, price)
	} else {
		http.Error(w, fmt.Sprintf("Item %q does not exist", item), http.StatusNotFound)
	}
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	patternHandler := map[string]func(w http.ResponseWriter, r *http.Request) {
		"/list": db.list,
		"/price": db.price,
		"/update": db.update,
		"/create": db.create,
		"/delete": db.delete,
	}
	for pattern, handler := range patternHandler {
		http.HandleFunc(pattern, handler)
	}
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
