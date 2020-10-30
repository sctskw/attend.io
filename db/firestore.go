package db

type FirestoreClient interface {
	DatabaseClient
}

type firestoreClient struct {
}

func NewFirestoreClient() FirestoreClient {
	return &firestoreClient{}
}

func (c *firestoreClient) Connect() {

}

func (c *firestoreClient) Ping() {

}

func (c *firestoreClient) FetchTalksById(id string) {

}
