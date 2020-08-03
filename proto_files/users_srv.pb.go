// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto_files/users_srv.proto

package users

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UserRegistrationRequest struct {
	FirstName            string    `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string    `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	EmailId              string    `protobuf:"bytes,3,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	Password             string    `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword      string    `protobuf:"bytes,5,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
	Tags                 string    `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	Location             *Location `protobuf:"bytes,7,opt,name=Location,proto3" json:"Location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserRegistrationRequest) Reset()         { *m = UserRegistrationRequest{} }
func (m *UserRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*UserRegistrationRequest) ProtoMessage()    {}
func (*UserRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{0}
}

func (m *UserRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegistrationRequest.Unmarshal(m, b)
}
func (m *UserRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegistrationRequest.Marshal(b, m, deterministic)
}
func (m *UserRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegistrationRequest.Merge(m, src)
}
func (m *UserRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_UserRegistrationRequest.Size(m)
}
func (m *UserRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegistrationRequest proto.InternalMessageInfo

func (m *UserRegistrationRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserRegistrationRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserRegistrationRequest) GetEmailId() string {
	if m != nil {
		return m.EmailId
	}
	return ""
}

func (m *UserRegistrationRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserRegistrationRequest) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func (m *UserRegistrationRequest) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *UserRegistrationRequest) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type Location struct {
	Lat                  float32  `protobuf:"fixed32,1,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Lng                  float32  `protobuf:"fixed32,2,opt,name=Lng,proto3" json:"Lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Location) GetLng() float32 {
	if m != nil {
		return m.Lng
	}
	return 0
}

type UserRegistrationResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRegistrationResponse) Reset()         { *m = UserRegistrationResponse{} }
func (m *UserRegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*UserRegistrationResponse) ProtoMessage()    {}
func (*UserRegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{2}
}

