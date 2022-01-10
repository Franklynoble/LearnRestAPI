package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//import "go.mongodb.org/mongo-driver/bson/primitive"

//import "go.mongodb.org/mongo-driver/bson/primitive"

// create struct
// we defined json and bson name to them when our structs serialize.
// means that if we did not assign any value our struct's field, do not show this field name
// after serialize process. Also, bson is relevant to mongo-driver. it works for   creating fileter
// type Book struct {
// ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// Isbn   string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
// Title  string             `json:"title,omitempty" bson:"title,omitempty"`
// Author *Author            `json:"author,omitempty" bson:"author,om"`
// }
// type Author struct {
// FirstName string `json:"firstname,omitempty" bson:"firstname"`
// LastName  string `json:"lastname,omitempty" bson:"lastname"`
// }
//

type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Isbn   string             `json:"isbn,omitempty" bson:"isbn,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}

type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
