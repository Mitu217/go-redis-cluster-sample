package main

import (
	"runtime"
	"strings"
	. "testing"

	"github.com/Mitu217/go-redis-cluster-sample/goredis"
	"github.com/Mitu217/go-redis-cluster-sample/radix"
)

func BenchmarkSerialGetSet(b *B) {
	key := "foo"
	val := "bar"
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
		defer cli.Close()
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := cli.SetGet(key, val); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("go-redis", func(b *B) {
		cli, err := goredis.NewGoRedisClient(hosts)
		if err != nil {
			b.Fatal(err)
		}
		defer cli.Close()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := cli.SetGet(key, val); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkSerialGetSetLargeArgs(b *B) {
	key := strings.Repeat("foo", 24)
	val := strings.Repeat("bar", 4096)
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
		defer cli.Close()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := cli.SetGet(key, val); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("go-redis", func(b *B) {
		cli, err := goredis.NewGoRedisClient(hosts)
		if err != nil {
			b.Fatal(err)
		}
		defer cli.Close()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if err := cli.SetGet(key, val); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkParallelGetSet(b *B) {
	parallel := runtime.GOMAXPROCS(0)

	do := func(b *B, fn func() error) {
		b.ResetTimer()
		b.SetParallelism(parallel)
		b.RunParallel(func(pb *PB) {
			for pb.Next() {
				if err := fn(); err != nil {
					b.Fatal(err)
				}
			}
		})
	}

	key := "foo"
	val := "bar"
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
		defer cli.Close()

		do(b, func() error {
			return cli.SetGet(key, val)
		})
	})

	b.Run("goredis", func(b *B) {
		cli, err := goredis.NewGoRedisClient(hosts)
		if err != nil {
			b.Fatal(err)
		}
		defer cli.Close()

		do(b, func() error {
			return cli.SetGet(key, val)
		})
	})
}
