package parser

import (
	"fmt"
	"strconv"
	"strings"
)

type Flags struct {
	f flagF
	d flagD
	s bool
}

func (fs Flags) F() bool {
	return fs.f.Enabled()
}

func (fs Flags) FValue() int {
	return fs.f.Index()
}

func (fs Flags) D() bool {
	return fs.d.Enabled()
}

func (fs Flags) DValue() string {
	return fs.d.Sep()
}

func (fs Flags) S() bool {
	return fs.s
}

type flagF struct {
	enabled bool
	index   int
}

func (ff *flagF) Enabled() bool {
	return ff.enabled
}

func (ff *flagF) Index() int {
	return ff.index
}

func (ff *flagF) String() string {
	if ff.enabled {
		str := fmt.Sprintf("{index:%v}", ff.index)
		return str
	}
	return ""
}

func (ff *flagF) Set(value string) error {
	if value == "" {
		ff.enabled = false
		return fmt.Errorf("option requires an argument -- f\n")
	}

	if strings.HasPrefix(value, "-f") {
		value = strings.TrimPrefix(value, "-f")
	}

	if value == "" {
		ff.enabled = false
		return fmt.Errorf("option requires an argument -- f\n")
	}

	num, err := strconv.Atoi(value)
	if err != nil || num < 1 {
		return fmt.Errorf("values may not include zero or below\n")
	}
	ff.enabled = true
	ff.index = num

	return nil
}

type flagD struct {
	enabled bool
	sep     string
}

func (fd *flagD) Enabled() bool {
	return fd.enabled
}

func (fd *flagD) Sep() string {
	return fd.sep
}

func (fd *flagD) String() string {
	if fd.enabled {
		str := fmt.Sprintf("{sep:%v}", fd.sep)
		return str
	}
	return ""
}

func (fd *flagD) Set(value string) error {
	if value == "" {
		fd.enabled = false
		return fmt.Errorf("option requires an argument -- d\n")
	}

	if strings.HasPrefix(value, "-d") {
		value = strings.TrimPrefix(value, "-d")
	}

	if value == "" {
		fd.enabled = false
		return fmt.Errorf("option requires an argument -- d\n")
	}

	fd.enabled = true
	fd.sep = value

	return nil
}
