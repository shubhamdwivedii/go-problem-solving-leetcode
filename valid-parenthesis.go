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

// TOn SOn 

func isValid(s string) bool {
    var open []rune

    for _, r := range s {
        if isOpen(r) {
            open = append(open, r)
        } else {
            // Closing bracket but nothing to match
            if len(open) == 0 {
                return false
            }

            if match(open[len(open)-1], r) {
                open = open[:len(open)-1]
            } else {
                return false
            }
        }
    }

    return len(open) == 0
}

func isOpen(r rune) bool {
    return r == '[' || r == '(' || r == '{'
}

func match(open, close rune) bool {
    switch open {
    case '[':
        return close == ']'
    case '(':
        return close == ')'
    case '{':
        return close == '}'
    default:
        return false
    }
}