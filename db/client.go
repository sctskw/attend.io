package db

type DatabaseClient interface {
	Connect()
	Ping()
	FetchTalksById(id string)
}

func NewClient() DatabaseClient {
	return NewFirestoreClient()
}

func NewMockClient() DatabaseClient {
	return NewMemoryClient()
}
