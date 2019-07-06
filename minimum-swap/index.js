export function minSwaps(arr) {
  let count = 0
  for (let i = 0;i < arr.length;i++) {
    const val = arr[i]
    if (val === i + 1) continue
    for (let j = i + 1; j < arr.length; j++) {
      if (arr[j] !== i+1) continue
      arr[i] = arr[j]
      arr[j] = val
      count ++
    }
  }
  return count
}