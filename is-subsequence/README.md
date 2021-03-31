# Is Subsequence
__Easy__  


Given two strings s and t, check if s is a subsequence of t.

A __subsequence__ of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

 

__Example 1:__
```
Input: s = "abc", t = "ahbgdc"
Output: true
```  
__Example 2:__
```
Input: s = "axc", t = "ahbgdc"
Output: false
```  
 
__Constraints:__
```
0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters.
```

Follow up: If there are lots of incoming s, say s1, s2, ..., sk where k >= 109, and you want to check one by one to see if t has its subsequence. In this scenario, how would you change your code?