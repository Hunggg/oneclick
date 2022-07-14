package rediscache

import "time"

type CacheClient struct {

}

func NewCacheClient() *CacheClient {
	return &CacheClient{}
}

func (c *CacheClient) Set(key string, value interface{}, expires time.Duration) error {
	return nil
}

func (c *CacheClient) Get(key string) (interface{}, bool) {
	return nil, true
}

func (c *CacheClient) Delete(key string) error {
	return nil
}

func (c *CacheClient) LPush(key string, value interface{}, expires time.Duration) error {
	return nil
}

func (c *CacheClient) LGet(key string, index int) (interface{}, bool) {
	return nil, true
}

func (c *CacheClient) LGetAll(key string) ([]interface{}, bool) {
	return nil, true
}

func (c *CacheClient) LDelete(key string) error {
	return nil
}