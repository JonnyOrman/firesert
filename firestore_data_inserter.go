package firesert

import (
	"context"

	"cloud.google.com/go/firestore"
)

type FirestoreDataInserter[T any] struct {
	configuration Configuration
}

func NewFirestoreDataInserter[T any](configuration Configuration) *FirestoreDataInserter[T] {
	firestoreDataInserter := new(FirestoreDataInserter[T])
	firestoreDataInserter.configuration = configuration
	return firestoreDataInserter
}

func (dataInserter FirestoreDataInserter[T]) Insert(data T) {
	client, _ := firestore.NewClient(context.Background(), dataInserter.configuration.ProjectID)

	defer client.Close()

	collection := client.Collection(dataInserter.configuration.CollectionName)

	_, _, _ = collection.Add(context.Background(), data)
}
