// main package. this script is like sysrc for Freebsd
package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/alexflint/go-arg"
	"log"
	"os"
	"slices"
	"strings"
)

// TODO: Add -f/--flag for flag inputs
var args struct {
	Service string `arg:"positional, required" help:"What you want to add to rc.conf i.e. samba=YES"`
}

type NbSysRcMeta struct {
	Ect        string `toml:"etc"`
	EctRcd     string `toml:"ect_rcd"`
	ExampleRcd string `toml:"example_rcd"`
	RcConf     string `toml:"rc_conf"`
}

func main() {
	arg.MustParse(&args)

	// Make sure there is some type of input
	if args.Service == "" {
		log.Fatal("Missing Input: i.e. samba=YES")
	}

	// Check that the input is valid
	if !validInput(args.Service) {
		log.Fatal("Invalid entry. i.e. samba=YES")
	}

	var config NbSysRcMeta
	_, err := toml.DecodeFile(".nbsysrc.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	// Read the rc.conf file
	var rcFile = fmt.Sprintf("%s/%s", config.Ect, config.RcConf)
	rcContent, err := ReadRcConfigFile(rcFile)
	if err != nil {
		log.Fatal("Error reading rc.conf", err)
	}

	// if the requested service is already set, exit
	if serviceIsSet(fmt.Sprintf("%s", args.Service), rcContent) {
		fmt.Printf("Service is already set\n%s", args.Service)
		os.Exit(0)
	}
	fmt.Println("DONE")
}

func validInput(arg string) bool {
	validValues := []string{"NO", "YES"}
	if len(strings.Split(arg, "=")) != 2 {
		return false
	}
	var value = strings.ToUpper(strings.Split(arg, "=")[1])

	if !slices.Contains(validValues, value) {
		return false
	}
	return true
}

func serviceIsSet(service string, rcContents []string) bool {
	for _, line := range rcContents {
		if strings.ToLower(strings.TrimSpace(service)) == strings.ToLower(strings.TrimSpace(line)) {
			fmt.Printf("Service exists: %s", service)
			return true
		}
	}
	return false
}

func isUserRoot() bool {
	return os.Getuid() == 0
}
