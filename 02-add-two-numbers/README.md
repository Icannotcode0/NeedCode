# Add Two Numbers

LeetCode #2 — https://leetcode.com/problems/add-two-numbers/

## Problem

Two non-negative integers are given as linked lists, where each node holds a
single digit stored in **reverse order** (least significant digit first).
Add the two numbers and return the sum as a linked list in the same format.

```
l1 = 2 -> 4 -> 3   (represents 342)
l2 = 5 -> 6 -> 4   (represents 465)
sum = 342 + 465 = 807
=> 7 -> 0 -> 8
```

The reverse-order storage isn't arbitrary — it's what makes the problem
tractable without extra work: digit 0 of both lists is already the *ones*
place, so you can add the lists left-to-right exactly like you'd add numbers
right-to-left by hand, carry and all, without ever reversing anything.

## Approach — simulate grade-school addition, O(max(n, m)) time

Walk both lists at the same time, one node per list per step:

1. At each step, take `l1`'s digit (or `0` if that list has run out) and
   `l2`'s digit (or `0`), plus the `carry` from the previous step.
2. `sum = num1 + num2 + carry`. The digit to emit is `sum % 10`; the new
   carry is `sum / 10` (0 or 1, since two digits plus a carry never exceeds
   19).
3. Append a new node holding that digit to the result list.
4. Advance whichever of `l1`/`l2` still has nodes.
5. Stop when both lists are exhausted **and** there's no carry left over
   (a trailing carry, e.g. `5 + 5 = 10`, needs one more digit than either
   input list has).

This is a linear merge, structurally the same shape as merging two sorted
lists or a ripple-carry adder in hardware — the "carry" is the whole reason
the two lists can't just be summed independently node-by-node.

## This implementation

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    ans := &ListNode{Val: 0}
    dummyHead := ans
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        num1, num2, sum := 0, 0, 0
        if l1 != nil {
            num1 = l1.Val
        }
        if l2 != nil {
            num2 = l2.Val
        }

        sum = num1 + num2 + carry
        digit := sum % 10
        carry = sum / 10

        newNode := &ListNode{Val: digit, Next: nil}
        dummyHead.Next = newNode
        dummyHead = dummyHead.Next

        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }

    return ans.Next
}
```

### The dummy head technique

```go
ans := &ListNode{Val: 0}
dummyHead := ans
...
return ans.Next
```

Building a linked list node-by-node normally needs special-case code for
"the first node," because there's no previous node to attach it to yet.
The dummy-head trick sidesteps that: allocate one throwaway node (`ans`)
up front with a value that's never read, and always append via
`dummyHead.Next = newNode`. Every append — including the first — becomes
the same operation. `ans` itself stays pinned to the throwaway node the
whole time so that `ans.Next` at the end still points at the *real* first
node of the result; `dummyHead` is the one that walks forward.

### Loop condition: `l1 != nil || l2 != nil || carry != 0`

Three independent stopping conditions ORed together — the loop keeps going
as long as *any* of them still has work to contribute. This is what makes
the trailing-carry case correct: after both lists are exhausted, if
`carry == 1`, one more iteration runs with `num1 = num2 = 0`, producing a
final `1` digit (e.g. `[5] + [5]` → carry 1 after the first digit → loop
runs again → emits node `1` → result `[0, 1]` = 10).

### Complexity

- **Time**: O(max(n, m)) where `n`, `m` are the lengths of `l1`, `l2` — each
  node from each list is visited exactly once, plus at most one extra step
  for a final carry.
- **Space**: O(max(n, m) + 1) for the newly allocated result list (not
  counting the input lists, which are read but not mutated).

## Go-specific notes

- `&ListNode{Val: 0}` allocates a struct on the heap and yields a pointer to
  it in one expression — Go's escape analysis takes care of promoting it
  off the stack automatically since a pointer to it outlives the function.
- `l1 != nil` guards every field access (`l1.Val`, `l1.Next`) because Go
  will panic on a nil-pointer dereference; unlike a hash map's zero-value
  fallback, there's no silent default here — the check is mandatory, not a
  style choice.
- `dummyHead` and `ans` are two separate variables holding the *same*
  pointer value initially, then diverge as `dummyHead` is reassigned
  (`dummyHead = dummyHead.Next`). `ans` is never reassigned, which is what
  lets it still refer to the dummy node at the end for `ans.Next`.
