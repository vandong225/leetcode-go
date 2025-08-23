package stack

import "strings"

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	stack := make([]string, 0)

	for i := 0; i < len(dirs); i++ {
		action := dirs[i]
		if action == "" || action == "." {
			continue
		}
		if action == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		stack = append(stack, action)
	}

	return "/" + strings.Join(stack, "/")
}
