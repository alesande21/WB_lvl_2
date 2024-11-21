package cut

import "myCut/internal/parser"

type MyCut struct{}

func NewCut() *MyCut {
	return &MyCut{}
}

func (mc *MyCut) Start(flags parser.Flags, line parser.Input) error {

}
