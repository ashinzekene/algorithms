package algorithms

type TrieNode struct {
	End  bool
	Keys map[rune]*TrieNode
}

type Trie struct {
	Head *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Head: &TrieNode{
			Keys: make(map[rune]*TrieNode),
		},
	}
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		End:  false,
		Keys: make(map[rune]*TrieNode),
	}
}

func (t *Trie) Insert(word string) *Trie {
	currentTrie := t.Head
	for _, char := range word {
		if currentTrie.Keys[char] == nil {
			currentTrie.Keys[char] = NewTrieNode()
		}
		// currentTrie.Keys[char].Count++
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
	return count(currentTrie)
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
		if i == len(word)-1 {
			if len(currentTrie.Keys) == 0 {
				delete(currentTrie.Keys, char)
				break
			} else {
				currentTrie.End = false
			}
		}
		currentTrie = currentTrie.Keys[char]
	}
	return true
}

func count(t *TrieNode) int {
	result := 0
	if t.End {
		result++
	}
	for _, trie := range t.Keys {
		result += count(trie)
	}
	return result
}
