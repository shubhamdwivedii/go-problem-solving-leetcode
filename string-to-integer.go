func myAtoi(s string) int {
	num := 0
	mul := 1
	charEnc := false
	rangeL := int(math.Pow(float64(2), 31))
	rangeR := int(math.Pow(float64(2), 31) - 1)
	for _, ch := range s {
		if ch == ' ' {
			if charEnc {
				return num * mul
			}
			continue
		} else if ch == '-' {
			if charEnc {
				return num * mul
			}
			mul = -1
			charEnc = true
			continue
		} else if ch == '+' {
			if charEnc {
				return num * mul
			}
			charEnc = true
			continue
		} else if ch >= '0' && ch <= '9' {
			charEnc = true
			nu := int(ch)
			num *= 10
			num += nu - int('0')
			if num*mul <= -rangeL {
				return -rangeL
			}
			if num*mul >= rangeR {
				return rangeR
			}
		} else {
			return num * mul
		}
	}
	return num * mul
}