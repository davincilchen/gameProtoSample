// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverMessage.proto

package server

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type StatusCode int32

const (
	StatusCode_Ok StatusCode = 0
	//-----------------------
	StatusCode_InsufficientBalance StatusCode = 101
	//-----------------------
	StatusCode_ErrorRequest StatusCode = 200
	StatusCode_ErrorUser    StatusCode = 201
)

var StatusCode_name = map[int32]string{
	0:   "Ok",
	101: "InsufficientBalance",
	200: "ErrorRequest",
	201: "ErrorUser",
}

var StatusCode_value = map[string]int32{
	"Ok":                  0,
	"InsufficientBalance": 101,
	"ErrorRequest":        200,
	"ErrorUser":           201,
}

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}

func (x *StatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StatusCode_value, data, "StatusCode")
	if err != nil {
		return err
	}
	*x = StatusCode(value)
	return nil
}

func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{0}
}

type LoginRequest struct {
	// Types that are valid to be assigned to User:
	//	*LoginRequest_Normal_
	//	*LoginRequest_Dev_
	User                 isLoginRequest_User `protobuf_oneof:"User"`
	Game                 *string             `protobuf:"bytes,3,req,name=Game" json:"Game,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

type isLoginRequest_User interface {
	isLoginRequest_User()
}

type LoginRequest_Normal_ struct {
	Normal *LoginRequest_Normal `protobuf:"group,1,opt,name=Normal,json=normal,oneof"`
}

type LoginRequest_Dev_ struct {
	Dev *LoginRequest_Dev `protobuf:"group,2,opt,name=Dev,json=dev,oneof"`
}

func (*LoginRequest_Normal_) isLoginRequest_User() {}

func (*LoginRequest_Dev_) isLoginRequest_User() {}

func (m *LoginRequest) GetUser() isLoginRequest_User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *LoginRequest) GetNormal() *LoginRequest_Normal {
	if x, ok := m.GetUser().(*LoginRequest_Normal_); ok {
		return x.Normal
	}
	return nil
}

func (m *LoginRequest) GetDev() *LoginRequest_Dev {
	if x, ok := m.GetUser().(*LoginRequest_Dev_); ok {
		return x.Dev
	}
	return nil
}

func (m *LoginRequest) GetGame() string {
	if m != nil && m.Game != nil {
		return *m.Game
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LoginRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LoginRequest_OneofMarshaler, _LoginRequest_OneofUnmarshaler, _LoginRequest_OneofSizer, []interface{}{
		(*LoginRequest_Normal_)(nil),
		(*LoginRequest_Dev_)(nil),
	}
}

func _LoginRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LoginRequest)
	// User
	switch x := m.User.(type) {
	case *LoginRequest_Normal_:
		b.EncodeVarint(1<<3 | proto.WireStartGroup)
		if err := b.Marshal(x.Normal); err != nil {
			return err
		}
		b.EncodeVarint(1<<3 | proto.WireEndGroup)
	case *LoginRequest_Dev_:
		b.EncodeVarint(2<<3 | proto.WireStartGroup)
		if err := b.Marshal(x.Dev); err != nil {
			return err
		}
		b.EncodeVarint(2<<3 | proto.WireEndGroup)
	case nil:
	default:
		return fmt.Errorf("LoginRequest.User has unexpected type %T", x)
	}
	return nil
}

func _LoginRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LoginRequest)
	switch tag {
	case 1: // User.normal
		if wire != proto.WireStartGroup {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoginRequest_Normal)
		err := b.DecodeGroup(msg)
		m.User = &LoginRequest_Normal_{msg}
		return true, err
	case 2: // User.dev
		if wire != proto.WireStartGroup {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoginRequest_Dev)
		err := b.DecodeGroup(msg)
		m.User = &LoginRequest_Dev_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LoginRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LoginRequest)
	// User
	switch x := m.User.(type) {
	case *LoginRequest_Normal_:
		n += 1 // tag and wire
		n += proto.Size(x.Normal)
		n += 1 // tag and wire
	case *LoginRequest_Dev_:
		n += 1 // tag and wire
		n += proto.Size(x.Dev)
		n += 1 // tag and wire
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LoginRequest_Normal struct {
	Token                *string  `protobuf:"bytes,1,req,name=Token" json:"Token,omitempty"`
	Platform             *string  `protobuf:"bytes,2,req,name=Platform" json:"Platform,omitempty"`
	Demo                 *bool    `protobuf:"varint,3,req,name=demo" json:"demo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest_Normal) Reset()         { *m = LoginRequest_Normal{} }
func (m *LoginRequest_Normal) String() string { return proto.CompactTextString(m) }
func (*LoginRequest_Normal) ProtoMessage()    {}
func (*LoginRequest_Normal) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{0, 0}
}

