# Permutation Sequence
**Medium**

The set `[1,2,3,...,n]` contains a total of n! unique permutations.

By listing and labeling all of the permutations in order, we get the following sequence for n = 3:
```
"123"
"132"
"213"
"231"
"312"
"321"
```
Given n and k, return the kth permutation sequence.
**Note:**
- Given n will be between 1 and 9 inclusive.
- Given k will be between 1 and n! inclusive.
**Example 1:**
```
Input: n = 3, k = 3
Output: "213"
```
**Example 2:**
```
Input: n = 4, k = 9
Output: "2314"
```
3 * 2 * 1
4 * 3 * 2 * 1

0000 === 1
0010 === 2
0100 === 3
0110 === 4
0200 === 5
0210 === 6
1000 === 7
1010 === 8
1100 === 9
2000 === 13
3000 === 19

17

2110

1234
1243
1324
1342
1423
1432
2134
2143
2314 ==== 9


1
2
6

123
213

001000