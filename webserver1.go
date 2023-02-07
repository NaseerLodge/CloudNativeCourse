package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	fmt.Println("ITEM: ", item)

	if price, ok := db[item]; ok {
		fmt.Println("ITEM: ", item)
		fmt.Println("PRICE: ", price)
		fmt.Println("OK: ", ok)
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	priceFloat, _ := strconv.ParseFloat(price, 32)

	fmt.Println("ITEM: ", item)
	fmt.Println("PRICE: ", price)

	delete(db, item)
	db[item] = dollars(priceFloat)

}

func (db database) create(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	priceFloat, _ := strconv.ParseFloat(price, 32)

	fmt.Println("ITEM: ", item)
	fmt.Println("PRICE: ", price)

	db[item] = dollars(priceFloat)

}

func (db database) delete(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")

	delete(db, item)

}

/*
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	//convert string to float and get error status ok
	priceStr := req.URL.Query().Get("price")
	priceFloat, ok := strconv.ParseFloat(priceStr, 32)
	RWLock.Lock() //Works similarly to create, but instead it checks if the item already exists
	if _, itemExist := db[item]; itemExist {
		if ok == nil {
			db[item] = dollars(priceFloat) // changes the price
			RWLock.Unlock()
			fmt.Fprintf(w, "updated price of item %s: %s", item, dollars(priceFloat))

		} else {
			fmt.Fprintf(w, "Invalid price %s", priceStr)
			RWLock.Unlock()
		}
	} else {
		fmt.Fprintf(w, "%s does not exist in list", item)
		RWLock.Unlock()
	}
}
*/
