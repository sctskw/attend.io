package db

type DatabaseClient interface {
	Connect()
	Ping()
	FetchAll(collection string) (FetchAllResponse, error)
	FetchAllById(collection string, ids ...string) (FetchAllResponse, error)
	FetchById(collection, id string) (FetchOneResponse, error)
	FetchByField(collection, field, value string) (FetchOneResponse, error)
	Insert(collection string, b []byte) (FetchOneResponse, error)
	Update(collection, path string, b []byte) (FetchOneResponse, error)
	DeleteById(collection, id string) EmptyResponse
}

type FetchAllResponse = []FetchOneResponse
type FetchOneResponse = []byte
type EmptyResponse = []byte

func NewClient() DatabaseClient {
	c := NewFirestoreClient()
	c.Connect()
	return c
}

func NewMockClient() DatabaseClient {
	return NewMemoryClient()
}
