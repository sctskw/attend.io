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

func (f *firestoreClient) FetchAll(collection string) (results FetchAllResponse, err error) {
	talks := f.client.Collection(collection)
	docs, err := talks.Documents(context.Background()).GetAll()

	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		d := doc.Data()
		d["id"] = doc.Ref.ID //copy the id
		d["ref"] = doc.Ref.Path

		//convert to bytes
		b, _ := json.Marshal(d)

		results = append(results, b)
	}

	return results, nil
}

func (f *firestoreClient) FetchAllById(collection string, ids ...string) (results FetchAllResponse, err error) {

	//NOTE: not great, but Firestore is limited to 10 logical OR operators so we have to do this in code
	for _, id := range ids {

		item, err := f.FetchById(collection, id)

		if err != nil {
			continue
		}

		results = append(results, item)
	}

	return results, nil
}

func (f *firestoreClient) FetchById(collection, id string) (FetchOneResponse, error) {

	doc := f.client.Collection(collection).Doc(id)
	item, err := doc.Get(context.Background())

	if err != nil {
		return nil, err
	}

	d := item.Data()
	d["id"] = item.Ref.ID //copy the id
	d["ref"] = item.Ref.Path

	//convert to bytes
	b, err := json.Marshal(d)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (f *firestoreClient) FetchByField(collection, field, value string) (FetchOneResponse, error) {

	doc, err := f.client.Collection(collection).
		Where(field, "==", value).
		Limit(1).
		Documents(context.Background()).
		Next()

	if err != nil {
		return nil, err
	}

	d := doc.Data()
	d["id"] = doc.Ref.ID //copy the id
	d["ref"] = doc.Ref.Path

	//convert to bytes
	b, err := json.Marshal(d)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (f *firestoreClient) Insert(collection string, b []byte) (FetchOneResponse, error) {

	data := make(map[string]interface{}, 0)
	err := json.Unmarshal(b, &data)

	if err != nil {
		return nil, err
	}

	if _, exists := data["id"]; exists {
		delete(data, "id")
	}

	ref, _, err := f.client.Collection(collection).Add(context.Background(), data)

	if err != nil {
		return nil, err
	}

	return f.FetchById(collection, ref.ID)
}

func (f *firestoreClient) Update(collection, id string, b []byte) (FetchOneResponse, error) {

	ctx := context.Background()
	doc := f.client.Collection(collection).Doc(id)

	data := make(map[string]interface{}, 0)
	_ = json.Unmarshal(b, &data)

	if _, exists := data["id"]; exists {
		delete(data, "id")
	}

	_, err := doc.Set(ctx, data)

	if err != nil {
		return nil, err
	}

	return f.FetchById(collection, id)
}

func (f *firestoreClient) DeleteById(collection, id string) EmptyResponse {
	doc := f.client.Collection(collection).Doc(id)
	_, _ = doc.Delete(context.Background())
	return nil
}
