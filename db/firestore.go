package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"cloud.google.com/go/firestore"

	"github.com/google/uuid"
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

func SetCredentialsPath(path string) {
	f, _ := filepath.Abs(path)
	info, err := os.Stat(f)

	if err == nil && !info.IsDir() {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", f)
	}
}

func (f *firestoreClient) Connect() {
	//TODO: make this a config
	SetCredentialsPath(".gcloud/attend-io.creds.json")

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
	return f.FetchByField(collection, "refId", id)
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

	refId := uuid.New().String()

	//generate  a custom  internal id since  doc ids  are  unstable
	data["refId"] = refId

	ref, _, err := f.client.Collection(collection).Add(context.Background(), data)

	if err != nil {
		return nil, err
	}

	if ref == nil {
		return nil, errors.New("didnt create document object")
	}

	return f.FetchById(collection, refId)
}

func (f *firestoreClient) Update(collection, id string, b []byte) (FetchOneResponse, error) {

	ctx := context.Background()

	doc, err := f.getDocByRefId(ctx, collection, id)

	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{}, 0)
	_ = json.Unmarshal(b, &data)

	_, err = doc.Ref.Set(ctx, data)

	if err != nil {
		return nil, err
	}

	return f.FetchById(collection, id)
}

func (f *firestoreClient) DeleteById(collection, id string) EmptyResponse {
	ctx := context.Background()
	doc, err := f.getDocByRefId(ctx, collection, id)

	if err != nil {
		return nil
	}

	_, _ = doc.Ref.Delete(ctx)

	return nil
}

func (f *firestoreClient) getDocByRefId(ctx context.Context, collection, id string) (*firestore.DocumentSnapshot, error) {
	return f.client.Collection(collection).
		Where("refId", "==", id).
		Limit(1).
		Documents(ctx).
		Next()
}
