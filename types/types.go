package types

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

var (
	// Generic not found error
	NotFoundErr = errors.New("not found")
)

type StubMessage struct {
	Type string
}

func (m *StubMessage) Reset() { *m = StubMessage{} }

func (m *StubMessage) String() string { return proto.CompactTextString(m) }

func (*StubMessage) ProtoMessage() {}
