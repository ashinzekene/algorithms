def rotLeft(a, d):
    length = len(a)
    rota = d % length
    res = [0 for x in a]
    for x in range(length):
        res[(x - rota) % length] = a[x]
    return res
    
print(rotLeft([1,2,3,4,5,6,7,8,9], 10000000000000000000000000000000))