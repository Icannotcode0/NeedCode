package _2_add_two_numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	ans := &ListNode{Val: 0}
	dummyHead := ans
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		num1 := 0
		num2 := 0
		sum := 0
		if l1 != nil {
			num1 = l1.Val
		}
		if l2 != nil {
			num2 = l2.Val
		}

		sum = num1 + num2 + carry
		digit := sum % 10

		carry = sum / 10
		newNode := &ListNode{
			Val:  digit,
			Next: nil,
		}

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
