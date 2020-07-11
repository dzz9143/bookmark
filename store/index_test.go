package store

import "testing"

func TestIndex(t *testing.T) {
	idx := &Index{}

	idx.Add("search engine", "google", "bing", "baidu")
	idx.Add("microsoft", "bing")
	idx.Add("china", "baidu", "bing")

	s := idx.Query([]string{"china", "search"})

	if a, e := len(s), 2; e != a {
		t.Errorf("Test failed, expect: %v, actual: %v", e, a)
	}
}
