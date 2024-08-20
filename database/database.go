package database

import (
	"context"
	"log"
	"myapp/graph/model"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ConnectionString string = "mongodb+srv://mongo:mongo@cluster0.7xdb4.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

// Connect connects to the database.

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(ConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{client: client}
}

func (db *DB) GetJob(id string) *model.JobListing {
	var jobListing model.JobListing

	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	var jobListings []*model.JobListing

	return jobListings
}

func (db *DB) CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing {
	var jobListing model.JobListing
	return &jobListing
}

func (db *DB) UpdateJobListing(id string, jobInfo model.UpdateJobListingInput) *model.JobListing {
	var returnJobListing model.JobListing
	return &returnJobListing
}

func (db *DB) DeleteJobListing(id string) *model.DeleteJobResponse {
	return &model.DeleteJobResponse{DeletedJobID: &id}
}
