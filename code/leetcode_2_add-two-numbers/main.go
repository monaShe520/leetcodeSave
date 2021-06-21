package main

import "fmt"

import "encoding/json"

type ListNode struct {
	Val int
	Next *ListNode
}
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.

// 输入：l1 = [0], l2 = [0]
// 输出：[0]

// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1List := []int{}
	l2List := []int{}
	for {
		if l1 == nil && l2 == nil {
			break
		} 
		if l2 == nil {
			l2List = append(l2List, 0)
		} else {
			l2List = append(l2List, l2.Val)
			l2 = l2.Next
		}
		if l1 == nil {
			l1List = append(l1List, 0)
		} else {
			l1List = append(l1List, l1.Val)
			l1 = l1.Next
		}
	}
	step := 0 // 是否进位标志
	ln := &ListNode{} // 最终返回的链表
	mark := ln // 最新node标志
	for index, l1Item := range l1List {
		var itemRes int
		if l1Item + l2List[index] + step >= 10 {
			itemRes = l1Item + l2List[index] + step - 10
			step = 1
		} else {
			itemRes =  l1Item + l2List[index] + step
			step = 0
		}
		mark.Val = itemRes
		mark.Next = &ListNode{}
		if index == len(l1List) -1 {
			if step == 1 {
				mark.Next = &ListNode{
					Val: 1,
				}
			} else {
				mark.Next = nil
			}
		}
		mark = mark.Next
	}
	if step == 1 {
		mark.Val = 1
		mark.Next = nil
	} else {
		mark = nil
	}
	return ln
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	res := addTwoNumbers(l1,l2)
	resM, _ := json.Marshal(res)
	fmt.Println(string(resM))
}