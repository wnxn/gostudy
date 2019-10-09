package main

import (
	"fmt"
	"log"
	"net/http"
)
type dollars float32

type database map[string]dollars

func(d dollars)String()string{
	return fmt.Sprintf("$%.2f", d)
}
func main() {
	db :=database{"shoes": 40, "socks":20}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:9000", mux))
}

func(db database)list(w http.ResponseWriter, req *http.Request){
	for item, price := range db{
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database)price(w http.ResponseWriter, req *http.Request){
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found: %s\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
