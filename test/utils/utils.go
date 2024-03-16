package utils

func Plus(v *int) int {
	*v++
	return *v
}

func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}
