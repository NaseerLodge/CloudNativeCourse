package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe(":8000", mux))
}

var RWLock sync.RWMutex

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	RWLock.Lock()
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	RWLock.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	RWLock.Lock()
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	RWLock.Unlock()
}

func (db database) update(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	priceFloat, _ := strconv.ParseFloat(price, 32)

	RWLock.Lock()
	if _, ok := db[item]; ok {
		delete(db, item)
		db[item] = dollars(priceFloat)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	RWLock.Unlock()

}

func (db database) create(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	priceFloat, _ := strconv.ParseFloat(price, 32)

	RWLock.Lock()
	db[item] = dollars(priceFloat)
	RWLock.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")

	//delete(db, item)

	RWLock.Lock()
	if _, ok := db[item]; ok {
		delete(db, item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
	RWLock.Unlock()

}

func main() {
	// create a mongo client
	client, err := mongo.NewClient(
		options.Client().ApplyURI(mongodbEndpoint),
	)
	checkError(err)

	// Connect to mongo
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// Disconnect
	defer client.Disconnect(ctx)

	// select collection from database
	col := client.Database("blog").Collection("posts")

	// Insert one
	res, err := col.InsertOne(ctx, &Post{
		ID:        primitive.NewObjectID(),
		Title:     "post",
		Tags:      []string{"mongodb"},
		Body:      `MongoDB is a NoSQL database`,
		CreatedAt: time.Now(),
	})
	fmt.Printf("inserted id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())

	// filter posts tagged as mongodb
	filter := bson.M{"tags": bson.M{"$elemMatch": bson.M{"$eq": "mongodb"}}}

	// find one document
	var p Post
	if col.FindOne(ctx, filter).Decode(&p); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("post: %+v\n", p)

}
