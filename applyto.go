package extensions

func ApplyToNewSlice[T comparable, K comparable](target *[]T, f FBack[T, K]) []K {
	lex := make([]K, len(*target))
	for k, v := range *target {
		lex[k] = f(v)
	}
	return lex
}

func ApplyToNewMap[K1 comparable, V1 comparable, V2 comparable](
	target *map[K1]V1,
	f FBack[V1, V2],
) map[K1]V2 {
	lex := make(map[K1]V2)
	for k, v := range *target {
		lex[k] = f(v)
	}
	return lex
}
