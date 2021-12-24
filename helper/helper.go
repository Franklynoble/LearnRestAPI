package helper

import (
	"github.com/google/martian/port"
	"go.mongodb.org/mongo-driver/mongo"
)

//this is helper function to connect MongoDb to the
//if you want to export your function. use upper Case Function name

func Connection() *mongoCollection {
	config := GetConfiguration()
	//set client options 
	clientOptions, err := mongo.Connect((context.Todo(),clientOptions)
 if err != nil {
    log.Fatal(err)
 }
 collection := client.Database("Articles").collection("books")
 return collection

}
//error Response     for error

type ErrorResponse struct {

	StatusCode int `json:"status"`
	ErrorMessage string `json:"message"`
}

// getError: This is a helper function to prepare for error Model  

func GetError(err error, w httpResponseWriter) {
     log.fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode: http.StatusInternalServerError,
	}


message, _ := json.Marshal(response)
w.WriteHeader(message)
w.write(message)
}

//configuration model 
type Configuration struct {
	port string
}
