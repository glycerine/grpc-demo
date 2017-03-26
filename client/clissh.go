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
	"net"
	//"strings"
	"time"

	tun "github.com/glycerine/sshego"
)

/*
 localhost-to-localhost on my laptop, over SSH with strong crypto:
 min and max observed:
 elap time to send 512 MB was 13.484078385s => 37.971 MB/sec
 elap time to send 512 MB was 10.078974477s => 50.799 MB/sec
*/
func clientSshMain(trustNewServer bool, rsaPrivateKeyPath, knownHostsPath, username, host, destHostPort string, serverExternalPort int64) (func(string, time.Duration) (net.Conn, error), error) {

	dc := tun.DialConfig{
		ClientKnownHostsPath: knownHostsPath,
		Mylogin:              username,
		RsaPath:              rsaPrivateKeyPath,
		TotpUrl:              "",
		Pw:                   "",
		Sshdhost:             host,
		Sshdport:             serverExternalPort,
		DownstreamHostPort:   destHostPort,
		Verbose:              false,

		TofuAddIfNotKnown: trustNewServer,
	}

	f := func(addr string, dur time.Duration) (net.Conn, error) {
		channelToTcpServer, _, err := dc.Dial()

		/*
			if strings.HasPrefix(err.Error(), "good: added "+
				"previously unknown sshd host") ||
				strings.Contains(err.Error(), "Re-run without -new") {
				// for prevention of MITM, enforce that
				// client to not do TOFU every time.
				if trustNewServer {
					dc.TofuAddIfNotKnown = false
					channelToTcpServer, _, err = dc.Dial()
				}
			}
		*/

		return channelToTcpServer, err
	}

	return f, nil
}
