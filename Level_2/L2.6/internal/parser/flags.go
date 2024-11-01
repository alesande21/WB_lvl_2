package parser

type Flag struct {
	a bool
	b bool
	c bool
	i bool
	v bool
	f bool
	n bool
}

func (f Flag) A() bool {
	return f.a
}

func (f Flag) B() bool {
	return f.b
}

func (f Flag) C() bool {
	return f.c
}

func (f Flag) I() bool {
	return f.i
}

func (f Flag) V() bool {
	return f.v
}

func (f Flag) F() bool {
	return f.f
}

func (f Flag) N() bool {
	return f.n
}
