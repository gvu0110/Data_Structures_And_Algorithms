# Sorting

Elementary sorts:
- Bubble sort
- Insertion sort
- Selection sort

More efficient sorts:
- Merge sort
- Quick sort

### Bubble sort
- Bubbling up the largest value using multiple passthrough
- Mostly used for education purposes, not used in real life
- Time complexity: O(n^2)
- Space complexity: O(1)

### Selection sort
- Scanning the list of items for the smallest element and then swapping that element for the one in the first position
- Mostly used for education purposes, not used in real life
- Time complexity: O(n^2)
- Space complexity: O(1)

### Insertion sort
- Building the final sorted array one item at a time
- It is so useful for times when you're pretty sure the list is almost sorted (or it's already sorted)
- The best case scenario: O(n) when the list is almost sorted

### Merge sort
- Divide and conquer
- Divides the input list into two halves, calls itself for the two halves, and then merges the two sorted halves
- Time complexity: O(n log(n))
- Space complexity: O(n)

### Stable vs unstable algorithms
- A sorting algorithm is said to be stable if two objects with equal keys appear in the same order in sorted output as they appear in the input
- Stable: Insertion sort, Merge sort, Bubble sort
- Unstable: Selection sort, Heap sort, Quick sort

### Quick sort
- Divide and conquer
- Uses pivoting technique to break the input list into smaller lists, and the smaller lists use the pivoting technique until they are sorted
- Once we pick the pivot, all members that are less than the pivot come to its left, and all members that are greater than the pivot come to its right
- Quick sort usually the fastest on average, but it has worst case O(n^2) if you can't guarantee the pivot is good

### Radix sort + Counting sort

### Sorting interview
```
//#1 - Sort 10 schools around your house by distance:
Small input => Insertion

//#2 - eBay sorts listings by the current Bid amount:
Bid is mostly in specific range (probably $1-100) => radix + counting

//#3 - Sport scores on ESPN
Quick sort to save the memory space

//#4 - Massive database (can't fit all into memory) needs to sort through past year's user data
Avoid the worst case of quick sort and don't need to worry about the memory => merge sort

//#5 - Almost sorted Udemy review data needs to update and add 2 new reviews
Insertion

//#6 - Temperature Records for the past 50 years in Canada
If the temperature is integer => radix + counting
Else => quick sort

//#7 - Large user name database needs to be sorted. Data is very random.
Merge

//#8 - You want to teach sorting for the first time
Bubble sort + Selection sort
```

### What sorting algorithms in languages

### 