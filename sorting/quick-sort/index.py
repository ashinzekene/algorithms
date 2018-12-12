def quick_sort(A, start = None, end = None):
    """
    Sorts elements of a list in ascending order
    """
    start = 0 if start == None else start
    end = len(A) if end == None else end
    def partition(A, start, end):
        """
        Puts an element(last element) of a list in 
        place (its ordered) position
        """
        pivot = A[end - 1]
        pIndex = start
        for i in range(start, end):
            if A[i] <= pivot:
                temp = A[i]
                A[i] = A[pIndex]
                A[pIndex] = temp
                pIndex += 1
        return pIndex
    
    if (start >= end):
        return
    partitionIndex = partition(A, start, end)
    quick_sort(A, start, partitionIndex - 1)
    quick_sort(A, partitionIndex, end)
    return A


print(quick_sort([1,4,5,2,6,7,9,8,3,24,72,14,13,55,42,16,31,30,29,28,27,20, 10, 12, 11]))
