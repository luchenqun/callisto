package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	common "github.com/forbole/bdjuno/v4/types/common"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ActionType defines the operations you can execute in greenfield storage network
type ActionType int32

const (
	ACTION_UNSPECIFIED           ActionType = 0
	ACTION_UPDATE_BUCKET_INFO    ActionType = 1
	ACTION_DELETE_BUCKET         ActionType = 2
	ACTION_CREATE_OBJECT         ActionType = 3
	ACTION_DELETE_OBJECT         ActionType = 4
	ACTION_COPY_OBJECT           ActionType = 5
	ACTION_GET_OBJECT            ActionType = 6
	ACTION_EXECUTE_OBJECT        ActionType = 7
	ACTION_LIST_OBJECT           ActionType = 8
	ACTION_UPDATE_GROUP_MEMBER   ActionType = 9
	ACTION_DELETE_GROUP          ActionType = 10
	ACTION_UPDATE_OBJECT_INFO    ActionType = 11
	ACTION_UPDATE_GROUP_EXTRA    ActionType = 12
	ACTION_UPDATE_GROUP_INFO     ActionType = 13
	ACTION_UPDATE_OBJECT_CONTENT ActionType = 14
	ACTION_TYPE_ALL              ActionType = 99
)

var ActionType_name = map[int32]string{
	0:  "ACTION_UNSPECIFIED",
	1:  "ACTION_UPDATE_BUCKET_INFO",
	2:  "ACTION_DELETE_BUCKET",
	3:  "ACTION_CREATE_OBJECT",
	4:  "ACTION_DELETE_OBJECT",
	5:  "ACTION_COPY_OBJECT",
	6:  "ACTION_GET_OBJECT",
	7:  "ACTION_EXECUTE_OBJECT",
	8:  "ACTION_LIST_OBJECT",
	9:  "ACTION_UPDATE_GROUP_MEMBER",
	10: "ACTION_DELETE_GROUP",
	11: "ACTION_UPDATE_OBJECT_INFO",
	12: "ACTION_UPDATE_GROUP_EXTRA",
	13: "ACTION_UPDATE_GROUP_INFO",
	14: "ACTION_UPDATE_OBJECT_CONTENT",
	99: "ACTION_TYPE_ALL",
}

var ActionType_value = map[string]int32{
	"ACTION_UNSPECIFIED":           0,
	"ACTION_UPDATE_BUCKET_INFO":    1,
	"ACTION_DELETE_BUCKET":         2,
	"ACTION_CREATE_OBJECT":         3,
	"ACTION_DELETE_OBJECT":         4,
	"ACTION_COPY_OBJECT":           5,
	"ACTION_GET_OBJECT":            6,
	"ACTION_EXECUTE_OBJECT":        7,
	"ACTION_LIST_OBJECT":           8,
	"ACTION_UPDATE_GROUP_MEMBER":   9,
	"ACTION_DELETE_GROUP":          10,
	"ACTION_UPDATE_OBJECT_INFO":    11,
	"ACTION_UPDATE_GROUP_EXTRA":    12,
	"ACTION_UPDATE_GROUP_INFO":     13,
	"ACTION_UPDATE_OBJECT_CONTENT": 14,
	"ACTION_TYPE_ALL":              99,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}

func (ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{0}
}

// Effect define the effect of the operation permission, include Allow or deny
type Effect int32

const (
	EFFECT_UNSPECIFIED Effect = 0
	EFFECT_ALLOW       Effect = 1
	EFFECT_DENY        Effect = 2
)

var Effect_name = map[int32]string{
	0: "EFFECT_UNSPECIFIED",
	1: "EFFECT_ALLOW",
	2: "EFFECT_DENY",
}

var Effect_value = map[string]int32{
	"EFFECT_UNSPECIFIED": 0,
	"EFFECT_ALLOW":       1,
	"EFFECT_DENY":        2,
}

func (x Effect) String() string {
	return proto.EnumName(Effect_name, int32(x))
}

func (Effect) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{1}
}

// PrincipalType refers to the identity type of system users or entities.
// In Greenfield, it usually refers to accounts or groups.
type PrincipalType int32

