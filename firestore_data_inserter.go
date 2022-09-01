package firesert

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/jonnyorman/fireworks"
)

type FirestoreDataInserter[T any] struct {
	configuration fireworks.Configuration
}

func NewFirestoreDataInserter[T any](configuration fireworks.Configuration) *FirestoreDataInserter[T] {
	this := new(FirestoreDataInserter[T])

	this.configuration = configuration

	return this
}

func (this FirestoreDataInserter[T]) Insert(data T) {
	client, _ := firestore.NewClient(context.Background(), this.configuration.ProjectID)

	defer client.Close()

	collection := client.Collection(this.configuration.CollectionName)

	_, _, _ = collection.Add(context.Background(), data)
}
