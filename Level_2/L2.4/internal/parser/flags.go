package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type Flag struct {
	k flagK
	n bool
	r bool
	u bool
	m bool
	b bool
	c bool
	h bool
}

func (f Flag) K() bool {
	return f.k.enabled
}

func (f Flag) Col() int {
	return f.k.col
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

type flagK struct {
	enabled bool
	col     int
}

func (fk *flagK) String() string {
	if fk.enabled {
		return strconv.Itoa(fk.col)
	}
	return ""
}

func (fk *flagK) Set(value string) error {
	if value == "" {
		fk.enabled = true
		fk.col = 1
	}

	if strings.HasPrefix(value, "-k") {
		value = strings.TrimPrefix(value, "-k")
	}

	if value == "" {
		fk.enabled = true
		fk.col = 1
		return nil
	}

	col, err := strconv.Atoi(value)
	if err != nil || col < 1 {
		return fmt.Errorf("-> strconv.Atoi: ошибка при преобразовании числа флага k: %s", err)
	}

	fk.enabled = true
	fk.col = col

	return nil
}
