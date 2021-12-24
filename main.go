package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//
// func HealthCheck(w http.ResponseWriter, r *http.Request) {
// log.Println("entering health check end point")
// w.WriteHeader(http.StatusOK)
// fmt.Fprintf(w, "API is up and running")
// }
//

func main() {
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
