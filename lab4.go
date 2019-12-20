package main

import (
	"./engine"
	"bufio"
	"strings"
	"os"
	"strconv"
)

func parse(commandLine string) engine.Command {
	parts := strings.Fields(commandLine)
	if parts[0] == "printc" {
		if len(parts) == 3 {
			count, err := strconv.Atoi(parts[1])
			symbol := parts[2]
			if err != nil {
					return &engine.PrintCommand{Arg: "SYNTAX ERROR: " + err.Error()}
			} else {
				if (len(symbol) == 1) {
						return &engine.PrintcCommand{Count: count, Symbol: symbol}
				} else {
						return &engine.PrintCommand{Arg: "SYNTAX ERROR: illegal symbol argument"}
				}
			}
		} else {
			return &engine.PrintCommand{Arg: "SYNTAX ERROR: illegal printc argument"}
		}
	} else if parts[0] == "print" {
		if len(parts) == 2 {
			return &engine.PrintCommand{Arg: parts[1]}
		} else {
			return &engine.PrintCommand{Arg: "SYNTAX ERROR: print arguments"}
		}
	} else {
		return &engine.PrintCommand{Arg: "SYNTAX ERROR: unexpected command"}
	}
}

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open("./examples.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
