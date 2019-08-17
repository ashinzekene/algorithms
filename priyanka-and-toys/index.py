def toys(arr):
  toysArr = sorted(arr)
  mins = []
  for x in toysArr:
    if len(mins) == 0:
      mins.append(x)
    else:
      if x > mins[-1] + 4:
        mins.append(x)
  return len(mins)