func (m *UserRegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRegistrationResponse.Unmarshal(m, b)
}
func (m *UserRegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRegistrationResponse.Marshal(b, m, deterministic)
}
func (m *UserRegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRegistrationResponse.Merge(m, src)
}
func (m *UserRegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_UserRegistrationResponse.Size(m)
}
func (m *UserRegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserRegistrationResponse proto.InternalMessageInfo

func (m *UserRegistrationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UserRegistrationResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type UserLoginRequest struct {
	EmailId              string    `protobuf:"bytes,1,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	Password             string    `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Location             *Location `protobuf:"bytes,3,opt,name=Location,proto3" json:"Location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{3}
}

func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginRequest.Unmarshal(m, b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginRequest.Marshal(b, m, deterministic)
}
func (m *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(m, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return xxx_messageInfo_UserLoginRequest.Size(m)
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

func (m *UserLoginRequest) GetEmailId() string {
	if m != nil {
		return m.EmailId
	}
	return ""
}

func (m *UserLoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserLoginRequest) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type UserLoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{4}
}

func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResponse.Unmarshal(m, b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResponse.Marshal(b, m, deterministic)
}
func (m *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(m, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return xxx_messageInfo_UserLoginResponse.Size(m)
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

func (m *UserLoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *UserLoginResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type ListUserRequests struct {
	All                  string   `protobuf:"bytes,1,opt,name=all,proto3" json:"all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUserRequests) Reset()         { *m = ListUserRequests{} }
func (m *ListUserRequests) String() string { return proto.CompactTextString(m) }
func (*ListUserRequests) ProtoMessage()    {}
func (*ListUserRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{5}
}

func (m *ListUserRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserRequests.Unmarshal(m, b)
}
func (m *ListUserRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserRequests.Marshal(b, m, deterministic)
}
func (m *ListUserRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserRequests.Merge(m, src)
}
func (m *ListUserRequests) XXX_Size() int {
	return xxx_messageInfo_ListUserRequests.Size(m)
}
func (m *ListUserRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserRequests.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserRequests proto.InternalMessageInfo

func (m *ListUserRequests) GetAll() string {
	if m != nil {
		return m.All
	}
	return ""
}

type ListUserResponse struct {
	FirstName            string    `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string    `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	EmailId              string    `protobuf:"bytes,3,opt,name=email_id,json=emailId,proto3" json:"email_id,omitempty"`
	Password             string    `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ConfirmPassword      string    `protobuf:"bytes,5,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
	Tags                 string    `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	Location             *Location `protobuf:"bytes,7,opt,name=Location,proto3" json:"Location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListUserResponse) Reset()         { *m = ListUserResponse{} }
func (m *ListUserResponse) String() string { return proto.CompactTextString(m) }
func (*ListUserResponse) ProtoMessage()    {}
func (*ListUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{6}
}

func (m *ListUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUserResponse.Unmarshal(m, b)
}
func (m *ListUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUserResponse.Marshal(b, m, deterministic)
}
func (m *ListUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUserResponse.Merge(m, src)
}
func (m *ListUserResponse) XXX_Size() int {
	return xxx_messageInfo_ListUserResponse.Size(m)
}
func (m *ListUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUserResponse proto.InternalMessageInfo

func (m *ListUserResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *ListUserResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *ListUserResponse) GetEmailId() string {
	if m != nil {
		return m.EmailId
	}
	return ""
}

func (m *ListUserResponse) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ListUserResponse) GetConfirmPassword() string {
	if m != nil {
		return m.ConfirmPassword
	}
	return ""
}

func (m *ListUserResponse) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *ListUserResponse) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

type CreateRoleRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoleRequest) Reset()         { *m = CreateRoleRequest{} }
func (m *CreateRoleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoleRequest) ProtoMessage()    {}
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{7}
}

func (m *CreateRoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleRequest.Unmarshal(m, b)
}
func (m *CreateRoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleRequest.Merge(m, src)
}
func (m *CreateRoleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoleRequest.Size(m)
}
func (m *CreateRoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleRequest proto.InternalMessageInfo

func (m *CreateRoleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRoleRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateRoleResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	RoleId               string   `protobuf:"bytes,2,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoleResponse) Reset()         { *m = CreateRoleResponse{} }
func (m *CreateRoleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRoleResponse) ProtoMessage()    {}
func (*CreateRoleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{8}
}

func (m *CreateRoleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoleResponse.Unmarshal(m, b)
}
func (m *CreateRoleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoleResponse.Marshal(b, m, deterministic)
}
func (m *CreateRoleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoleResponse.Merge(m, src)
}
func (m *CreateRoleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRoleResponse.Size(m)
}
func (m *CreateRoleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoleResponse proto.InternalMessageInfo

func (m *CreateRoleResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreateRoleResponse) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

type CreatePermissionRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePermissionRequest) Reset()         { *m = CreatePermissionRequest{} }
func (m *CreatePermissionRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePermissionRequest) ProtoMessage()    {}
func (*CreatePermissionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{9}
}

func (m *CreatePermissionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePermissionRequest.Unmarshal(m, b)
}
func (m *CreatePermissionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePermissionRequest.Marshal(b, m, deterministic)
}
func (m *CreatePermissionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePermissionRequest.Merge(m, src)
}
func (m *CreatePermissionRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePermissionRequest.Size(m)
}
func (m *CreatePermissionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePermissionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePermissionRequest proto.InternalMessageInfo

func (m *CreatePermissionRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreatePermissionRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreatePermissionResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	PermissionId         string   `protobuf:"bytes,2,opt,name=permission_id,json=permissionId,proto3" json:"permission_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePermissionResponse) Reset()         { *m = CreatePermissionResponse{} }
func (m *CreatePermissionResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePermissionResponse) ProtoMessage()    {}
func (*CreatePermissionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{10}
}

func (m *CreatePermissionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePermissionResponse.Unmarshal(m, b)
}
func (m *CreatePermissionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePermissionResponse.Marshal(b, m, deterministic)
}
func (m *CreatePermissionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePermissionResponse.Merge(m, src)
}
func (m *CreatePermissionResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePermissionResponse.Size(m)
}
func (m *CreatePermissionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePermissionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePermissionResponse proto.InternalMessageInfo

func (m *CreatePermissionResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreatePermissionResponse) GetPermissionId() string {
	if m != nil {
		return m.PermissionId
	}
	return ""
}

type CreateTagRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTagRequest) Reset()         { *m = CreateTagRequest{} }
func (m *CreateTagRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTagRequest) ProtoMessage()    {}
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{11}
}

func (m *CreateTagRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTagRequest.Unmarshal(m, b)
}
func (m *CreateTagRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTagRequest.Marshal(b, m, deterministic)
}
func (m *CreateTagRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTagRequest.Merge(m, src)
}
func (m *CreateTagRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTagRequest.Size(m)
}
func (m *CreateTagRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTagRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTagRequest proto.InternalMessageInfo

func (m *CreateTagRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateTagRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateTagResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	TagId                string   `protobuf:"bytes,2,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTagResponse) Reset()         { *m = CreateTagResponse{} }
func (m *CreateTagResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTagResponse) ProtoMessage()    {}
func (*CreateTagResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fa5c1fd726f0fc2, []int{12}
}

func (m *CreateTagResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTagResponse.Unmarshal(m, b)
}
func (m *CreateTagResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTagResponse.Marshal(b, m, deterministic)
}
func (m *CreateTagResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTagResponse.Merge(m, src)
}
func (m *CreateTagResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTagResponse.Size(m)
}
func (m *CreateTagResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTagResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTagResponse proto.InternalMessageInfo

func (m *CreateTagResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreateTagResponse) GetTagId() string {
	if m != nil {
		return m.TagId
	}
	return ""
}

func init() {
	proto.RegisterType((*UserRegistrationRequest)(nil), "users.UserRegistrationRequest")
	proto.RegisterType((*Location)(nil), "users.Location")
	proto.RegisterType((*UserRegistrationResponse)(nil), "users.UserRegistrationResponse")
	proto.RegisterType((*UserLoginRequest)(nil), "users.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "users.UserLoginResponse")
	proto.RegisterType((*ListUserRequests)(nil), "users.ListUserRequests")
	proto.RegisterType((*ListUserResponse)(nil), "users.ListUserResponse")
	proto.RegisterType((*CreateRoleRequest)(nil), "users.CreateRoleRequest")
	proto.RegisterType((*CreateRoleResponse)(nil), "users.CreateRoleResponse")
	proto.RegisterType((*CreatePermissionRequest)(nil), "users.CreatePermissionRequest")
	proto.RegisterType((*CreatePermissionResponse)(nil), "users.CreatePermissionResponse")
	proto.RegisterType((*CreateTagRequest)(nil), "users.CreateTagRequest")
	proto.RegisterType((*CreateTagResponse)(nil), "users.CreateTagResponse")
}

func init() {
	proto.RegisterFile("proto_files/users_srv.proto", fileDescriptor_9fa5c1fd726f0fc2)
}

var fileDescriptor_9fa5c1fd726f0fc2 = []byte{
	// 606 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x55, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xee, 0x36, 0xcd, 0xd7, 0xb4, 0x55, 0x53, 0xeb, 0xad, 0xb2, 0x4d, 0xf5, 0x42, 0xb5, 0xe5,
	0x50, 0x84, 0x14, 0xa4, 0xf2, 0x07, 0x80, 0x22, 0xc1, 0x4a, 0x69, 0xa9, 0x96, 0xe4, 0x00, 0x97,
	0x95, 0xd9, 0x9d, 0x2c, 0x16, 0xfb, 0x11, 0x6c, 0x27, 0xfc, 0x0c, 0x7e, 0x2b, 0x37, 0x8e, 0xc8,
	0xf6, 0x7e, 0x18, 0x42, 0x14, 0x04, 0x37, 0x6e, 0x9e, 0x67, 0x66, 0x1f, 0x3f, 0xcf, 0x78, 0xec,
	0x85, 0xb3, 0x05, 0x2f, 0x64, 0x11, 0xce, 0x59, 0x8a, 0xe2, 0xf1, 0x52, 0x20, 0x17, 0xa1, 0xe0,
	0xab, 0xb1, 0x46, 0x49, 0x5b, 0x03, 0xde, 0x37, 0x07, 0x86, 0x33, 0x81, 0x3c, 0xc0, 0x84, 0x09,
	0xc9, 0xa9, 0x64, 0x45, 0x1e, 0xe0, 0xa7, 0x25, 0x0a, 0x49, 0xfe, 0x07, 0x98, 0x33, 0x2e, 0x64,
	0x98, 0xd3, 0x0c, 0x5d, 0xe7, 0xdc, 0xb9, 0xec, 0x07, 0x7d, 0x8d, 0xdc, 0xd2, 0x0c, 0xc9, 0x19,
	0xf4, 0x53, 0x5a, 0x65, 0x77, 0x75, 0xb6, 0xa7, 0x00, 0x9d, 0x3c, 0x85, 0x1e, 0x66, 0x94, 0xa5,
	0x21, 0x8b, 0xdd, 0x96, 0xce, 0x75, 0x75, 0xec, 0xc7, 0x64, 0x04, 0xbd, 0x05, 0x15, 0xe2, 0x73,
	0xc1, 0x63, 0x77, 0xcf, 0x7c, 0x56, 0xc5, 0xe4, 0x21, 0x0c, 0xa2, 0x22, 0x9f, 0x33, 0x9e, 0x85,
	0x75, 0x4d, 0x5b, 0xd7, 0x1c, 0x95, 0xf8, 0x5d, 0x55, 0x4a, 0x60, 0x4f, 0xd2, 0x44, 0xb8, 0x1d,
	0x9d, 0xd6, 0x6b, 0xf2, 0x08, 0x7a, 0x93, 0x22, 0xd2, 0x26, 0xdc, 0xee, 0xb9, 0x73, 0xb9, 0x7f,
	0x75, 0x34, 0xd6, 0x3e, 0xc7, 0x15, 0x1c, 0xd4, 0x05, 0xde, 0xb8, 0x29, 0x26, 0x03, 0x68, 0x4d,
	0xa8, 0xd4, 0x1e, 0x77, 0x03, 0xb5, 0xd4, 0x48, 0x9e, 0x68, 0x5f, 0x0a, 0xc9, 0x13, 0xef, 0x06,
	0xdc, 0xf5, 0x4e, 0x89, 0x45, 0x91, 0x0b, 0x24, 0x2e, 0x74, 0xc5, 0x32, 0x8a, 0x50, 0x08, 0xcd,
	0xd1, 0x0b, 0xaa, 0x90, 0x0c, 0xa1, 0xab, 0x14, 0xa8, 0x3e, 0x98, 0x1e, 0x75, 0x54, 0xe8, 0xc7,
	0xde, 0x0a, 0x06, 0x8a, 0x6e, 0x52, 0x24, 0xac, 0xee, 0xb8, 0xdd, 0x35, 0x67, 0x73, 0xd7, 0x76,
	0x7f, 0xea, 0x9a, 0x6d, 0xbb, 0xb5, 0xcd, 0xf6, 0x2d, 0x1c, 0x5b, 0xfb, 0x96, 0xfa, 0xff, 0x83,
	0xb6, 0x2c, 0x3e, 0x62, 0x5e, 0xee, 0x6a, 0x02, 0x72, 0x01, 0x87, 0x1c, 0xe7, 0x1c, 0xc5, 0x87,
	0xd0, 0x64, 0xcd, 0xc6, 0x07, 0x25, 0x38, 0x55, 0x98, 0xf7, 0x00, 0x06, 0x13, 0x26, 0xa4, 0x69,
	0x8d, 0xb6, 0x21, 0x54, 0xf3, 0x68, 0x9a, 0x96, 0x64, 0x6a, 0xe9, 0x7d, 0x75, 0xec, 0xb2, 0x72,
	0xd7, 0x7f, 0x7c, 0xc0, 0x7c, 0x38, 0xbe, 0xe6, 0x48, 0x25, 0x06, 0x45, 0x8a, 0xd5, 0x11, 0x13,
	0xd8, 0xb3, 0xdc, 0xea, 0x35, 0x39, 0x87, 0xfd, 0x18, 0x45, 0xc4, 0xd9, 0x42, 0x13, 0x1b, 0xab,
	0x36, 0xe4, 0xbd, 0x04, 0x62, 0x53, 0xfd, 0xce, 0xd4, 0xf1, 0x22, 0x45, 0x6b, 0xea, 0x54, 0xe8,
	0xc7, 0xde, 0x6b, 0x18, 0x1a, 0xa2, 0x3b, 0xe4, 0x19, 0x13, 0xc2, 0xba, 0xee, 0x7f, 0xa6, 0xec,
	0x2d, 0xb8, 0xeb, 0x84, 0x5b, 0xf5, 0x5d, 0xc0, 0xe1, 0xa2, 0xae, 0x6f, 0x54, 0x1e, 0x34, 0xa0,
	0x1f, 0x7b, 0xaf, 0x60, 0x60, 0xa8, 0xa7, 0x34, 0xf9, 0x3b, 0x91, 0x2f, 0xaa, 0x93, 0xd0, 0x4c,
	0x5b, 0xd5, 0x9d, 0x40, 0x47, 0xd2, 0xa4, 0x91, 0xd5, 0x96, 0x34, 0xf1, 0xe3, 0xab, 0x2f, 0x2d,
	0x38, 0x51, 0xf3, 0x7b, 0x43, 0x73, 0x9a, 0x60, 0x86, 0xb9, 0x7c, 0x83, 0x7c, 0xc5, 0x22, 0x24,
	0x33, 0x73, 0x97, 0xed, 0xa7, 0x81, 0xdc, 0x2b, 0x07, 0x63, 0xc3, 0xeb, 0x3a, 0xba, 0xbf, 0x31,
	0x6f, 0xf4, 0x79, 0x3b, 0xe4, 0x29, 0xf4, 0xeb, 0xab, 0x4a, 0x86, 0x56, 0xbd, 0xfd, 0x68, 0x8c,
	0xdc, 0xf5, 0x44, 0xcd, 0x70, 0x0d, 0xd0, 0xcc, 0x0d, 0xa9, 0x2a, 0xd7, 0xa6, 0x72, 0x74, 0xfa,
	0x8b, 0x4c, 0x4d, 0x32, 0xab, 0xce, 0xa1, 0x39, 0xe2, 0xda, 0xdd, 0x86, 0x61, 0xaa, 0xdd, 0x6d,
	0x9a, 0x0d, 0x6f, 0x87, 0x3c, 0xab, 0xb4, 0x4d, 0xd5, 0xcd, 0x1a, 0xfe, 0xf0, 0x41, 0x73, 0xe2,
	0x23, 0x77, 0x3d, 0x51, 0x51, 0x3c, 0xef, 0xbe, 0x33, 0xbf, 0xb1, 0xf7, 0x1d, 0xfd, 0x53, 0x7b,
	0xf2, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x4f, 0xef, 0xad, 0xf3, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserManagementServiceClient is the client API for UserManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserManagementServiceClient interface {
	UserRegistration(ctx context.Context, in *UserRegistrationRequest, opts ...grpc.CallOption) (*UserRegistrationResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error)
	CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error)
	CreateTags(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error)
}

type userManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementServiceClient(cc grpc.ClientConnInterface) UserManagementServiceClient {
	return &userManagementServiceClient{cc}
}

func (c *userManagementServiceClient) UserRegistration(ctx context.Context, in *UserRegistrationRequest, opts ...grpc.CallOption) (*UserRegistrationResponse, error) {
	out := new(UserRegistrationResponse)
	err := c.cc.Invoke(ctx, "/users.UserManagementService/UserRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/users.UserManagementService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*CreateRoleResponse, error) {
	out := new(CreateRoleResponse)
	err := c.cc.Invoke(ctx, "/users.UserManagementService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) CreatePermission(ctx context.Context, in *CreatePermissionRequest, opts ...grpc.CallOption) (*CreatePermissionResponse, error) {
	out := new(CreatePermissionResponse)
	err := c.cc.Invoke(ctx, "/users.UserManagementService/CreatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServiceClient) CreateTags(ctx context.Context, in *CreateTagRequest, opts ...grpc.CallOption) (*CreateTagResponse, error) {
	out := new(CreateTagResponse)
	err := c.cc.Invoke(ctx, "/users.UserManagementService/CreateTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServiceServer is the server API for UserManagementService service.
type UserManagementServiceServer interface {
	UserRegistration(context.Context, *UserRegistrationRequest) (*UserRegistrationResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	CreateRole(context.Context, *CreateRoleRequest) (*CreateRoleResponse, error)
	CreatePermission(context.Context, *CreatePermissionRequest) (*CreatePermissionResponse, error)
	CreateTags(context.Context, *CreateTagRequest) (*CreateTagResponse, error)
}

// UnimplementedUserManagementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserManagementServiceServer struct {
}

func (*UnimplementedUserManagementServiceServer) UserRegistration(ctx context.Context, req *UserRegistrationRequest) (*UserRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegistration not implemented")
}
func (*UnimplementedUserManagementServiceServer) UserLogin(ctx context.Context, req *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedUserManagementServiceServer) CreateRole(ctx context.Context, req *CreateRoleRequest) (*CreateRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (*UnimplementedUserManagementServiceServer) CreatePermission(ctx context.Context, req *CreatePermissionRequest) (*CreatePermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePermission not implemented")
}
func (*UnimplementedUserManagementServiceServer) CreateTags(ctx context.Context, req *CreateTagRequest) (*CreateTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTags not implemented")
}

func RegisterUserManagementServiceServer(s *grpc.Server, srv UserManagementServiceServer) {
	s.RegisterService(&_UserManagementService_serviceDesc, srv)
}

func _UserManagementService_UserRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).UserRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserManagementService/UserRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).UserRegistration(ctx, req.(*UserRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserManagementService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserManagementService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_CreatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).CreatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserManagementService/CreatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).CreatePermission(ctx, req.(*CreatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementService_CreateTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTagRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServiceServer).CreateTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserManagementService/CreateTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServiceServer).CreateTags(ctx, req.(*CreateTagRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "users.UserManagementService",
	HandlerType: (*UserManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegistration",
			Handler:    _UserManagementService_UserRegistration_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _UserManagementService_UserLogin_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _UserManagementService_CreateRole_Handler,
		},
		{
			MethodName: "CreatePermission",
			Handler:    _UserManagementService_CreatePermission_Handler,
		},
		{
			MethodName: "CreateTags",
			Handler:    _UserManagementService_CreateTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto_files/users_srv.proto",
}
