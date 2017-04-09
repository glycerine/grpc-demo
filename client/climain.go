package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"google.golang.org/grpc"
)

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
	cfg.SkipEncryption = true

	err := myflags.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("%s command line flag error: '%s'", ProgramName, err)
	}

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

	var opts []grpc.DialOption
	if cfg.UseTLS {
		cfg.setupTLS(&opts)
	} else if cfg.SkipEncryption {
		// no encryption
		opts = append(opts, grpc.WithInsecure())
		p("client configured to skip encryption.")
	} else {
		cfg.setupSSH(&opts)
	}

	serverAddr := fmt.Sprintf("%v:%v", cfg.ServerHost, cfg.ServerPort)

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// SendFile
	c := newClient(conn)
	isBcastSet := false
	myID := "test-client-0"
	data := []byte("hello peer, it is nice to meet you!!")
	err = c.runSendFile("file1", data, 3, isBcastSet, myID)
	panicOn(err)

	data2 := []byte("second set of data should be kept separate!")
	err = c.runSendFile("file2", data2, 3, isBcastSet, myID)
	panicOn(err)

	n := 1 << 29 // test with 512MB file. Works with up to 1MB or 2MB chunks.

	p("generating test data of size %v bytes", n)
	data3 := SequentialPayload(int64(n))
	//chunkSz := 1 << 22 // 4MB // GRPC will fail with EOF.
	chunkSz := 1 << 20

	c2done := make(chan struct{})

	overlap := false

	// overlap two sends to different paths
	go func() {
		if overlap {
			time.Sleep(10 * time.Millisecond)
			p("after 10msec of sleep, comencing bigfile3...")

			c2 := newClient(conn)
			t0 := time.Now()

			err = c2.runSendFile("bigfile3", data3, chunkSz, isBcastSet, myID)
			t1 := time.Now()
			panicOn(err)
			mb := float64(len(data3)) / float64(1<<20)
			elap := t1.Sub(t0)
			p("c2: elap time to send %v MB was %v => %.03f MB/sec", mb, elap, mb/(float64(elap)/1e9))
		}
		close(c2done)
	}()

	t0 := time.Now()
	err = c.runSendFile("bigfile4", data3, chunkSz, isBcastSet, myID)
	t1 := time.Now()
	panicOn(err)
	mb := float64(len(data3)) / float64(1<<20)
	elap := t1.Sub(t0)
	p("c: elap time to send %v MB was %v => %.03f MB/sec", mb, elap, mb/(float64(elap)/1e9))

	<-c2done
}
