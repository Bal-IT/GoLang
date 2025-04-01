package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB []Book
var booksCollection *mongo.Collection
var opts *options.ClientOptions
var ctx context.Context

type Book struct {
	Id        string `json:"id" bson:"-"`
	Author    string `json:"author" bson:"author"`
	Country   string `json:"country" bson:"country"`
	ImageLink string `json:"imageLink" bson:"imageLink"`
	Language  string `json:"language" bson:"language"`
	Link      string `json:"link" bson:"link"`
	Title     string `json:"title" bson:"title"`
}

func init() {
	ctx = context.TODO()
	opts = options.Client().ApplyURI("mongodb://10.118.36.178:27017")

	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}

	// Close current connection
	// defer client.Disconnect(ctx)
	booksDb := client.Database("booksDb")
	booksCollection = booksDb.Collection("books")
}

// FindBookByID function find book by id
func FindBookByID(id string) (Book, bool) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return Book{}, false
	}

	sr := booksCollection.FindOne(ctx, bson.M{"_id": objectID})

	var b Book
	err = sr.Decode(&b)
	if err != nil {
		log.Print(err)
		return Book{}, false
		//log.Fatal(err)
	}
	b.Id = id

	return b, true
}

// FindBookByID function find book by id
func GetAllBooks() ([]Book, bool) {
	/*
		limitOpts := options.Find()
		limitOpts.SetSkip(3)
		limitOpts.SetLimit(5)
	*/
	allBooks, err := booksCollection.Find(ctx, bson.M{} /*, limitOpts */)
	if err != nil {
		log.Print(err)
		return nil, false
	}

	var bAll []Book
	if err := allBooks.All(ctx, &bAll); err != nil {
		log.Print(err)
		return nil, false
	}

	return bAll, true
}

func DeleteBookById(id string) bool {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return false
	}

	_, b := FindBookByID(id)
	if !b {
		log.Print("book whith id " + id + "not found")
		return false
	}

	_, err = booksCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

func CreateBook(book Book) bool {
	_, err := booksCollection.InsertOne(ctx, book)
	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

func UpdateBookById(id string, book Book) bool {

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return false
	}

	_, b := FindBookByID(id)
	if !b {
		log.Print("book whith id " + id + "not found")
		return false
	}

	_, err = booksCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},

		bson.D{
			{Key: "$set", Value: bson.M{"author": book.Author}},
			{Key: "$set", Value: bson.M{"country": book.Country}},
			{Key: "$set", Value: bson.M{"imageLink": book.ImageLink}},
			{Key: "$set", Value: bson.M{"language": book.Language}},
			{Key: "$set", Value: bson.M{"link": book.Link}},
			{Key: "$set", Value: bson.M{"title": book.Title}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return true
}
