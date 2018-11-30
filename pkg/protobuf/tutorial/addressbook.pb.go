// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

package tutorial

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}

var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

type Person struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id    int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// Types that are valid to be assigned to Contrace:
	//	*Person_Str
	//	*Person_In
	Contrace             isPerson_Contrace     `protobuf_oneof:"contrace"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated          *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type isPerson_Contrace interface {
	isPerson_Contrace()
}

type Person_Str struct {
	Str string `protobuf:"bytes,6,opt,name=str,proto3,oneof"`
}

type Person_In struct {
	In int64 `protobuf:"varint,7,opt,name=in,proto3,oneof"`
}

func (*Person_Str) isPerson_Contrace() {}

func (*Person_In) isPerson_Contrace() {}

func (m *Person) GetContrace() isPerson_Contrace {
	if m != nil {
		return m.Contrace
	}
	return nil
}

func (m *Person) GetStr() string {
	if x, ok := m.GetContrace().(*Person_Str); ok {
		return x.Str
	}
	return ""
}

func (m *Person) GetIn() int64 {
	if x, ok := m.GetContrace().(*Person_In); ok {
		return x.In
	}
	return 0
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *Person) GetLastUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Person) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Person_OneofMarshaler, _Person_OneofUnmarshaler, _Person_OneofSizer, []interface{}{
		(*Person_Str)(nil),
		(*Person_In)(nil),
	}
}

func _Person_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Person)
	// contrace
	switch x := m.Contrace.(type) {
	case *Person_Str:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Str)
	case *Person_In:
		b.EncodeVarint(7<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.In))
	case nil:
	default:
		return fmt.Errorf("Person.Contrace has unexpected type %T", x)
	}
	return nil
}

func _Person_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Person)
	switch tag {
	case 6: // contrace.str
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Contrace = &Person_Str{x}
		return true, err
	case 7: // contrace.in
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Contrace = &Person_In{int64(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Person_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Person)
	// contrace
	switch x := m.Contrace.(type) {
	case *Person_Str:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Str)))
		n += len(x.Str)
	case *Person_In:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.In))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Person_PhoneNumber struct {
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=tutorial.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb1a68c9dd6d429, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor_1eb1a68c9dd6d429) }

var fileDescriptor_1eb1a68c9dd6d429 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x9b, 0x9f, 0xe6, 0x6b, 0x4f, 0x3e, 0x4a, 0x3c, 0x88, 0x84, 0x20, 0x18, 0xba, 0x0a,
	0x08, 0x53, 0xa8, 0x82, 0x2b, 0x17, 0x16, 0x0a, 0x15, 0xad, 0x2d, 0x43, 0x8b, 0x4b, 0x99, 0x34,
	0x63, 0x0d, 0x4d, 0x32, 0x43, 0x32, 0x59, 0xf4, 0xa2, 0xbd, 0x07, 0xc9, 0x24, 0x15, 0x11, 0x77,
	0xe7, 0xe7, 0x21, 0x79, 0x9f, 0x33, 0x70, 0xc6, 0x92, 0xa4, 0xe4, 0x55, 0x15, 0x0b, 0x71, 0x20,
	0xb2, 0x14, 0x4a, 0xe0, 0x40, 0xd5, 0x4a, 0x94, 0x29, 0xcb, 0x82, 0xab, 0xbd, 0x10, 0xfb, 0x8c,
	0x4f, 0xf4, 0x3c, 0xae, 0xdf, 0x27, 0x2a, 0xcd, 0x79, 0xa5, 0x58, 0x2e, 0x5b, 0x74, 0xfc, 0x69,
	0x82, 0xb3, 0xe6, 0x65, 0x25, 0x0a, 0x44, 0xb0, 0x0b, 0x96, 0x73, 0xdf, 0x08, 0x8d, 0x68, 0x48,
	0x75, 0x8d, 0x23, 0x30, 0xd3, 0xc4, 0x37, 0x43, 0x23, 0xea, 0x53, 0x33, 0x4d, 0xf0, 0x1c, 0xfa,
	0x3c, 0x67, 0x69, 0xe6, 0x5b, 0x1a, 0x6a, 0x1b, 0x44, 0xb0, 0x2a, 0x55, 0xfa, 0x4e, 0x33, 0x5b,
	0xf4, 0x68, 0xd3, 0xa0, 0x07, 0x66, 0x5a, 0xf8, 0xff, 0x42, 0x23, 0xb2, 0x16, 0x3d, 0x6a, 0xa6,
	0x05, 0xde, 0x82, 0x23, 0x3f, 0x44, 0xc1, 0x2b, 0xdf, 0x0e, 0xad, 0xc8, 0x9d, 0x5e, 0x92, 0x53,
	0x4c, 0xd2, 0x26, 0x20, 0xeb, 0x66, 0xfd, 0x52, 0xe7, 0x31, 0x2f, 0x69, 0xc7, 0xe2, 0x3d, 0xfc,
	0xcf, 0x58, 0xa5, 0xde, 0x6a, 0x99, 0x30, 0xc5, 0x13, 0xbf, 0x1f, 0x1a, 0x91, 0x3b, 0x0d, 0x48,
	0x2b, 0x46, 0x4e, 0x62, 0x64, 0x73, 0x12, 0xa3, 0x6e, 0xc3, 0x6f, 0x5b, 0x3c, 0xd8, 0x82, 0xfb,
	0xe3, 0xab, 0x78, 0x01, 0x4e, 0xa1, 0xab, 0xce, 0xb2, 0xeb, 0x90, 0x80, 0xad, 0x8e, 0x92, 0x6b,
	0xd3, 0xd1, 0x34, 0xf8, 0x3b, 0xd9, 0xe6, 0x28, 0x39, 0xd5, 0xdc, 0xf8, 0x1a, 0x86, 0xdf, 0x23,
	0x04, 0x70, 0x96, 0xab, 0xd9, 0xe3, 0xf3, 0xdc, 0xeb, 0xe1, 0x00, 0xec, 0xc5, 0x6a, 0x39, 0xf7,
	0x8c, 0xa6, 0x7a, 0x5d, 0xd1, 0x27, 0xcf, 0x9c, 0x01, 0x0c, 0x76, 0xa2, 0x50, 0x25, 0xdb, 0xf1,
	0xf1, 0x1d, 0xb8, 0x0f, 0xed, 0x7b, 0xcd, 0x84, 0x38, 0x60, 0x04, 0x8e, 0xe4, 0x42, 0x66, 0xcd,
	0xd5, 0x9b, 0x9b, 0x78, 0xbf, 0xff, 0x4c, 0xbb, 0x7d, 0xec, 0x68, 0xd3, 0x9b, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbf, 0x97, 0x09, 0xab, 0xef, 0x01, 0x00, 0x00,
}
