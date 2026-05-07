# Alternating Characters 
**Easy**

https://www.hackerrank.com/challenges/alternating-characters

You are given a string containing characters $A$ and $B$ only.  Your task is to change it into a string such that there are no matching adjacent characters. To do this, you are allowed to delete zero or more characters in the string.  

Your task is to find the minimum number of required deletions.

**Example**  
$s = AABAAB$  

Remove an $A$ at positions $0$ and $3$ to make $s = ABAB$ in $2$ deletions.

**Function Description**

Complete the *alternatingCharacters* function in the editor below.  

alternatingCharacters has the following parameter(s):

- *string s*: a string  

**Returns**  

- *int:* the minimum number of deletions required

**Input Format**

The first line contains an integer $q$, the number of queries.  
The next $q$ lines each contain a string $s$ to analyze.

**Constraints**

- $ 1 \le q \le 10$  
- $ 1 \le \text{ length of s }\le 10^5$
- Each string $s$ will consist only of characters $A$ and $B$.
