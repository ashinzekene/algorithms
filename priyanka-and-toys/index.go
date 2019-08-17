package main

import (
	"sort"
)

func toys(w []int32) int32 {
	sort.Slice(w, func(i, j int) bool {
		return i < j
	})
	mins := make([]int32, 0)
	for i, v := range w {
		if i == 0 {
			mins = append([]int32{v}, mins...)
		} else {
			if v > mins[0]+4 {
				mins = append([]int32{v}, mins...)
			}
		}
	}
	return int32(len(mins))
}
