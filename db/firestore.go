package db

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
)

type FirestoreClient interface {
	DatabaseClient
}

type firestoreClient struct {
	client *firestore.Client
}

func NewFirestoreClient() FirestoreClient {
	return &firestoreClient{}
}

func (f *firestoreClient) Connect() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "attend-io-294107")
	if err != nil {
		panic(fmt.Sprintf("couldn't connect to firestore db: %s", err.Error()))
	}

	f.client = client
}

func (f *firestoreClient) Ping() {

}

func (f *firestoreClient) FetchAll(collection string) (results FetchAllResponse) {
	talks := f.client.Collection(collection)
	docs, err := talks.Documents(context.Background()).GetAll()

	if err != nil {
		panic(err)
	}

	for _, doc := range docs {
		d := doc.Data()
		d["id"] = doc.Ref.ID //copy the id

		//convert to bytes
		b, _ := json.Marshal(d)

		results = append(results, b)
	}

	return results
}

func (f *firestoreClient) FetchAllById(collection string, ids ...string) (results FetchAllResponse) {

	//NOTE: not great, but Firestore is limited to 10 logical OR operators so we have to do this in code
	for _, id := range ids {
		results = append(results, f.FetchById(collection, id))
	}

	return results
}

func (f *firestoreClient) FetchById(collection, id string) FetchOneResponse {

	doc, err := f.client.Collection(collection).Doc(id).Get(context.Background())

	if err != nil {
		return nil
	}

	d := doc.Data()
	d["id"] = doc.Ref.ID //copy the id

	//convert to bytes
	b, err := json.Marshal(d)

	if err != nil {
		panic(err)
	}

	return b
}

func (f *firestoreClient) FetchByField(collection, field, value string) FetchOneResponse {

	doc, err := f.client.Collection(collection).
		Where(field, "==", value).
		Limit(1).
		Documents(context.Background()).
		Next()

	if err != nil {
		return nil
	}

	d := doc.Data()
	d["id"] = doc.Ref.ID //copy the id

	//convert to bytes
	b, err := json.Marshal(d)

	if err != nil {
		panic(err)
	}

	return b
}
