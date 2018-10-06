package numutil

// Max returns the max int in a slice of ints.
func Max(xs ...int) int {
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}
