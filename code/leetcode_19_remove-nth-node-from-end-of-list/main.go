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

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	
	res := removeNthFromEnd(l1, 2)
	resM, _ := json.Marshal(res)
	fmt.Println(string(resM))
}