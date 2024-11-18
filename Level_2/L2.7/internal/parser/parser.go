package parser

import (
	"flag"
	"os"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFlags() (*Flag, *FilePath, error) {
	var f Flag
	var err error
	os.Args, _ = p.checkFlagK(&f.k)
	flag.Var(&f.k, "k", "указание колонки для сортировки (слова в строке могут выступать "+
		"в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&f.n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&f.r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&f.u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&f.m, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&f.b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&f.c, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&f.h, "h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()

	filePath, err := p.checkFilePath()
	return &f, &filePath, err
}
