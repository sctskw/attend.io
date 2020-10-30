package db

type DatabaseClient interface {
	Connect()
	Ping()
	FetchAll(collection string) FetchAllResponse
	FetchAllById(collection string, ids ...string) FetchAllResponse
	FetchById(collection, id string) FetchOneResponse
	FetchByField(collection, field, value string) FetchOneResponse
}

type FetchAllResponse = []FetchOneResponse
type FetchOneResponse = []byte

func NewClient() DatabaseClient {
	c := NewFirestoreClient()
	c.Connect()
	return c
}

func NewMockClient() DatabaseClient {
	return NewMemoryClient()
}
