package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//this is helper function to connect MongoDb to the
//if you want to export your function. use upper Case Function name

func Connection() *mongo.Collection {
	//config := GetConfiguration()
	//set client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0")
	//mongo.Connect((context.(),clientOptions)
	//connect to Mongo Db

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Connected to Mongo!")

	collection := client.Database("Articles").Collection("books")
	return collection

}

//error Response     for error

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// getError: This is a helper function to prepare for error Model

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{

		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)
	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

//configuration model
type Configuration struct {
	Port       string
	Connection string
}

//GetConfiguration method basically populate configuration from .env Configuration Model
// func GetConfiguration() Configuration {
//
// err := gotoenv.Load("./.env")
// if err != nil {
// log.Fatal("Error loading .env file")
// }
// configuration := Configuration{
// os.Getenv("PORT"),
// os.Getenv("CONNECTION"),
// }
// return configuration
//
//}
