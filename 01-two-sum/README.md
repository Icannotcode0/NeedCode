# Two Sum

LeetCode #1 — https://leetcode.com/problems/two-sum/

## Problem

Given an array of integers `nums` and an integer `target`, return the indices
of the two numbers that add up to `target`. Each input has exactly one
solution, and the same element can't be used twice.

```
nums = [2, 7, 11, 15], target = 9
=> [0, 1]   (2 + 7 == 9)
```

## Approaches

### 1. Brute force — O(n²) time, O(1) space

Check every pair `(i, j)` and test whether `nums[i] + nums[j] == target`.
Correct but wasteful: for each element you re-scan the rest of the array
looking for its complement.

### 2. Hash table — O(n) time, O(n) space

The key insight: "does a complement exist" is a *lookup* question, and a hash
map turns lookup from O(n) (linear scan) into O(1) (amortized). Trade space
for time — store what you've seen so you never have to re-scan for it.

There are two common ways to use the map:

- **One-pass**: walk the array once. At index `i`, check if
  `target - nums[i]` is already in the map; if not, insert `nums[i] -> i` and
  continue. This finds a match as soon as the *second* element of a valid
  pair is reached, and never needs a second loop.
- **Two-pass**: first build the map for the entire array (`value -> index`),
  then walk the array again checking each element's complement against the
  now-complete map.

## This implementation

```go
func twoSum(nums []int, target int) []int {
    recorder := make(map[int]int)
    for i, num := range nums {
        recorder[num] = i
    }
    ans := []int{}

    for i := range nums {
        complement := target - nums[i]
        val, ok := recorder[complement]
        if !ok || val == i {
            continue
        }
        ans = append(ans, []int{i, val}...)
        break
    }
    return ans
}
```

This is the **two-pass** variant:

1. First loop builds `recorder`, mapping each value to the index it last
   occurred at.
2. Second loop looks up each element's complement in the completed map.

### The comma-ok idiom

```go
val, ok := recorder[complement]
```

Indexing a Go map always returns two values when written this way: the
value (`val`) and a `bool` (`ok`) saying whether the key was actually
present. This matters because a missing key silently returns the map's zero
value (`0` for `int`) rather than erroring — without checking `ok`, a
missing complement of `0` would be indistinguishable from a genuine `0` in
`recorder`.

### Why `val == i` is checked

A map key is unique, so if a value appears more than once in `nums`,
`recorder[value]` only remembers the *last* index it was seen at, overwriting
earlier ones. The `val == i` check exists to reject a "complement" that is
actually the same element looked at twice (`nums[i]`'s only recorded index
being `i` itself) — e.g. `nums = [3, 2, 4]`, `target = 6`: the complement of
`3` is `3`, and `recorder[3] == 0 == i`, so it's correctly skipped rather
than answering `[0, 0]`.

For true duplicate pairs this still works: `nums = [3, 2, 3]`, `target = 6`
→ `recorder[3] = 2` (last occurrence). At `i = 0`, `complement = 3`,
`val = 2 != i`, so `[0, 2]` is returned correctly.

### Complexity

- **Time**: O(n) — two linear passes over `nums`; map insert/lookup is O(1)
  amortized.
- **Space**: O(n) — the map holds up to `n` entries.

### One-pass vs. two-pass trade-off

The one-pass version is strictly better here (single loop, can exit as soon
as a match is found, uses no more memory) and is the version most commonly
recommended. The two-pass version above is still O(n) overall and easier to
reason about (the map is always fully built before any lookup happens), but
it does unnecessary extra work in the common case, since the map gets fully
built even when a match could have been found and returned early during
construction.

## Go-specific notes

- `make(map[int]int)` allocates an empty, ready-to-use map (map zero value
  `nil` is *not* usable — you can't assign into a `nil` map).
- `[]int{}` initializes a non-nil, zero-length slice; `append` grows it as
  needed.
- `append(ans, []int{i, val}...)` splices the two-element slice into `ans`
  using the `...` spread; equivalent to
  `ans = append(ans, i, val)` here.
