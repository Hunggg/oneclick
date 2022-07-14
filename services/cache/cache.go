package cache

import "time"

type Cache interface {
	Set(key string, value interface{}, expires time.Duration) error
	Get(key string) (interface{}, bool)
	Delete(key string) error

	LPush(key string, value interface{}, expires time.Duration) error
	LGet(key string, index int) (interface{}, bool)
	LGetAll(key string) ([]interface{}, bool)
	LDelete(key string) error
}