package store

import "testing"

func TestData(t *testing.T) {
	d := NewData()
	d.Add("foo", "bar")
	d.Add("hello", "world")

	s := d.Get([]string{"foo", "hello"})
	if a, e := len(s), 2; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	foo := s[0]

	if a, e := foo, [2]string{"foo", "bar"}; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	hello := s[1]

	if a, e := hello, [2]string{"hello", "world"}; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}

	s = d.Get([]string{})

}
