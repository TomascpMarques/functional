package extensions

/* Anonymous functions types*/
type FNonce func()
type F[T comparable] func(T)
type FBack[T comparable, K comparable] func(T) K
