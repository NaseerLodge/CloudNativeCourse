// Example use of Go mongo-driver
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

const (
	mongodbEndpoint = "mongodb://172.17.0.2:27017" // Find this from the Mongo container
)

var RWLock sync.RWMutex
var col *mongo.Collection
var ctx context.Context

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type Post struct {
	ID        primitive.ObjectID `bson:"_id"`
	Clothing  string             `bson:"clothing"`
	Price     dollars            `bson:"price"`
	Tags      string             `bson:"tags"`
	CreatedAt time.Time          `bson:"created_at"`
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

	col = client.Database("blog").Collection("posts")

	mux := http.NewServeMux()
	mux.HandleFunc("/list", list)
	mux.HandleFunc("/create", create)
	log.Fatal(http.ListenAndServe(":8000", mux))

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func list(w http.ResponseWriter, req *http.Request) {

	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		fmt.Println(episode)
	}

}

func create(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	priceFloat, _ := strconv.ParseFloat(price, 32)

	fmt.Println("ITEM : ", item)
	fmt.Println("PRICE: ", priceFloat)

	res, err := col.InsertOne(ctx, &Post{
		ID:        primitive.NewObjectID(),
		Clothing:  item,
		Price:     dollars(priceFloat),
		Tags:      "clothing",
		CreatedAt: time.Now(),
	})

	if err == nil {
		fmt.Printf("inserted id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())
	}

}

func delete(w http.ResponseWriter, req *http.Request) {

	item := req.URL.Query().Get("item")
	result, err := col.DeleteOne(ctx, bson.M{"clothing": item})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
}
