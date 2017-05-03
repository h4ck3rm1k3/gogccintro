package main

import "github.com/valyala/ybc/bindings/go/ybc"
import "time"

//var cache ybc.Cacher

type Cache struct {
	Cache ybc.Cacher
}
func (c *Cache) Close()  {
	c.Cache.Close()
}

func (c *Cache) Add(key [] byte, bytes [] byte )  {
	//c.Cache.
	MaxTtl := time.Hour * 24 * 365 * 100
	err := c.Cache.Set(key, bytes, MaxTtl)
	if err != nil {
		panic(err)
	}
}

func LoadCache() (*Cache) {
	config := ybc.Config{
		DataFile :  ".go-memcached.data",
		IndexFile : ".go-memcached.index",
		MaxItemsCount:   ybc.SizeT(1000),
		DataFileSize:    ybc.SizeT(10) * ybc.SizeT(1024*1024*1024),
		//HotItemsCount:   ybc.SizeT(*hotItemsCount),
		//HotDataSize:     ybc.SizeT(*hotDataSize),
		//DeHashtableSize: *deHashtableSize,
		//SyncInterval:    syncInterval_,
	}
	cache, err := config.OpenCache(true)
	if err != nil {
		panic("err")
	}
	
	return &Cache{Cache:cache}
}
