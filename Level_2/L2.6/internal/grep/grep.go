package grep

import (
	"bufio"
	"fmt"
	"grep/internal/parser"
	"os"
	"regexp"
	"strings"
)

type Grep struct {
	flags *parser.Flag
}

func NewGrep(flags *parser.Flag) *Grep {
	return &Grep{flags: flags}
}

func (g *Grep) Run() error {
	if g.flags == nil {
		return fmt.Errorf(": флаги отсутствуют")
	}

}

func (g *Grep) compilePattern(args []string, reg string) (*regexp.Regexp, error) {
	var regPattern string
	var err error
	if g.flags.F() && len(args) > 0 {
		regPattern, err = g.loadPatternFromFile(args[0])
		if err != nil {
			return nil, fmt.Errorf("-> g.loadPatternFromFile%s", err)
		}
		args = args[1:]
	} else if len(args) > 0 {
		regPattern = args[0]
		args = args[1:]
	} else {
		return nil, fmt.Errorf(": нет доступного регулярного выражения")
	}

	if g.flags.I() {
		regPattern = "(?i)" + regPattern
	}

	pattern, err := regexp.Compile(regPattern)
	if err != nil {
		return nil, fmt.Errorf("-> regexp.Compile: ошибка компиляции шаблона `%s`: %w", regPattern, err)
	}

	return pattern, nil
}

func (g *Grep) loadPatternFromFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("-> os.Open: ошибка открытия файла: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var b strings.Builder
	for scanner.Scan() {
		b.WriteString(scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		return "", fmt.Errorf("-> scanner.Scan: ошибка при сканировании документа: %s", err)
	}

	return b.String(), nil
}

func (g *Grep) implProcess(argv []string, pattern *regexp.Regexp) {
	var file *os.File
	var err error
	var curFile int = 0

	for curFile < len(argv) {
		file, err = os.Open(argv[curFile])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s: No such file or directory\n", "MyGrep", argv[curFile])
			curFile++
			continue
		}

	}
}

func (g *Grep) implFlags(argv []string, pattern *regexp.Regexp) {

}
