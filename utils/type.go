package utils

type CollectionType interface {
	Compare(interface{}) int
}

type IntType int

func (t IntType) Compare(b interface{}) int {
	bInt := b.(IntType)
	if t == bInt {
		return 0
	} else if t < bInt {
		return -1
	} else {
		return 1
	}
}
