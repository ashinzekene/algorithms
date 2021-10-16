package algorithms

func maxProductThree(A []int) int {
    sort.Ints(A)
    l := len(A)
    positiveMax := A[l-1] * A[l-2]* A[l-3]
    if A[0] >= 0 {
        return positiveMax
    }
    negativeMax := A[0] * A[1] * A[l-1]
    if positiveMax > negativeMax {
        return positiveMax
    }
    return negativeMax
}
