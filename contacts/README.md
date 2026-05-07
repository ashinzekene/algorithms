# Contacts
**Medium**

https://www.hackerrank.com/challenges/contacts

We're going to make our own *Contacts* application! The application must perform two types of operations:

1. `add name`, where $name$ is a string denoting a contact name. This must store $name$ as a new contact in the application.  
2. `find partial`, where $partial$ is a string denoting a partial name to search the application for. It must count the number of contacts starting with $partial$ and print the count on a new line.

Given $n$ sequential *add* and *find* operations, perform each operation in order.

**Example**   
Operations are requested as follows:

<pre>
add ed
add eddie
add edward
find ed
add edwina
find edw
find a
</pre>

Three $add$ operations include the names 'ed', 'eddie', and 'edward'.  Next, $\texttt{find ed}$ matches all $3$ names.  Note that it matches and counts the entire name 'ed'.  Next, add 'edwina' and then find 'edw'.  $2$ names match: 'edward' and 'edwina'.  In the final operation, there are $0$ names that start with 'a'.  Return the array $[3, 2, 0]$.  

**Function Description**   

Complete the *contacts* function below.  

*contacts* has the following parameters:

- *string queries[n]:* the operations to perform   

**Returns**  

- *int[]:* the results of each find operation

**Input Format**

The first line contains a single integer, $n$, the number of operations to perform (the size of $queries[]$).  
Each of the following $n$ lines contains a string, $queries[i]$.

**Constraints**

- $1 \le n \le 10^5$  
- $1 \le \texttt{length of }name \le 21$  
- $1 \le \texttt{length of }partial \le 21$  
- $name$ and $partial$ contain lowercase English letters only.
- The input does not have any duplicate $name$ for the $add$ operation.
