package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintHelp(t *testing.T) {
	var buff bytes.Buffer
	PrintHelp(&buff)
	strHelp := buff.String()
	if !strings.HasPrefix(strHelp, "USAGE") {
		t.Fatal("help message doesn't start with USAGE")
	}
}
