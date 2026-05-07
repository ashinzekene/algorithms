# The Love-Letter Mystery
**Easy**

https://www.hackerrank.com/challenges/the-love-letter-mystery

James found a love letter that his friend Harry has written to his girlfriend. James is a prankster, so he decides to meddle with the letter. He changes all the words in the letter into [palindromes](https://en.wikipedia.org/wiki/Palindrome).   

To do this, he follows two rules:  

1. He can only reduce the value of a letter by $1$, i.e. he can change _d_ to _c_, but he cannot change _c_ to _d_ or _d_ to _b_.  
2. The letter $a$ may not be reduced any further.  

Each reduction in the value of any letter is counted as a single operation. Find the minimum number of operations required to convert a given string into a palindrome.

**Example**   
$s = \texttt{cde}$   

The following two operations are performed:  _cd<strong>e</strong>_ &rarr; _cd<strong>d</strong>_ &rarr; _cdc_.  Return $2$.

**Function Description**  

Complete the *theLoveLetterMystery* function in the editor below.  

theLoveLetterMystery has the following parameter(s):  

- *string s*: the text of the letter   

**Returns**   

- *int:* the minimum number of operations

**Input Format**

The first line contains an integer $q$, the number of queries.  
The next $q$ lines will each contain a string $s$.

**Constraints**

$1 \le q \le 10$  
$1 \le$ | s | $\le 10^4$  
All strings are composed of lower case English letters, *ascii[a-z]*, with no spaces.
