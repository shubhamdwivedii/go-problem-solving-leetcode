func isValid(s string) bool {
	var stack []string

	valid := true
	for _, ch := range s {
		if strings.Contains("({[", string(ch)) {
			stack = append(stack, string(ch))
		} else if strings.Contains(")}]", string(ch)) {
			if len(stack) <= 0 {
				return false
			}
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch string(ch) {
			case ")":
				if popped == "(" {
					continue
				} else {
					valid = false
					break
				}
			case "}":
				if popped == "{" {
					continue
				} else {
					valid = false
					break
				}

			case "]":
				if popped == "[" {
					continue
				} else {
					valid = false
					break
				}
			default:
				valid = false
				break
			}
		} else {
			valid = false
			break
		}
	}

	if len(stack) > 0 {
		return false
	}
	return valid
}