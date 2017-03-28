package main

// code from github.com/glycerine/sshego is
// used under the following MIT license.
/*
The MIT License (MIT)

Portions Copyright (c) 2016 Jason E. Aten, Ph.D.
Portions Copyright (c) 2015 Rackspace, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import (
	"flag"
	"fmt"
	"log"
	"os"

	tun "github.com/glycerine/sshego"
)

func serverSshMain(args []string, host string, securedPort, targetPort int) error {

	myflags := flag.NewFlagSet(ProgramName, flag.ContinueOnError)
	cfg := tun.NewSshegoConfig()
	cfg.DefineFlags(myflags)

	// ignore errors here
	myflags.Parse(args)

	if cfg.ShowVersion {
		fmt.Printf("\n%v\n", tun.SourceVersion())
		os.Exit(0)
	}

	cfg.EmbeddedSSHd.Addr = fmt.Sprintf("%v:%v", host, securedPort)
	cfg.SkipPassphrase = true
	cfg.SkipTOTP = true
	cfg.SkipRSA = false

	cfg.DirectTcp = false
	cfg.RemoteToLocal.Listen.Addr = ""
	cfg.LocalToRemote.Listen.Addr = ""

	err := cfg.ValidateConfig()
	if err != nil {
		log.Fatalf("%s command line flag error: '%s'", ProgramName, err)
	}
	//p("cfg = %#v", cfg)
	h, err := tun.NewKnownHosts(cfg.ClientKnownHostsPath, tun.KHJson)
	panicOn(err)
	cfg.KnownHosts = h

	if cfg.AddUser != "" {
		tun.AddUserAndExit(cfg)
	}

	if cfg.DelUser != "" {
		tun.DelUserAndExit(cfg)
	}

	log.Printf("grpc-demo/server/ssh.go is starting -esshd with addr: %s", cfg.EmbeddedSSHd.Addr)
	err = cfg.EmbeddedSSHd.ParseAddr()
	if err != nil {
		p("grpc-demo/server/ssh.go cfg.EmbeddedSSHd.ParseAddr() error = '%s'", err)
		return err
	}
	cfg.NewEsshd()
	p("grpc-demo/server/ssh.go about to call cfg.Esshd.Start()")
	go cfg.Esshd.Start()

	return nil
}