const (
	PRINCIPAL_TYPE_UNSPECIFIED  PrincipalType = 0
	PRINCIPAL_TYPE_GNFD_ACCOUNT PrincipalType = 1
	PRINCIPAL_TYPE_GNFD_GROUP   PrincipalType = 2
)

var PrincipalType_name = map[int32]string{
	0: "PRINCIPAL_TYPE_UNSPECIFIED",
	1: "PRINCIPAL_TYPE_GNFD_ACCOUNT",
	2: "PRINCIPAL_TYPE_GNFD_GROUP",
}

var PrincipalType_value = map[string]int32{
	"PRINCIPAL_TYPE_UNSPECIFIED":  0,
	"PRINCIPAL_TYPE_GNFD_ACCOUNT": 1,
	"PRINCIPAL_TYPE_GNFD_GROUP":   2,
}

func (x PrincipalType) String() string {
	return proto.EnumName(PrincipalType_name, int32(x))
}

func (PrincipalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{2}
}

type Statement struct {
	// effect define the impact of permissions, which can be Allow/Deny
	Effect Effect `protobuf:"varint,1,opt,name=effect,proto3,enum=greenfield.permission.Effect" json:"effect,omitempty"`
	// action_type define the operation type you can act. greenfield defines a set of permission
	// that you can specify in a permissionInfo. see ActionType enum for detail.
	Actions []ActionType `protobuf:"varint,2,rep,packed,name=actions,proto3,enum=greenfield.permission.ActionType" json:"actions,omitempty"`
	// CAN ONLY USED IN bucket level. Support fuzzy match and limit to 5.
	// The sub-resource name must comply with the standard specified in the greenfield/types/grn.go file for Greenfield resource names.
	// If the sub-resources include 'grn:o:{bucket_name}/*' in the statement, it indicates that specific permissions is granted to all objects within the specified bucket.
	// If the sub-resources include 'grn:o:{bucket_name}/test_*' in the statement, it indicates that specific permissions is granted to all objects with the `test_` prefix within the specified bucket.
	// If the sub-resources is empty, when you need to operate(excluding CreateObject) a specified subresource, it will be denied because it cannot match any subresource.
	Resources []string `protobuf:"bytes,3,rep,name=resources,proto3" json:"resources,omitempty"`
	// expiration_time defines how long the permission is valid. If not explicitly specified, it means it will not expire.
	ExpirationTime *time.Time `protobuf:"bytes,4,opt,name=expiration_time,json=expirationTime,proto3,stdtime" json:"expiration_time,omitempty"`
	// limit_size defines the total data size that is allowed to operate. If not explicitly specified, it means it will not limit.
	LimitSize *common.UInt64Value `protobuf:"bytes,5,opt,name=limit_size,json=limitSize,proto3" json:"limit_size,omitempty"`
}

func (m *Statement) Reset()         { *m = Statement{} }
func (m *Statement) String() string { return proto.CompactTextString(m) }
func (*Statement) ProtoMessage()    {}
func (*Statement) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{0}
}
func (m *Statement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Statement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Statement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Statement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statement.Merge(m, src)
}
func (m *Statement) XXX_Size() int {
	return m.Size()
}
func (m *Statement) XXX_DiscardUnknown() {
	xxx_messageInfo_Statement.DiscardUnknown(m)
}

var xxx_messageInfo_Statement proto.InternalMessageInfo

func (m *Statement) GetEffect() Effect {
	if m != nil {
		return m.Effect
	}
	return EFFECT_UNSPECIFIED
}

func (m *Statement) GetActions() []ActionType {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *Statement) GetResources() []string {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Statement) GetExpirationTime() *time.Time {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *Statement) GetLimitSize() *common.UInt64Value {
	if m != nil {
		return m.LimitSize
	}
	return nil
}