func (m *LoginRequest_Normal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest_Normal.Unmarshal(m, b)
}
func (m *LoginRequest_Normal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest_Normal.Marshal(b, m, deterministic)
}
func (m *LoginRequest_Normal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest_Normal.Merge(m, src)
}
func (m *LoginRequest_Normal) XXX_Size() int {
	return xxx_messageInfo_LoginRequest_Normal.Size(m)
}
func (m *LoginRequest_Normal) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest_Normal.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest_Normal proto.InternalMessageInfo

func (m *LoginRequest_Normal) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

func (m *LoginRequest_Normal) GetPlatform() string {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return ""
}

func (m *LoginRequest_Normal) GetDemo() bool {
	if m != nil && m.Demo != nil {
		return *m.Demo
	}
	return false
}

type LoginRequest_Dev struct {
	Id                   *string  `protobuf:"bytes,1,req,name=Id" json:"Id,omitempty"`
	Password             *string  `protobuf:"bytes,2,req,name=Password" json:"Password,omitempty"`
	ForcedEnd            *bool    `protobuf:"varint,3,req,name=ForcedEnd" json:"ForcedEnd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest_Dev) Reset()         { *m = LoginRequest_Dev{} }
func (m *LoginRequest_Dev) String() string { return proto.CompactTextString(m) }
func (*LoginRequest_Dev) ProtoMessage()    {}
func (*LoginRequest_Dev) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{0, 1}
}

func (m *LoginRequest_Dev) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest_Dev.Unmarshal(m, b)
}
func (m *LoginRequest_Dev) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest_Dev.Marshal(b, m, deterministic)
}
func (m *LoginRequest_Dev) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest_Dev.Merge(m, src)
}
func (m *LoginRequest_Dev) XXX_Size() int {
	return xxx_messageInfo_LoginRequest_Dev.Size(m)
}
func (m *LoginRequest_Dev) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest_Dev.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest_Dev proto.InternalMessageInfo

func (m *LoginRequest_Dev) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *LoginRequest_Dev) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

func (m *LoginRequest_Dev) GetForcedEnd() bool {
	if m != nil && m.ForcedEnd != nil {
		return *m.ForcedEnd
	}
	return false
}

type GameStartRequest struct {
	GameCommand          []string `protobuf:"bytes,1,rep,name=GameCommand" json:"GameCommand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameStartRequest) Reset()         { *m = GameStartRequest{} }
func (m *GameStartRequest) String() string { return proto.CompactTextString(m) }
func (*GameStartRequest) ProtoMessage()    {}
func (*GameStartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{1}
}

func (m *GameStartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStartRequest.Unmarshal(m, b)
}
func (m *GameStartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStartRequest.Marshal(b, m, deterministic)
}
func (m *GameStartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStartRequest.Merge(m, src)
}
func (m *GameStartRequest) XXX_Size() int {
	return xxx_messageInfo_GameStartRequest.Size(m)
}
func (m *GameStartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GameStartRequest proto.InternalMessageInfo

func (m *GameStartRequest) GetGameCommand() []string {
	if m != nil {
		return m.GameCommand
	}
	return nil
}

type ResultRequest struct {
	GameCommand          []string `protobuf:"bytes,1,rep,name=GameCommand" json:"GameCommand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResultRequest) Reset()         { *m = ResultRequest{} }
func (m *ResultRequest) String() string { return proto.CompactTextString(m) }
func (*ResultRequest) ProtoMessage()    {}
func (*ResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{2}
}

func (m *ResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultRequest.Unmarshal(m, b)
}
func (m *ResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultRequest.Marshal(b, m, deterministic)
}
func (m *ResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultRequest.Merge(m, src)
}
func (m *ResultRequest) XXX_Size() int {
	return xxx_messageInfo_ResultRequest.Size(m)
}
func (m *ResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResultRequest proto.InternalMessageInfo

func (m *ResultRequest) GetGameCommand() []string {
	if m != nil {
		return m.GameCommand
	}
	return nil
}

type GameEndRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameEndRequest) Reset()         { *m = GameEndRequest{} }
func (m *GameEndRequest) String() string { return proto.CompactTextString(m) }
func (*GameEndRequest) ProtoMessage()    {}
func (*GameEndRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{3}
}

func (m *GameEndRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameEndRequest.Unmarshal(m, b)
}
func (m *GameEndRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameEndRequest.Marshal(b, m, deterministic)
}
func (m *GameEndRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameEndRequest.Merge(m, src)
}
func (m *GameEndRequest) XXX_Size() int {
	return xxx_messageInfo_GameEndRequest.Size(m)
}
func (m *GameEndRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GameEndRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GameEndRequest proto.InternalMessageInfo

type ToServer struct {
	// Types that are valid to be assigned to Message:
	//	*ToServer_LoginRequest
	//	*ToServer_GameStartRequest
	//	*ToServer_ResultRequest
	//	*ToServer_GameEndRequest
	Message              isToServer_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ToServer) Reset()         { *m = ToServer{} }
func (m *ToServer) String() string { return proto.CompactTextString(m) }
func (*ToServer) ProtoMessage()    {}
func (*ToServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{4}
}

func (m *ToServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer.Unmarshal(m, b)
}
func (m *ToServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer.Marshal(b, m, deterministic)
}
func (m *ToServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer.Merge(m, src)
}
func (m *ToServer) XXX_Size() int {
	return xxx_messageInfo_ToServer.Size(m)
}
func (m *ToServer) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer proto.InternalMessageInfo

type isToServer_Message interface {
	isToServer_Message()
}

type ToServer_LoginRequest struct {
	LoginRequest *LoginRequest `protobuf:"bytes,1,opt,name=LoginRequest,oneof"`
}

type ToServer_GameStartRequest struct {
	GameStartRequest *GameStartRequest `protobuf:"bytes,2,opt,name=GameStartRequest,oneof"`
}

type ToServer_ResultRequest struct {
	ResultRequest *ResultRequest `protobuf:"bytes,3,opt,name=ResultRequest,oneof"`
}

type ToServer_GameEndRequest struct {
	GameEndRequest *GameEndRequest `protobuf:"bytes,4,opt,name=GameEndRequest,oneof"`
}

func (*ToServer_LoginRequest) isToServer_Message() {}

func (*ToServer_GameStartRequest) isToServer_Message() {}

func (*ToServer_ResultRequest) isToServer_Message() {}

func (*ToServer_GameEndRequest) isToServer_Message() {}

func (m *ToServer) GetMessage() isToServer_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ToServer) GetLoginRequest() *LoginRequest {
	if x, ok := m.GetMessage().(*ToServer_LoginRequest); ok {
		return x.LoginRequest
	}
	return nil
}

func (m *ToServer) GetGameStartRequest() *GameStartRequest {
	if x, ok := m.GetMessage().(*ToServer_GameStartRequest); ok {
		return x.GameStartRequest
	}
	return nil
}

func (m *ToServer) GetResultRequest() *ResultRequest {
	if x, ok := m.GetMessage().(*ToServer_ResultRequest); ok {
		return x.ResultRequest
	}
	return nil
}

func (m *ToServer) GetGameEndRequest() *GameEndRequest {
	if x, ok := m.GetMessage().(*ToServer_GameEndRequest); ok {
		return x.GameEndRequest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ToServer) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ToServer_OneofMarshaler, _ToServer_OneofUnmarshaler, _ToServer_OneofSizer, []interface{}{
		(*ToServer_LoginRequest)(nil),
		(*ToServer_GameStartRequest)(nil),
		(*ToServer_ResultRequest)(nil),
		(*ToServer_GameEndRequest)(nil),
	}
}

func _ToServer_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ToServer)
	// Message
	switch x := m.Message.(type) {
	case *ToServer_LoginRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LoginRequest); err != nil {
			return err
		}
	case *ToServer_GameStartRequest:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GameStartRequest); err != nil {
			return err
		}
	case *ToServer_ResultRequest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResultRequest); err != nil {
			return err
		}
	case *ToServer_GameEndRequest:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GameEndRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ToServer.Message has unexpected type %T", x)
	}
	return nil
}

func _ToServer_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ToServer)
	switch tag {
	case 1: // Message.LoginRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoginRequest)
		err := b.DecodeMessage(msg)
		m.Message = &ToServer_LoginRequest{msg}
		return true, err
	case 2: // Message.GameStartRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GameStartRequest)
		err := b.DecodeMessage(msg)
		m.Message = &ToServer_GameStartRequest{msg}
		return true, err
	case 3: // Message.ResultRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ResultRequest)
		err := b.DecodeMessage(msg)
		m.Message = &ToServer_ResultRequest{msg}
		return true, err
	case 4: // Message.GameEndRequest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GameEndRequest)
		err := b.DecodeMessage(msg)
		m.Message = &ToServer_GameEndRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ToServer_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ToServer)
	// Message
	switch x := m.Message.(type) {
	case *ToServer_LoginRequest:
		s := proto.Size(x.LoginRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToServer_GameStartRequest:
		s := proto.Size(x.GameStartRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToServer_ResultRequest:
		s := proto.Size(x.ResultRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToServer_GameEndRequest:
		s := proto.Size(x.GameEndRequest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LoginResponse struct {
	StatusCode           *StatusCode `protobuf:"varint,1,req,name=StatusCode,enum=main.StatusCode" json:"StatusCode,omitempty"`
	Balance              *uint64     `protobuf:"varint,2,opt,name=Balance" json:"Balance,omitempty"`
	GameData             *string     `protobuf:"bytes,3,opt,name=GameData" json:"GameData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{5}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetStatusCode() StatusCode {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return StatusCode_Ok
}

func (m *LoginResponse) GetBalance() uint64 {
	if m != nil && m.Balance != nil {
		return *m.Balance
	}
	return 0
}

func (m *LoginResponse) GetGameData() string {
	if m != nil && m.GameData != nil {
		return *m.GameData
	}
	return ""
}

type GameStartResponse struct {
	StatusCode           *StatusCode `protobuf:"varint,1,req,name=StatusCode,enum=main.StatusCode" json:"StatusCode,omitempty"`
	Balance              *uint64     `protobuf:"varint,2,opt,name=Balance" json:"Balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GameStartResponse) Reset()         { *m = GameStartResponse{} }
func (m *GameStartResponse) String() string { return proto.CompactTextString(m) }
func (*GameStartResponse) ProtoMessage()    {}
func (*GameStartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{6}
}

func (m *GameStartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameStartResponse.Unmarshal(m, b)
}
func (m *GameStartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameStartResponse.Marshal(b, m, deterministic)
}
func (m *GameStartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameStartResponse.Merge(m, src)
}
func (m *GameStartResponse) XXX_Size() int {
	return xxx_messageInfo_GameStartResponse.Size(m)
}
func (m *GameStartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GameStartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GameStartResponse proto.InternalMessageInfo

func (m *GameStartResponse) GetStatusCode() StatusCode {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return StatusCode_Ok
}

func (m *GameStartResponse) GetBalance() uint64 {
	if m != nil && m.Balance != nil {
		return *m.Balance
	}
	return 0
}

type ResultResponse struct {
	StatusCode           *StatusCode `protobuf:"varint,1,req,name=StatusCode,enum=main.StatusCode" json:"StatusCode,omitempty"`
	Result               *string     `protobuf:"bytes,2,opt,name=Result" json:"Result,omitempty"`
	BonusInfo            *string     `protobuf:"bytes,3,opt,name=BonusInfo" json:"BonusInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ResultResponse) Reset()         { *m = ResultResponse{} }
func (m *ResultResponse) String() string { return proto.CompactTextString(m) }
func (*ResultResponse) ProtoMessage()    {}
func (*ResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{7}
}

func (m *ResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultResponse.Unmarshal(m, b)
}
func (m *ResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultResponse.Marshal(b, m, deterministic)
}
func (m *ResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultResponse.Merge(m, src)
}
func (m *ResultResponse) XXX_Size() int {
	return xxx_messageInfo_ResultResponse.Size(m)
}
func (m *ResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResultResponse proto.InternalMessageInfo

func (m *ResultResponse) GetStatusCode() StatusCode {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return StatusCode_Ok
}

func (m *ResultResponse) GetResult() string {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return ""
}

func (m *ResultResponse) GetBonusInfo() string {
	if m != nil && m.BonusInfo != nil {
		return *m.BonusInfo
	}
	return ""
}

type GameEndResponse struct {
	StatusCode           *StatusCode `protobuf:"varint,1,req,name=StatusCode,enum=main.StatusCode" json:"StatusCode,omitempty"`
	Balance              *uint64     `protobuf:"varint,2,opt,name=Balance" json:"Balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GameEndResponse) Reset()         { *m = GameEndResponse{} }
func (m *GameEndResponse) String() string { return proto.CompactTextString(m) }
func (*GameEndResponse) ProtoMessage()    {}
func (*GameEndResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{8}
}

func (m *GameEndResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameEndResponse.Unmarshal(m, b)
}
func (m *GameEndResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameEndResponse.Marshal(b, m, deterministic)
}
func (m *GameEndResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameEndResponse.Merge(m, src)
}
func (m *GameEndResponse) XXX_Size() int {
	return xxx_messageInfo_GameEndResponse.Size(m)
}
func (m *GameEndResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GameEndResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GameEndResponse proto.InternalMessageInfo

func (m *GameEndResponse) GetStatusCode() StatusCode {
	if m != nil && m.StatusCode != nil {
		return *m.StatusCode
	}
	return StatusCode_Ok
}

func (m *GameEndResponse) GetBalance() uint64 {
	if m != nil && m.Balance != nil {
		return *m.Balance
	}
	return 0
}

type TimeOut struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeOut) Reset()         { *m = TimeOut{} }
func (m *TimeOut) String() string { return proto.CompactTextString(m) }
func (*TimeOut) ProtoMessage()    {}
func (*TimeOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{9}
}

func (m *TimeOut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeOut.Unmarshal(m, b)
}
func (m *TimeOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeOut.Marshal(b, m, deterministic)
}
func (m *TimeOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeOut.Merge(m, src)
}
func (m *TimeOut) XXX_Size() int {
	return xxx_messageInfo_TimeOut.Size(m)
}
func (m *TimeOut) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeOut.DiscardUnknown(m)
}

var xxx_messageInfo_TimeOut proto.InternalMessageInfo

type ToClient struct {
	// Types that are valid to be assigned to Message:
	//	*ToClient_LoginResponse
	//	*ToClient_GameStartResponse
	//	*ToClient_ResultResponse
	//	*ToClient_GameEndResponse
	//	*ToClient_TimeOut
	Message              isToClient_Message `protobuf_oneof:"Message"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ToClient) Reset()         { *m = ToClient{} }
func (m *ToClient) String() string { return proto.CompactTextString(m) }
func (*ToClient) ProtoMessage()    {}
func (*ToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0b54f7f181ffd53, []int{10}
}

func (m *ToClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient.Unmarshal(m, b)
}
func (m *ToClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient.Marshal(b, m, deterministic)
}
func (m *ToClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient.Merge(m, src)
}
func (m *ToClient) XXX_Size() int {
	return xxx_messageInfo_ToClient.Size(m)
}
func (m *ToClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient proto.InternalMessageInfo

type isToClient_Message interface {
	isToClient_Message()
}

type ToClient_LoginResponse struct {
	LoginResponse *LoginResponse `protobuf:"bytes,1,opt,name=LoginResponse,oneof"`
}

type ToClient_GameStartResponse struct {
	GameStartResponse *GameStartResponse `protobuf:"bytes,2,opt,name=GameStartResponse,oneof"`
}

type ToClient_ResultResponse struct {
	ResultResponse *ResultResponse `protobuf:"bytes,3,opt,name=ResultResponse,oneof"`
}

type ToClient_GameEndResponse struct {
	GameEndResponse *GameEndResponse `protobuf:"bytes,4,opt,name=GameEndResponse,oneof"`
}

type ToClient_TimeOut struct {
	TimeOut *TimeOut `protobuf:"bytes,5,opt,name=TimeOut,oneof"`
}

func (*ToClient_LoginResponse) isToClient_Message() {}

func (*ToClient_GameStartResponse) isToClient_Message() {}

func (*ToClient_ResultResponse) isToClient_Message() {}

func (*ToClient_GameEndResponse) isToClient_Message() {}

func (*ToClient_TimeOut) isToClient_Message() {}

func (m *ToClient) GetMessage() isToClient_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ToClient) GetLoginResponse() *LoginResponse {
	if x, ok := m.GetMessage().(*ToClient_LoginResponse); ok {
		return x.LoginResponse
	}
	return nil
}

func (m *ToClient) GetGameStartResponse() *GameStartResponse {
	if x, ok := m.GetMessage().(*ToClient_GameStartResponse); ok {
		return x.GameStartResponse
	}
	return nil
}

func (m *ToClient) GetResultResponse() *ResultResponse {
	if x, ok := m.GetMessage().(*ToClient_ResultResponse); ok {
		return x.ResultResponse
	}
	return nil
}

func (m *ToClient) GetGameEndResponse() *GameEndResponse {
	if x, ok := m.GetMessage().(*ToClient_GameEndResponse); ok {
		return x.GameEndResponse
	}
	return nil
}

func (m *ToClient) GetTimeOut() *TimeOut {
	if x, ok := m.GetMessage().(*ToClient_TimeOut); ok {
		return x.TimeOut
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ToClient) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ToClient_OneofMarshaler, _ToClient_OneofUnmarshaler, _ToClient_OneofSizer, []interface{}{
		(*ToClient_LoginResponse)(nil),
		(*ToClient_GameStartResponse)(nil),
		(*ToClient_ResultResponse)(nil),
		(*ToClient_GameEndResponse)(nil),
		(*ToClient_TimeOut)(nil),
	}
}

func _ToClient_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ToClient)
	// Message
	switch x := m.Message.(type) {
	case *ToClient_LoginResponse:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LoginResponse); err != nil {
			return err
		}
	case *ToClient_GameStartResponse:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GameStartResponse); err != nil {
			return err
		}
	case *ToClient_ResultResponse:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResultResponse); err != nil {
			return err
		}
	case *ToClient_GameEndResponse:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GameEndResponse); err != nil {
			return err
		}
	case *ToClient_TimeOut:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimeOut); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ToClient.Message has unexpected type %T", x)
	}
	return nil
}

