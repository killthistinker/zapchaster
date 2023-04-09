package mapper

func Map[T, U any](items []T, f func(T) U) []U {
	mapped := make([]U, len(items))
	for i, item := range items {
		mapped[i] = f(item)
	}
	return mapped
}
