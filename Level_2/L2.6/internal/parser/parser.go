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
	var c bool
	flag.BoolVar(&f.a, "A", false, "печатать +N строк после совпадения")
	flag.BoolVar(&f.b, "B", false, "печатать +N строк до совпадения")
	flag.BoolVar(&c, "C", false, "(A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&f.c, "c", false, "количество строк")
	flag.BoolVar(&f.i, "i", false, "игнорировать регистр")
	flag.BoolVar(&f.v, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&f.f, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&f.n, "n", false, "напечатать номер строки")
	flag.Parse()

	if c == true {
		f.a = true
		f.b = true
	}

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

func (f Flag) String() string {
	return fmt.Sprintf("{A:%v, B:%v, c:%v, i:%v, v:%v, F:%v, n:%v}", f.a, f.b, f.c, f.i, f.v, f.f, f.n)
}
