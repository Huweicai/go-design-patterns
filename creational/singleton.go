package creational

import "sync"

type singleMap map[string]string

var (
	once     sync.Once
	instance singleMap
)

func NewSingleMap() singleMap {
	once.Do(func() {
		instance = make(map[string]string)
	})
	return instance
}
