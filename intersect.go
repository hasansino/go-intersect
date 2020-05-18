/*
Package intersect provides a simple way to find intersection
between two int64 or string arrays.

Simple usage:
    var (
        a = []int64{1,2,3}
        b = []int64{3,5,6}
    )
    c := intersect.Int64(a, b)
*/
package intersect

// Int64 finds intersection between two []int64 arrays
// Complexity O(n+m)
func Int64(a, b []int64) []int64 {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	var small, big []int64
	if len(a) < len(b) {
		small, big = a, b
	} else {
		small, big = b, a
	}

	var (
		result  = make([]int64, 0, len(small))
		hashMap = make(map[int64]struct{}, len(small))
	)

	for i := 0; i < len(small); i++ {
		if _, exist := hashMap[small[i]]; !exist {
			hashMap[small[i]] = struct{}{}
		}
	}

	for i := 0; i < len(big); i++ {
		if _, exist := hashMap[big[i]]; exist {
			result = append(result, big[i])
		}
	}

	return result
}

// Int64 finds intersection between two []string arrays
// Complexity O(n+m)
func String(a, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}

	var small, big []string
	if len(a) < len(b) {
		small, big = a, b
	} else {
		small, big = b, a
	}

	var (
		result  = make([]string, 0, len(small))
		hashMap = make(map[string]struct{}, len(small))
	)

	for i := 0; i < len(small); i++ {
		if _, exist := hashMap[small[i]]; !exist {
			hashMap[small[i]] = struct{}{}
		}
	}

	for i := 0; i < len(big); i++ {
		if _, exist := hashMap[big[i]]; exist {
			result = append(result, big[i])
		}
	}

	return result
}
