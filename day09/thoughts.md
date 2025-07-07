# My thought process 09

## Problem

Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.
For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.
Follow-up: Can you do this in O(N) time and constant space?

## Thoughts

- use brute force (sadly couldn't find a way to actual do it, later learned i needed to figure out a way to get all possible valid subsets)
- use index parity strategy
- Final solution using dynamic programming and Sliding State Variables

## Index Parity Strategy

### Idea

- To avoid adjacent elements, consider only elements at either even indices or odd indices.
- Compute the sum of elements at even indices.
- Compute the sum of elements at odd indices.
- Return the larger of the two sums.

### Example 1: [2, 4, 6, 2, 5]

- Even indices: 0, 2, 4 → 2 + 6 + 5 = 13  
- Odd indices: 1, 3 → 4 + 2 = 6  
- Result: 13 (correct)

### Example 2: [5, 1, 1, 5]

- Even indices: 0, 2 → 5 + 1 = 6  
- Odd indices: 1, 3 → 1 + 5 = 6  
- Optimal solution: 5 + 5 = 10 (indices 0 and 3)  
- Result: 6 (incorrect)

### Strengths

- Ensures non-adjacency by construction.
- Simple and fast (O(n) time, O(1) space).

### Limitations

- Not guaranteed to find the optimal sum.
- Misses valid combinations that mix even and odd indices.
- Too rigid for general cases.

## Dynamic Programming

### Intuition

At each element, we face two choices:

- **Include** it → skip the next one
- **Exclude** it → move to the next

We use two variables to track the best possible sums:

- `prev1`: best sum so far (excluding current)
- `prev2`: best sum from two steps ahead (used when including current)

We iterate from the end of the list to the beginning. At each step:

```go
current = max(data[i] + prev2, prev1)
```

- `data[i] + prev2`: include current, add non-adjacent sum
- `prev1`: skip current

Then we update:

```go
prev2 = prev1
prev1 = current
```

At the end, `prev1` holds the final result.

### Complexity

Time: O(n) — one pass through the slice

Space: O(1) — only two extra variables
