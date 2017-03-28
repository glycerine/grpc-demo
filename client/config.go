package main

import (
	"flag"
	"fmt"
	"os"
)

type ClientConfig struct {
	Ssh                bool // use Tls if false
	AllowNewServer     bool // only give once to prevent MITM.
	CertPath           string
	KeyPath            string
	ServerHost         string // ip address
	ServerPort         int
	ServerInternalHost string // ip address
	ServerInternalPort int
	ServerHostOverride string

	Username             string
	PrivateKeyPath       string
	ClientKnownHostsPath string

	CpuProfilePath string
}

func (c *ClientConfig) DefineFlags(fs *flag.FlagSet) {
	fs.BoolVar(&c.AllowNewServer, "new", false, "allow new server host key to be recognized and stored in known-hosts")
	fs.BoolVar(&c.Ssh, "ssh", false, "Use SSH for security (default is TLS)")
	fs.StringVar(&c.CertPath, "cert_file", "testdata/server1.pem", "The TLS cert file")
	fs.StringVar(&c.KeyPath, "key_file", "testdata/server1.key", "The TLS key file")
	fs.StringVar(&c.ServerHost, "host", "127.0.0.1", "host IP address or name to connect to")
	fs.IntVar(&c.ServerPort, "port", 10000, "The exteral server port")
	fs.StringVar(&c.ServerInternalHost, "ihost", "127.0.0.1", "internal host IP address or name to connect to")
	fs.IntVar(&c.ServerInternalPort, "iport", 10001, "The internal server port")
	fs.StringVar(&c.ServerHostOverride, "server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")

	user := os.Getenv("USER")
	fs.StringVar(&c.Username, "user", user, "username for sshd login (default is $USER)")

	home := os.Getenv("HOME")
	fs.StringVar(&c.PrivateKeyPath, "key", home+"/.ssh/.sshego.sshd.db/users/"+user+"/id_rsa", "private key for sshd login")
	fs.StringVar(&c.ClientKnownHostsPath, "known-hosts", home+"/.ssh/.sshego.cli.known.hosts", "path to our own known-hosts file, for sshd login")

	fs.StringVar(&c.CpuProfilePath, "cpuprofile", "", "write cpu profile to file")
}

func (c *ClientConfig) ValidateConfig() error {

	if !c.Ssh {
		if c.KeyPath == "" {
			return fmt.Errorf("must provide -key_file under TLS")
		}
		if !fileExists(c.KeyPath) {
			return fmt.Errorf("-key_path '%s' does not exist", c.KeyPath)
		}

		if c.CertPath == "" {
			return fmt.Errorf("must provide -key_file under TLS")
		}
		if !fileExists(c.CertPath) {
			return fmt.Errorf("-cert_path '%s' does not exist", c.CertPath)
		}
	}

	return nil
}
