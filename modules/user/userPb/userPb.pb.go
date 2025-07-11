// Version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: modules/user/userPb/userPb.proto

package userPb_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Structures
type UserProfile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{0}
}

func (x *UserProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserProfile) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetUserProfileReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserProfileReq) Reset() {
	*x = GetUserProfileReq{}
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserProfileReq) ProtoMessage() {}

func (x *GetUserProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserProfileReq.ProtoReflect.Descriptor instead.
func (*GetUserProfileReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserProfileReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateUserReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CredentialSearchReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CredentialSearchReq) Reset() {
	*x = CredentialSearchReq{}
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CredentialSearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialSearchReq) ProtoMessage() {}

func (x *CredentialSearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialSearchReq.ProtoReflect.Descriptor instead.
func (*CredentialSearchReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{3}
}

func (x *CredentialSearchReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CredentialSearchReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type FindOneUserProfileToRefreshReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FindOneUserProfileToRefreshReq) Reset() {
	*x = FindOneUserProfileToRefreshReq{}
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FindOneUserProfileToRefreshReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneUserProfileToRefreshReq) ProtoMessage() {}

func (x *FindOneUserProfileToRefreshReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneUserProfileToRefreshReq.ProtoReflect.Descriptor instead.
func (*FindOneUserProfileToRefreshReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneUserProfileToRefreshReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_modules_user_userPb_userPb_proto protoreflect.FileDescriptor

const file_modules_user_userPb_userPb_proto_rawDesc = "" +
	"\n" +
	" modules/user/userPb/userPb.proto\"f\n" +
	"\vUserProfile\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x1d\n" +
	"\n" +
	"created_at\x18\x04 \x01(\tR\tcreatedAt\"+\n" +
	"\x11GetUserProfileReq\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\tR\x06userId\"U\n" +
	"\rCreateUserReq\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"G\n" +
	"\x13CredentialSearchReq\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"8\n" +
	"\x1eFindOneUserProfileToRefreshReq\x12\x16\n" +
	"\x06userId\x18\x01 \x01(\tR\x06userId2\xf7\x01\n" +
	"\x0fUserGrpcService\x122\n" +
	"\x0eGetUserProfile\x12\x12.GetUserProfileReq\x1a\f.UserProfile\x12*\n" +
	"\n" +
	"CreateUser\x12\x0e.CreateUserReq\x1a\f.UserProfile\x126\n" +
	"\x10CredentialSearch\x12\x14.CredentialSearchReq\x1a\f.UserProfile\x12L\n" +
	"\x1bFindOneUserProfileToRefresh\x12\x1f.FindOneUserProfileToRefreshReq\x1a\f.UserProfileB/Z-user-manages/modules/user/userPb/userPb.protob\x06proto3"

var (
	file_modules_user_userPb_userPb_proto_rawDescOnce sync.Once
	file_modules_user_userPb_userPb_proto_rawDescData []byte
)

func file_modules_user_userPb_userPb_proto_rawDescGZIP() []byte {
	file_modules_user_userPb_userPb_proto_rawDescOnce.Do(func() {
		file_modules_user_userPb_userPb_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_modules_user_userPb_userPb_proto_rawDesc), len(file_modules_user_userPb_userPb_proto_rawDesc)))
	})
	return file_modules_user_userPb_userPb_proto_rawDescData
}

var file_modules_user_userPb_userPb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_modules_user_userPb_userPb_proto_goTypes = []any{
	(*UserProfile)(nil),                    // 0: UserProfile
	(*GetUserProfileReq)(nil),              // 1: GetUserProfileReq
	(*CreateUserReq)(nil),                  // 2: CreateUserReq
	(*CredentialSearchReq)(nil),            // 3: CredentialSearchReq
	(*FindOneUserProfileToRefreshReq)(nil), // 4: FindOneUserProfileToRefreshReq
}
var file_modules_user_userPb_userPb_proto_depIdxs = []int32{
	1, // 0: UserGrpcService.GetUserProfile:input_type -> GetUserProfileReq
	2, // 1: UserGrpcService.CreateUser:input_type -> CreateUserReq
	3, // 2: UserGrpcService.CredentialSearch:input_type -> CredentialSearchReq
	4, // 3: UserGrpcService.FindOneUserProfileToRefresh:input_type -> FindOneUserProfileToRefreshReq
	0, // 4: UserGrpcService.GetUserProfile:output_type -> UserProfile
	0, // 5: UserGrpcService.CreateUser:output_type -> UserProfile
	0, // 6: UserGrpcService.CredentialSearch:output_type -> UserProfile
	0, // 7: UserGrpcService.FindOneUserProfileToRefresh:output_type -> UserProfile
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_user_userPb_userPb_proto_init() }
func file_modules_user_userPb_userPb_proto_init() {
	if File_modules_user_userPb_userPb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_modules_user_userPb_userPb_proto_rawDesc), len(file_modules_user_userPb_userPb_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_user_userPb_userPb_proto_goTypes,
		DependencyIndexes: file_modules_user_userPb_userPb_proto_depIdxs,
		MessageInfos:      file_modules_user_userPb_userPb_proto_msgTypes,
	}.Build()
	File_modules_user_userPb_userPb_proto = out.File
	file_modules_user_userPb_userPb_proto_goTypes = nil
	file_modules_user_userPb_userPb_proto_depIdxs = nil
}
