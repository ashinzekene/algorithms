# Climbing the leaderboard

https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem


Solution

100 100 50 40 40 20 10
5 25 50 120

100 100 50 40 40 20 10
5 25 50 90

```
- lastIndex = -1
- position = 0
- looping alice from right to left
- looping scores from lastIndex to right
- lastindex is -1 and alice >= score at lastIndex+1,  position++, add position to result, add last index
- lastindex is not -1 and Score at lastIndex+1 == score at lastindex, add last index
- lastindex is not -1 and Score at lastIndex+1 != score at lastindex and alice >= score at lastindex+1, position++, add position to result, add last index
- lastindex is not -1 and Score at lastIndex+1 != score at lastindex and alice < score at lastindex+1, add last index
```

pos = 1
li = 0
res = []
for i = ali.length-1;i >=0;i++ {
  alice = a[i]
  for j = li; j < sc.length; j++ {
    if alice >= sc[j] {
      res[i] = pos
      break
      li = j
    } else if (li == len(sc)) {
      for k = i;k < 0; ++ {
        res[k] = pos+1
      }
    } else if (sc[li] > sc[li+1]) {
      pos++
    }
  }
}