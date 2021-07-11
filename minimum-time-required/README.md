## Minimum Time Required

https://www.hackerrank.com/challenges/minimum-time-required/  
**Medium**  
You are planning production for an order. You have a number of machines that each have a fixed number of days to produce an item. Given that all the machines operate simultaneously, determine the minimum number of days to produce the required order.

For example, you have to produce `goal = 10` items. You have three machines that take `machines = [2, 3, 2]` days to produce an item. The following is a schedule of items produced:

```
Day Production  Count
2   2               2
3   1               3
4   2               5
6   3               8
8   2              10
```

It takes 8 days to produce 10 items using these machines.
  
**Function Description**  
Complete the minimumTime function in the editor below. It should return an integer representing the minimum number of days required to complete the order.

minimumTime has the following parameter(s):

- machines: an array of integers representing days to produce one item per machine
- goal: an integer, the number of items required to complete the order

**Input Format** 

The first line consist of two integers `n` and `goal`, the size of `machines` and the target production.
The next line `n` contains  space-separated integers,`machines[i]` .

**Output Format**

Return the minimum time required to produce `goal` items considering all machines work simultaneously.