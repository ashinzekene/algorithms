# Priyanka and Toys
https://www.hackerrank.com/challenges/priyanka-and-toys/problem
**Easy**

Priyanka works for an international toy company that ships by container. Her task is to the determine the lowest cost way to combine her orders for shipping. She has a list of item weights. The shipping company has a requirement that all items loaded in a container must weigh less than or equal to 4 units plus the weight of the minimum weight item. All items meeting that requirement will be shipped in one container.

What is the smallest number of containers that can be contracted to ship the items based on the given list of weights?

For example, there are items with weights `w = [1, 2, 3, 4, 5, 10, 11, 12, 13]`. This can be broken into two containers: `[1, 2, 3, 4, 5]` and `[10, 11, 12 ,13]`. Each container will contain items weighing within `4` units of the minimum weight item

**Function Description**  

Complete the toys function in the editor below. It should return the minimum number of containers required to ship.

toys has the following parameter(s):

w: an array of integers that represent the weights of each order to ship


**Sample Input:**  
```
8
1 2 3 21 7 12 14 21
```
**Sample Output:**  
```
4
```