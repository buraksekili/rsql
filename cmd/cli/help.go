package cli

import (
	"fmt"
	"io"
)

func PrintHelp(w io.Writer) {
	fmt.Fprintf(w, "USAGE\n"+
		"\trsql -h, --help				: displays usage message and exits\n"+
		"\trsql -f <FNAME>, --envfile <FNAME>	: reads environment file to establish MySQL connection.\n"+
		"COMMANDS\n"+
		"\tadd <TABLE>\t: adds data to <TABLE>\n"+
		"\tinfo <TABLE>\t: displays the column informations of the <TABLE>\n"+
		"\tdisplay <TABLE>\t: displays the data of the <TABLE>\n"+
		"\ttables\t\t: displays available tables under the <DB> specified by user\n"+
		"\tq, exit <TABLE>\t: exits the program\n")
}
