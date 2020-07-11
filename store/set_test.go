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

	if a, e := s.Len(), 2; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	s.Add("foo") // add a duplicate value will not change anything

	if a, e := s.Len(), 2; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	s.Del("foo")

	if a, e := s.Exist("foo"), false; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	if a, e := s.Len(), 1; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	s.Del("boo") // delete a not exist value will not change anything

	if a, e := s.Exist("boo"), false; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	if a, e := s.Len(), 1; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	s.Del("bar") // delete the last value in set

	if a, e := s.Exist("bar"), false; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	if a, e := s.Len(), 0; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	for i := 0; i < 3; i++ {
		s.Add(string(i))
	}

	arr := s.ToSlice()
	if a, e := stringSliceEqual(arr, []string{"0", "1", "2"}), false; e != a {
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

	s := Intersect(&a, &b)
	if a, e := s.Exist("foo"), true; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	if a, e := s.Len(), 1; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}
}
