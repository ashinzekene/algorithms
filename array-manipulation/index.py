def arrayManipulation(n, queries):
    arr = [0] * n
    max_n = 0
    cur = 0
    for q in queries:
        arr[q[0]-1] += q[2]
        if q[1] < n:
            arr[q[1]] += -q[2]
        print(arr)
    for elem in arr:
        cur += elem
        if cur > max_n:
            max_n = cur
    return max_n
    
    
    
print(arrayManipulation(10, [
[2, 6, 8],
[3, 5, 7],
[1, 8, 1],
[5, 9, 15],
]))