package grep

import (
	"bufio"
	"flag"
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

	args := flag.Args()

	pattern, args, err := g.compilePattern(args)
	if err != nil {
		return fmt.Errorf("-> g.compilePattern%s", err)
	}
	g.implProcess(args, pattern)

	return nil
}

func (g *Grep) compilePattern(args []string) (*regexp.Regexp, []string, error) {
	var regPattern string
	var err error
	if g.flags.F() && len(args) > 0 {
		regPattern, err = g.loadPatternFromFile(args[0])
		if err != nil {
			return nil, args, fmt.Errorf("-> g.loadPatternFromFile%s", err)
		}
		args = args[1:]
	} else if len(args) > 0 {
		regPattern = args[0]
		args = args[1:]
	} else {
		return nil, args, fmt.Errorf(": нет доступного регулярного выражения")
	}

	if g.flags.I() {
		regPattern = "(?i)" + regPattern
	}

	pattern, err := regexp.Compile(regPattern)
	if err != nil {
		return nil, args, fmt.Errorf("-> regexp.Compile: ошибка компиляции шаблона `%s`: %w", regPattern, err)
	}

	return pattern, args, nil
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

		g.implFlags(argv, pattern, file, curFile)

		file.Close()
		curFile++

	}
}

func (g *Grep) implFlags(argv []string, pattern *regexp.Regexp, file *os.File, curFile int) {
	var (
		n int = 0
		k int = 1
		//s   int  = 0
		ret bool = false
		//err error
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if g.flags.V() && g.flags.C() {
			n++
		} else {
			ret, _ = regexp.MatchString(pattern.String(), line)

			if g.flags.I() {
				line = strings.ToLower(line)
			}

			if g.flags.V() {
				if !ret {
					if !g.flags.C() {
						g.printResult(argv, line, len(argv)-1, curFile, k)
					} else {
						n++
					}
				}
			} else {
				if ret {
					if !g.flags.C() {
						g.printResult(argv, line, len(argv)-1, curFile, k)
					} else {
						n++
					}
				}
			}
		}
		k++
	}

	g.printResultCLV(argv, len(argv)-1, curFile, ret, n)
}

func (g *Grep) printResultCLV(argv []string, nFiles int, curFile int, ret bool, n int) {
	if g.flags.C() {
		if nFiles > 1 {
			fmt.Printf("%s", argv[curFile])
			fmt.Printf(":")
		}
		fmt.Printf("%d\n", n)
	}

	if g.flags.V() && ret {
		fmt.Printf("%s\n", argv[curFile])
	}
}

func (g *Grep) printResult(argv []string, line string, nFiles int, curFile int, k int) {
	if nFiles > 1 {
		fmt.Printf("%s", argv[curFile])
		fmt.Printf(":")
	}

	if g.flags.N() {
		fmt.Fprintf(os.Stdout, "%d:", k)
	}
	fmt.Fprintf(os.Stdout, "%s\n", line)
}
