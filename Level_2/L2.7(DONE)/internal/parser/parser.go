package parser

import (
	"fmt"
	"strings"
)

type Input struct {
	text string
}

func (i *Input) Text() string {
	return i.text
}

func (i *Input) SetText(text string) {
	i.text = text
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseFlags(args []string) (*Flags, error) {
	var fs Flags
	var err error
	fs.DefD()
	for _, arg := range args {
		if strings.HasPrefix(arg, "-f") {
			err = fs.f.Set(arg)
			if err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(arg, "-d") {
			err = fs.d.Set(arg)
			if err != nil {
				return nil, err
			}
		} else if strings.HasPrefix(arg, "-s") {
			fs.s = true
		} else {
			return nil, fmt.Errorf("таких флагов не существует. используйте -f -d -s")
		}

	}

	if !fs.f.enabled {
		return nil, fmt.Errorf("%v", "недостаточно флагов. usage: cut -f -otherFlags [text]\n")
	}

	return &fs, nil
}

func (fs Flags) String() string {
	return fmt.Sprintf("{f:%v, d:%v, s:%v}", fs.f, fs.d, fs.s)
}
