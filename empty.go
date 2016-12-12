package gomill

type Empty struct{}

func NewEmpty() *Empty {
	return new(Empty)
}
