package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, _ = mongo.NewClient(options.Client().ApplyURI(ENV_DATABASE_URL))
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
var db = client.Database("mo1394_lightning")
var links = db.Collection("links")

func mongodbInit() {
	var err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func takeMe(ID string) {
	ctx, errrrrrr := context.WithTimeout(context.Background(), 10*time.Second)
	if errrrrrr != nil {
		fmt.Println(errrrrrr)
	}

	var link bson.M
	if err := links.FindOne(ctx, bson.M{"ID": ID}).Decode(&link); err != nil {
		log.Print(err)
	}
	fmt.Println(link["Link"])
}

func PostLink(req LinkResponse) {
	ctx, errrrrrr := context.WithTimeout(context.Background(), 10*time.Second)
	if errrrrrr != nil {
		fmt.Println(errrrrrr)
	}

	newLink, err := links.InsertOne(ctx, bson.D{
		{Key: "ID", Value: req.ID},
		{Key: "Link", Value: req.Link},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newLink)
}
