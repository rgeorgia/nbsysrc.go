// rc_conf.go - structs and methods for rc.conf file
package main

import (
	"bufio"
	"os"
)

const (
	DEFAULTETCPATH    = "/etc/"
	DEFAULTRCCONFFILE = "rc.conf"
	DEFAULTRCD        = "/etc/rc.d/"
	DEFAULTECAMPLERCD = "/usr/pkg/share/example/rc.d"
)

// ReadRcConfigFile reads a file and returns the content.
// In this case the file name should be /etc/rc.conf
func ReadRcConfigFile(filename string) ([]string, error) {
	var fileLines []string

	readFile, err := os.Open(filename)
	if err != nil {
		return fileLines, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	err = readFile.Close()
	if err != nil {
		return nil, err
	}
	return fileLines, nil
}

func WriteRcConfFile() {}

// Consider moving the following to its own package

func RcdDirectoryListing()    {}
func EtcRcdDirectoryListing() {}
