package main

import "strings"

import "fmt"


var numMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var res []string

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	numsString := []string{}
	for _, s := range strings.Split(digits, "") {
		numsString = append(numsString, numMap[s])
	}
	res = []string{}
	if len(numsString) == 1 {
		res = strings.Split(numsString[len(numsString) - 1], "")
		return res
	}
	recursion(numsString, 0, "")
	return res
}

func recursion(numsString []string, index int, end string) {
	if index == len(numsString) {
		res = append(res, end)
		return
	}
	for _, s := range strings.Split(numsString[index], "") {
		recursion(numsString, index + 1, end + s)
	}
}


func main() {
	fmt.Println(letterCombinations("234"))
}