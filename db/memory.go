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

func (c *memoryClient) FetchAllTalks() FetchAllResponse {
	return nil
}

func (c *memoryClient) FetchTalkById(id string) FetchOneResponse {
	return nil
}

func (c *memoryClient) LoadFromFile(path string) {

}
