def sockMerchant(n, ar):
    res = {}
    ans = 0
    for x in ar:
        if x in res:
            res[x] += 1
        else:
            res[x] = 1
    
    for x in res.values():
        ans += (x // 2)

    return ans