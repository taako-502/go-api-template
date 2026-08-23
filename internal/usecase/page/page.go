package page

type Page[T any] struct {
	Items []T
	Total int
}

// Map はページング情報を維持したまま、要素を別の型へ変換する。
func (p Page[T]) Map[U any](convert func(T) U) Page[U] {
	items := make([]U, len(p.Items))
	for i, item := range p.Items {
		items[i] = convert(item)
	}
	return Page[U]{Items: items, Total: p.Total}
}
