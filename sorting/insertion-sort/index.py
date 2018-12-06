def insertionSort(arr):
    for x in range(1, len(arr)):
        item = arr[x]
        i = x - 1
        while i >= 0 and item < arr[i]:
            arr[i + 1] = arr[i]
            i -= 1
        arr[i + 1] = item
    return arr
    
print(insertionSort([1,9,8,4,6,7,3,12,5,18,2,22]))