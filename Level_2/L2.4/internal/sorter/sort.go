package sorter

import (
	"fmt"
	"io"
	"os"
	"sort_command/internal/parser"
)

type Sort struct{}

func NewSorter() *Sort {
	return &Sort{}
}

func (s *Sort) Run(flags *parser.Flag, filePath *parser.FilePath) error {

	if flags == nil {
		return fmt.Errorf(": flags == nil ")
	}

	if filePath == nil {
		return fmt.Errorf(": filePath == nil ")
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

	if *flags.N() == true {

	}

	return nil
}
