// rc_conf.go - structs and methods for rc.conf file
package main

import (
	"bufio"
	"os"
)

const (
	DefaultEtcPath    = "/etc/"
	DefaultRcConfFile = "rc.conf"
	DefaultRcD        = "/etc/rc.d/"
	DefaultExampleRcD = "/usr/pkg/share/example/rc.d"
)

// ReadRcConfigFile reads a file and returns the content.
// In this case the file name should be /etc/rc.conf
func ReadRcConfigFile(filename string) ([]string, error) {
	var fileLines []string

	rcFileContent, err := os.Open(filename)
	if err != nil {
		return fileLines, err
	}
	fileScanner := bufio.NewScanner(rcFileContent)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	err = rcFileContent.Close()
	if err != nil {
		return nil, err
	}
	return fileLines, nil
}

func WriteRcConfFile() {}
