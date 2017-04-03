// gRPC server
package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime/pprof"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"

	"github.com/glycerine/grpc-demo/api"
	pb "github.com/glycerine/grpc-demo/streambigfile"
)

const ProgramName = "server"

func main() {

	myflags := flag.NewFlagSet(ProgramName, flag.ExitOnError)
	cfg := &ServerConfig{
		SkipEncryption: true,
	}
	cfg.DefineFlags(myflags)

	sshegoCfg := setupSshFlags(myflags)

	args := os.Args[1:]
	err := myflags.Parse(args)

	if cfg.CpuProfilePath != "" {
		f, err := os.Create(cfg.CpuProfilePath)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	err = cfg.ValidateConfig()
	if err != nil {
		log.Fatalf("%s command line flag error: '%s'", ProgramName, err)
	}

	var gRpcBindPort int
	var gRpcHost string
	if cfg.UseTLS {
		// use TLS
		gRpcBindPort = cfg.ExternalLsnPort
		gRpcHost = cfg.Host

		p("gRPC with TLS listening on %v:%v", gRpcHost, gRpcBindPort)

	} else {
		// SSH will take the external, gRPC will take the internal.
		gRpcBindPort = cfg.InternalLsnPort
		gRpcHost = "127.0.0.1" // local only, behind the SSHD

		p("external SSHd listening on %v:%v, internal gRPC service listening on 127.0.0.1:%v", cfg.Host, cfg.ExternalLsnPort, cfg.InternalLsnPort)

	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", gRpcHost, gRpcBindPort))
	if err != nil {
		utclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	if cfg.UseTLS {
		// use TLS
		creds, err := credentials.NewServerTLSFromFile(cfg.CertPath, cfg.KeyPath)
		if err != nil {
			utclog.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	} else {
		// use SSH
		err = serverSshMain(sshegoCfg, cfg.Host,
			cfg.ExternalLsnPort, cfg.InternalLsnPort)
		panicOn(err)
	}

	peer := NewPeerMemoryOnly()

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPeerServer(grpcServer, NewPeerServerClass(peer, cfg))
	grpcServer.Serve(lis)
}

func NewPeerMemoryOnly() *PeerMemoryOnly {
	return &PeerMemoryOnly{}
}

type PeerMemoryOnly struct{}

func (peer *PeerMemoryOnly) LocalGet(key []byte, includeValue bool) (ki *api.KeyInv, err error) {
	return nil, fmt.Errorf("unimplimeneted")
}

func (peer *PeerMemoryOnly) LocalSet(ki *api.KeyInv) error {
	return nil
}
