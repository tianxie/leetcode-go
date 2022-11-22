package jzoffer001

type solution interface {
	divide(a int, b int) int
}

type solution1 struct{}

func (s solution1) divide(a int, b int) int {
	return 0
}
