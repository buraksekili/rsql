package main

import (
	"os"
	"strings"
)

type Operation interface{}

type HelpOp struct{}

type EnvFileOp struct {
	Filename string
}

type InvalidOp struct {
	Error string
}

type UnknownOp struct {
	Error string
}

func parseFlag() Operation {
	flags := os.Args[1:]

	if len(flags) == 0 {
		return nil
	}

	if len(flags) > 2 {
		return InvalidOp{Error: "Too many flags entered."}
	}

	op := strings.Trim(flags[0], " ")
	if !strings.HasPrefix(op, "-") {
		return UnknownOp{op}
	}

	if op == "-h" || op == "--help" {
		return HelpOp{}
	}

	if op == "-f" || op == "--envfile" {
		if len(flags) < 2 {
			return InvalidOp{Error: "you need to specify file name."}
		}
		fname := flags[1]
		return EnvFileOp{Filename: fname}
	}
	if strings.HasPrefix(op, "-f") && strings.Contains(op, "=") {
		i := strings.Index(op, "=")
		fname := op[i+1:]
		return EnvFileOp{fname}
	}
	if strings.HasPrefix(op, "--envfile") && strings.Contains(op, "=") {
		i := strings.Index(op, "=")
		fname := op[i+1:]
		return EnvFileOp{fname}
	}

	return UnknownOp{Error: "unsupported flag"}
}
