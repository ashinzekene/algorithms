def rotLeft(a, d):
    length = len(a)
    rota = d % length
    res = [0 for x in a]
    for x in range(length):
        res[(x - rota) % length] = a[x]
    return res
    
print(rotLeft([1,2,3,4,5,6,7,8,9], 10000000000000000000000000000000))
1
4,3,2
100,100,90,80,60

25,60,70,75,78,80,97