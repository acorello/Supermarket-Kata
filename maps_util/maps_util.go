package maps_util

func Merge[K comparable, V any](combine func(key K, leftVal, rightVal V) V, result, input map[K]V) {
	// TODO: test inputs are NOT modified
	for k, v := range input {
		if vr, found := result[k]; found {
			v = combine(k, vr, v)
		}
		result[k] = v
	}
}
