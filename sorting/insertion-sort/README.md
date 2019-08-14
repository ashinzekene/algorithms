# Insertion sort  

## Time Complexity  
O(N<sup>2</sup>)

## Instructions
Starting from the second element, the element checks from the beginning of the array checking if any element is greater than it and them swapping it if it is.
Example: `[2, 5, 3, 1, 4, 8, 7]`
1. Compare `5` and `2`. Is 2 is smaller, so continue.
`[2, 5, 3, 1, 4, 8, 7]`
2. Compare `3` and `5`, `5` isn't smaller, so swap `3` with `5`, stop there since `2` is less than `3`
`[2, 3, 5, 1, 4, 8, 7]`
2. Compare `1` and `5`, `5` isn't smaller, so swap `1` with `5`, swap, 3 and 2 too
`[1, 2, 3, 5, 4, 8, 7]`
2. Compare `4` and `5`, `5` isn't smaller, so swap `4` with `5`, stop there since `3` is less than `4`
`[1, 2, 3, 4, 5, 8, 7]`
2. Compare `8` and `5`, `5`. 8 is not smaller so leave
`[1, 2, 3, 4, 5, 8, 7]`
2. Compare `7` and `8`, `5`. 7 is smaller so swap with 8, stop there since `5` is less than `7`
`[1, 2, 3, 4, 5, 8, 7]`
