import heapq
import math

def networkDelayTime(times, n, k):
    adjList = {i: {} for i in range(1, n+1)}

    for u, v, w in times:
        if u in adjList:
            adjList[u][v] = w

    delays = [math.inf] * (n + 1)
    delays[k] = 0

    minHeap = [(0, k)]

    while minHeap:
        delay, node = heapq.heappop(minHeap)

        if delay > delays[node]:
            continue

        for child in adjList[node]:
            newDelay = delay + adjList[node][child]
            if newDelay < delays[child]:
                delays[child] = newDelay
                heapq.heappush(minHeap, (newDelay, child))

    maxDelay = max(delays[1:])  # Exclude delays[0] as it's not used

    return maxDelay if maxDelay != math.inf else -1