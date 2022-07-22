package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						},
					},
				},
			},
		},
	}

	result := partition(head, 3)
	for result != nil {
		println(result.Val)
		result = result.Next
	}
}

func partition(head *ListNode, x int) *ListNode {
	left, right := &ListNode{}, &ListNode{}

	small, great, point := left, right, head
	left.Next = head

	for point != nil {
		if point.Val < x {
			small.Next = point
			small = small.Next
		} else {
			great.Next = point
			great = great.Next
		}
		point = point.Next
	}
	great.Next = nil
	small.Next = right.Next

	return left.Next
}
