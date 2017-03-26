package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"hash"
	"log"
	"os"
	"time"

	"github.com/glycerine/blake2b" // vendor https://github.com/dchest/blake2b"
	pb "github.com/glycerine/grpc-demo/streambigfile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

type client struct {
	hasher    hash.Hash
	nextChunk int64
}

func newClient() *client {
	h, err := blake2b.New(nil)
	panicOn(err)
	return &client{
		hasher: h,
	}
}

func (c *client) startNewFile() {
	c.hasher.Reset()
	c.nextChunk = 0
}

func (c *client) runSendFile(client pb.PeerClient, path string, data []byte, maxChunkSize int) error {
	p("client runSendFile() starting")

	c.startNewFile()
	stream, err := client.SendFile(context.Background())
	if err != nil {
		grpclog.Fatalf("%v.SendFile(_) = _, %v", client, err)
	}
	n := len(data)
	numChunk := n / maxChunkSize
	if n%maxChunkSize > 0 {
		numChunk++
	}
	nextByte := 0
	lastChunk := numChunk - 1
	p("client sees %v chunks of size ~ %v bytes", numChunk, maxChunkSize)
	for i := 0; i < numChunk; i++ {
		sendLen := intMin(maxChunkSize, n-(i*maxChunkSize))
		chunk := data[nextByte:(nextByte + sendLen)]
		nextByte += sendLen

		var nk pb.BigFileChunk
		nk.Filepath = path
		nk.SizeInBytes = int64(sendLen)
		nk.SendTime = uint64(time.Now().UnixNano())

		// checksums
		c.hasher.Write(chunk)
		nk.Blake2B = blake2bOfBytes(chunk)
		nk.Blake2BCumulative = []byte(c.hasher.Sum(nil))

		nk.Data = chunk
		nk.ChunkNumber = c.nextChunk
		c.nextChunk++
		nk.IsLastChunk = (i == lastChunk)

		//p("client, on chunk %v of '%s', checksum='%x', and cumul='%x'", nk.ChunkNumber, nk.Filepath, nk.Blake2B, nk.Blake2BCumulative)

		if err := stream.Send(&nk); err != nil {
			panic(err)
			//grpclog.Fatalf("%v.Send() = %v", stream, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		grpclog.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	compared := bytes.Compare(reply.WholeFileBlake2B, []byte(c.hasher.Sum(nil)))
	grpclog.Printf("Reply saw checksum: '%x' match: %v; size sent = %v, size received = %v", reply.WholeFileBlake2B, compared == 0, len(data), reply.SizeInBytes)
	if int64(len(data)) != reply.SizeInBytes {
		panic("size mismatch")
	}
	return nil
}

func blake2bOfBytes(by []byte) []byte {
	h, err := blake2b.New(nil)
	panicOn(err)
	h.Write(by)
	return []byte(h.Sum(nil))
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SequentialPayload(n int64) []byte {
	if n%8 != 0 {
		panic(fmt.Sprintf("n == %v must be a multiple of 8; has remainder %v", n, n%8))
	}

	k := uint64(n / 8)
	by := make([]byte, n)
	j := uint64(0)
	for i := uint64(0); i < k; i++ {
		j = i * 8
		binary.LittleEndian.PutUint64(by[j:j+8], j)
	}
	return by
}

const ProgramName = "client"

func main() {

	myflags := flag.NewFlagSet(ProgramName, flag.ContinueOnError)
	cfg := &ClientConfig{}
	cfg.DefineFlags(myflags)
	err := myflags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("%s command line flag error: '%s'", ProgramName, err)
	}

	err = cfg.ValidateConfig()
	if err != nil {
		log.Fatalf("%s command line flag error: '%s'", ProgramName, err)
	}

	var opts []grpc.DialOption
	if cfg.Tls {
		// with -tls flag
		cfg.setupTLS(&opts)
	} else {
		// default to ssh
		cfg.setupSSH(&opts)
	}

	serverAddr := fmt.Sprintf("%v:%v", cfg.ServerHost, cfg.ServerPort)

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPeerClient(conn)

	// SendFile
	c := newClient()
	data := []byte("hello peer, it is nice to meet you!!")
	err = c.runSendFile(client, "file1", data, 3)
	panicOn(err)

	data2 := []byte("second set of data should be kept separate!")
	err = c.runSendFile(client, "file2", data2, 3)
	panicOn(err)

	n := 1 << 29 // test with 512MB file. Works with up to 1MB or 2MB chunks.

	p("generating test data of size %v bytes", n)
	data3 := SequentialPayload(int64(n))
	//chunkSz := 1 << 22 // 4MB // GRPC will fail with EOF.
	chunkSz := 1 << 20
	t0 := time.Now()
	err = c.runSendFile(client, "bigfile3", data3, chunkSz)
	t1 := time.Now()
	panicOn(err)
	mb := float64(len(data3)) / float64(1<<20)
	elap := t1.Sub(t0)
	p("elap time to send %v MB was %v => %.03f MB/sec", mb, elap, mb/(float64(elap)/1e9))
}

func (cfg *ClientConfig) setupTLS(opts *[]grpc.DialOption) {
	var sn string
	if cfg.ServerHostOverride != "" {
		sn = cfg.ServerHostOverride
	}
	var creds credentials.TransportCredentials
	if cfg.CertPath != "" {
		var err error
		creds, err = credentials.NewClientTLSFromFile(cfg.CertPath, sn)
		if err != nil {
			grpclog.Fatalf("Failed to create TLS credentials %v", err)
		}
	} else {
		creds = credentials.NewClientTLSFromCert(nil, sn)
	}
	*opts = append(*opts, grpc.WithTransportCredentials(creds))
}

func (cfg *ClientConfig) setupSSH(opts *[]grpc.DialOption) {

	destAddr := fmt.Sprintf("%v:%v", cfg.ServerInternalHost, cfg.ServerInternalPort)

	dialer, err := clientSshMain(cfg.AllowNewServer, cfg.PrivateKeyPath, cfg.ClientKnownHostsPath, cfg.Username, cfg.ServerHost, destAddr, int64(cfg.ServerPort))
	panicOn(err)

	*opts = append(*opts, grpc.WithDialer(dialer))

	// have to do this too, since we are using an SSH tunnel
	// that grpc doesn't know about:
	*opts = append(*opts, grpc.WithInsecure())
}
