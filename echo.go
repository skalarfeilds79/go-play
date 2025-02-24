package main

import (
	"os";
	"flag";	// command line option parser
)

var omitNewline = flag.Bool("n", false, "don't print final newline")

const (
	Space	= " ";
	Newline	= "\n";
)

func main() {
	flag.Parse();	// Scans the arg list and sets up flags
	s := "";	// var s = ""; or even var s string = "";
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i);
	}
	if !*omitNewline {
		s += Newline
	}
	os.Stdout.WriteString(s);
}
