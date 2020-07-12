package main

import "bookmark/store"

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
