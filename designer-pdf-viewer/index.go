package algorithms

import "fmt"

func main() {
	// fmt.Println("a"[0])
	fmt.Println(designerPdfViewer(
		[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		"abc",
	))
}

func designerPdfViewer(h []int32, word string) int32 {
	maxHeight := 1
	for _, letter := range word {
		index := letter - 97
		if int(h[index]) > maxHeight {
			maxHeight = int(h[index])
		}
	}
	return int32(len(word) * maxHeight)
}
