# Luck Balance
**Easy**

https://www.hackerrank.com/challenges/luck-balance

Lena is preparing for an important coding competition that is preceded by a number of sequential preliminary contests. Initially, her luck balance is 0. She believes in "saving luck", and wants to check her theory. Each contest is described by two integers, $L[i]$ and $T[i]$:

- $L[i]$ is the amount of luck associated with a contest. If Lena *wins* the contest, her luck balance will *decrease* by $L[i]$; if she *loses* it, her luck balance will *increase* by $L[i]$.
- $T[i]$ denotes the contest's *importance rating*. It's equal to $1$ if the contest is *important*, and it's equal to $0$ if it's *unimportant*.

If Lena loses no more than $k$ *important* contests, what is the maximum amount of luck she can have after competing in all the preliminary contests? This value *may* be negative.

**Example**  
$k=2$   
$L = [5, 1, 4]$   
$T = [1, 1, 0]$  


```
Contest		L[i]	T[i]
1		5	1
2		1	1
3		4	0
```

If Lena loses all of the contests, her will be $5+1+4=10$.  Since she is allowed to lose $2$ important contests, and there are only $2$ important contests, she can lose all three contests to maximize her luck at $10$.  

If $k=1$, she has to win at least $1$ of the $2$ important contests.  She would choose to win the lowest value important contest worth $1$.  Her final luck will be $5+4-1=8$.  

**Function Description**  

Complete the *luckBalance* function in the editor below.  

luckBalance has the following parameter(s):  

- *int k*: the number of important contests Lena can lose  
- *int contests[n][2]:* a 2D array of integers where each $contests[i]$ contains two integers that represent the luck balance and importance of the $i^{th}$ contest   

**Returns**   

- *int:*  the maximum luck balance achievable

**Input Format**

The first line contains two space-separated integers $n$ and $k$, the number of preliminary contests and the maximum number of important contests Lena can lose.  
Each of the next $n$ lines contains two space-separated integers, $L[i]$ and $T[i]$, the contest's luck balance and its importance rating.

**Constraints**

- $1 \le n \le 100$
- $0 \le k \le N$
- $1 \le L[i] \le 10^4$
- $T[i] \in \{0,1\}$
