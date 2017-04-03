package api

import (
	"fmt"
	"time"
)

//go:generate msgp

type LocalGetSet interface {
	LocalGet(key []byte, includeValue bool) (ki *KeyInv, err error)
	LocalSet(ki *KeyInv) error
}

type Peerface interface {
	LocalGetSet
	BcastGet(key []byte, includeValue bool, timeout time.Duration, who string) (kis []*KeyInv, err error)
	GetLatest(key []byte, includeValue bool) (ki *KeyInv, err error)
}

// KeyInv supplies the keys and their
// peer location (Who) and their timestamps
// (When) while optionally (but not necessarily)
// providing their data Val.
//
// The includeValue flag in the
// calls below determines if we return the Val
// on Get calls. Val must always be provided
// on Set.
//
type KeyInv struct {
	Key     []byte
	Who     string
	When    time.Time
	Size    int64
	Blake2b []byte
	Val     []byte
}

func (ki *KeyInv) String() string {
	return fmt.Sprintf(`{Key:"%s", Who:"%s", When:"%s", Size:%v, Blake2b:"%x"}`,
		string(ki.Key), ki.Who, ki.When.UTC(), ki.Size, ki.Blake2b)
}

type BcastGetRequest struct {
	FromID string

	// Key specifies the key to query and return the value of.
	Key []byte

	// Who should be left empty to get all replies.
	// Otherwise only the peer whose name matches will reply.
	Who string

	// IncludeValue when false returns the timestamp and size without
	// the whole (big) value.
	IncludeValue bool

	ReplyGrpcHost  string
	ReplyGrpcXPort int
	ReplyGrpcIPort int
}

type BcastGetReply struct {
	FromID string
	Ki     *KeyInv
	Err    string
}

type BcastSetRequest struct {
	FromID string
	Ki     *KeyInv
}

type BcastSetReply struct {
	Err string
}
