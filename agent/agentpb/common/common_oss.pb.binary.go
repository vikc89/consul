// +build !consulent

// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: agent/agentpb/common/common_oss.proto

package common

import (
	"github.com/golang/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *EnterpriseMeta) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *EnterpriseMeta) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
