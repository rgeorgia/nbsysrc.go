// main package. this script is like sysrc for Freebsd
package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"log"
	"os"
)

var args struct {
	Service map[string]string `arg:"positional, required" help:"What you want to add to rc.conf i.e. samba=YES"`
}

func main() {
	arg.MustParse(&args)
	if args.Service == nil {
		log.Fatal("You must enter a service")
	}

	fmt.Printf("Key value [%v]\n", args.Service)
	fmt.Printf("Am I root? %v\n", isUserRoot())
}

func isUserRoot() bool {
	return os.Getuid() == 0
}
