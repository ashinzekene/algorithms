# Bubble Sort

## Time Complexity  
O(N<sup>2</sup>)

## How it works

Bubble sort compares two values swapping them if the one on the left is smaller beginning from the first item in the array to the last item. For example. Given the array `[2, 5, 3, 1, 4, 8, 7]`, bubble sort works thus:
1. Compares `2` and `5`, no sorting occurs
`[2, 5, 3, 1, 4, 8, 7]`
2. Compares `5` and `3`, they will be swapped, hence we have
`[2, 3, 5, 1, 4, 8, 7]`
2. Compares `5` and `1`, they will be swapped, hence we have
`[2, 3, 1, 5, 4, 8, 7]`
2. Next is `5` and `4`, they will be swapped, hence we have
`[2, 3, 1, 4, 5, 8, 7]`
2. Next is `5` and `8`, they are already sorted
`[2, 3, 1, 4, 5, 8, 7]`
2. Next is `8` and `7`, they will be swapped, hence we have
`[2, 3, 1, 4, 5, 7, 8]`

This continues till the end of the array is reached, then this process is repeated till the whole array is sorted