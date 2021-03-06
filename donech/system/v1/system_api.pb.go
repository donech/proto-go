// Code generated by protoc-gen-go. DO NOT EDIT.
// source: donech/system/v1/system_api.proto

package systemv1

import (
	context "context"
	fmt "fmt"
	v1 "github.com/donech/donech-/gen/go/donech/util/v1"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//SaveUserRequest.
type SaveUserRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveUserRequest) Reset()         { *m = SaveUserRequest{} }
func (m *SaveUserRequest) String() string { return proto.CompactTextString(m) }
func (*SaveUserRequest) ProtoMessage()    {}
func (*SaveUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{0}
}

func (m *SaveUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveUserRequest.Unmarshal(m, b)
}
func (m *SaveUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveUserRequest.Marshal(b, m, deterministic)
}
func (m *SaveUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveUserRequest.Merge(m, src)
}
func (m *SaveUserRequest) XXX_Size() int {
	return xxx_messageInfo_SaveUserRequest.Size(m)
}
func (m *SaveUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveUserRequest proto.InternalMessageInfo

func (m *SaveUserRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *SaveUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SaveUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

//SaveUserResponse.
type SaveUserResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveUserResponse) Reset()         { *m = SaveUserResponse{} }
func (m *SaveUserResponse) String() string { return proto.CompactTextString(m) }
func (*SaveUserResponse) ProtoMessage()    {}
func (*SaveUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{1}
}

func (m *SaveUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveUserResponse.Unmarshal(m, b)
}
func (m *SaveUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveUserResponse.Marshal(b, m, deterministic)
}
func (m *SaveUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveUserResponse.Merge(m, src)
}
func (m *SaveUserResponse) XXX_Size() int {
	return xxx_messageInfo_SaveUserResponse.Size(m)
}
func (m *SaveUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveUserResponse proto.InternalMessageInfo

func (m *SaveUserResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SaveUserResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//UsersRequest.
type UsersRequest struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Pager                *v1.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UsersRequest) Reset()         { *m = UsersRequest{} }
func (m *UsersRequest) String() string { return proto.CompactTextString(m) }
func (*UsersRequest) ProtoMessage()    {}
func (*UsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{2}
}

func (m *UsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersRequest.Unmarshal(m, b)
}
func (m *UsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersRequest.Marshal(b, m, deterministic)
}
func (m *UsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersRequest.Merge(m, src)
}
func (m *UsersRequest) XXX_Size() int {
	return xxx_messageInfo_UsersRequest.Size(m)
}
func (m *UsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UsersRequest proto.InternalMessageInfo

func (m *UsersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UsersRequest) GetPager() *v1.Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

//UsersResponse.
type UsersResponse struct {
	Code                 int32                 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string                `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 []*UsersResponse_Data `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Pager                *v1.Pager             `protobuf:"bytes,4,opt,name=pager,proto3" json:"pager,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UsersResponse) Reset()         { *m = UsersResponse{} }
func (m *UsersResponse) String() string { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()    {}
func (*UsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{3}
}

func (m *UsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersResponse.Unmarshal(m, b)
}
func (m *UsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersResponse.Marshal(b, m, deterministic)
}
func (m *UsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersResponse.Merge(m, src)
}
func (m *UsersResponse) XXX_Size() int {
	return xxx_messageInfo_UsersResponse.Size(m)
}
func (m *UsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UsersResponse proto.InternalMessageInfo

func (m *UsersResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *UsersResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *UsersResponse) GetData() []*UsersResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UsersResponse) GetPager() *v1.Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

//Data.
type UsersResponse_Data struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Account              string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	CreatedTime          string   `protobuf:"bytes,4,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersResponse_Data) Reset()         { *m = UsersResponse_Data{} }
func (m *UsersResponse_Data) String() string { return proto.CompactTextString(m) }
func (*UsersResponse_Data) ProtoMessage()    {}
func (*UsersResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{3, 0}
}

func (m *UsersResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersResponse_Data.Unmarshal(m, b)
}
func (m *UsersResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersResponse_Data.Marshal(b, m, deterministic)
}
func (m *UsersResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersResponse_Data.Merge(m, src)
}
func (m *UsersResponse_Data) XXX_Size() int {
	return xxx_messageInfo_UsersResponse_Data.Size(m)
}
func (m *UsersResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_UsersResponse_Data proto.InternalMessageInfo

func (m *UsersResponse_Data) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UsersResponse_Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UsersResponse_Data) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UsersResponse_Data) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

//CheckAccountRequest.
type CheckAccountRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAccountRequest) Reset()         { *m = CheckAccountRequest{} }
func (m *CheckAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CheckAccountRequest) ProtoMessage()    {}
func (*CheckAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{4}
}

func (m *CheckAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAccountRequest.Unmarshal(m, b)
}
func (m *CheckAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAccountRequest.Marshal(b, m, deterministic)
}
func (m *CheckAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAccountRequest.Merge(m, src)
}
func (m *CheckAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CheckAccountRequest.Size(m)
}
func (m *CheckAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAccountRequest proto.InternalMessageInfo

func (m *CheckAccountRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

//CheckAccountResponse.
type CheckAccountResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckAccountResponse) Reset()         { *m = CheckAccountResponse{} }
func (m *CheckAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CheckAccountResponse) ProtoMessage()    {}
func (*CheckAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{5}
}

func (m *CheckAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckAccountResponse.Unmarshal(m, b)
}
func (m *CheckAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckAccountResponse.Marshal(b, m, deterministic)
}
func (m *CheckAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckAccountResponse.Merge(m, src)
}
func (m *CheckAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CheckAccountResponse.Size(m)
}
func (m *CheckAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckAccountResponse proto.InternalMessageInfo

func (m *CheckAccountResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CheckAccountResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

//CurrentUserRequest.
type CurrentUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentUserRequest) Reset()         { *m = CurrentUserRequest{} }
func (m *CurrentUserRequest) String() string { return proto.CompactTextString(m) }
func (*CurrentUserRequest) ProtoMessage()    {}
func (*CurrentUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{6}
}

func (m *CurrentUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentUserRequest.Unmarshal(m, b)
}
func (m *CurrentUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentUserRequest.Marshal(b, m, deterministic)
}
func (m *CurrentUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentUserRequest.Merge(m, src)
}
func (m *CurrentUserRequest) XXX_Size() int {
	return xxx_messageInfo_CurrentUserRequest.Size(m)
}
func (m *CurrentUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentUserRequest proto.InternalMessageInfo

//CurrentUserResponse.
type CurrentUserResponse struct {
	Code                 int32                     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string                    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *CurrentUserResponse_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CurrentUserResponse) Reset()         { *m = CurrentUserResponse{} }
func (m *CurrentUserResponse) String() string { return proto.CompactTextString(m) }
func (*CurrentUserResponse) ProtoMessage()    {}
func (*CurrentUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{7}
}

func (m *CurrentUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentUserResponse.Unmarshal(m, b)
}
func (m *CurrentUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentUserResponse.Marshal(b, m, deterministic)
}
func (m *CurrentUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentUserResponse.Merge(m, src)
}
func (m *CurrentUserResponse) XXX_Size() int {
	return xxx_messageInfo_CurrentUserResponse.Size(m)
}
func (m *CurrentUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentUserResponse proto.InternalMessageInfo

func (m *CurrentUserResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CurrentUserResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *CurrentUserResponse) GetData() *CurrentUserResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

//Data.
type CurrentUserResponse_Data struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Title                string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Phone                string   `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	CreatedTime          string   `protobuf:"bytes,7,opt,name=created_time,json=createdTime,proto3" json:"created_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrentUserResponse_Data) Reset()         { *m = CurrentUserResponse_Data{} }
func (m *CurrentUserResponse_Data) String() string { return proto.CompactTextString(m) }
func (*CurrentUserResponse_Data) ProtoMessage()    {}
func (*CurrentUserResponse_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_f449b9bea0d81ab6, []int{7, 0}
}

func (m *CurrentUserResponse_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrentUserResponse_Data.Unmarshal(m, b)
}
func (m *CurrentUserResponse_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrentUserResponse_Data.Marshal(b, m, deterministic)
}
func (m *CurrentUserResponse_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrentUserResponse_Data.Merge(m, src)
}
func (m *CurrentUserResponse_Data) XXX_Size() int {
	return xxx_messageInfo_CurrentUserResponse_Data.Size(m)
}
func (m *CurrentUserResponse_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrentUserResponse_Data.DiscardUnknown(m)
}

var xxx_messageInfo_CurrentUserResponse_Data proto.InternalMessageInfo

func (m *CurrentUserResponse_Data) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CurrentUserResponse_Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CurrentUserResponse_Data) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *CurrentUserResponse_Data) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CurrentUserResponse_Data) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CurrentUserResponse_Data) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *CurrentUserResponse_Data) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func init() {
	proto.RegisterType((*SaveUserRequest)(nil), "donech.system.v1.SaveUserRequest")
	proto.RegisterType((*SaveUserResponse)(nil), "donech.system.v1.SaveUserResponse")
	proto.RegisterType((*UsersRequest)(nil), "donech.system.v1.UsersRequest")
	proto.RegisterType((*UsersResponse)(nil), "donech.system.v1.UsersResponse")
	proto.RegisterType((*UsersResponse_Data)(nil), "donech.system.v1.UsersResponse.Data")
	proto.RegisterType((*CheckAccountRequest)(nil), "donech.system.v1.CheckAccountRequest")
	proto.RegisterType((*CheckAccountResponse)(nil), "donech.system.v1.CheckAccountResponse")
	proto.RegisterType((*CurrentUserRequest)(nil), "donech.system.v1.CurrentUserRequest")
	proto.RegisterType((*CurrentUserResponse)(nil), "donech.system.v1.CurrentUserResponse")
	proto.RegisterType((*CurrentUserResponse_Data)(nil), "donech.system.v1.CurrentUserResponse.Data")
}

func init() { proto.RegisterFile("donech/system/v1/system_api.proto", fileDescriptor_f449b9bea0d81ab6) }

var fileDescriptor_f449b9bea0d81ab6 = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xed, 0xa4, 0x3f, 0x93, 0xfe, 0x44, 0xdb, 0xb4, 0x32, 0x56, 0x05, 0xcd, 0xaa, 0x45,
	0x55, 0x91, 0x6c, 0x25, 0x5c, 0x2a, 0x84, 0x90, 0xfa, 0xc3, 0x81, 0x13, 0x91, 0x0b, 0x1c, 0xa0,
	0x52, 0xb4, 0xd8, 0xab, 0xc4, 0x6a, 0xec, 0x35, 0xf6, 0xc6, 0xc0, 0x09, 0x09, 0x89, 0x27, 0xe0,
	0x19, 0x10, 0xef, 0xc2, 0x15, 0x1e, 0x81, 0x07, 0x41, 0xfb, 0x93, 0xd6, 0xa9, 0x03, 0x21, 0x27,
	0x7b, 0x66, 0x67, 0xe6, 0x9b, 0xfd, 0xbe, 0x99, 0x85, 0x76, 0xc8, 0x12, 0x1a, 0x0c, 0xbd, 0xfc,
	0x63, 0xce, 0x69, 0xec, 0x15, 0x1d, 0xfd, 0xd7, 0x27, 0x69, 0xe4, 0xa6, 0x19, 0xe3, 0x0c, 0x35,
	0x55, 0x88, 0xab, 0x0e, 0xdc, 0xa2, 0xe3, 0xdc, 0xd1, 0x49, 0x63, 0x1e, 0x8d, 0x44, 0x8a, 0xf8,
	0xaa, 0x60, 0x67, 0x77, 0xc0, 0xd8, 0x60, 0x44, 0x3d, 0x92, 0x46, 0x1e, 0x49, 0x12, 0xc6, 0x09,
	0x8f, 0x58, 0x92, 0xab, 0x53, 0xfc, 0x06, 0x36, 0x2f, 0x48, 0x41, 0x5f, 0xe6, 0x34, 0xf3, 0xe9,
	0xbb, 0x31, 0xcd, 0x39, 0xb2, 0x61, 0x99, 0x04, 0x01, 0x1b, 0x27, 0xdc, 0x36, 0xf6, 0x8c, 0xc3,
	0x55, 0x7f, 0x62, 0x22, 0x04, 0xb5, 0x84, 0xc4, 0xd4, 0x36, 0xa5, 0x5b, 0xfe, 0x23, 0x07, 0x56,
	0x52, 0x92, 0xe7, 0xef, 0x59, 0x16, 0xda, 0x96, 0xf4, 0x5f, 0xdb, 0xf8, 0x18, 0x9a, 0x37, 0xc5,
	0xf3, 0x94, 0x25, 0x39, 0x15, 0x35, 0x02, 0x16, 0x52, 0x59, 0xba, 0xee, 0xcb, 0x7f, 0xd4, 0x04,
	0x2b, 0xce, 0x07, 0xba, 0xac, 0xf8, 0xc5, 0xcf, 0x61, 0x4d, 0x64, 0xe5, 0x93, 0x9e, 0x26, 0xc8,
	0x46, 0x09, 0xf9, 0x01, 0xd4, 0x53, 0x32, 0xa0, 0x99, 0xcc, 0x6b, 0x74, 0xb7, 0x5d, 0xcd, 0x8a,
	0xbc, 0x7b, 0xd1, 0x71, 0x7b, 0xe2, 0xd0, 0x57, 0x31, 0xf8, 0x8b, 0x09, 0xeb, 0xba, 0xe2, 0x22,
	0x8d, 0xa0, 0x63, 0xa8, 0x85, 0x84, 0x13, 0xdb, 0xda, 0xb3, 0x0e, 0x1b, 0xdd, 0x7d, 0xf7, 0x36,
	0xf3, 0xee, 0x54, 0x51, 0xf7, 0x9c, 0x70, 0xe2, 0xcb, 0x8c, 0x9b, 0xf6, 0x6a, 0xf3, 0xdb, 0x73,
	0x06, 0x50, 0x13, 0xa9, 0x68, 0x03, 0xcc, 0x28, 0x94, 0x2d, 0x59, 0xbe, 0x19, 0x85, 0x33, 0x19,
	0x2f, 0xe9, 0x63, 0x4d, 0xeb, 0xd3, 0x86, 0xb5, 0x20, 0xa3, 0x84, 0xd3, 0xb0, 0xcf, 0xa3, 0x98,
	0x4a, 0xe4, 0x55, 0xbf, 0xa1, 0x7d, 0x2f, 0xa2, 0x98, 0x62, 0x0f, 0xb6, 0xce, 0x86, 0x34, 0xb8,
	0x3a, 0x51, 0x29, 0x73, 0x35, 0xc7, 0x8f, 0xa1, 0x35, 0x9d, 0xb0, 0x90, 0x8e, 0x2d, 0x40, 0x67,
	0xe3, 0x2c, 0xa3, 0x09, 0x2f, 0x4d, 0x18, 0xfe, 0x66, 0xc2, 0xd6, 0x94, 0x7b, 0x21, 0x49, 0x9e,
	0x5c, 0x4b, 0x22, 0x78, 0x3d, 0xaa, 0x4a, 0x32, 0xa3, 0x74, 0x49, 0x18, 0xe7, 0xbb, 0xb1, 0x00,
	0xd9, 0x3b, 0xb0, 0x44, 0x0a, 0xc2, 0x49, 0xa6, 0xb9, 0xd6, 0x16, 0x6a, 0x41, 0x9d, 0xc6, 0x24,
	0x1a, 0x69, 0x8e, 0x95, 0x21, 0xbc, 0x3c, 0xe2, 0x23, 0x6a, 0xd7, 0x95, 0x57, 0x1a, 0xc2, 0x9b,
	0x0e, 0x59, 0x42, 0xed, 0x25, 0xe5, 0x95, 0x46, 0x45, 0xac, 0xe5, 0x8a, 0x58, 0xdd, 0x5f, 0x16,
	0xac, 0x5e, 0xc8, 0x6b, 0x9d, 0xf4, 0x9e, 0xa1, 0x2b, 0x58, 0x99, 0x6c, 0x13, 0x6a, 0x57, 0x6f,
	0x7d, 0x6b, 0x8d, 0x1d, 0xfc, 0xaf, 0x10, 0xc5, 0x0a, 0xb6, 0x3f, 0xff, 0xfc, 0xfd, 0xd5, 0x44,
	0x78, 0x5d, 0xbe, 0x0e, 0xe2, 0xdd, 0x10, 0xd3, 0xfc, 0xc8, 0x38, 0x42, 0x9f, 0x60, 0xad, 0x2c,
	0x3b, 0x3a, 0x98, 0x41, 0x73, 0x75, 0x8e, 0x9c, 0xfb, 0xf3, 0xc2, 0x34, 0xf0, 0x9e, 0x04, 0x76,
	0xf0, 0xf6, 0x04, 0x38, 0x10, 0x51, 0x9e, 0x1e, 0x3a, 0xd1, 0x40, 0x1f, 0xea, 0x72, 0xb5, 0xd0,
	0xdd, 0xbf, 0xee, 0x9c, 0x82, 0xbc, 0x37, 0x67, 0x27, 0xf1, 0xb6, 0xc4, 0xda, 0x44, 0xd3, 0x97,
	0x44, 0x1f, 0xa0, 0x51, 0x1a, 0x14, 0xb4, 0x3f, 0x67, 0x8e, 0x14, 0xd8, 0xc1, 0x7f, 0x4d, 0x1b,
	0xde, 0x95, 0x90, 0x3b, 0xa8, 0x55, 0x86, 0xf4, 0x02, 0x15, 0x79, 0xfa, 0x14, 0x5a, 0x01, 0x8b,
	0x2b, 0x95, 0x4e, 0x37, 0xb4, 0xd6, 0x69, 0xd4, 0x13, 0x6f, 0x73, 0xcf, 0x78, 0xbd, 0xa2, 0x0e,
	0x8b, 0xce, 0x0f, 0xb3, 0x79, 0x2e, 0xc3, 0x2f, 0x55, 0xcc, 0xe5, 0xab, 0xce, 0xdb, 0x25, 0xf9,
	0x82, 0x3f, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x0d, 0x51, 0x6f, 0x31, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SystemAPIClient is the client API for SystemAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SystemAPIClient interface {
	//SaveUser.
	SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error)
	//CheckAccount.
	CheckAccount(ctx context.Context, in *CheckAccountRequest, opts ...grpc.CallOption) (*CheckAccountResponse, error)
	//Users.
	Users(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	//CurrentUser.
	CurrentUser(ctx context.Context, in *CurrentUserRequest, opts ...grpc.CallOption) (*CurrentUserResponse, error)
}

type systemAPIClient struct {
	cc *grpc.ClientConn
}

func NewSystemAPIClient(cc *grpc.ClientConn) SystemAPIClient {
	return &systemAPIClient{cc}
}

func (c *systemAPIClient) SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error) {
	out := new(SaveUserResponse)
	err := c.cc.Invoke(ctx, "/donech.system.v1.SystemAPI/SaveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemAPIClient) CheckAccount(ctx context.Context, in *CheckAccountRequest, opts ...grpc.CallOption) (*CheckAccountResponse, error) {
	out := new(CheckAccountResponse)
	err := c.cc.Invoke(ctx, "/donech.system.v1.SystemAPI/CheckAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemAPIClient) Users(ctx context.Context, in *UsersRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, "/donech.system.v1.SystemAPI/Users", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemAPIClient) CurrentUser(ctx context.Context, in *CurrentUserRequest, opts ...grpc.CallOption) (*CurrentUserResponse, error) {
	out := new(CurrentUserResponse)
	err := c.cc.Invoke(ctx, "/donech.system.v1.SystemAPI/CurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemAPIServer is the server API for SystemAPI service.
type SystemAPIServer interface {
	//SaveUser.
	SaveUser(context.Context, *SaveUserRequest) (*SaveUserResponse, error)
	//CheckAccount.
	CheckAccount(context.Context, *CheckAccountRequest) (*CheckAccountResponse, error)
	//Users.
	Users(context.Context, *UsersRequest) (*UsersResponse, error)
	//CurrentUser.
	CurrentUser(context.Context, *CurrentUserRequest) (*CurrentUserResponse, error)
}

func RegisterSystemAPIServer(s *grpc.Server, srv SystemAPIServer) {
	s.RegisterService(&_SystemAPI_serviceDesc, srv)
}

func _SystemAPI_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemAPIServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/donech.system.v1.SystemAPI/SaveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemAPIServer).SaveUser(ctx, req.(*SaveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemAPI_CheckAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemAPIServer).CheckAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/donech.system.v1.SystemAPI/CheckAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemAPIServer).CheckAccount(ctx, req.(*CheckAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemAPI_Users_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemAPIServer).Users(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/donech.system.v1.SystemAPI/Users",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemAPIServer).Users(ctx, req.(*UsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemAPI_CurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemAPIServer).CurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/donech.system.v1.SystemAPI/CurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemAPIServer).CurrentUser(ctx, req.(*CurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SystemAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "donech.system.v1.SystemAPI",
	HandlerType: (*SystemAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _SystemAPI_SaveUser_Handler,
		},
		{
			MethodName: "CheckAccount",
			Handler:    _SystemAPI_CheckAccount_Handler,
		},
		{
			MethodName: "Users",
			Handler:    _SystemAPI_Users_Handler,
		},
		{
			MethodName: "CurrentUser",
			Handler:    _SystemAPI_CurrentUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "donech/system/v1/system_api.proto",
}
