package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trie"
)

func Contacts(queries [][]string) []int32 {
	trie := NewTrie()
	result := make([]int32, 0)
	for _, query := range queries {
		op := query[0]
		if op == "add" {
			trie.Insert(query[1])
		} else if op == "find" {
			count := trie.Count(query[1])
			result = append(result, int32(count))
		}
	}
	return result
}
