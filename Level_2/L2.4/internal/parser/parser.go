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
	f.k = flag.Bool("k", false, "указание колонки для сортировки (слова в строке могут выступать "+
		"в качестве колонок, по умолчанию разделитель — пробел)")
	f.n = flag.Bool("n", false, "сортировать по числовому значению")
	f.r = flag.Bool("r", false, "сортировать в обратном порядке")
	f.u = flag.Bool("u", false, "не выводить повторяющиеся строки")
	f.m = flag.Bool("M", false, "сортировать по названию месяца")
	f.b = flag.Bool("b", false, "игнорировать хвостовые пробелы")
	f.c = flag.Bool("c", false, "проверять отсортированы ли данные")
	f.h = flag.Bool("h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()

	filePath, err := p.checkFilePath()
	return &f, &filePath, err
}

/*
_, err := os.Stat(filePath)
	if !(err == nil || !os.IsNotExist(err)) {
		return fmt.Errorf("-> os.Stat: файл по указаному пути не найден %s", filePath)
	}

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return fmt.Errorf("-> os.OpenFile: ошибка при открытии файла %s: %w", filePath, err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("-> io.ReadAll: ошибка при чтении файла %s: %w", filePath, err)
	}

	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return fmt.Errorf("-> yaml.Unmarshal: ошибка при кодировании файла: %w", err)
	}

	err = cleanenv.UpdateEnv(c)
	if err != nil {
		return fmt.Errorf("-> cleanenv.UpdateEnv: ошибка при обновлении параметроа из переменныз окружения%w", err)
	}

	err = c.validate()
	if err != nil {
		return fmt.Errorf("-> c.validate%w", err)
	}

	return nil
*/

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
	return fmt.Sprintf("{k:%v, n:%v, r:%v, u:%v, M:%v, b:%v, c:%v, h:%v}", *f.k, *f.n, *f.r, *f.u, *f.m, *f.b, *f.c, *f.h)
}
