package main

import (
	"bookmark/store"
	"fmt"
	"os"
)

// Bookmark - contains a reverse index & data store
type Bookmark struct {
	index store.Index
	data  store.Data
}

// NewBookmark - create new bookmark
func NewBookmark() *Bookmark {
	return &Bookmark{
		index: store.Index{},
		data:  store.Data{},
	}
}

// Add - store the k, v data & compile the revert index
func (b *Bookmark) Add(k, v string, tags ...string) {
	b.data.Add(k, v)
	b.index.Add(k, k)
	for _, t := range tags {
		b.index.Add(t, k)
	}
}

// Query - get data match the keys
func (b *Bookmark) Query(search []string) []store.KeyValue {
	k := b.index.Query(search)
	kv := b.data.Get(k)
	return kv
}

func main() {
	b := NewBookmark()

	b.Add("google", "www.google.com", "search")
	b.Add("baidu", "www.baidu.com", "search", "china")
	b.Add("bing", "www.bing.com", "search", "microsoft")
	b.Add("cnbing", "cn.bing.com", "search", "microsoft", "china")

	r := b.Query(os.Args[1:])
	if len(r) > 0 {
		fmt.Println("Search result:")
	}
	for i, kv := range r {
		fmt.Printf("[#%d] %s: %s\n", i, kv.Key(), kv.Value())
	}
}
