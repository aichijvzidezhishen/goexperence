package tree

type TrieNode struct {
	children      map[rune]*TrieNode
	isEndOfWord   bool
	faildNodeLink *TrieNode
	outPutLink    *TrieNode
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		if current.children[char] == nil {
			current.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		current = current.children[char]
	}
	current.isEndOfWord = true
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return current.isEndOfWord
}

func (t *Trie) StartsWith(prefix string) bool {
	current := t.root
	for _, char := range prefix {
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return true
}
