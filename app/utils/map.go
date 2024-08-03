package utils

/**
 * @author eyesYeager
 * @date 2023/2/19 14:24
 */

func Contains[T string | uint64 | uint32](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
