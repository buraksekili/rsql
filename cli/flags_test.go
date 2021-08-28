package cli

import (
	"reflect"
	"testing"
)

func TestParseFlag(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected Operation
	}{
		{
			name:     "nil flag",
			args:     nil,
			expected: ConnInfoInput{},
		},
		{
			name:     "empty flag",
			args:     []string{},
			expected: ConnInfoInput{},
		},
		{
			name:     "unknown flag (no - before flag); f",
			args:     []string{"f"},
			expected: UnknownOp{Error: "f"},
		},
		{
			name:     "env file flag; -f=file",
			args:     []string{"-f=lol.txt"},
			expected: EnvFileOp{Filename: "lol.txt"},
		},
		{
			name:     "env file flag; -f file",
			args:     []string{"-f", "lol.txt"},
			expected: EnvFileOp{Filename: "lol.txt"},
		},
		{
			name:     "env file flag; --envfile=file",
			args:     []string{"--envfile=lol.txt"},
			expected: EnvFileOp{Filename: "lol.txt"},
		},
		{
			name:     "env file flag; --envfile file",
			args:     []string{"--envfile", "lol.txt"},
			expected: EnvFileOp{Filename: "lol.txt"},
		},
		{
			name:     "invalid -f flag; no file specified",
			args:     []string{"-f"},
			expected: InvalidOp{Error: "you need to specify file name."},
		},
		{
			name:     "invalid -f flag; no file specified after =",
			args:     []string{"-f="},
			expected: InvalidOp{Error: "you need to specify file name."},
		},
		{
			name:     "invalid --envfile file flag; no file specified",
			args:     []string{"--envfile"},
			expected: InvalidOp{Error: "you need to specify file name."},
		},
		{
			name:     "invalid --envfile file flag; no file specified after =",
			args:     []string{"--envfile="},
			expected: InvalidOp{Error: "you need to specify file name."},
		},
		{
			name:     "help operator -h",
			args:     []string{"-h"},
			expected: HelpOp{},
		},
		{
			name:     "help operator --help",
			args:     []string{"--help"},
			expected: HelpOp{},
		},
		{
			name:     "invalid number of arguments",
			args:     []string{"-f", "lol.txt", "dsad"},
			expected: InvalidOp{Error: "Too many flags entered."},
		},
		{
			name:     "unknown flag",
			args:     []string{"-x"},
			expected: UnknownOp{Error: "unsupported flag"},
		},
	}
	
	for _, test := range tests {
		op := ParseFlag(test.args)
		if !reflect.DeepEqual(test.expected, op) {
			t.Fatalf("%s, got=%T expected=%T", test.name, op, test.expected)
		}
	}
}
