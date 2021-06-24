package main

import "fmt"

import "encoding/json"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	headC := head
	index := []*ListNode{}
	for {
		if headC != nil {
			index = append(index, headC)
			headC = headC.Next
		} else {
			break
		}
	}
	if n == 1 {
		index[len(index) - n - 1].Next = nil
		return head
	}
	if n == len(index) {
		return head.Next
	}
	index[len(index) - n - 1].Next = index[len(index) - n +1]
	return head
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	first := 0
	firstHead := head
	lastHead := head
	last := -n
	for {
		if last >= 0 {
			lastHead = lastHead.Next
		}
		last = last + 1
		first = first + 1
		firstHead = firstHead.Next
		if firstHead.Next == nil {
			if last < 0 {
				return head.Next
			}
			lastHead.Next = lastHead.Next.Next
			break
		}
	}
	return head

}
func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	
	res := removeNthFromEnd(l1, 2)
	resM, _ := json.Marshal(res)
	fmt.Println(string(resM))
}