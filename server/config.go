package main

import (
	"flag"
	"fmt"
	"net"
)

type ServerConfig struct {
	Host string // ip address

	// by default we use TLS
	Ssh bool

	CertPath string
	KeyPath  string

	ExternalLsnPort int
	InternalLsnPort int
}

func (c *ServerConfig) DefineFlags(fs *flag.FlagSet) {
	fs.BoolVar(&c.Ssh, "ssh", false, "Use SSH instead of the default TLS.")
	fs.StringVar(&c.CertPath, "cert_file", "testdata/server1.pem", "The TLS cert file")
	fs.StringVar(&c.KeyPath, "key_file", "testdata/server1.key", "The TLS key file")
	fs.StringVar(&c.Host, "host", "127.0.0.1", "host IP address or name to bind")
	fs.IntVar(&c.ExternalLsnPort, "xport", 10000, "The exteral server port")
	fs.IntVar(&c.InternalLsnPort, "iport", 10001, "The internal server port")
}

func (c *ServerConfig) ValidateConfig() error {

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

	if c.Ssh {
		lsn, err := net.Listen("tcp", fmt.Sprintf(":%v", c.InternalLsnPort))
		if err != nil {
			return fmt.Errorf("internal port %v already bound", c.InternalLsnPort)
		}
		lsn.Close()
	}

	lsnX, err := net.Listen("tcp", fmt.Sprintf(":%v", c.ExternalLsnPort))
	if err != nil {
		return fmt.Errorf("external port %v already bound", c.ExternalLsnPort)
	}
	lsnX.Close()

	return nil
}
