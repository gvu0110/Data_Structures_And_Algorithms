# Arrays

### Static vs dynamic array
1. Static array
- Array
```
a := [5]int{10,20,30,40,50}
```
2. Dynamic array
- Slice
- Append
growslice() essentially carries out the following actions:
- Calculate the new cap of the slice after growing:
    + For slices with len <1024, it will double cap;
    + For slices with len >= 1024, new cap will be (1.25 * cap);
- Determine the memory size of the new slice.
- Allocate memory for the new slice (and clear some memory).
- Copy the data from the old memory address to the new memory address.

### Implementing an array

#### Two sum
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to target. You may assume that each input would have exactly one solution, and you may not use the same element twice. You can return the answer in any order.
```go
```
#### Maximum subarray
Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
```go
```
#### Move zeroes
```go
```
#### Contains duplicate
```go
```
#### Rotate array
```go
```

### Array reviews
Pros:
- Fast lookups via index
- Fast push/pop
- Ordered

Cons:
- Slow inserts
- Slow deletes
- If using static array, it's fixed size