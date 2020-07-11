package store

// Set - set of strings, should only contain unique strings
type Set map[string]bool

// Intersect - get the intersect between two sets
func Intersect(a, b *Set) *Set {
	if a == nil || b == nil {
		return nil
	}

	s := &Set{}
	for k := range *a {
		s.Add(k)
	}

	for k := range *s {
		if !b.Exist(k) {
			s.Del(k)
		}
	}

	return s
}

// NewSet - create a new set
func NewSet() Set {
	return Set{}
}

// Add - add a string `k` into the set
func (s *Set) Add(k string) {
	(*s)[k] = true
}

// Merge - merge one set into the other
func (s *Set) Merge(other *Set) {
	if other == nil {
		return
	}

	for v, exist := range *other {
		if exist {
			s.Add(v)
		}
	}
}

// Del - delete a string `k` from the set
func (s *Set) Del(k string) {
	_, ok := (*s)[k]
	if ok {
		delete(*s, k)
	}
}

// Len - get length of set
func (s *Set) Len() int {
	return len(*s)
}

// Exist - check if a string `k` exists in the set
func (s *Set) Exist(k string) bool {
	return (*s)[k]
}

// ToSlice - convert the string set to string array
func (s *Set) ToSlice() []string {
	var col []string
	for k, ok := range *s {
		if ok {
			col = append(col, k)
		}
	}

	return col
}
