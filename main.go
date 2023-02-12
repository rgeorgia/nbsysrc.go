package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

var args struct {
	Service map[string]string `arg:"positional, required" help:"What you want to add to rc.conf i.e. samba=YES"`
	Verbose bool              `arg:"-v,--verbose" help:"verbosity level"`
}

func main() {
	arg.MustParse(&args)
	fmt.Println("Input:", args.Service)
}