// Principal define the roles that can be grant permissions to. Currently, it can be account or group.
type Principal struct {
	Type PrincipalType `protobuf:"varint,1,opt,name=type,proto3,enum=greenfield.permission.PrincipalType" json:"type,omitempty"`
	// When the type is an account, its value is sdk.AccAddress().String();
	// when the type is a group, its value is math.Uint().String()
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_33a4d646aee30990, []int{1}
}
func (m *Principal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return m.Size()
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

func (m *Principal) GetType() PrincipalType {
	if m != nil {
		return m.Type
	}
	return PRINCIPAL_TYPE_UNSPECIFIED
}

func (m *Principal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("greenfield.permission.ActionType", ActionType_name, ActionType_value)
	proto.RegisterEnum("greenfield.permission.Effect", Effect_name, Effect_value)
	proto.RegisterEnum("greenfield.permission.PrincipalType", PrincipalType_name, PrincipalType_value)
	proto.RegisterType((*Statement)(nil), "greenfield.permission.Statement")
	proto.RegisterType((*Principal)(nil), "greenfield.permission.Principal")
}

func init() {
	proto.RegisterFile("greenfield/permission/common.proto", fileDescriptor_33a4d646aee30990)
}

var fileDescriptor_33a4d646aee30990 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xcb, 0x6e, 0xda, 0x4c,
	0x14, 0xc6, 0x40, 0x48, 0x38, 0x49, 0x88, 0xff, 0x49, 0xf2, 0xd7, 0xa1, 0x89, 0xa1, 0x51, 0x17,
	0x28, 0x52, 0xb1, 0x4a, 0x2f, 0xaa, 0xd4, 0x95, 0x31, 0x03, 0x72, 0x4b, 0x6c, 0xcb, 0x98, 0x36,
	0x69, 0x17, 0x16, 0xa1, 0x03, 0xb5, 0x84, 0xb1, 0x65, 0x9b, 0x34, 0xc9, 0x13, 0x74, 0x99, 0x77,
	0xe8, 0x13, 0xf4, 0x21, 0x2a, 0x65, 0x99, 0x65, 0x57, 0x6d, 0x95, 0xbc, 0x48, 0x65, 0x7b, 0xb8,
	0x35, 0x64, 0x83, 0x98, 0xef, 0x36, 0xe7, 0x9c, 0x99, 0x31, 0xec, 0xf7, 0x3d, 0x42, 0x86, 0x3d,
	0x8b, 0x0c, 0x3e, 0x09, 0x2e, 0xf1, 0x6c, 0xcb, 0xf7, 0x2d, 0x67, 0x28, 0x74, 0x1d, 0xdb, 0x76,
	0x86, 0x65, 0xd7, 0x73, 0x02, 0x07, 0x6d, 0x4f, 0x35, 0xe5, 0xa9, 0x26, 0xbf, 0xd3, 0x75, 0x7c,
	0xdb, 0xf1, 0xcd, 0x48, 0x24, 0xc4, 0x8b, 0xd8, 0x91, 0xdf, 0xea, 0x3b, 0x7d, 0x27, 0xc6, 0xc3,
	0x7f, 0x14, 0x2d, 0xf4, 0x1d, 0xa7, 0x3f, 0x20, 0x42, 0xb4, 0x3a, 0x19, 0xf5, 0x84, 0xc0, 0xb2,
	0x89, 0x1f, 0x74, 0x6c, 0x77, 0x22, 0x98, 0x16, 0x13, 0x57, 0x20, 0x7c, 0xf1, 0x3a, 0xae, 0x4b,
	0xbc, 0x58, 0xb0, 0xff, 0x3d, 0x09, 0xd9, 0x56, 0xd0, 0x09, 0x88, 0x4d, 0x86, 0x01, 0x7a, 0x01,
	0x19, 0xd2, 0xeb, 0x91, 0x6e, 0xc0, 0x31, 0x45, 0xa6, 0x94, 0xab, 0xec, 0x95, 0x17, 0x16, 0x5a,
	0xc6, 0x91, 0x48, 0xa7, 0x62, 0xf4, 0x1a, 0x96, 0x3b, 0xdd, 0xc0, 0x72, 0x86, 0x3e, 0x97, 0x2c,
	0xa6, 0x4a, 0xb9, 0xca, 0xa3, 0x7b, 0x7c, 0x62, 0xa4, 0x32, 0xce, 0x5d, 0xa2, 0x8f, 0x1d, 0x68,
	0x17, 0xb2, 0x1e, 0xf1, 0x9d, 0x91, 0xd7, 0x25, 0x3e, 0x97, 0x2a, 0xa6, 0x4a, 0x59, 0x7d, 0x0a,
	0xa0, 0x43, 0xd8, 0x20, 0x67, 0xae, 0xe5, 0x75, 0x42, 0xb1, 0x19, 0xb6, 0xc7, 0xa5, 0x8b, 0x4c,
	0x69, 0xb5, 0x92, 0x2f, 0xc7, 0xbd, 0x97, 0xc7, 0xbd, 0x97, 0x8d, 0x71, 0xef, 0xd5, 0x95, 0xab,
	0x5f, 0x05, 0xe6, 0xf2, 0x77, 0x81, 0xd1, 0x73, 0x53, 0x73, 0x48, 0x23, 0x09, 0x60, 0x60, 0xd9,
	0x56, 0x60, 0xfa, 0xd6, 0x05, 0xe1, 0x96, 0xa2, 0x24, 0x7e, 0xb6, 0x58, 0x7a, 0x4c, 0x6d, 0x79,
	0x18, 0xbc, 0x7c, 0xfe, 0xae, 0x33, 0x18, 0x91, 0x6a, 0x3a, 0x4c, 0xd3, 0xb3, 0x91, 0xaf, 0x65,
	0x5d, 0x90, 0xfd, 0x8f, 0x90, 0xd5, 0x3c, 0x6b, 0xd8, 0xb5, 0xdc, 0xce, 0x00, 0xbd, 0x82, 0x74,
	0x70, 0xee, 0x12, 0x3a, 0xb0, 0xc7, 0xf7, 0x34, 0x3e, 0xd1, 0x47, 0xbd, 0x47, 0x0e, 0xb4, 0x05,
	0x4b, 0xa7, 0xe1, 0x06, 0x5c, 0xb2, 0xc8, 0x94, 0xb2, 0x7a, 0xbc, 0x38, 0xf8, 0x91, 0x02, 0x98,
	0x8e, 0x09, 0xfd, 0x0f, 0x48, 0x94, 0x0c, 0x59, 0x55, 0xcc, 0xb6, 0xd2, 0xd2, 0xb0, 0x24, 0xd7,
	0x65, 0x5c, 0x63, 0x13, 0x68, 0x0f, 0x76, 0xc6, 0xb8, 0x56, 0x13, 0x0d, 0x6c, 0x56, 0xdb, 0xd2,
	0x5b, 0x6c, 0x98, 0xb2, 0x52, 0x57, 0x59, 0x06, 0x71, 0xb0, 0x45, 0xe9, 0x1a, 0x6e, 0xe2, 0x09,
	0xcd, 0x26, 0x67, 0x18, 0x49, 0xc7, 0xa1, 0x51, 0xad, 0xbe, 0xc1, 0x92, 0xc1, 0xa6, 0xee, 0x7a,
	0x28, 0x93, 0x9e, 0x29, 0x42, 0x52, 0xb5, 0xe3, 0x31, 0xbe, 0x84, 0xb6, 0xe1, 0x3f, 0x8a, 0x37,
	0xb0, 0x31, 0x86, 0x33, 0x68, 0x07, 0xb6, 0x29, 0x8c, 0x8f, 0xb0, 0xd4, 0x9e, 0x26, 0x2d, 0xcf,
	0x24, 0x35, 0xe5, 0xd6, 0xc4, 0xb2, 0x82, 0x78, 0xc8, 0xcf, 0xb7, 0xd3, 0xd0, 0xd5, 0xb6, 0x66,
	0x1e, 0xe2, 0xc3, 0x2a, 0xd6, 0xd9, 0x2c, 0x7a, 0x00, 0x9b, 0xf3, 0xb5, 0x45, 0x3c, 0x0b, 0x77,
	0xe7, 0x10, 0x47, 0xc6, 0x73, 0x58, 0xbd, 0x4b, 0xc7, 0xb9, 0xf8, 0xc8, 0xd0, 0x45, 0x76, 0x0d,
	0xed, 0x02, 0xb7, 0x88, 0x8e, 0xcc, 0xeb, 0xa8, 0x08, 0xbb, 0x0b, 0xb3, 0x25, 0x55, 0x31, 0xb0,
	0x62, 0xb0, 0x39, 0xb4, 0x09, 0x1b, 0x54, 0x61, 0x1c, 0x6b, 0xd8, 0x14, 0x9b, 0x4d, 0xb6, 0x9b,
	0x4f, 0x7f, 0xfd, 0xc6, 0x27, 0x0e, 0x64, 0xc8, 0xc4, 0xaf, 0x24, 0xec, 0x19, 0xd7, 0xeb, 0xa1,
	0x71, 0xfe, 0x08, 0x59, 0x58, 0xa3, 0xb8, 0xd8, 0x6c, 0xaa, 0xef, 0x59, 0x06, 0x6d, 0xc0, 0x2a,
	0x45, 0x6a, 0x58, 0x39, 0x66, 0x93, 0x34, 0x6a, 0x04, 0xeb, 0x73, 0xf7, 0x27, 0x9c, 0x96, 0xa6,
	0xcb, 0x8a, 0x24, 0x6b, 0x62, 0x33, 0xde, 0x79, 0x3e, 0xb9, 0x00, 0x0f, 0xff, 0xe1, 0x1b, 0x4a,
	0xbd, 0x66, 0x8a, 0x92, 0xa4, 0xb6, 0x15, 0x83, 0x65, 0xc2, 0xb1, 0x2c, 0x12, 0xc4, 0x43, 0xa5,
	0xdb, 0x56, 0x1b, 0x57, 0x37, 0x3c, 0x73, 0x7d, 0xc3, 0x33, 0x7f, 0x6e, 0x78, 0xe6, 0xf2, 0x96,
	0x4f, 0x5c, 0xdf, 0xf2, 0x89, 0x9f, 0xb7, 0x7c, 0xe2, 0xc3, 0x93, 0xbe, 0x15, 0x7c, 0x1e, 0x9d,
	0x84, 0x8f, 0x45, 0x20, 0xa7, 0xb6, 0xe3, 0xd3, 0xdf, 0xd3, 0xa7, 0x15, 0xe1, 0x6c, 0xf6, 0xa3,
	0x17, 0xde, 0x73, 0xff, 0x24, 0x13, 0x3d, 0xd1, 0x67, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x4e,
	0xe1, 0x32, 0x33, 0x1a, 0x05, 0x00, 0x00,
}

