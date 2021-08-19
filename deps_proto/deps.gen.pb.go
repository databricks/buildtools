// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: deps_proto/deps.proto

package deps_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Dependency_Kind int32

const (
	Dependency_EXPLICIT   Dependency_Kind = 0
	Dependency_IMPLICIT   Dependency_Kind = 1
	Dependency_UNUSED     Dependency_Kind = 2
	Dependency_INCOMPLETE Dependency_Kind = 3
)

// Enum value maps for Dependency_Kind.
var (
	Dependency_Kind_name = map[int32]string{
		0: "EXPLICIT",
		1: "IMPLICIT",
		2: "UNUSED",
		3: "INCOMPLETE",
	}
	Dependency_Kind_value = map[string]int32{
		"EXPLICIT":   0,
		"IMPLICIT":   1,
		"UNUSED":     2,
		"INCOMPLETE": 3,
	}
)

func (x Dependency_Kind) Enum() *Dependency_Kind {
	p := new(Dependency_Kind)
	*p = x
	return p
}

func (x Dependency_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Dependency_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_deps_proto_deps_proto_enumTypes[0].Descriptor()
}

func (Dependency_Kind) Type() protoreflect.EnumType {
	return &file_deps_proto_deps_proto_enumTypes[0]
}

func (x Dependency_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Dependency_Kind) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Dependency_Kind(num)
	return nil
}

// Deprecated: Use Dependency_Kind.Descriptor instead.
func (Dependency_Kind) EnumDescriptor() ([]byte, []int) {
	return file_deps_proto_deps_proto_rawDescGZIP(), []int{1, 0}
}

type SourceLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Line   *int32  `protobuf:"varint,2,opt,name=line" json:"line,omitempty"`
	Column *int32  `protobuf:"varint,3,opt,name=column" json:"column,omitempty"`
}

func (x *SourceLocation) Reset() {
	*x = SourceLocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deps_proto_deps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceLocation) ProtoMessage() {}

func (x *SourceLocation) ProtoReflect() protoreflect.Message {
	mi := &file_deps_proto_deps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceLocation.ProtoReflect.Descriptor instead.
func (*SourceLocation) Descriptor() ([]byte, []int) {
	return file_deps_proto_deps_proto_rawDescGZIP(), []int{0}
}

func (x *SourceLocation) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *SourceLocation) GetLine() int32 {
	if x != nil && x.Line != nil {
		return *x.Line
	}
	return 0
}

func (x *SourceLocation) GetColumn() int32 {
	if x != nil && x.Column != nil {
		return *x.Column
	}
	return 0
}

type Dependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path     *string           `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Kind     *Dependency_Kind  `protobuf:"varint,2,req,name=kind,enum=blaze_deps.Dependency_Kind" json:"kind,omitempty"`
	Location []*SourceLocation `protobuf:"bytes,3,rep,name=location" json:"location,omitempty"`
}

func (x *Dependency) Reset() {
	*x = Dependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deps_proto_deps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dependency) ProtoMessage() {}

func (x *Dependency) ProtoReflect() protoreflect.Message {
	mi := &file_deps_proto_deps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dependency.ProtoReflect.Descriptor instead.
func (*Dependency) Descriptor() ([]byte, []int) {
	return file_deps_proto_deps_proto_rawDescGZIP(), []int{1}
}

func (x *Dependency) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *Dependency) GetKind() Dependency_Kind {
	if x != nil && x.Kind != nil {
		return *x.Kind
	}
	return Dependency_EXPLICIT
}

func (x *Dependency) GetLocation() []*SourceLocation {
	if x != nil {
		return x.Location
	}
	return nil
}

type Dependencies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dependency       []*Dependency `protobuf:"bytes,1,rep,name=dependency" json:"dependency,omitempty"`
	RuleLabel        *string       `protobuf:"bytes,2,opt,name=rule_label,json=ruleLabel" json:"rule_label,omitempty"`
	Success          *bool         `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
	ContainedPackage []string      `protobuf:"bytes,4,rep,name=contained_package,json=containedPackage" json:"contained_package,omitempty"`
}

func (x *Dependencies) Reset() {
	*x = Dependencies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_deps_proto_deps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dependencies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dependencies) ProtoMessage() {}

