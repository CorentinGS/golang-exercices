package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	point := head
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		point.Next = &ListNode{Val: carry % 10}
		carry /= 10
		point = point.Next
	}
	return head.Next
}