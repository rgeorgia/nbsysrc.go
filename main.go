// main package. this script is like sysrc for Freebsd
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alexflint/go-arg"
)

var args struct {
	Service      map[string]string `arg:"positional, required" help:"What you want to add to rc.conf i.e. samba=YES"`
	LocalEtc     string            `arg:"-e" default:"/etc/" help:"Root path to alternate etc dir. Used for local testing"`
	LocalExample string            `arg:"-x" default:"/usr/pkg/share/example/" help:"Root path to alternate etc dir. Used for local testing"`
	Verbose      bool              `arg:"-v,--verbose" help:"verbosity level"`
}

func main() {
	arg.MustParse(&args)
	if args.Service == nil {
		log.Fatal("You must enter a service")
	}

	rcContent, err := ReadRcConfigFile(args.LocalEtc + DefaultRcConfFile)
	if err != nil {
		log.Fatal(err)
	}

	var rcMap = make(map[string]string)

	for _, line := range rcContent {
		if !strings.HasPrefix(line, "#") && len(line) > 0 && strings.Contains(line, "=") {
			rcParts := strings.Split(line, "=")
			rcMap[rcParts[0]] = rcParts[1]
		}
	}

	fmt.Println("slim is here ", rcMap["slim"])
	fmt.Printf("Key value [%v]\n", args.Service)
}

func isUserRoot() bool {
	return os.Getuid() == 0
}
