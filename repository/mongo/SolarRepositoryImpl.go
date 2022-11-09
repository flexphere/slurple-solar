package sqlite

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/flexphere/slurple-solar/repository"
)

type SolarRepositoryImpl struct {
	db      *mongo.Client
	context context.Context
}

func (r *SolarRepositoryImpl) Connect() {
	uri := os.Getenv("DATABASE")
	r.context = context.TODO()
	db, err := mongo.Connect(r.context, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := db.Ping(r.context, readpref.Primary()); err != nil {
		panic(err)
	}

	r.db = db
}

func (r *SolarRepositoryImpl) Disconnect() {
	if err := r.db.Disconnect(r.context); err != nil {
		panic(err)
	}
}

func (r *SolarRepositoryImpl) SaveRecords(records []repository.SolarRecord) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for _, record := range records {
		doc, err := toDoc(record)
		if err != nil {
			log.Fatal(err)
		}

		filter := bson.D{{"ts", record.TS}}
		update := bson.D{{"$set", doc}}
		opts := options.Update().SetUpsert(true)
		r.db.Database("slurple").Collection("solar").UpdateOne(ctx, filter, update, opts)
	}
}

func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