func (m *Statement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Statement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Statement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LimitSize != nil {
		{
			size, err := m.LimitSize.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.ExpirationTime != nil {
		n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.ExpirationTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintCommon(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Resources) > 0 {
		for iNdEx := len(m.Resources) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Resources[iNdEx])
			copy(dAtA[i:], m.Resources[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.Resources[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Actions) > 0 {
		dAtA4 := make([]byte, len(m.Actions)*10)
		var j3 int
		for _, num := range m.Actions {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintCommon(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x12
	}
	if m.Effect != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Effect))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Principal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Principal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Principal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Statement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Effect != 0 {
		n += 1 + sovCommon(uint64(m.Effect))
	}
	if len(m.Actions) > 0 {
		l = 0
		for _, e := range m.Actions {
			l += sovCommon(uint64(e))
		}
		n += 1 + sovCommon(uint64(l)) + l
	}
	if len(m.Resources) > 0 {
		for _, s := range m.Resources {
			l = len(s)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if m.ExpirationTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.ExpirationTime)
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.LimitSize != nil {
		l = m.LimitSize.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func (m *Principal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCommon(uint64(m.Type))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Statement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Statement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Statement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Effect", wireType)
			}
			m.Effect = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Effect |= Effect(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v ActionType
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= ActionType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Actions = append(m.Actions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCommon
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCommon
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Actions) == 0 {
					m.Actions = make([]ActionType, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v ActionType
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= ActionType(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Actions = append(m.Actions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ExpirationTime == nil {
				m.ExpirationTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.ExpirationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitSize", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LimitSize == nil {
				m.LimitSize = &common.UInt64Value{}
			}
			if err := m.LimitSize.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Principal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Principal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Principal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= PrincipalType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
