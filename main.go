package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
)

var args struct {
	Service      map[string]string `arg:"positional, required" help:"What you want to add to rc.conf i.e. samba=YES"`
	LocalEtc     string            `arg:"-e" default:"/etc/" help:"Root path to alternate etc dir. Used for local testing`
	LocalExample string            `arg:"-x" default:"/usr/pkg/share/example/" help:"Root path to alternate etc dir. Used for local testing`
	Verbose      bool              `arg:"-v,--verbose" help:"verbosity level"`
}

func main() {
	arg.MustParse(&args)
	fmt.Println("Input:", args.Service)
	fmt.Printf("Local Etc: %s\n", args.LocalEtc)
	fmt.Printf("Local Example: %s\n", args.LocalExample)
	fmt.Printf("Is user root? %v", is_user_root())
}

func is_user_root() bool {
	if os.Geteuid() == 0 {
		return true
	}
	return false
}
