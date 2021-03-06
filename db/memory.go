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

func (c *memoryClient) FetchAll(collection string) (FetchAllResponse, error) {
	return nil, nil
}

func (c *memoryClient) FetchAllById(collection string, ids ...string) (FetchAllResponse, error) {
	return nil, nil
}

func (c *memoryClient) FetchById(collection, id string) (FetchOneResponse, error) {
	return nil, nil
}

func (c *memoryClient) FetchByField(collection, field, value string) (FetchOneResponse, error) {
	return nil, nil
}

func (c *memoryClient) Insert(collection string, b []byte) (FetchOneResponse, error) {
	return nil, nil
}

func (c *memoryClient) Update(collection, id string, b []byte) (FetchOneResponse, error) {
	return nil, nil
}

func (c *memoryClient) DeleteById(collection, id string) EmptyResponse {
	return nil
}
