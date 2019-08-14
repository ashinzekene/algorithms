import time

def report_time(func, *args):
    st = time.time()
    for _ in range(9999):
        func(*args)
    total = round(time.time() - st, 5)
    print('-------------------------------------')
    print('{} took {} seconds'.format(func.__name__, total))
    return func(*args)

def bubble_sort(arr):
    end = 1
    isUnsorted = True
    while isUnsorted:
        isUnsorted = False
        for i in range(len(arr) - end):
            if arr[i] > arr[i + 1]:
                isUnsorted = True
                temp = arr[i]
                arr[i] = arr[i + 1]
                arr[i + 1] = temp
        end += 1
    return arr

print(
   report_time(bubble_sort, [1,4,5,2,6,7,9,8,3,24,72,14,13,55,42,31,30,29,28,27,20])
)
