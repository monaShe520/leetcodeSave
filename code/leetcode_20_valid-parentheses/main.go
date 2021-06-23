package main

import "fmt"

import "strings"

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 示例 1：
// 输入：s = "()"
// 输出：true
// 示例 2：
// 输入：s = "()[]{}"
// 输出：true
// 示例 3：
// 输入：s = "(]"
// 输出：false
// 示例 4：
// 输入：s = "([)]"
// 输出：false
// 示例 5：
// 输入：s = "{[]}"
// 输出：true
// 提示：
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成
type Stack struct {
	list	[]int
	num		int
}

func (s *Stack) len() int {
	return s.num
}

func (s *Stack) top() int {
	return s.list[s.len()-1]
}

func (s *Stack) push(a int) {
	s.list = append(s.list, a)
	s.num = s.num + 1
}

func (s *Stack) pop() int {
	end := s.list[s.len()-1]
	s.list = s.list[:s.len()-1]
	s.num = s.num - 1
	return end
}

func isValid(s string) bool {
	
	slist := strings.Split(s, "")
	intlist := []int{}
	for _, r := range slist {
		switch string(r) {
			case "{":
				intlist = append(intlist, -1)
			case "}":
				intlist = append(intlist, 1)
			case "[":
				intlist = append(intlist, -2)
			case "]":
				intlist = append(intlist, 2)
			case "(":
				intlist = append(intlist, -3)
			case ")":
				intlist = append(intlist, 3)
		}

	}
	stack := &Stack{}
	stack.push(intlist[0])
	for i:=1;i<len(intlist); i++ {
		if stack.len() == 0 {
			stack.push(intlist[i])
			continue
		} else {
			if stack.top() < 0 &&stack.top() + intlist[i] == 0 {
				stack.pop()
			} else {
				stack.push(intlist[i])
			}
		}
	}
	return stack.len() == 0
}

func main() {
	fmt.Println(isValid("{[]}"))
}