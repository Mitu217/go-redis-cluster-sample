package main

import (
	. "testing"

	"github.com/Mitu217/go-redis-cluster-sample/goredis"
	"github.com/Mitu217/go-redis-cluster-sample/radix"
)

func BenchmarkSerialGetSet(b *B) {
	hosts := []string{
		"127.0.0.1:7000",
		"127.0.0.1:7001",
		"127.0.0.1:7002",
	}

	b.Run("radix", func(b *B) {
		cli, err := radix.NewRadixClient(hosts)
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := cli.SetGet("foo", "bar"); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("go-redis", func(b *B) {
		cli, err := goredis.NewGoRedisClient(hosts)
		if err != nil {
			b.Fatal(err)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := cli.SetGet("foo", "bar"); err != nil {
				b.Fatal(err)
			}
		}
	})
}
