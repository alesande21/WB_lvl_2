package parser

type Flag struct {
	k bool
	n bool
	r bool
	u bool
	m bool
	b bool
	c bool
	h bool
}

func (f Flag) K() bool {
	return f.k
}

func (f Flag) N() bool {
	return f.n
}

func (f Flag) R() bool {
	return f.r
}

func (f Flag) U() bool {
	return f.u
}

func (f Flag) M() bool {
	return f.m
}

func (f Flag) B() bool {
	return f.b
}

func (f Flag) C() bool {
	return f.c
}

func (f Flag) H() bool {
	return f.h
}
