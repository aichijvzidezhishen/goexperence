package tree

import (
	"fmt"
	"testing"
)

func TestTernaryTrie_Insert(t *testing.T) {
	trie := NewTernaryTrie()
	trie.Insert("abc")
	trie.Insert("abcd")
	trie.Insert("abcde")
	trie.Insert("abcdef")
	trie.Insert("abcdefg")
	trie.Insert("abcdefgh")
	trie.Insert("abcdefghi")
	trie.Insert("abcdefghij")
	trie.Insert("abcdefghijk")
	trie.Insert("abcdefghijkl")
	trie.Insert("abcdefghijklm")
	trie.Insert("abcdefghijklmn")

	fmt.Println(trie.Search("abca"))
}
