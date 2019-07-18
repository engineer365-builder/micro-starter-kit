// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/account/proto/account/account.proto

package go_micro_srv_account

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type User struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            string   `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type UserRequest struct {
	Id        *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  *wrappers.StringValue `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName *wrappers.StringValue `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  *wrappers.StringValue `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	// string email = 5 [(validate.rules).string.email = true]; // Works
	// google.protobuf.StringValue email = 5 [(validate.rules).string.email = true]; // Not Working
	Email                *wrappers.StringValue `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{1}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetId() *wrappers.UInt32Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UserRequest) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *UserRequest) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *UserRequest) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *UserRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type UserResponse struct {
	Result               *User    `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetResult() *User {
	if m != nil {
		return m.Result
	}
	return nil
}

type UserExistResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserExistResponse) Reset()         { *m = UserExistResponse{} }
func (m *UserExistResponse) String() string { return proto.CompactTextString(m) }
func (*UserExistResponse) ProtoMessage()    {}
func (*UserExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{3}
}

func (m *UserExistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserExistResponse.Unmarshal(m, b)
}
func (m *UserExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserExistResponse.Marshal(b, m, deterministic)
}
func (m *UserExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserExistResponse.Merge(m, src)
}
func (m *UserExistResponse) XXX_Size() int {
	return xxx_messageInfo_UserExistResponse.Size(m)
}
func (m *UserExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserExistResponse proto.InternalMessageInfo

func (m *UserExistResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type UserListQuery struct {
	Limit     *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page      *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort      *wrappers.StringValue `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Username  *wrappers.StringValue `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	FirstName *wrappers.StringValue `protobuf:"bytes,5,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  *wrappers.StringValue `protobuf:"bytes,6,opt,name=lastName,proto3" json:"lastName,omitempty"`
	// string email = 7 [(validate.rules).string.email = true]; // Works
	// google.protobuf.StringValue email = 7 [(validate.rules).string.email = true]; // Not Working
	Email                *wrappers.StringValue `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserListQuery) Reset()         { *m = UserListQuery{} }
func (m *UserListQuery) String() string { return proto.CompactTextString(m) }
func (*UserListQuery) ProtoMessage()    {}
func (*UserListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{4}
}

func (m *UserListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListQuery.Unmarshal(m, b)
}
func (m *UserListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListQuery.Marshal(b, m, deterministic)
}
func (m *UserListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListQuery.Merge(m, src)
}
func (m *UserListQuery) XXX_Size() int {
	return xxx_messageInfo_UserListQuery.Size(m)
}
func (m *UserListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_UserListQuery proto.InternalMessageInfo

func (m *UserListQuery) GetLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *UserListQuery) GetPage() *wrappers.UInt32Value {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *UserListQuery) GetSort() *wrappers.StringValue {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *UserListQuery) GetUsername() *wrappers.StringValue {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *UserListQuery) GetFirstName() *wrappers.StringValue {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *UserListQuery) GetLastName() *wrappers.StringValue {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *UserListQuery) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

type UserListResponse struct {
	Results              []*User  `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total                uint32   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListResponse) Reset()         { *m = UserListResponse{} }
func (m *UserListResponse) String() string { return proto.CompactTextString(m) }
func (*UserListResponse) ProtoMessage()    {}
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{5}
}

func (m *UserListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListResponse.Unmarshal(m, b)
}
func (m *UserListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListResponse.Marshal(b, m, deterministic)
}
func (m *UserListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListResponse.Merge(m, src)
}
func (m *UserListResponse) XXX_Size() int {
	return xxx_messageInfo_UserListResponse.Size(m)
}
func (m *UserListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserListResponse proto.InternalMessageInfo

func (m *UserListResponse) GetResults() []*User {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *UserListResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type Profile struct {
	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tz     string `protobuf:"bytes,2,opt,name=tz,proto3" json:"tz,omitempty"`
	Avatar string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender string `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	// GenderType gender = 4; // enum
	// google.protobuf.Timestamp birthday = 5;
	UserId               uint32   `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            string   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{6}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Profile) GetTz() string {
	if m != nil {
		return m.Tz
	}
	return ""
}

func (m *Profile) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Profile) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Profile) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Profile) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Profile) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type ProfileRequest struct {
	Id                   *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Tz                   *wrappers.StringValue `protobuf:"bytes,3,opt,name=tz,proto3" json:"tz,omitempty"`
	Avatar               *wrappers.StringValue `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender               *wrappers.StringValue `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProfileRequest) Reset()         { *m = ProfileRequest{} }
func (m *ProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ProfileRequest) ProtoMessage()    {}
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{7}
}

func (m *ProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileRequest.Unmarshal(m, b)
}
func (m *ProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileRequest.Marshal(b, m, deterministic)
}
func (m *ProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileRequest.Merge(m, src)
}
func (m *ProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ProfileRequest.Size(m)
}
func (m *ProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileRequest proto.InternalMessageInfo

func (m *ProfileRequest) GetId() *wrappers.UInt32Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ProfileRequest) GetUserId() *wrappers.UInt32Value {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *ProfileRequest) GetTz() *wrappers.StringValue {
	if m != nil {
		return m.Tz
	}
	return nil
}

func (m *ProfileRequest) GetAvatar() *wrappers.StringValue {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *ProfileRequest) GetGender() *wrappers.StringValue {
	if m != nil {
		return m.Gender
	}
	return nil
}

type ProfileResponse struct {
	Result               *Profile `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileResponse) Reset()         { *m = ProfileResponse{} }
func (m *ProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ProfileResponse) ProtoMessage()    {}
func (*ProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{8}
}

func (m *ProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileResponse.Unmarshal(m, b)
}
func (m *ProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileResponse.Marshal(b, m, deterministic)
}
func (m *ProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileResponse.Merge(m, src)
}
func (m *ProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ProfileResponse.Size(m)
}
func (m *ProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileResponse proto.InternalMessageInfo

func (m *ProfileResponse) GetResult() *Profile {
	if m != nil {
		return m.Result
	}
	return nil
}

type ProfileListQuery struct {
	Limit                *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	UserId               *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Gender               *wrappers.StringValue `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ProfileListQuery) Reset()         { *m = ProfileListQuery{} }
func (m *ProfileListQuery) String() string { return proto.CompactTextString(m) }
func (*ProfileListQuery) ProtoMessage()    {}
func (*ProfileListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{9}
}

func (m *ProfileListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileListQuery.Unmarshal(m, b)
}
func (m *ProfileListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileListQuery.Marshal(b, m, deterministic)
}
func (m *ProfileListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileListQuery.Merge(m, src)
}
func (m *ProfileListQuery) XXX_Size() int {
	return xxx_messageInfo_ProfileListQuery.Size(m)
}
func (m *ProfileListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileListQuery proto.InternalMessageInfo

func (m *ProfileListQuery) GetLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *ProfileListQuery) GetPage() *wrappers.UInt32Value {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *ProfileListQuery) GetSort() *wrappers.StringValue {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *ProfileListQuery) GetUserId() *wrappers.UInt32Value {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *ProfileListQuery) GetGender() *wrappers.StringValue {
	if m != nil {
		return m.Gender
	}
	return nil
}

type ProfileListResponse struct {
	Results              []*Profile `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total                uint32     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProfileListResponse) Reset()         { *m = ProfileListResponse{} }
func (m *ProfileListResponse) String() string { return proto.CompactTextString(m) }
func (*ProfileListResponse) ProtoMessage()    {}
func (*ProfileListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f36848fb0b6d28e2, []int{10}
}

func (m *ProfileListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileListResponse.Unmarshal(m, b)
}
func (m *ProfileListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileListResponse.Marshal(b, m, deterministic)
}
func (m *ProfileListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileListResponse.Merge(m, src)
}
func (m *ProfileListResponse) XXX_Size() int {
	return xxx_messageInfo_ProfileListResponse.Size(m)
}
func (m *ProfileListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileListResponse proto.InternalMessageInfo

func (m *ProfileListResponse) GetResults() []*Profile {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ProfileListResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.account.User")
	proto.RegisterType((*UserRequest)(nil), "go.micro.srv.account.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "go.micro.srv.account.UserResponse")
	proto.RegisterType((*UserExistResponse)(nil), "go.micro.srv.account.UserExistResponse")
	proto.RegisterType((*UserListQuery)(nil), "go.micro.srv.account.UserListQuery")
	proto.RegisterType((*UserListResponse)(nil), "go.micro.srv.account.UserListResponse")
	proto.RegisterType((*Profile)(nil), "go.micro.srv.account.Profile")
	proto.RegisterType((*ProfileRequest)(nil), "go.micro.srv.account.ProfileRequest")
	proto.RegisterType((*ProfileResponse)(nil), "go.micro.srv.account.ProfileResponse")
	proto.RegisterType((*ProfileListQuery)(nil), "go.micro.srv.account.ProfileListQuery")
	proto.RegisterType((*ProfileListResponse)(nil), "go.micro.srv.account.ProfileListResponse")
}

func init() {
	proto.RegisterFile("srv/account/proto/account/account.proto", fileDescriptor_f36848fb0b6d28e2)
}

var fileDescriptor_f36848fb0b6d28e2 = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x18, 0xad, 0x1d, 0xc7, 0x49, 0xa7, 0xa4, 0x84, 0xa1, 0x82, 0x28, 0xb4, 0x28, 0x84, 0xd2, 0x86,
	0xaa, 0x71, 0x4a, 0x4a, 0x84, 0x8a, 0x00, 0x29, 0x29, 0x50, 0x55, 0xa2, 0x15, 0x75, 0x09, 0x48,
	0x54, 0x50, 0x4d, 0xed, 0xa9, 0x19, 0xc9, 0xb1, 0xc3, 0x78, 0x1c, 0xda, 0x22, 0x24, 0x2e, 0xb8,
	0xe2, 0x6a, 0x1f, 0x60, 0xaf, 0xf7, 0x0d, 0x56, 0xda, 0x17, 0xd8, 0xc7, 0xd8, 0xdb, 0x7d, 0x86,
	0x55, 0xae, 0x56, 0xf6, 0x8c, 0x9d, 0x9f, 0xa6, 0x89, 0x77, 0x37, 0x7b, 0xb1, 0x57, 0xc9, 0x37,
	0xdf, 0x8f, 0xe6, 0x3b, 0xe7, 0xcc, 0x31, 0xd8, 0xf4, 0x68, 0xaf, 0x86, 0x0c, 0xc3, 0xf5, 0x1d,
	0x56, 0xeb, 0x52, 0x97, 0xb9, 0x71, 0x24, 0x7e, 0xb5, 0xf0, 0x14, 0xae, 0x58, 0xae, 0xd6, 0x21,
	0x06, 0x75, 0x35, 0x8f, 0xf6, 0x34, 0x91, 0x2b, 0x7e, 0x68, 0xb9, 0xae, 0x65, 0x63, 0xde, 0x79,
	0xe1, 0x5f, 0xd6, 0xfe, 0xa2, 0xa8, 0xdb, 0xc5, 0xd4, 0xe3, 0x5d, 0xc5, 0xa6, 0x45, 0xd8, 0x1f,
	0xfe, 0x85, 0x66, 0xb8, 0x9d, 0x1a, 0x76, 0x7a, 0xee, 0x75, 0x97, 0xba, 0x57, 0xd7, 0xbc, 0xdc,
	0xa8, 0x5a, 0xd8, 0xa9, 0xf6, 0x90, 0x4d, 0x4c, 0xc4, 0x70, 0xed, 0xd6, 0x1f, 0x3e, 0xa2, 0xfc,
	0x58, 0x02, 0x4a, 0xdb, 0xc3, 0x14, 0x2e, 0x03, 0x99, 0x98, 0x05, 0xa9, 0x24, 0x55, 0x72, 0xba,
	0x4c, 0x4c, 0x58, 0x04, 0x59, 0xdf, 0xc3, 0xd4, 0x41, 0x1d, 0x5c, 0x90, 0x4b, 0x52, 0x65, 0x51,
	0x8f, 0x63, 0xb8, 0x0a, 0x16, 0x2f, 0x09, 0xf5, 0xd8, 0x71, 0x90, 0x4c, 0x85, 0xc9, 0xc1, 0x41,
	0xd0, 0x69, 0x23, 0x91, 0x54, 0x78, 0x67, 0x14, 0xc3, 0x15, 0x90, 0xc6, 0x1d, 0x44, 0xec, 0x42,
	0x3a, 0x4c, 0xf0, 0x00, 0xae, 0x01, 0x60, 0x50, 0x8c, 0x18, 0x36, 0xcf, 0x11, 0x2b, 0xa8, 0x7c,
	0xa0, 0x38, 0x69, 0xb2, 0x20, 0xed, 0x77, 0xcd, 0x28, 0x9d, 0xe1, 0x69, 0x71, 0xd2, 0x64, 0x5f,
	0xa6, 0x9e, 0xb5, 0xa4, 0xf2, 0x53, 0x19, 0x2c, 0x05, 0x7b, 0xe8, 0xf8, 0x4f, 0x1f, 0x7b, 0x0c,
	0x36, 0xe2, 0x75, 0x96, 0xea, 0xab, 0x1a, 0xc7, 0x51, 0x8b, 0x70, 0xd4, 0xda, 0x87, 0x0e, 0xdb,
	0xad, 0xff, 0x8c, 0x6c, 0x1f, 0xb7, 0x32, 0xfd, 0x96, 0xb2, 0x25, 0x57, 0xa4, 0x70, 0xeb, 0x5f,
	0xc7, 0xb6, 0x9e, 0xd4, 0x7c, 0xca, 0x28, 0x71, 0x2c, 0xde, 0x5c, 0xea, 0xb7, 0xd6, 0xe8, 0x07,
	0x79, 0xa5, 0x90, 0xaf, 0xfc, 0x2b, 0xd7, 0xe1, 0xef, 0x67, 0xa8, 0x7a, 0xb3, 0x53, 0xdd, 0x3b,
	0xaf, 0xfe, 0xf6, 0xf7, 0xee, 0xf6, 0x67, 0x8d, 0x7f, 0xd6, 0x87, 0x50, 0xdb, 0x1f, 0x47, 0x6d,
	0xd6, 0xf0, 0xe0, 0x66, 0x54, 0xce, 0xa7, 0x86, 0xc1, 0x6d, 0x8e, 0x81, 0x9b, 0x78, 0xc6, 0x80,
	0x83, 0xaf, 0x86, 0x39, 0x98, 0xd5, 0x9f, 0xed, 0xb7, 0xd2, 0xff, 0x4b, 0x72, 0x7e, 0x41, 0x70,
	0x55, 0x3e, 0x00, 0x6f, 0x71, 0x9c, 0xbd, 0xae, 0xeb, 0x78, 0x18, 0xd6, 0x81, 0x4a, 0xb1, 0xe7,
	0xdb, 0x4c, 0x80, 0x5d, 0xd4, 0x26, 0x49, 0x59, 0x0b, 0x7b, 0x44, 0x25, 0x67, 0x6c, 0x07, 0xbc,
	0x13, 0x1c, 0x7e, 0x77, 0x45, 0x3c, 0x16, 0x4f, 0x7b, 0x6f, 0x64, 0x5a, 0x76, 0xb4, 0xe3, 0x49,
	0x0a, 0xe4, 0x82, 0x96, 0x1f, 0x88, 0xc7, 0x4e, 0x7c, 0x4c, 0xaf, 0xe1, 0xd7, 0x20, 0x6d, 0x93,
	0x0e, 0x61, 0x89, 0x88, 0x5e, 0xec, 0xb7, 0xd4, 0x2d, 0xa5, 0x60, 0x56, 0x24, 0x9d, 0x77, 0xc1,
	0x3d, 0xa0, 0x74, 0x91, 0x75, 0x37, 0xd3, 0x13, 0x65, 0x12, 0xb6, 0xc0, 0x1d, 0xa0, 0x78, 0x2e,
	0x65, 0x49, 0x78, 0xd4, 0xc3, 0xca, 0x11, 0x69, 0x29, 0xaf, 0x53, 0x5a, 0xe9, 0x39, 0x48, 0x4b,
	0x7d, 0x45, 0x69, 0x65, 0x5e, 0x46, 0x5a, 0x06, 0xc8, 0x47, 0xf4, 0xc6, 0x82, 0xf8, 0x1c, 0x64,
	0xb8, 0x04, 0xbc, 0x82, 0x54, 0x4a, 0xcd, 0xd0, 0x57, 0x54, 0x1a, 0xd8, 0x0c, 0x73, 0x19, 0xb2,
	0x43, 0x66, 0x73, 0x3a, 0x0f, 0xb8, 0x88, 0x1e, 0x4a, 0x20, 0xf3, 0x23, 0x75, 0x2f, 0x89, 0x8d,
	0x6f, 0x79, 0xde, 0x32, 0x90, 0xd9, 0x8d, 0x70, 0x3b, 0x99, 0xdd, 0x04, 0x6a, 0x44, 0x3d, 0xc4,
	0x10, 0x15, 0x26, 0x27, 0xa2, 0xe0, 0xdc, 0xc2, 0x8e, 0x89, 0xa9, 0xf0, 0x37, 0x11, 0xc1, 0xf7,
	0x41, 0x26, 0xa0, 0xe4, 0x9c, 0x98, 0x21, 0x80, 0x39, 0x5d, 0x0d, 0xc2, 0x43, 0x73, 0xcc, 0xe0,
	0x32, 0xd3, 0x0d, 0x2e, 0x3b, 0xd1, 0xe0, 0x1e, 0xc8, 0x60, 0x59, 0xdc, 0x3b, 0xf2, 0xb8, 0xed,
	0xa4, 0x1e, 0x17, 0x2e, 0xd7, 0x18, 0x5c, 0x2e, 0x81, 0xde, 0xe3, 0xab, 0x6f, 0x87, 0x98, 0x24,
	0x91, 0x79, 0x80, 0xd8, 0x37, 0x31, 0x62, 0x4a, 0x42, 0x05, 0xd0, 0xd4, 0x3d, 0x49, 0x8a, 0x91,
	0x6d, 0xc6, 0xc8, 0x26, 0x51, 0xf1, 0x52, 0xbf, 0x95, 0xa5, 0xaa, 0x2e, 0x1d, 0xe9, 0xd2, 0xf7,
	0x11, 0x09, 0xe5, 0x23, 0xf0, 0x76, 0x8c, 0x93, 0x10, 0x51, 0x63, 0xcc, 0xa3, 0xd6, 0x26, 0x6b,
	0x28, 0x6a, 0x1b, 0x31, 0x9d, 0x47, 0x32, 0xc8, 0x8b, 0xc4, 0x9b, 0xe9, 0x3b, 0x43, 0xbc, 0x2b,
	0x2f, 0xc0, 0xfb, 0x1c, 0x98, 0x20, 0xe0, 0xdd, 0x21, 0xe4, 0x62, 0x36, 0xbe, 0x18, 0x7f, 0xd2,
	0x33, 0xe8, 0x48, 0xf0, 0xaa, 0xeb, 0xff, 0x29, 0xfc, 0xf3, 0x7f, 0x8a, 0x69, 0x8f, 0x18, 0x18,
	0xb6, 0x41, 0x3a, 0xfc, 0xb0, 0xc0, 0x8f, 0xa6, 0xd8, 0x05, 0x7f, 0x46, 0xc5, 0xcd, 0xbb, 0x4b,
	0x46, 0x3e, 0x4e, 0xe5, 0x05, 0xd8, 0x06, 0x4a, 0xb0, 0x0a, 0xfc, 0xf8, 0xee, 0x96, 0x58, 0x24,
	0xc5, 0x8d, 0xe9, 0x45, 0x43, 0x63, 0x8f, 0x41, 0xea, 0x00, 0x27, 0xba, 0x6b, 0x79, 0x5a, 0x49,
	0x3c, 0xef, 0x04, 0xa8, 0xfb, 0xa1, 0xb9, 0xcc, 0x75, 0x64, 0x3b, 0x34, 0xa4, 0xb9, 0x8e, 0xfc,
	0x16, 0xdb, 0x78, 0x8e, 0x23, 0xeb, 0xf7, 0x07, 0x26, 0x19, 0x29, 0xe1, 0x4c, 0x50, 0xb6, 0x31,
	0x55, 0x64, 0x03, 0xd6, 0x3e, 0x9d, 0x59, 0x37, 0xb4, 0xc2, 0x4f, 0x9c, 0xb8, 0xf5, 0xe9, 0x02,
	0x16, 0x2b, 0x7c, 0x32, 0xa3, 0x2a, 0x9e, 0xfa, 0x4b, 0x4c, 0xdf, 0x7c, 0x07, 0x5f, 0xa8, 0xe1,
	0xdb, 0xdd, 0x7d, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xa1, 0xbc, 0x20, 0x9e, 0x97, 0x0c, 0x00, 0x00,
}
