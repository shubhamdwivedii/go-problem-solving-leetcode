func maxArea(h []int) int {
	ptrL := 0
	ptrR := len(h) - 1

	maxArea := 0

	for ptrL < ptrR {
		length := ptrR - ptrL
		height := int(math.Min(float64(h[ptrL]), float64(h[ptrR])))
		area := length * height
		if area > maxArea {
			maxArea = area
		}

		if h[ptrL] < h[ptrR] {
			ptrL++
		} else {
			ptrR--
		}
	}

	return maxArea
}