package main

import "fmt"
import "encoding/json"

type ListNode struct {
	Val int
	Next *ListNode
}
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	index1 := l1
	index2 := l2
	res := &ListNode{}
	index := res
	for{
		if index1 == nil && index2 == nil {
			index = nil
			break
		}
		if index1 == nil {
			index.Val = index2.Val
			index.Next = index2.Next
			break
		}
		if index2 == nil {
			index.Val = index1.Val
			index.Next = index1.Next
			break
		}
		if index1.Val <= index2.Val {
			index.Val = index1.Val
			index.Next = &ListNode{}
			index = index.Next
			index1 = index1.Next
		} else {
			index.Val = index2.Val
			index.Next = &ListNode{}
			index = index.Next
			index2 = index2.Next
		}

	}
	if l1 == nil && l2 == nil {
		return nil
	}
	return res
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	res := mergeTwoLists(l1,l2)
	resM, _ := json.Marshal(res)
	fmt.Println(string(resM))
}