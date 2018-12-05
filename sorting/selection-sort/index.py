import time

def report_time(func, *args):
    st = time.time()
    for _ in range(9999):
        func(*args)
    total = round(time.time() - st, 5)
    print('-------------------------------------')
    print('{} took {} seconds'.format(func.__name__, total))
    return func(*args)

def selection_sort(arr):
    for i in range(len(arr)):
        mini = arr[i]
        min_index = i
        # Get the smallest element
        for x in range(i, len(arr)):
            if arr[x] < mini:
                mini = arr[x]
                min_index = x
        # Swap it with the index of the current array iteration
        temp = arr[i]
        arr[i] = mini
        arr[min_index] = temp
        
    return arr
    
print(
    report_time(selection_sort, [1,4,5,2,6,7,9,8,3,24,72,14,13,55,42,31,30,29,28,27,20])
)