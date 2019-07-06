def minimumBribes(q):
    bribes = 0
    for x in range(len(q)):
        if q[x] > (x + 3):
            print('Too chaotic')
            return
        for j in range(max(0, q[x]-2), x):
            if q[j] > q[x]:
                bribes += 1
    print(bribes)

print(minimumBribes([2, 1, 5, 3, 4]))
print(minimumBribes([2, 5, 1, 3, 4]))
