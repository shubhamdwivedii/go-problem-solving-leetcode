func reverseString(s []byte) {
	ptrL := 0
	ptrR := len(s) - 1

	for ptrL < ptrR {
		// swap
		s[ptrL], s[ptrR] = s[ptrR], s[ptrL]
		ptrL++
		ptrR--
	}

}