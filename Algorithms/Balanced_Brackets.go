package main

import (
	"fmt"
)

/*
if len(s)%2 == 0 {
	brac := map[string]string{
		// "{": "{",
		// "(": "(",
		// "[": "[",
		"}": "{",
		")": "(",
		"]": "[",
	}
	j := len(s) - 1

	for i := 0; i < len(s)/2; i++ {
		l := string(s[i])
		r := string(s[j])
		fmt.Println(l, brac[r])
		if !(l == brac[r]) {
			return "NO"
		}
		j--
	}
	return "YES"
}
return "NO"
*/
func isBalanced(s string) string {
	var stack []string
	for _, v := range s {

		switch string(v) {
		case "{":
		case "[":
		case "(":
			stack = append(stack, string(v))
			break
		case "}":
			if stack[0] != "{" {
				stack = stack[1:]
				fmt.Println(stack)
				return "NO"
			}
			break
		case ")":
			if stack[0] != "(" {
				stack = stack[1:]
				return "NO"
			}
			break
		case "]":
			if stack[0] != "[" {
				stack = stack[1:]
				return "NO"
			}
			break
		default:
			return "YES"
		}
	}
	return "YES"
}

func main() {

	fmt.Println(isBalanced("{{[[(())]]}}"))
}
