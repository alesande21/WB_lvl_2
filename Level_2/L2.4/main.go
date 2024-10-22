package main

import (
	"flag"
	"fmt"
)

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

func (f *Flag) ParseFlags() {
	f.k = *flag.Bool("k", false, "указание колонки для сортировки (слова в строке могут выступать "+
		"в качестве колонок, по умолчанию разделитель — пробел)")
	f.n = *flag.Bool("n", false, "сортировать по числовому значению")
	f.r = *flag.Bool("r", false, "сортировать в обратном порядке")
	f.u = *flag.Bool("u", false, "не выводить повторяющиеся строки")
	f.m = *flag.Bool("M", false, "сортировать по названию месяца")
	f.b = *flag.Bool("b", false, "игнорировать хвостовые пробелы")
	f.c = *flag.Bool("c", false, "проверять отсортированы ли данные")
	f.h = *flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()
}

func (f *Flag) Strings() string {
	return "asdasda"
}

func main() {

	var fl Flag

	fl.ParseFlags()
	fmt.Println(fl)
}
