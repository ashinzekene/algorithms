## Swapping Nodes in a Linked List

https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
__Medium__  

You are given the head of a linked list, and an integer k.

Return the head of the linked list after swapping the values of the kth node from the beginning and the kth node from the end (the list is 1-indexed).


__Example 1:__  
```
Input: head = [1,2,3,4,5], k = 2
Output: [1,4,3,2,5]
```

__Example 2:__  
```
Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
Output: [7,9,6,6,8,7,3,0,9,5]
```

__Example 3:__  
```
Input: head = [1], k = 1
Output: [1]
```

__Example 4:__  
```
Input: head = [1,2], k = 1
Output: [2,1]
```

__Example 5:__  
```
Input: head = [1,2,3], k = 2
Output: [1,2,3]
``` 

__Constraints:__  
```
The number of nodes in the list is n.
1 <= k <= n <= 105
0 <= Node.val <= 100
```
