package store

import (
	"strings"
)

// Index - is a k:v store, v is a Set
type Index map[string]Set

// NewIndex - create a new index
func NewIndex() Index {
	return Index{}
}

// Add - add a group of tags into the index
func (i Index) Add(k string, values ...string) {
	s, ok := i[k]
	if !ok {
		s = NewSet()
		i[k] = s
	}

	for _, v := range values {
		s.Add(v)
	}
}

// Query - query if a set of keys contains in this index
func (i Index) Query(search []string) []string {
	var res Set

	for _, s := range search {
		cur := NewSet()
		for key, set := range i {
			if strings.Contains(key, s) {
				cur.Merge(set)
			}
		}

		if res == nil {
			res = cur
		} else {
			res = Intersect(res, cur)
		}
	}

	if res == nil {
		return []string{}
	}

	return res.ToSlice()
}
