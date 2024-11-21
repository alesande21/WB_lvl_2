package parser

import (
	"fmt"
	"strings"
)

type Input struct {
	text string
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFlags(args []string) (*Flags, *Input, error) {
	var fs Flags
	var err error
	var end int

	for i, arg := range args {
		end = i + 1
		if strings.HasPrefix(arg, "-f") {
			err = fs.f.Set(arg)
			if err != nil {
				return nil, nil, err
			}
		} else if strings.HasPrefix(arg, "-d") {
			err = fs.d.Set(arg)
			if err != nil {
				return nil, nil, err
			}
		} else if strings.HasPrefix(arg, "-s") {
			fs.s = true
		} else {
			end -= 1
			break
		}

	}

	var in Input
	in.text = strings.Join(args[end:], " ")

	if !fs.f.enabled {
		return nil, nil, fmt.Errorf("%v", "недостаточно флагов. usage: cut -f -otherFlags [text]\n")
	}

	return &fs, &in, nil
}

func (fs Flags) String() string {
	return fmt.Sprintf("{f:%v, d:%v, s:%v}", fs.f, fs.d, fs.s)
}
