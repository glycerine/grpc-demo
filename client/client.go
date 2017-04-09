// gRPC client
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/glycerine/blake2b" // vendor https://github.com/dchest/blake2b"
	pb "github.com/glycerine/grpc-demo/streambigfile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"hash"
)

type client struct {
	hasher     hash.Hash
	nextChunk  int64
	peerClient pb.PeerClient
}

func newClient(conn *grpc.ClientConn) *client {
	h, err := blake2b.New(nil)
	panicOn(err)
	return &client{
		hasher:     h,
		peerClient: pb.NewPeerClient(conn),
	}
}

func (c *client) startNewFile() {
	c.hasher.Reset()
	c.nextChunk = 0
}

func (c *client) runSendFile(path string, data []byte, maxChunkSize int, isBcastSet bool, myID string) error {
	//p("client runSendFile(path='%s') starting", path)

	startOfRunSendFile := time.Now().UTC()
	startOfRunSendFileNanoUint64 := uint64(startOfRunSendFile.UnixNano())

	c.startNewFile()
	stream, err := c.peerClient.SendFile(context.Background())
	if err != nil {
		log.Fatalf("%v.SendFile(_) = _, %v", c.peerClient, err)
	}
	n := len(data)
	numChunk := n / maxChunkSize
	if n%maxChunkSize > 0 {
		numChunk++
	}
	nextByte := 0
	lastChunk := numChunk - 1
	//p("'%s' client sees %v chunks of size ~ %v bytes", path, numChunk, intMin(n, maxChunkSize))
	for i := 0; i < numChunk; i++ {
		sendLen := intMin(maxChunkSize, n-(i*maxChunkSize))
		chunk := data[nextByte:(nextByte + sendLen)]
		nextByte += sendLen

		var nk pb.BigFileChunk
		nk.IsBcastSet = isBcastSet
		nk.Filepath = path
		nk.SizeInBytes = int64(sendLen)
		nk.SendTime = uint64(time.Now().UnixNano())
		nk.OriginalStartSendTime = startOfRunSendFileNanoUint64

		// checksums
		c.hasher.Write(chunk)
		nk.Blake2B = blake2bOfBytes(chunk)
		nk.Blake2BCumulative = []byte(c.hasher.Sum(nil))

		nk.Data = chunk
		nk.ChunkNumber = c.nextChunk
		c.nextChunk++
		nk.IsLastChunk = (i == lastChunk)

		//		if nk.ChunkNumber%100 == 0 {
		//p("client, on chunk %v of '%s', checksum='%x', and cumul='%x'", nk.ChunkNumber, nk.Filepath, nk.Blake2B, nk.Blake2BCumulative)
		//		}

		if err := stream.Send(&nk); err != nil {
			// EOF?
			if err == io.EOF {
				if !nk.IsLastChunk {
					panic(fmt.Sprintf("'%s' we got io.EOF before "+
						"the last chunk! At: %v of %v", path, nk.ChunkNumber, numChunk))
				} else {
					break
				}
			}
			panic(err)
			//log.Fatalf("%v.Send() = %v", stream, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		// EOF ??
		log.Printf("%v.CloseAndRecv() got error %v, want %v. reply=%v", stream, err, nil, reply)
		return err
	}

	compared := bytes.Compare(reply.WholeFileBlake2B, []byte(c.hasher.Sum(nil)))
	log.Printf("%s client.runSendFile got from stream.CloseAndRecv() a Reply with checksum: '%x'; checksum matches the sent data: %v; size sent = %v, size received = %v. startOfRunSendFile='%v'.", myID, reply.WholeFileBlake2B, compared == 0, len(data), reply.SizeInBytes, startOfRunSendFile)

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

func (cfg *ClientConfig) ClientSendFile(path string, data []byte, isBcastSet bool, myID string) error {

	var opts []grpc.DialOption
	if cfg.SkipEncryption {
		opts = append(opts, grpc.WithInsecure())
	} else {
		if cfg.UseTLS {
			cfg.setupTLS(&opts)
		} else {
			cfg.setupSSH(&opts)
		}
	}

	serverAddr := fmt.Sprintf("%v:%v", cfg.ServerHost, cfg.ServerPort)

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return err
	}
	defer conn.Close()

	// SendFile
	c := newClient(conn)

	chunkSz := 1 << 20

	t0 := time.Now()
	err = c.runSendFile(path, data, chunkSz, isBcastSet, myID)
	t1 := time.Now()
	elap := t1.Sub(t0)
	if err != nil {
		log.Printf("%s ClientSendFile: c.runSendFile sees error '%v' after elap %v", myID, err, elap)
		return err
	}
	mb := float64(len(data)) / float64(1<<20)
	_ = mb
	_ = elap
	log.Printf("%s ClientSendFile: elap time to runSendFile(path='%s', len(data)=%v) on %v MB was %v => %.03f MB/sec", myID, path, len(data), mb, elap, mb/(float64(elap)/1e9))
	return nil
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
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
	} else {
		creds = credentials.NewClientTLSFromCert(nil, sn)
	}
	*opts = append(*opts, grpc.WithTransportCredentials(creds))
}

func (cfg *ClientConfig) setupSSH(opts *[]grpc.DialOption) {

	destAddr := fmt.Sprintf("%v:%v", cfg.ServerInternalHost, cfg.ServerInternalPort)

	dialer, err := clientSshMain(cfg.AllowNewServer, cfg.TestAllowOneshotConnect, cfg.PrivateKeyPath, cfg.ClientKnownHostsPath, cfg.Username, cfg.ServerHost, destAddr, int64(cfg.ServerPort))
	panicOn(err)

	*opts = append(*opts, grpc.WithDialer(dialer))

	// have to do this too, since we are using an SSH tunnel
	// that grpc doesn't know about:
	*opts = append(*opts, grpc.WithInsecure())
}
