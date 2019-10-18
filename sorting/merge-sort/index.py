def merge_sort(A):
    """
    A function that sorts elements of an list
    """
    def merge(L, R, arr):
        """
        This merges the left (L) and right (R) parts
        of list (arr) in order
        """
        l = r = i = 0
        print(len(L), len(R), len(arr))
        while l < len(L) and r < len(R):
            if L[l] < R[r]:
                arr[i] = L[l]
                l += 1
            else:
                arr[i] = R[r]
                r += 1
            i += 1
        while l < len(L): 
            arr[i] = L[l]
            l += 1
            i += 1
        while r < len(R):
            arr[i] = R[r]
            r += 1
            i += 1
        return arr
    if (len(A) < 2):
        return A
    mid = len(A) // 2
    L = merge_sort(A[:mid])
    R = merge_sort(A[mid:])
    merge(L, R, A)
    return A
    
    
print(merge_sort([1,4,5,2,6,7,9,8,3,24,72,14,13,55,42,16,31,30,29,28,27,20, 10, 12, 11]))