package engine

import (
	"fmt"
)

type Command interface {
	Execute(handler Handler)
}

type PrintCommand struct {
	Arg string
}

func (p *PrintCommand) Execute(loop Handler) {
	fmt.Println(p.Arg)
}

type PrintcCommand struct {
	Count int
	Symbol string
}

func (printcArgs *PrintcCommand) Execute(loop Handler) {
	res := ""

	for i := 0; i < printcArgs.Count; i++{
		res += printcArgs.Symbol
	}

	loop.Post(&PrintCommand{Arg: res})
}

type finishCommand func (handler Handler)

func (c finishCommand) Execute(handler Handler) {
	c(handler)
}
