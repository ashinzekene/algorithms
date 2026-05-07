# Equalize the Array
**Easy**

https://www.hackerrank.com/challenges/equality-in-a-array

Given an array of integers, determine the minimum number of elements to delete to leave only elements of equal value.
  
**Example**  

$arr = [1, 2, 2, 3]$  

Delete the $2$ elements $1$ and $3$ leaving $arr = [2, 2]$. If both twos plus either the $1$ or the $3$ are deleted, it takes $3$ deletions to leave either $[3]$ or $[1]$.  The minimum number of deletions is $2$.

**Function Description**  

Complete the *equalizeArray* function in the editor below.  

equalizeArray has the following parameter(s):  

- *int arr[n]:* an array of integers   

**Returns**  

- *int:* the minimum number of deletions required

**Input Format**

The first line contains an integer $n$, the number of elements in $arr$.  		
The next line contains $n$ space-separated integers $arr[i]$.

**Constraints**

- $1 \leq n \leq 100$
- $1 \le arr[i] \le 100$
