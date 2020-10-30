package db

type MemoryClient interface {
	DatabaseClient
	LoadFromFile(path string)
}

type memoryClient struct{}

func NewMemoryClient() FirestoreClient {
	return &memoryClient{}
}

func (c *memoryClient) Connect() {

}

func (c *memoryClient) Ping() {

}

func (c *memoryClient) FetchTalksById(id string) {

}
