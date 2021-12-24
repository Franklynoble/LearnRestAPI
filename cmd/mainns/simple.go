package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Persons []Person `json:"persons"`
}
type Art struct {
	// Id      int    `json:"id"`
	// Title   string `json:"Title"`
	// Desc    string `json:"desc"`
	// Content string `json:"content"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}


type Issue struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	Title       string             `bson:"title"`
	Code        string             `bson:"code"`
	Description string             `bson:"description"`
	Completed   bool               `bson:"completed"`
}

//	Art := Article{Id: 1, Title: "Robert Greene", Desc: "Authour, and entrepreneur", Content: "books and  perosnality traits"}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")

	// client, err := mongo.Connect(context.TODO(), clientOptions)

	// if err != nil {
		// log.Fatal(err)
	// }
	check the connections and
	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
		// log.Fatal(err)
	// }
	// fmt.Println("before after", reflect.TypeOf(client).String())
	// fmt.Println("Connected to MongoDB!")
	// Earticless := client.Database("Articles").Collection("Employee")

	// e, err := Earticless.Find(Earticless.FindOne(context.TODO() bson..Id)
	inserts, err = Earticless.(context.TODO(), Earticless)
	if err != nil {
	fmt.Println("Error", err.Error())
	}

	var result Article // copy
	// err = Earticless.FindOne(context.TODO(), bson.D{}).Decode(&result)
	// if err != nil {

		// log.Fatal(err)

		// fmt.Println("Inserted a single document  Error: ", inserts.Current)

	// } else {
		// fmt.Println("FindOne() result:", result)
		// fmt.Println("findOne() ", result.Title)
	// }

	collection := client.Database("Articles").Collection("Earticles").Collection("
	")

	log.Println("starting API server")
	create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	specify endpoints
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	http.Handle("/", router)

	start and listen to requests

	http.ListenAndServe(":3000", router) uncomment to test

	// handleRequests()
}

// func prepareResponse() []Person {
	// var persons []Person

	// var person Person
	// person.Id = 1
	// person.FirstName = "Issac"
	// person.LastName = "N"
	// persons = append(persons, person)

	// person.Id = 2
	// person.FirstName = "Albert"
	// person.LastName = "E"
	// persons = append(persons, person)

	// person.Id = 3
	// person.FirstName = "Thomas"
	// person.LastName = "E"
	// persons = append(persons, person)
	// return persons
}

/*
// let's declare a global Article array
// that we can then populate in our main function
// to simulate a database
*/

// var Articles []Article

// func Persons(w http.ResponseWriter, r *http.Request) {
	// log.Println("entering persons end point")
	// var response Response
	// persons := prepareResponse()

	// response.Persons = persons

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// jsonResponse, err := json.Marshal(response)
	// if err != nil {
		// return
	// }

	// w.Write(jsonResponse)
}

// func homePage(w http.ResponseWriter, req *http.Request) {

	// fmt.Fprint(w, "Welcome to the Home Page!")
	// fmt.Fprintf(w, "End point Hit: homePage")
}

// func handleRequests() {

	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/allart", allArticle)
	// myRouter.HandleFunc("/allart/{Id}", returnSingleArticl)
	// myRouter.HandleFunc("/allarts/{id}", returnSingleArticle)
	// myRouter.HandleFunc("/delete/{id}", deleteArticle).Methods("DELETE")
	// myRouter.HandleFunc("/createnew", createNewArticle).Methods("POST")
	// myRouter.HandleFunc("/updateart/{Id}", updateArticle).Methods("PATCH")
	// log.Fatal(http.ListenAndServe(":2020", myRouter))

}

// func allArticle(w http.ResponseWriter, req *http.Request) {

	// s, err := json.Marshal(Articles)
	// if err != nil {
		// log.Print(err)
	// }
	// fmt.Fprintf(w, string(s))
	// req.Header.Set("contentType", "application/json")
	// req.Body.Read(s)
}

// func returnSingleArticl(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["Id"]

	loop over all of our articles
	if the articles.id equals the key we pass in
	return the article encoded as json
	
	for _, article := range Articles {
		if article.Id == key {
	fmt.Fprintln(w, "key: "+key)
	log.Println(w, "key: "+key)
	fmt.Println(article)
	json.NewEncoder(w).Encode(article)
	
	} else {
	
	http.Error(w, "Error!!", http.StatusInternalServerError)
	
	}
	}
	
	}
		}
}

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
		key := vars["id"]

	Loop over all of our Articles
	if the article.Id equals the key we pass in
	return the article encoded as JSON
	// for _, article := range Articles {
			if article.Id == key {
		// json.NewEncoder(w).Encode(article)
		}
	// }
}

// func deleteArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]

	k, err := strconv.Atoi(key)

	  fg []article :=  Articles
	var fg [] int
	// fmt.Println(Articles)
	// for index, article := range Articles {

			if article.Id == key {
		k, err := strconv.Atoi(article.Id)
		if err != nil {
			log.Println("Failed to convert")
		}

		// Articles = append(Articles[:index], Articles[index+1]) // copy(Articles[k:], Articles[k+1:])

		// fmt.Println("deleted from id", article.Id, "")

		// json.NewEncoder(w).Encode(Articles)

			}

	// }
}

// func createNewArticle(w http.ResponseWriter, r *http.Request) {
	get  the body of our POST  request
	return the String response containing the request body

	// reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	// var article Article // copy
	// json.Unmarshal(reqBody, &article)

	update our global Articles array to include

	our  new Article
	// Articles = append(Articles, article)

	// json.NewEncoder(w).Encode(article)

}

// func updateArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
		key := vars["Id"]

	// var articleupdate Article // copy

	// reqBody, err := ioutil.ReadAll(r.Body)

	// if err != nil {
		// fmt.Fprint(w, "Error")
	// }

	// json.Unmarshal(reqBody, &articleupdate)

	// for k, article := range Articles {

			if article.Id == key {

		// article.Title = articleupdate.Title
		// article.Desc = articleupdate.Desc
		// article.Content = articleupdate.Content

		// Articles = append(Articles[:k], article)
		// fmt.Println(w, "edited from Id:", article.Id)
		// json.NewEncoder(w).Encode(article)

		}

	// }

}
