package parser

import (
	"flag"
	"fmt"
	"os"
)

type FilePath struct {
	Path string
}

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

func (p Parser) checkFilePath() (FilePath, error) {
	args := flag.Args()
	if len(args) != 1 {
		return FilePath{}, fmt.Errorf("-> parseFilePath: введено неверное количество аргументов: %d вместо ожидаемых: 1", len(args))
	}

	_, err := os.Stat(args[0])
	if !(err == nil || !os.IsNotExist(err)) {
		return FilePath{}, fmt.Errorf("-> parseFilePath-> os.Stat: файл по указаному пути не найден %s", args[0])
	}

	return FilePath{args[0]}, nil
}

func (p *Parser) checkFlagK(fk *flagK) ([]string, error) {
	args := os.Args

	if len(args) != 3 {
		return os.Args, nil
	}

	//fmt.Println("TUT")

	newArgs := os.Args[:1]

	err := fk.Set(os.Args[1])
	if err != nil {
		return os.Args, nil
	}

	newArgs = append(newArgs, os.Args[2])

	return newArgs, nil
}

func (f Flag) String() string {
	return fmt.Sprintf("{k:%v, n:%v, r:%v, u:%v, M:%v, b:%v, c:%v, h:%v}", f.k, f.n, f.r, f.u, f.m, f.b, f.c, f.h)
}
