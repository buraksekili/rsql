package cli

import (
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

type ConnInfoInput struct{}

type UnknownOp struct {
	Error string
}

func ParseFlag(flags []string) Operation {
	if len(flags) == 0 {
		return ConnInfoInput{}
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
		return EnvFileOp{Filename: flags[1]}
	}

	if strings.HasPrefix(op, "-f") && strings.Contains(op, "=") {
		i := strings.Index(op, "=")
		fname := op[i+1:]
		if len(fname) == 0 {
			return InvalidOp{Error: "you need to specify file name."}
		}
		return EnvFileOp{fname}
	}

	if strings.HasPrefix(op, "--envfile") && strings.Contains(op, "=") {
		i := strings.Index(op, "=")
		fname := op[i+1:]
		if len(fname) == 0 {
			return InvalidOp{Error: "you need to specify file name."}
		}
		return EnvFileOp{fname}
	}

	return UnknownOp{Error: "unsupported flag"}
}
