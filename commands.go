package main

import (
  "./engine"
	"fmt"
)

type PrintCommand struct {
	Arg string
}

func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.Arg)
}

type PrintcCommand struct {
	Count int
	Symbol string
}

func (printcArgs *PrintcCommand) Execute(loop engine.Handler) {
	res := ""

	for i := 0; i < printcArgs.Count; i++{
		res += printcArgs.Symbol
	}

	loop.Post(&PrintCommand{Arg: res})
}
