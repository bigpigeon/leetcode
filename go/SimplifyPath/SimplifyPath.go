/*
71. Simplify Path

Given an absolute path for a file (Unix-style), simplify it.

For example,
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"
Corner Cases:
Did you consider the case where path = "/../"?
In this case, you should return "/".
Another corner case is the path might contain multiple slashes '/' together, such as "/home//foo/".
In this case, you should ignore redundant slashes and return "/home/foo".
*/
package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	if len(path) == 0 {
		return ""
	}
	//split with /
	sourceList := strings.Split(path, "/")
	pathList := []string{}
	for _, v := range sourceList {
		if len(v) == 0 {
			continue
		}
		if v == ".." {
			if len(pathList) > 0 {
				pathList = pathList[:len(pathList)-1]
			}
		} else if v == "." {

		} else {
			pathList = append(pathList, v)
		}
	}
	s := "/" + strings.Join(pathList, "/")
	return s
}

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/home/../"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a/./b/../..///c/"))
	fmt.Println(simplifyPath("/../"))
}
