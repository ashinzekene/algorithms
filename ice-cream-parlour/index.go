package algorithms

import "fmt"

type Node struct {
	value int32
	next  *Node
}

func whatFlavors(cost []int32, money int32) string {
	values := make(map[int32]*Node)
	for i, price := range cost {
		if values[price] == nil {
			values[price] = &Node{int32(i + 1), nil}
		} else {
			val := values[price]
			for val.next != nil {
				val = val.next
			}
			val.next = &Node{int32(i + 1), nil}
		}
	}
	for price, indexNode := range values {
		index := indexNode.value
		complimentNode := values[money-price]
		for complimentNode != nil {
			complimentIndex := complimentNode.value
			if complimentIndex != index {
				if index < complimentIndex {
					return fmt.Sprintf("%d %d", index, complimentIndex)
				} else {
					return fmt.Sprintf("%d %d", complimentIndex, index)
				}
			}
			complimentNode = complimentNode.next
		}
	}
	return ""
}
