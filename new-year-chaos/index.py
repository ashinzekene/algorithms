# Better Algorithm needed. Current -55%

def minimumBribes(q):
    bribes = 0
    def noOfGreater(n, arr):
        c = 0
        for x in arr:
            if n > x:
                c += 1
                if c > 2:
                    return -1
        return c
        
    for x in range(len(q)):
        if q[x] > (x + 3):
            return 'Too chaotic'
        no_of_bribes = noOfGreater(q[x], q[x:])
        if(no_of_bribes > -1):
            bribes += no_of_bribes
        else:
            return 'Too chaotic'
    return bribes

print(minimumBribes([2, 1, 5, 3, 4]))
print(minimumBribes([2, 5, 1, 3, 4]))
