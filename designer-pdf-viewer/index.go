package algorithms

func DesignerPdfViewer(h []int32, word string) int32 {
	maxHeight := 1
	for _, letter := range word {
		index := letter - 97
		if int(h[index]) > maxHeight {
			maxHeight = int(h[index])
		}
	}
	return int32(len(word) * maxHeight)
}
