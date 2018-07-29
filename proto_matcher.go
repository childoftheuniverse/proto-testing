package proto_testing

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/proto"
)

type protoMatcher struct {
	msg proto.Message
}

/*
ProtoEq creates a gomock matcher which can be used to determine whether two
protocol buffers have the same data content.
*/
func ProtoEq(m proto.Message) gomock.Matcher {
	return protoMatcher{msg: m}
}

func (p protoMatcher) Matches(msg interface{}) bool {
	m, ok := msg.(proto.Message)
	if !ok {
		return false
	}

	return proto.Equal(m, p.msg)
}

func (p protoMatcher) String() string {
	return fmt.Sprintf("is equal to %s", p.msg)
}
