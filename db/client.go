package db

type DatabaseClient interface {
	Connect()
	Ping()
	FetchAllTalks() FetchAllResponse
	FetchTalkById(id string) FetchOneResponse
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
