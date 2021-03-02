package main

import (
	"fmt"
	"io"
)

func printHelp(w io.Writer) {
	fmt.Fprintf(w, "USAGE\n"+
		"rsql -h, --help				: displays usage message and exits\n"+
		"rsql -f <FNAME>, --envfile <FNAME>	: reads environment file to establish MySQL connection.\n"+
		"COMMANDS\n"+
		"add <TABLE>\t: adds data to <TABLE>\n"+
		"info <TABLE>\t: displays the column informations of the <TABLE>\n"+
		"display <TABLE>\t: displays the data of the <TABLE>\n"+
		"q, exit <TABLE>\t: exits the program\n")
}
