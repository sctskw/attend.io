package db

type MemoryClient interface {
	DatabaseClient
	LoadFromFile(path string)
}

type memoryClient struct{}

func NewMemoryClient() MemoryClient {
	return &memoryClient{}
}

func (c *memoryClient) Connect() {

}

func (c *memoryClient) Ping() {

}

func (c *memoryClient) LoadFromFile(path string) {

}

func (c *memoryClient) FetchAll(collection string) FetchAllResponse {
	return nil
}

func (c *memoryClient) FetchAllById(collection string, ids ...string) FetchAllResponse {
	return nil
}

func (c *memoryClient) FetchById(collection, id string) FetchOneResponse {
	return nil
}

func (c *memoryClient) FetchByField(collection, field, value string) FetchOneResponse {
	return nil
}
