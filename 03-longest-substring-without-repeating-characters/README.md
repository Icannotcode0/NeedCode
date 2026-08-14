# Longest Substring Without Repeating Characters

LeetCode #3 — https://leetcode.com/problems/longest-substring-without-repeating-characters/

## Problem

Given a string `s`, find the length of the longest substring that contains
no repeated characters.

```
s = "abcabcbb" => 3   ("abc")
s = "bbbbb"     => 1   ("b")
s = "pwwkew"    => 3   ("wke")
```

## Approach — sliding window, O(n) time

Brute force checks every substring for uniqueness: O(n²) or O(n³) depending
on how the uniqueness check is done. The key observation that gets this down
to O(n): once a window `s[left:right]` is duplicate-free and you extend it
by one character on the right and find a repeat, you never need to move
`left` back — only forward, and only as far as just past the *previous*
occurrence of the character that caused the collision. `left` and `right`
each only ever move forward, so the whole scan is a single linear pass with
two pointers, not a pointer restarted from scratch for every window.

A hash map tracks how many times each character currently appears inside
the window `[left, right]`. As long as every count is `≤ 1`, the window is
valid; the answer is the max window width seen.

## This implementation

```go
func lengthOfLongestSubstring(s string) int {
    length := 0
    left := 0
    recorder := make(map[string]int)
    for right := range s {
        rightString := string(s[right])
        recorder[rightString]++

        for recorder[rightString] > 1 {
            recorder[string(s[left])]--
            left++
        }
        length = max(length, right-left+1)
        right++
    }
    return length
}
```

1. `right` scans forward one character at a time; that character's count in
   `recorder` is incremented.
2. If incrementing pushed *that specific character's* count above 1, the
   window now contains a duplicate — shrink from `left` (decrementing counts
   as characters leave the window) until the count drops back to 1. Because
   only the character just added could possibly be the duplicate, checking
   `recorder[rightString] > 1` alone is enough; no other count needs
   checking.
3. `length` is updated to the widest valid window seen so far.

### `right++` at the end of the loop body does nothing

```go
for right := range s {
    ...
    right++
}
```

`right` here is the loop variable of a `for ... range` clause, not a
C-style `for i := 0; i < n; i++`. On each iteration Go's `range` clause
itself computes the next value of `right` (the next byte index into `s`)
and overwrites whatever the body did to it. The `right++` at the bottom is
dead code — it mutates a value that gets discarded before it's ever read
again. This is a common trap coming from C-style loops: in a `range` loop,
reassigning the loop variable inside the body never affects iteration, only
`break`/`continue`/the range source itself does.

### Ranging over a `string` walks *bytes of runes*, not bytes

`for right := range s` iterates once per *rune* (Unicode code point) in
`s`, with `right` set to the **byte offset** where that rune starts — not
a plain 0,1,2,... byte counter. For this implementation, that distinction
is silently papered over by `string(s[right])`, which takes a single
*byte* at that offset and wraps it back into a string. This is only
correct because the problem's test inputs are ASCII, where every rune is
exactly one byte. For a multi-byte UTF-8 character (e.g. "café" or any
CJK/emoji input), `right` would jump by 2–4 on that step, and
`string(s[right])` would grab one raw byte of a multi-byte rune rather
than the whole character — silently corrupting both the map key and the
`right-left+1` length math. Worth knowing as a boundary of this solution
rather than a bug to fix blindly, since LeetCode's judge only exercises
ASCII/basic multilingual plane inputs for this problem.

### Complexity

- **Time**: O(n) — `right` advances once per character; `left` advances at
  most once per character across the whole run (never resets backward), so
  the two pointers together do O(n) total work, not O(n) per outer step.
- **Space**: O(min(n, Σ)) where Σ is the size of the character set — the map
  holds at most one entry per distinct character currently in the window.

## Go-specific notes

- `recorder[rightString]++` relies on the map zero-value rule: reading an
  absent key returns `0` (the zero value for `int`) rather than erroring,
  so the first occurrence of a character can be incremented without a
  separate "insert if missing" step.
- Map entries are only ever decremented, never deleted with `delete()`, once
  a character's count returns to `0`. That's fine here — a `0` entry and a
  never-present key behave identically for both `recorder[x]++` and
  `recorder[x] > 1` — just something to notice if the map's size vs. window
  size were ever compared directly.
- `max` is hand-written here (`func max(a, b int) int`) rather than using
  Go's builtin `max` — the builtin `max`/`min` were only added in Go 1.21,
  so this shadowing pattern is common in code targeting older toolchains or
  simply predates the builtin's adoption.