func _ToClient_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ToClient)
	switch tag {
	case 1: // Message.LoginResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LoginResponse)
		err := b.DecodeMessage(msg)
		m.Message = &ToClient_LoginResponse{msg}
		return true, err
	case 2: // Message.GameStartResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GameStartResponse)
		err := b.DecodeMessage(msg)
		m.Message = &ToClient_GameStartResponse{msg}
		return true, err
	case 3: // Message.ResultResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ResultResponse)
		err := b.DecodeMessage(msg)
		m.Message = &ToClient_ResultResponse{msg}
		return true, err
	case 4: // Message.GameEndResponse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GameEndResponse)
		err := b.DecodeMessage(msg)
		m.Message = &ToClient_GameEndResponse{msg}
		return true, err
	case 5: // Message.TimeOut
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TimeOut)
		err := b.DecodeMessage(msg)
		m.Message = &ToClient_TimeOut{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ToClient_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ToClient)
	// Message
	switch x := m.Message.(type) {
	case *ToClient_LoginResponse:
		s := proto.Size(x.LoginResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToClient_GameStartResponse:
		s := proto.Size(x.GameStartResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToClient_ResultResponse:
		s := proto.Size(x.ResultResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToClient_GameEndResponse:
		s := proto.Size(x.GameEndResponse)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ToClient_TimeOut:
		s := proto.Size(x.TimeOut)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("main.StatusCode", StatusCode_name, StatusCode_value)
	proto.RegisterType((*LoginRequest)(nil), "main.LoginRequest")
	proto.RegisterType((*LoginRequest_Normal)(nil), "main.LoginRequest.Normal")
	proto.RegisterType((*LoginRequest_Dev)(nil), "main.LoginRequest.Dev")
	proto.RegisterType((*GameStartRequest)(nil), "main.GameStartRequest")
	proto.RegisterType((*ResultRequest)(nil), "main.ResultRequest")
	proto.RegisterType((*GameEndRequest)(nil), "main.GameEndRequest")
	proto.RegisterType((*ToServer)(nil), "main.ToServer")
	proto.RegisterType((*LoginResponse)(nil), "main.LoginResponse")
	proto.RegisterType((*GameStartResponse)(nil), "main.GameStartResponse")
	proto.RegisterType((*ResultResponse)(nil), "main.ResultResponse")
	proto.RegisterType((*GameEndResponse)(nil), "main.GameEndResponse")
	proto.RegisterType((*TimeOut)(nil), "main.TimeOut")
	proto.RegisterType((*ToClient)(nil), "main.ToClient")
}

func init() { proto.RegisterFile("serverMessage.proto", fileDescriptor_f0b54f7f181ffd53) }

var fileDescriptor_f0b54f7f181ffd53 = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x4d, 0xd3, 0xae, 0x5b, 0xbe, 0x6d, 0x25, 0x73, 0xc7, 0x16, 0x26, 0x2e, 0xa6, 0x5c, 0x8d,
	0x5d, 0x54, 0x30, 0xb8, 0x40, 0x42, 0x9a, 0x44, 0x7f, 0x58, 0x2a, 0xc1, 0x3a, 0xb9, 0xe5, 0x12,
	0x21, 0xab, 0x71, 0xa7, 0x6a, 0x8d, 0x3d, 0xec, 0xa4, 0xe3, 0x39, 0x78, 0x0e, 0x1e, 0x84, 0xbd,
	0x15, 0xb2, 0xe3, 0xa6, 0xf9, 0xe9, 0x0d, 0x08, 0xee, 0xf2, 0xfd, 0x1d, 0xfb, 0x3b, 0xe7, 0x38,
	0xd0, 0x96, 0x54, 0x2c, 0xa9, 0xf8, 0x44, 0xa5, 0x24, 0xb7, 0xb4, 0x73, 0x2f, 0x78, 0xcc, 0x51,
	0x23, 0x22, 0x73, 0xe6, 0xff, 0xb4, 0x61, 0xef, 0x23, 0xbf, 0x9d, 0x33, 0x4c, 0xbf, 0x25, 0x54,
	0xc6, 0xe8, 0x35, 0x34, 0x19, 0x17, 0x11, 0x59, 0x78, 0xb5, 0xd3, 0xda, 0x19, 0x5c, 0x3c, 0xeb,
	0xa8, 0xbe, 0x4e, 0xbe, 0xa7, 0x73, 0xad, 0x1b, 0x02, 0x0b, 0x9b, 0x56, 0x74, 0x0e, 0xf5, 0x90,
	0x2e, 0x3d, 0x5b, 0x4f, 0x1c, 0x6d, 0x98, 0xe8, 0xd3, 0x65, 0x60, 0x61, 0xd5, 0x84, 0x10, 0x34,
	0xae, 0x48, 0x44, 0xbd, 0xfa, 0xa9, 0x7d, 0xe6, 0x60, 0xfd, 0x7d, 0x72, 0x0d, 0xcd, 0x14, 0x13,
	0x1d, 0xc2, 0xd6, 0x84, 0xdf, 0x51, 0xe6, 0xd5, 0x74, 0x39, 0x0d, 0xd0, 0x09, 0xec, 0xdc, 0x2c,
	0x48, 0x3c, 0xe3, 0x22, 0xf2, 0x6c, 0x5d, 0xc8, 0x62, 0x85, 0x17, 0xd2, 0x88, 0x6b, 0xbc, 0x1d,
	0xac, 0xbf, 0x4f, 0x46, 0x50, 0xef, 0xd3, 0x25, 0x6a, 0x81, 0x3d, 0x0c, 0x0d, 0x92, 0x3d, 0x0c,
	0x35, 0x0c, 0x91, 0xf2, 0x81, 0x8b, 0x30, 0x83, 0x31, 0x31, 0x7a, 0x0e, 0xce, 0x07, 0x2e, 0xa6,
	0x34, 0x1c, 0xb0, 0xd0, 0x60, 0xad, 0x13, 0xdd, 0x26, 0x34, 0x3e, 0x4b, 0x2a, 0xfc, 0x37, 0xe0,
	0xaa, 0x0b, 0x8f, 0x63, 0x22, 0xe2, 0x15, 0x63, 0xa7, 0xb0, 0xab, 0x72, 0x3d, 0x1e, 0x45, 0x84,
	0xa9, 0xe3, 0xea, 0x67, 0x0e, 0xce, 0xa7, 0xfc, 0x57, 0xb0, 0x8f, 0xa9, 0x4c, 0x16, 0x7f, 0x30,
	0xe2, 0x42, 0x4b, 0x85, 0x03, 0x16, 0x9a, 0x19, 0xff, 0x87, 0x0d, 0x3b, 0x13, 0x3e, 0xd6, 0x4a,
	0xa2, 0xb7, 0x45, 0xd5, 0xb4, 0x56, 0xbb, 0x17, 0xa8, 0xca, 0x7c, 0x60, 0xe1, 0xa2, 0xbe, 0xfd,
	0xea, 0x06, 0x5a, 0xb7, 0xdd, 0x95, 0x6e, 0xe5, 0x6a, 0x60, 0xe1, 0xea, 0xce, 0xef, 0x4a, 0x1b,
	0x79, 0x75, 0x0d, 0xd1, 0x4e, 0x21, 0x0a, 0xa5, 0xc0, 0xc2, 0xa5, 0xed, 0x2f, 0xcb, 0xbb, 0x79,
	0x0d, 0x3d, 0x7d, 0xb8, 0xbe, 0xc0, 0xba, 0x16, 0x58, 0xb8, 0xd4, 0xdd, 0x75, 0x60, 0xdb, 0x58,
	0xd9, 0x7f, 0x80, 0x7d, 0xb3, 0x9d, 0xbc, 0xe7, 0x4c, 0x52, 0xf4, 0x12, 0x60, 0x1c, 0x93, 0x38,
	0x91, 0x3d, 0x1e, 0x52, 0x2d, 0x7d, 0xeb, 0xc2, 0x4d, 0x71, 0xd7, 0x79, 0x9c, 0xeb, 0x41, 0x1e,
	0x6c, 0x77, 0xc9, 0x82, 0xb0, 0x29, 0xd5, 0x3c, 0x34, 0xf0, 0x2a, 0x54, 0x76, 0x51, 0x27, 0xf7,
	0x49, 0x4c, 0xf4, 0x7e, 0x0e, 0xce, 0x62, 0xff, 0x2b, 0x1c, 0xe4, 0x48, 0xf9, 0xf7, 0x87, 0xfb,
	0xdf, 0xa1, 0xb5, 0x62, 0xed, 0xaf, 0xd1, 0x8f, 0xa0, 0x99, 0x62, 0x68, 0x70, 0x07, 0x9b, 0x48,
	0x79, 0xbd, 0xcb, 0x59, 0x22, 0x87, 0x6c, 0xc6, 0xcd, 0x66, 0xeb, 0x84, 0xff, 0x05, 0x9e, 0x64,
	0x84, 0xff, 0x87, 0xc5, 0x1c, 0xd8, 0x9e, 0xcc, 0x23, 0x3a, 0x4a, 0x62, 0xff, 0x51, 0x5b, 0xba,
	0xb7, 0x98, 0x53, 0xa6, 0x2d, 0x55, 0x90, 0xd2, 0x78, 0xba, 0x5d, 0xf0, 0x74, 0x5a, 0x52, 0x96,
	0x2a, 0xca, 0x7e, 0xb5, 0x41, 0x0e, 0x63, 0xeb, 0xe3, 0x8a, 0xad, 0x33, 0x90, 0x0d, 0x12, 0x5e,
	0x96, 0x69, 0x37, 0xce, 0x3e, 0x2c, 0x3a, 0x3b, 0x83, 0x28, 0x8b, 0xf4, 0xbe, 0x42, 0x9e, 0x31,
	0xf7, 0xd3, 0x92, 0xb9, 0x33, 0x84, 0x0a, 0xd9, 0x2f, 0x32, 0x82, 0xbc, 0x2d, 0x3d, 0xba, 0x9f,
	0x8e, 0x9a, 0x64, 0x60, 0xe1, 0x55, 0x3d, 0xf7, 0x12, 0xce, 0x6f, 0xf2, 0x12, 0xa1, 0x26, 0xd8,
	0xa3, 0x3b, 0xd7, 0x42, 0xc7, 0xd0, 0x1e, 0x32, 0x99, 0xcc, 0x66, 0xf3, 0xa9, 0x22, 0xd9, 0x68,
	0xe0, 0x52, 0x74, 0x00, 0x7b, 0x03, 0x21, 0xb8, 0x30, 0x6f, 0xca, 0xfd, 0x55, 0x43, 0x2d, 0x70,
	0x74, 0x4a, 0xfd, 0xe8, 0xdc, 0xc7, 0xda, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xcd, 0x38,
	0x66, 0x36, 0x06, 0x00, 0x00,
}
