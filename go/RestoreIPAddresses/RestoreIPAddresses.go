/*
93. Restore IP Addresses

Given a string containing only digits, restore it by returning all possible valid IP address combinations.

For example:
Given "25525511135",

return ["255.255.11.135", "255.255.111.35"]. (Order does not matter)
*/
package main

import (
	"fmt"
	"strconv"
)

func restoreIpPart(s string, part int) []string {
	if part < 1 || len(s) == 0 {
		return []string{}
	}

	solution := []string{}
	if s[0] == '0' {
		if part == 1 {
			if len(s) > 1 {
				return []string{}
			}
			return []string{"0"}
		} else {
			for _, sub := range restoreIpPart(s[1:], part-1) {
				solution = append(solution, "0."+sub)
			}
		}
	} else {
		if part == 1 {
			if d, _ := strconv.Atoi(s); d <= 255 {
				return []string{s}
			}
			return []string{}
		}
		for i := 1; i <= len(s)-part+1; i++ {

			if d, _ := strconv.Atoi(s[:i]); d > 255 {
				break
			}
			for _, sub := range restoreIpPart(s[i:], part-1) {
				solution = append(solution, s[:i]+"."+sub)
			}
		}
	}
	//	fmt.Println(s, solution, part)
	return solution
}

func restoreIpAddresses(s string) []string {
	return restoreIpPart(s, 4)
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("010010"))
}
