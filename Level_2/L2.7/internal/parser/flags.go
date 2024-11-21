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

type flagF struct {
	enabled bool
	start   int
	end     int
}

func (ff *flagF) String() string {
	if ff.enabled {
		str := fmt.Sprintf("{start:%v, end:%v}", ff.start, ff.end)
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
	ff.start = num

	return nil
}

type flagD struct {
	enabled bool
	sep     string
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
