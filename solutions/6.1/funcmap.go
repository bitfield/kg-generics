package funcmap

type FuncMap[T, U any] map[string]func(T) U

func (fm FuncMap[T, U]) Apply(name string, val T) U {
	return fm[name](val)
}
