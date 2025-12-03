package math

import "math"

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func GCD(a, b int) int {
	aInner := a
	bInner := b
	for bInner != 0 {
		t := bInner
		bInner = aInner % bInner
		aInner = t
	}
	return aInner
}

func Abs(a int) int {
	res := a
	if a < 0 {
		res *= -1
	}
	return res
}

func LCM(a, b int) int {
	gcd := GCD(a, b)
	return Abs(a) * (Abs(b) / gcd)
}

func Permutations[T any](items []T) [][]T {
	var helper func([]T, int)
	res := [][]T{}

	helper = func(arr []T, n int) {
		if n == 1 {
			tmp := make([]T, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[1]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {

					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(items, len(items))
	return res
}
