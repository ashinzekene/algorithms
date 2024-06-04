package utils

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a ...T) T {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

func Min[T constraints.Ordered](a ...T) T {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}

func UniqueStr(items []string) []string {
	holder := make(map[string]int)
	newItems := make([]string, 0)
	for i := 0; i < len(items); i++ {
		item := items[i]
		if holder[item] == 0 {
			holder[item] = 1
			newItems = append(newItems, item)
		}
	}
	return newItems
}

func UniqueInt(items []int) []int {
	holder := make(map[int]int)
	newItems := make([]int, 0)
	for i := 0; i < len(items); i++ {
		item := items[i]
		if holder[item] == 0 {
			holder[item] = 1
			newItems = append(newItems, item)
		}
	}
	return newItems
}
