def minSwaps(arr):
  count = 0
  for i in range(len(arr)):
    val = arr[i]
    if val == i + 1:
        continue
    for j in range(i + 1, len(arr)):
      if arr[j] != i+1:
          continue
      arr[i] = arr[j]
      arr[j] = val
      count += 1
  return count
