# Picking Numbers
**Easy**

https://www.hackerrank.com/challenges/picking-numbers

Given an array of integers, find the longest subarray where the absolute difference between any two elements is less than or equal to $1$.  

**Example**  

$a=[1,1,2,2,4,4,5,5,5]$  

There are two subarrays meeting the criterion: $[1,1,2,2]$ and $[4,4,5,5,5]$.  The maximum length subarray has $5$ elements.

**Function Description**  

Complete the *pickingNumbers* function in the editor below.  

pickingNumbers has the following parameter(s):  

- *int a[n]:* an array of integers  

**Returns**  

- *int:*  the length of the longest subarray that meets the criterion

**Input Format**

The first line contains a single integer $n$, the size of the array $a$.   	
The second line contains $n$ space-separated integers, each an $a[i]$.

**Constraints**

- $2 \le n \le 100$
- $0 < a[i] < 100$
- The answer will be $\ge 2$.
