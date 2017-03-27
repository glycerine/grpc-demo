# grpc-demo

One of the unusually nice features of the gRPC system is that it
supports streaming of messages. Although each individual message
in a stream cannot be more than 1-2MB (or else we get strange EOF
failures back from the gRPC library), we can readily transfer
any large number of 1MB chunks, resulting in a large file
transfer.

In short, this library demonstrates how to use gRPC http://www.grpc.io/ in
Golang (Go) to stream large files (of arbitrary size) between
two endpoints.

For security, we demonstrate two choices. We show how to do the
transfer using TLS and using an embedded SSH tunnel. The ssh client and server are
compiled into the binaries. You don't need to run a separate sshd on
your host or docker container.

Under flag `-ssh` to both client and server, we setup an SSH tunnel using https://github.com/glycerine/sshego
and 4096-bit RSA keys. Key exchange is done with `kexAlgoCurve25519SHA256`.

By default, we will use TLS.
We assume that the necessary certs for your server host have been
generated in the testdata/ directory, for example https://github.com/glycerine/grpc-demo/tree/master/testdata

For verifying the integrity of the transfer, we use Blake2B cryptographic hashes (https://blake2.net/) on both the inidividual chunks, and on the complete, cumulative transfer.

## installation

0) pre-requisite: you should have the latest protoc (protocol buffer compiler) installed and in your PATH. This example was developed with v3.0.0. https://github.com/google/protobuf/releases/tag/v3.0.0

1) get and compile and the source:

~~~
$ go get -t -u -v github.com/glycerine/grpc-demo/...
$ cd ~/go/src/github.com/glycerine/grpc-demo
$ make
~~~

## code

The protocol buffers definition is here https://github.com/glycerine/grpc-demo/blob/master/streambigfile/sbf.proto

See the https://github.com/glycerine/grpc-demo/blob/master/client/client.go file for the client side code.

See the https://github.com/glycerine/grpc-demo/blob/master/server/server.go file for the server side code.

## running the example code

For TLS, we use the pre-generated test certificate/public/private key
that come with the gRPC examples. You'll want to replace these
(in testdata/) with your own. There are plenty of tools on the
internet for generating your own certs.

For SSH, there are two preliminaries, which we will illustrate
below. First, the client needs an account with a public/private
key pair, and they need to be copied to the client's host.
Second, the first time (and only the first time) you run the ssh client
against a new server, you must provide the `-new` flag to the client.
This lets the client learn the server's identity, and to verify
that there has been no MITM thereafter.

# running the demo client and server with TLS

tls server
----------

at the console:

~~~

# server
$ cd ~/go/src/github.com/glycerine/grpc-demo
$ ./server/server

~~~

tls client
----------

Leave the server running, and open a separate terminal to start the client:

~~~

# client
$ cd ~/go/src/github.com/glycerine/grpc-demo
$ ./client/client

 client runSendFile() starting

 client sees 12 chunks of size ~ 3 bytes
2017/03/27 02:26:05 Reply saw checksum: 'e431e228e0eb7fe6618f47b9995e6a0a021aad88dafb8ec2cf24862f0cee99b7a7b8edbc8d87f7ef2ad0a47658e055bfdfe4312835a7e12b8067753681a81047' match: true; size sent = 36, size received = 36

 client runSendFile() starting

 client sees 15 chunks of size ~ 3 bytes
2017/03/27 02:26:05 Reply saw checksum: 'd89061b56be9db1df09874e12a7a25e884e6da3229a67687242561b4228267c34f18133b8dcb0882fd6ff06eea265be383b1bad8f3904b2263c015c96b07869e' match: true; size sent = 43, size received = 43

 generating test data of size 536870912 bytes

 client runSendFile() starting

 client sees 512 chunks of size ~ 1048576 bytes
2017/03/27 02:26:09 Reply saw checksum: '5255b783ed27354a28c035785dbcefc249212099e84096f5a561764af5690c72c31262e87b3b04c963a312f2df13982ec9eadd9be823fe7dba8c0a8b3af1cb1b' match: true; size sent = 536870912, size received = 536870912

 elap time to send 512 MB was 4.589398305s => 111.561 MB/sec
 $

~~~

# running the demo client and server with SSH

ssh server
-----------

First a user account with a public/private key pair must
be generated, then the server's host key must be
accepted and stored.

~~~

# server
$ cd ~/go/src/github.com/glycerine/grpc-demo
$ ./server/server -adduser $USER ## and follow the prompts:
...
Enter the email address for 'me' (for backups/recovery): me@devnull
...
Do you want to backup the passphase to email 'me@devnull'? [y/n]:n
...
Corresponding to 'me'/'me@devnull', enter the first and last name (e.g. 'John Q. Smith'). This helps identify the account during maintenance. First and last name: test acct
...

Your new RSA Private key is here (on host my-laptop.local):
/Users/me/.ssh/.sshego.sshd.db/users/me/id_rsa

Your new RSA Public key is here (on host my-laptop.local):
/Users/me/.ssh/.sshego.sshd.db/users/me/id_rsa.pub
$
$ # good, that generated an account. now start the server:
$ ./server/server -ssh

~~~

ssh client
------------

1) Leave the server running.

2) open a separate terminal

3) Copy /Users/me/.ssh/.sshego.sshd.db/users/me/id_rsa* from the
   server host to the client host (assuming they are distinct hosts).
   Keep the same directory structure. You can use the command line
   flags to change the default locations, but keep it simple the
   first time.

4) On the client host, start the client twice as follows (the 2x is only
necessary the first time):


~~~

# client
$ cd ~/go/src/github.com/glycerine/grpc-demo
$
$ ./client/client -ssh -new # must give -new the first time, to store the server's host key.
$
$ ./client/client -ssh # now that host key is stored, must start without -new.
$ # ... watch as three file transfers in a row happen, as the client/client.go code dictates

 client runSendFile() starting

 client sees 12 chunks of size ~ 3 bytes
2017/03/27 02:24:38 Reply saw checksum: 'e431e228e0eb7fe6618f47b9995e6a0a021aad88dafb8ec2cf24862f0cee99b7a7b8edbc8d87f7ef2ad0a47658e055bfdfe4312835a7e12b8067753681a81047' match: true; size sent = 36, size received = 36

 client runSendFile() starting

 client sees 15 chunks of size ~ 3 bytes
2017/03/27 02:24:38 Reply saw checksum: 'd89061b56be9db1df09874e12a7a25e884e6da3229a67687242561b4228267c34f18133b8dcb0882fd6ff06eea265be383b1bad8f3904b2263c015c96b07869e' match: true; size sent = 43, size received = 43

 generating test data of size 536870912 bytes

 client runSendFile() starting

 client sees 512 chunks of size ~ 1048576 bytes
2017/03/27 02:24:48 Reply saw checksum: '5255b783ed27354a28c035785dbcefc249212099e84096f5a561764af5690c72c31262e87b3b04c963a312f2df13982ec9eadd9be823fe7dba8c0a8b3af1cb1b' match: true; size sent = 536870912, size received = 536870912

 elap time to send 512 MB was 9.885738243s => 51.792 MB/sec
$

~~~

author
--------

Jason E. Aten, Ph.D.

license
--------

MIT License; see the https://github.com/glycerine/grpc-demo/blob/master/LICENSE file.
