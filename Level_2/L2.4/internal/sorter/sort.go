package sorter

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"slices"
	"sort"
	"sort_command/internal/parser"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Sort struct {
	lines []string
	m     map[string]bool
	col   int
}

func NewSorter() *Sort {
	return &Sort{
		lines: make([]string, 0),
		m:     make(map[string]bool),
	}
}

func (s *Sort) LessForNumbers(i, j int) bool {
	numI, _ := s.extractLeadingNumbers(s.lines[i])
	numJ, _ := s.extractLeadingNumbers(s.lines[j])

	if numI == numJ {
		return strings.Compare(s.lines[i], s.lines[j]) == -1
	}
	return numI < numJ
}

func (s *Sort) extractLeadingNumbers(str string) (int64, error) {
	var b strings.Builder

	for i := 0; i < len(str); {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		if runeValue <= 47 || runeValue > 57 {
			break
		}
		b.WriteRune(runeValue)
		i += width
	}

	if b.Len() == 0 {
		return 0, nil
	}

	return strconv.ParseInt(b.String(), 10, 64)
}

func (s *Sort) extractLeadingNumbers2(str string) (int64, error) {
	re := regexp.MustCompile(`^\D*(\d+)`)
	matches := re.FindStringSubmatch(str)
	if len(matches) > 1 {
		return strconv.ParseInt(matches[1], 10, 64)
	}
	return 0, nil
}

func (s *Sort) sortByNumbers(lines []string) []string {
	//res := make([]string, len(lines))

	sort.Slice(lines, s.LessForNumbers)

	return lines
}

func (s *Sort) sortReverse(lines []string) []string {
	//res := make([]string, len(lines))

	sort.Strings(s.lines)
	slices.Reverse(s.lines)

	return lines
}

// LessForUnique сортировка уникальных значений
func (s *Sort) LessForUnique(i, j int) bool {

	comp := strings.Compare(s.lines[i], s.lines[j])

	if comp == 0 {
		s.m[s.lines[i]] = true
		return false
	}

	return comp == -1
}

// sortUnique сортировка уникальных значений
func (s *Sort) sortUnique(lines []string) []string {
	//res := make([]string, len(lines))

	sort.Slice(lines, s.LessForUnique)

	return lines
}

// LessForSortByColumn сортировка по столбцу
func (s *Sort) LessForSortByColumn(i, j int) bool {
	re := regexp.MustCompile(`\s+`)
	cleanedLine := re.ReplaceAllString(s.lines[i], " ")
	strI := strings.Split(cleanedLine, " ")

	cleanedLine = re.ReplaceAllString(s.lines[j], " ")
	strJ := strings.Split(cleanedLine, " ")

	//strI := strings.Fields(s.lines[i])
	//strJ := strings.Fields(s.lines[j])

	if len(strI) < s.col {
		return true
	}

	if len(strJ) < s.col {
		return false
	}

	//comp := strings.Compare(strI[s.col], strJ[s.col])

	fmt.Printf("strI:%s, strJ:%s\n", strI[s.col], strJ[s.col])

	return strI[s.col] < strJ[s.col]
}

func (s *Sort) getStr(slice []string, i int) string {
	if i < 0 || i >= len(slice) {
		return ""
	}

	//fmt.Println(slice[i])

	return slice[i]
}

// sortByColumn сортировка по столбцу
func (s *Sort) sortByColumn(col int, lines []string) []string {
	s.col = col - 1
	sort.Slice(lines, s.LessForSortByColumn)

	return lines
}

func (s *Sort) Run(flags *parser.Flag, filePath *parser.FilePath) error {

	if flags == nil {
		return fmt.Errorf(": flags == nil ")
	}

	if filePath == nil {
		return fmt.Errorf(": filePath == nil ")
	}

	file, err := os.OpenFile(filePath.Path, os.O_RDONLY, 0666)
	if err != nil {
		return fmt.Errorf("-> os.OpenFile: ошибка при открытии файла %s: %w", filePath, err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("-> io.ReadAll: ошибка при чтении файла %s: %w", filePath, err)
	}

	s.lines = strings.Split(string(buf), "\n")
	//fmt.Println("До: ")
	//for _, l := range lines {
	//	fmt.Println(l)
	//}

	if flags.N() {
		s.sortByNumbers(s.lines)
	} else if flags.R() {
		s.sortReverse(s.lines)
	} else if flags.U() {
		s.sortUnique(s.lines)
	} else if flags.K() {
		s.sortByColumn(flags.Col(), s.lines)
	}
	//fmt.Println("После: ")

	for _, l := range s.lines {
		if flags.U() {
			val, ok := s.m[l]
			if ok && val == false {
				continue
			} else if ok {
				s.m[l] = false
			}

		}
		fmt.Printf("%s\n", l)
	}

	return nil
}
