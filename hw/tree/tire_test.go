package tree

import "testing"

func TestTire(t *testing.T) {
	trie := NewTrie()
	keys := []string{"her", "say", "she", "shr", "he"}
	for _, v := range keys {
		trie.Insert(v)
	}

	if trie.Search("a") {
		t.Log("search succ")
	} else {
		t.Log("search faild ")
	}

}
