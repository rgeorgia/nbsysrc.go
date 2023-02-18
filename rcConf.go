// rcConf.go - structs and methods for rc.conf file
package main

const (
	DefaultEtcPath    = "/etc/"
	DefaultRcConfFile = "rc.conf"
	DefaultRcD        = "/etc/rc.d/"
	DefaultExampleRcD = "/usr/pkg/share/example/rc.d"
)

// ReadFile reads a file and returns the content.
// In this case the file name should be /etc/rc.conf
// Read the rc.conf file. if the line starts with # ignore it
// If the line contains an =, split it and stick it in a map
// rc.conf has the following format
/*
# If this is not set to YES, the system will drop into single-user mode.
#
rc_configured=YES

# Add local overrides below.
#
dhcpcd=YES
dhcpcd_flags="-qM bge0"
sshd=YES
ntpd=YES
lvm=YES
wscons=YES
dbus=YES
avahidaemon=YES
rpcbind=YES
famd=YES
hal=YES
zfs=YES
*/
func ReadFile(filename string) (string, error) {
	return filename, nil
}
