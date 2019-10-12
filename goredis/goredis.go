package goredis

import (
	"fmt"

	redis "github.com/go-redis/redis/v7"
)

type GoRedisClient struct {
	cli *redis.ClusterClient
}

func NewGoRedisClient(addrs []string) (*GoRedisClient, error) {
	cli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
	})
	return &GoRedisClient{cli}, nil
}

func (c GoRedisClient) Close() error {
	return c.cli.Close()
}

func (c GoRedisClient) SetGet(key string, val string) error {
	if err := c.Set(key, val); err != nil {
		return err
	}
	got, err := c.Get(key)
	if err != nil {
		return err
	}
	if got != val {
		return fmt.Errorf("got wrong value")
	}
	return nil
}

func (c GoRedisClient) Set(key string, val string) error {
	return c.cli.Set(key, val, 0).Err()
}

func (c GoRedisClient) Get(key string) (string, error) {
	got, err := c.cli.Get(key).Result()
	if err != nil {
		return "", err
	}
	return got, nil
}
