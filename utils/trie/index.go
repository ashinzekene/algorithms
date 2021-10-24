package algorithms

type TrieNode struct {
	Count int
	End   bool
	Keys  map[rune]*TrieNode
}

type Trie struct {
	Head *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Head: &TrieNode{
			Count: 0,
			Keys:  make(map[rune]*TrieNode),
		},
	}
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Count: 0,
		End:   false,
		Keys:  make(map[rune]*TrieNode),
	}
}

func (t *Trie) Insert(word string) *Trie {
	currentTrie := t.Head
	for _, char := range word {
		if currentTrie.Keys[char] == nil {
			currentTrie.Keys[char] = NewTrieNode()
		}
		currentTrie.Keys[char].Count++
		currentTrie = currentTrie.Keys[char]
	}
	currentTrie.End = true
	return t
}

func (t *Trie) Count(substr string) int {
	currentTrie := t.Head
	for _, char := range substr {
		if currentTrie.Keys[char] == nil {
			return 0
		}
		currentTrie = currentTrie.Keys[char]
	}
	return currentTrie.Count
}

func (t *Trie) Find(substr string) bool {
	currentTrie := t.Head
	for _, char := range substr {
		if currentTrie.Keys[char] == nil {
			return false
		}
		currentTrie = currentTrie.Keys[char]
	}
	return currentTrie.End
}

func (t *Trie) Remove(word string) bool {
	currentTrie := t.Head
	if !t.Find(word) {
		return false
	}
	for i, char := range word {
		currentTrie.Keys[char].Count--
		if i == len(word)-1 && currentTrie.Keys[char].Count == 0 {
			delete(currentTrie.Keys, char)
			break
		}
		currentTrie = currentTrie.Keys[char]
	}
	return true
}
