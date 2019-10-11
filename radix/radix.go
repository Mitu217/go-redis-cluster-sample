package radix

import (
	"fmt"

	radix "github.com/mediocregopher/radix/v3"
)

type RadixClient struct {
	cli radix.Client
}

func NewRadixClient(hosts []string) (*RadixClient, error) {
	cluster, err := radix.NewCluster(hosts)
	if err != nil {
		return nil, err
	}
	return &RadixClient{cluster}, nil
}

func (c RadixClient) Close() error {
	return c.cli.Close()
}

func (c RadixClient) SetGet(key string, val string) error {
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

func (c RadixClient) Set(key string, val string) error {
	if err := c.cli.Do(radix.Cmd(nil, "SET", key, val)); err != nil {
		return err
	}
	return nil
}

func (c RadixClient) Get(key string) (string, error) {
	var out string
	if err := c.cli.Do(radix.Cmd(&out, "GET", key)); err != nil {
		return "", err
	}
	return out, nil
}
