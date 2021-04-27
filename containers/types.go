package containers

type Container interface {
	Key() interface{}
	Less(container Container) bool
	Validate(interface{}) Container
}

type IntContainer int

func (c IntContainer) Key() interface{} {
	return int(c)
}

func (c IntContainer) Less(value Container) bool {
	y := value.(IntContainer)
	return int(c) < int(y)
}

func (IntContainer) Validate(x interface{}) Container {
	converted, ok := x.(int)
	if !ok {
		panic("Type conversion failed")
	}
	return IntContainer(converted)
}

func (c IntContainer) ToNative() int {
	return int(c)
}

type StringContainer string

func (c StringContainer) Key() interface{} {
	return string(c)
}

func (c StringContainer) Less(value Container) bool {
	x := string(c)
	y := string(value.(StringContainer))
	return x < y
}

func (StringContainer) Validate(x interface{}) Container {
	converted, ok := x.(string)
	if !ok {
		panic("Type conversion failed")
	}
	return StringContainer(converted)
}
