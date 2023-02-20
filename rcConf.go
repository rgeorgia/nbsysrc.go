// rcConf.go - structs and methods for rc.conf file
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

// ReadFile reads a file and returns the content.
// In this case the file name should be /etc/rc.conf
func ReadFile(filename string) ([]string, error) {
	var fileLines []string

	rcFileContent, err := os.Open(filename)
	if err != nil {
		return fileLines, err
	}
	defer rcFileContent.Close()
	fileScanner := bufio.NewScanner(rcFileContent)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	return fileLines, nil
}