func (x *Dependencies) ProtoReflect() protoreflect.Message {
	mi := &file_deps_proto_deps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dependencies.ProtoReflect.Descriptor instead.
func (*Dependencies) Descriptor() ([]byte, []int) {
	return file_deps_proto_deps_proto_rawDescGZIP(), []int{2}
}

func (x *Dependencies) GetDependency() []*Dependency {
	if x != nil {
		return x.Dependency
	}
	return nil
}

func (x *Dependencies) GetRuleLabel() string {
	if x != nil && x.RuleLabel != nil {
		return *x.RuleLabel
	}
	return ""
}

func (x *Dependencies) GetSuccess() bool {
	if x != nil && x.Success != nil {
		return *x.Success
	}
	return false
}

func (x *Dependencies) GetContainedPackage() []string {
	if x != nil {
		return x.ContainedPackage
	}
	return nil
}

var File_deps_proto_deps_proto protoreflect.FileDescriptor

var file_deps_proto_deps_proto_rawDesc = []byte{
	0x0a, 0x15, 0x64, 0x65, 0x70, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x70,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6c, 0x61, 0x7a, 0x65, 0x5f, 0x64,
	0x65, 0x70, 0x73, 0x22, 0x50, 0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x22, 0xc9, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x6c, 0x61, 0x7a, 0x65, 0x5f, 0x64,
	0x65, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x36, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x6c,
	0x61, 0x7a, 0x65, 0x5f, 0x64, 0x65, 0x70, 0x73, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3e, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x50,
	0x4c, 0x49, 0x43, 0x49, 0x54, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4d, 0x50, 0x4c, 0x49,
	0x43, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x55, 0x53, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x03, 0x22, 0xac, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62, 0x6c, 0x61, 0x7a, 0x65, 0x5f, 0x64,
	0x65, 0x70, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x75,
	0x6c, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x75, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64,
	0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x42, 0x2a, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x6c, 0x69,
	0x62, 0x2e, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_deps_proto_deps_proto_rawDescOnce sync.Once
	file_deps_proto_deps_proto_rawDescData = file_deps_proto_deps_proto_rawDesc
)

func file_deps_proto_deps_proto_rawDescGZIP() []byte {
	file_deps_proto_deps_proto_rawDescOnce.Do(func() {
		file_deps_proto_deps_proto_rawDescData = protoimpl.X.CompressGZIP(file_deps_proto_deps_proto_rawDescData)
	})
	return file_deps_proto_deps_proto_rawDescData
}

var file_deps_proto_deps_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_deps_proto_deps_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_deps_proto_deps_proto_goTypes = []interface{}{
	(Dependency_Kind)(0),   // 0: blaze_deps.Dependency.Kind
	(*SourceLocation)(nil), // 1: blaze_deps.SourceLocation
	(*Dependency)(nil),     // 2: blaze_deps.Dependency
	(*Dependencies)(nil),   // 3: blaze_deps.Dependencies
}
var file_deps_proto_deps_proto_depIdxs = []int32{
	0, // 0: blaze_deps.Dependency.kind:type_name -> blaze_deps.Dependency.Kind
	1, // 1: blaze_deps.Dependency.location:type_name -> blaze_deps.SourceLocation
	2, // 2: blaze_deps.Dependencies.dependency:type_name -> blaze_deps.Dependency
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_deps_proto_deps_proto_init() }
func file_deps_proto_deps_proto_init() {
	if File_deps_proto_deps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_deps_proto_deps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceLocation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deps_proto_deps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dependency); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_deps_proto_deps_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dependencies); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_deps_proto_deps_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_deps_proto_deps_proto_goTypes,
		DependencyIndexes: file_deps_proto_deps_proto_depIdxs,
		EnumInfos:         file_deps_proto_deps_proto_enumTypes,
		MessageInfos:      file_deps_proto_deps_proto_msgTypes,
	}.Build()
	File_deps_proto_deps_proto = out.File
	file_deps_proto_deps_proto_rawDesc = nil
	file_deps_proto_deps_proto_goTypes = nil
	file_deps_proto_deps_proto_depIdxs = nil
}
