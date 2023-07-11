package util

func SliceContain[T comparable](s []T, target T) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func AppendIfMissing[T comparable](slice []T, elem T) []T {
	for _, s := range slice {
		if s == elem {
			return slice
		}
	}
	return append(slice, elem)
}

func Filter[T comparable](slice []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, s := range slice {
		if f(s) {
			result = append(result, s)
		}
	}
	return result
}

func SliceUnion[T comparable](slices ...[]T) []T {
	unionMap := make(map[T]bool)
	for _, slice := range slices {
		for _, elem := range slice {
			unionMap[elem] = true
		}
	}
	union := make([]T, 0, len(unionMap))
	for elem := range unionMap {
		union = append(union, elem)
	}
	return union
}

func RemoveElements[T comparable](s []T, elems ...T) []T {
	m := make(map[T]bool)
	for _, elem := range elems {
		m[elem] = true
	}
	result := make([]T, 0, len(s))
	for _, v := range s {
		if !m[v] {
			result = append(result, v)
		}
	}
	return result
}
