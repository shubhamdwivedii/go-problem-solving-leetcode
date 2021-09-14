func fizzBuzz(n int) []string {
	out := make([]string, n)

	for i := 0; i < n; i++ {
		t := i + 1
		if t%3 == 0 && t%5 == 0 {
			out[i] = "FizzBuzz"
			continue
		}
		if t%3 == 0 {
			out[i] = "Fizz"
			continue
		}
		if t%5 == 0 {
			out[i] = "Buzz"
			continue
		}
		out[i] = fmt.Sprintf("%v", t)
	}

	return out

}