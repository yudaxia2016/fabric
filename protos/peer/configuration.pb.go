// Code generated by protoc-gen-go.
// source: peer/configuration.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// AnchorPeers simply represents list of anchor peers which is used in ConfigurationItem
type AnchorPeers struct {
	AnchorPeers []*AnchorPeer `protobuf:"bytes,1,rep,name=anchor_peers,json=anchorPeers" json:"anchor_peers,omitempty"`
}

func (m *AnchorPeers) Reset()                    { *m = AnchorPeers{} }
func (m *AnchorPeers) String() string            { return proto.CompactTextString(m) }
func (*AnchorPeers) ProtoMessage()               {}
func (*AnchorPeers) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *AnchorPeers) GetAnchorPeers() []*AnchorPeer {
	if m != nil {
		return m.AnchorPeers
	}
	return nil
}

// AnchorPeer message structure which provides information about anchor peer, it includes host name,
// port number and peer certificate.
type AnchorPeer struct {
	// DNS host name of the anchor peer
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// The port number
	Port int32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	// SSL certificate to be used to maintain mutual TLS
	// connection with anchor peer
	Cert []byte `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (m *AnchorPeer) Reset()                    { *m = AnchorPeer{} }
func (m *AnchorPeer) String() string            { return proto.CompactTextString(m) }
func (*AnchorPeer) ProtoMessage()               {}
func (*AnchorPeer) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func init() {
	proto.RegisterType((*AnchorPeers)(nil), "protos.AnchorPeers")
	proto.RegisterType((*AnchorPeer)(nil), "protos.AnchorPeer")
}

func init() { proto.RegisterFile("peer/configuration.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x8f, 0x31, 0x4f, 0xc6, 0x20,
	0x10, 0x86, 0x83, 0x55, 0x13, 0x69, 0x27, 0x26, 0x46, 0xd2, 0x09, 0x63, 0x52, 0x12, 0x8d, 0x3f,
	0x40, 0xe3, 0xe0, 0x68, 0x18, 0x5d, 0x0c, 0xc5, 0x6b, 0x21, 0xd1, 0x1e, 0x39, 0xe8, 0xf0, 0xfd,
	0xfb, 0x2f, 0xd0, 0xa1, 0x13, 0x0f, 0xef, 0xdd, 0x93, 0xbc, 0xc7, 0x65, 0x02, 0x20, 0xe3, 0x71,
	0x5b, 0xe2, 0xba, 0x93, 0x2b, 0x11, 0xb7, 0x29, 0x11, 0x16, 0x14, 0xf7, 0xed, 0xc9, 0xe3, 0x07,
	0xef, 0xdf, 0x36, 0x1f, 0x90, 0xbe, 0x00, 0x28, 0x8b, 0x57, 0x3e, 0xb8, 0xf6, 0xfd, 0xa9, 0x66,
	0x96, 0x4c, 0x75, 0xba, 0x7f, 0x16, 0x87, 0x94, 0xa7, 0x73, 0xd5, 0xf6, 0xee, 0xd4, 0xc6, 0x4f,
	0xce, 0xcf, 0x91, 0x10, 0xfc, 0x36, 0x60, 0x2e, 0x92, 0x29, 0xa6, 0x1f, 0x6c, 0xe3, 0x9a, 0x25,
	0xa4, 0x22, 0x6f, 0x14, 0xd3, 0x77, 0xb6, 0x71, 0xcd, 0x3c, 0x50, 0x91, 0x9d, 0x62, 0x7a, 0xb0,
	0x8d, 0xdf, 0x9f, 0xbe, 0x1f, 0xd7, 0x58, 0xc2, 0x3e, 0x4f, 0x1e, 0xff, 0x4d, 0xb8, 0x24, 0xa0,
	0x3f, 0xf8, 0x5d, 0x81, 0xcc, 0xe2, 0x66, 0x8a, 0xde, 0x1c, 0x4d, 0x4c, 0xad, 0x37, 0x1f, 0x47,
	0xbc, 0x5c, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x2c, 0x85, 0xf6, 0xe7, 0x00, 0x00, 0x00,
}
