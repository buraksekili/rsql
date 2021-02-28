package main

import (
	"fmt"
	"io"
)

func printHelp(w io.Writer) {
	fmt.Fprintf(w, "USAGE\n"+
		"rsql -h, --help				: displays usage message\n"+
		"rsql -f <FNAME>, --envfile <FNAME>	: reads environment file to establish MySQL connection.\n")
}
