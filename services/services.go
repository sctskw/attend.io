package services

import (
	"sync"

	"github.com/sctskw/attend.io/db"
)

var newOnce sync.Once
var withOnce sync.Once
var instance ServiceRegistry

func New() ServiceRegistry {
	newOnce.Do(func() {
		instance = NewServiceRegistry(db.NewClient())
	})
	return instance
}

//useful for testing
func NewWithClient(client db.DatabaseClient) ServiceRegistry {
	withOnce.Do(func() {
		instance = NewServiceRegistry(client)
	})
	return instance
}

//TODO
func NewMockRegistry() ServiceRegistry {
	return nil
}

func Get() ServiceRegistry {

	if instance == nil {
		panic("service registry has not been instantiated")
	}

	return instance
}
