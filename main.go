package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Franklynoble/LearnRestAPI/helpers"
	"github.com/Franklynoble/LearnRestAPI/models"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

//
// func HealthCheck(w http.ResponseWriter, r *http.Request) {
// log.Println("entering health check end point")
// w.WriteHeader(http.StatusOK)
// fmt.Fprintf(w, "API is up and running")
// }
//
func init() {
	//usersCollection := client.Database("testing").Collection("users")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connected! from  init")

}

var collection = helpers.Connection()

func main() {
	r := mux.NewRouter() //

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")        // this is working
	r.HandleFunc("api/books/{id}", updateBook).Methods("PUT")     // this is not working yet
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE") // this is working

	log.Fatal(http.ListenAndServe(":9000", r))

	//	config := helpers.Connection()

	//fmt.Println(config)

	//log.Fatal(http.ListenAndServe(config.Addr, nil))
}

func getAllDocuments() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))

	usersCollection := client.Database("Articles").Collection("users")
	// retrieve all the documents in a collection
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	// convert the cursor result to bson
	var results []bson.M
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	// display the documents retrieved
	fmt.Println("displaying all results in a collection")
	for _, result := range results {
		fmt.Println(result)
	}

}

func mainConnection() {

	//usersCollection := client.Database("testing").Collection("users")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
	usersCollection := client.Database("Articles").Collection("users")
	// insert a single document into a collection
	// create a bson.D object
	//	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	// insert the bson object using InsertOne()
	//	result, err := usersCollection.InsertOne(context.TODO(), user)
	// check for errors in the insertion
	//	if err != nil {
	//		panic(err)
	//}
	// display the id of the newly inserted object
	//fmt.Println(result.InsertedID)
	// insert multiple documents into a collection
	// create a slice of bson.D objects
	//users := []interface{}{
	//	bson.D{{"fullName", "User 2"}, {"age", 25}},
	//	bson.D{{"fullName", "User 3"}, {"age", 20}},
	//	bson.D{{"fullName", "User 4"}, {"age", 28}},
	//}
	// insert the bson object slice using InsertMany()
	//results, err := usersCollection.InsertMany(context.TODO(), users)
	// check for errors in the insertion
	//if err != nil {
	//	panic(err)
	//}
	// display the ids of the newly inserted objects
	//fmt.Println("Inserted object", results.InsertedIDs)
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}
	//var result bson.M
	fmt.Println("displaying all results from the search query")
	// retrieving the first document that match the filter
	var result bson.M
	// check for errors in the finding
	if err = usersCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		panic(err)
	}
	// display the document retrieved
	fmt.Println("displaying the first result from the search filter")
	fmt.Println(result)
	getAllDocuments()

}

func createBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&book)

	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helpers.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(result)

}

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//we created Book Array
	var book models.Book

	var params = mux.Vars(r)
	//convert from string to primitive Object

	id, _ := primitive.ObjectIDFromHex(params["id"])

	// if err != nil {
	// log.Println(err)
	// }
	//	we create filter. If it is unnecessary to sort data for you, we use bson.M{}

	// bson{}, we passed empty filter. SO we want to get all Data.

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helpers.GetError(err, w)
		fmt.Println("Error From get Method")
		return
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		fmt.Println("error can not decode to document")
	}
	// close the cursor once finished
	//  A Defer statement defers the execution of a function until the surrounding function returns
	//simply, run cur.close() process but after cur.next() finished .*/

}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var book models.Book

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&book)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"isbn", book.Isbn},
			{"title", book.Title},
			{"author", bson.D{
				{"firstname", book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)

	if err != nil {
		helpers.GetError(err, w)
		return
	}

	book.ID = id

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	// set Header
	w.Header().Set("Content-Type", "application/json")
	//get params

	var params = mux.Vars(r)

	//string to primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter.
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helpers.GetError(err, w)
		return
	}
	json.NewEncoder(w).Encode(deleteResult)

}

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//we created Book array
	var books []models.Book //

	//bson.M{} ,we passed empty filter. so want tp get all data

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helpers.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var book models.Book
		// & character returns the memory address of the Following variables
		err := cur.Decode(&book) // decode similar to deserialization process

		if err != nil {
			log.Fatal(err)
		}
		//add item to our array
		books = append(books, book)

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(books) // encode similar to serialize Process

}
