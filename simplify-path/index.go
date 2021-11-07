package algorithms

func SimplifyPath(path string) string {
	stack := make([]string, 0)
	currentStr := ""
	path += "/"
	for i := 0; i < len(path); i++ {
		char := path[i]
		if char == '/' {
			if currentStr == "." || currentStr == "" {
				currentStr = ""
				continue
			} else if currentStr == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, currentStr)
			}
			currentStr = ""
		} else {
			currentStr += string(char)
		}
	}
	result := ""
	for _, val := range stack {
		result += "/" + val
	}
	if result == "" {
		return "/"
	}
	return result
}
