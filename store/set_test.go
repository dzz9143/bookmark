package store

import (
	"testing"
)

func stringSliceEqual(a, b []string) bool {
	alen := len(a)
	if alen != len(b) {
		return false
	}

	for i := 0; i < alen; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestSet(t *testing.T) {
	s := NewSet()

	if s == nil {
		t.Errorf("Test failed, fail to create a new set\n")
	}

	s.Add("foo")
	s.Add("bar")

	if a, e := s.Exist("foo"), true; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	if a, e := s.Exist("boo"), false; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	s.Del("bar")

	if a, e := s.Exist("bar"), false; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	s.Del("boo")

	if a, e := s.Exist("boo"), false; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	other := NewSet()

	other.Add("boo")

	s.Merge(other)

	if a, e := s.Exist("boo"), true; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}
}

func TestSetIntersect(t *testing.T) {
	a := Set{}
	b := Set{}

	a.Add("foo")
	a.Add("bar")
	a.Add("boo")

	b.Add("foo")
	b.Add("hello")

	s := Intersect(a, b)
	if a, e := s.Exist("foo"), true; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}
}
