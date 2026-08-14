# NeedCode

My daily LeetCode practice, solved in Go.

Each problem lives in its own numbered directory containing the solution
and a `README.md` write-up of the algorithm, the approach behind it, and
the Go-specific details the solution relies on.

## Problems

| # | Problem | Solution | Notes |
|---|---------|----------|-------|
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | [`01-two-sum/2sum.go`](01-two-sum/2sum.go) | [README](01-two-sum/README.md) |
| 2 | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | [`02-add-two-numbers/add-two-numbers.go`](02-add-two-numbers/add-two-numbers.go) | [README](02-add-two-numbers/README.md) |
| 3 | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | [`03-longest-substring-without-repeating-characters/03.go`](03-longest-substring-without-repeating-characters/03.go) | [README](03-longest-substring-without-repeating-characters/README.md) |

## Layout

```
NN-problem-slug/
├── <solution>.go   # package <problem_slug>, one exported solving function
└── README.md       # problem statement, approach, and a walkthrough of this solution
```
