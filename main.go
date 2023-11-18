// main package. this script is like sysrc for Freebsd
package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/alexflint/go-arg"
	"log"
	"os"
)

var args struct {
	Service map[string]string `arg:"positional, required" help:"What you want to add to rc.conf i.e. samba=YES"`
}

type NbSysRcMeta struct {
	EctRcd     string `toml:"ect_rcd"`
	ExampleRcd string `toml:"example_rcd"`
	RcConf     string `toml:"rc_conf"`
}

func main() {
	arg.MustParse(&args)
	if args.Service == nil {
		log.Fatal("You must enter a service")
	}

	var config NbSysRcMeta
	_, err := toml.DecodeFile(".nbsysrc.toml", &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("/etc/rc.d is: %s\n", config.EctRcd)
	fmt.Printf("Key value [%s]\n", args.Service)
	fmt.Printf("Am I root? %v\n", isUserRoot())
}

func isUserRoot() bool {
	return os.Getuid() == 0
}

//func readRcConf(file_name string) (Content string, err error) {
//	readFile, err := os.Open(file_name)
//	if err != nil {
//		return "", err
//	}
//	defer readFile.Close()
//
//}
