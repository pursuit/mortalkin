// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.1
// source: mortalkin/game.proto

package mortalkin_proto

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

type PlayGamePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	CharacterId int64     `protobuf:"varint,2,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
	Position    *Position `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *PlayGamePayload) Reset() {
	*x = PlayGamePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mortalkin_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayGamePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayGamePayload) ProtoMessage() {}

func (x *PlayGamePayload) ProtoReflect() protoreflect.Message {
	mi := &file_mortalkin_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayGamePayload.ProtoReflect.Descriptor instead.
func (*PlayGamePayload) Descriptor() ([]byte, []int) {
	return file_mortalkin_game_proto_rawDescGZIP(), []int{0}
}

func (x *PlayGamePayload) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *PlayGamePayload) GetCharacterId() int64 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

func (x *PlayGamePayload) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

type GameNotif struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharacterOnNotifs       []*Character         `protobuf:"bytes,1,rep,name=character_on_notifs,json=characterOnNotifs,proto3" json:"character_on_notifs,omitempty"`
	CharacterPositionNotifs []*CharacterPosition `protobuf:"bytes,2,rep,name=character_position_notifs,json=characterPositionNotifs,proto3" json:"character_position_notifs,omitempty"`
}

func (x *GameNotif) Reset() {
	*x = GameNotif{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mortalkin_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameNotif) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameNotif) ProtoMessage() {}

func (x *GameNotif) ProtoReflect() protoreflect.Message {
	mi := &file_mortalkin_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameNotif.ProtoReflect.Descriptor instead.
func (*GameNotif) Descriptor() ([]byte, []int) {
	return file_mortalkin_game_proto_rawDescGZIP(), []int{1}
}

func (x *GameNotif) GetCharacterOnNotifs() []*Character {
	if x != nil {
		return x.CharacterOnNotifs
	}
	return nil
}

func (x *GameNotif) GetCharacterPositionNotifs() []*CharacterPosition {
	if x != nil {
		return x.CharacterPositionNotifs
	}
	return nil
}

type Character struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position *Position `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Character) Reset() {
	*x = Character{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mortalkin_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Character) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Character) ProtoMessage() {}

func (x *Character) ProtoReflect() protoreflect.Message {
	mi := &file_mortalkin_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Character.ProtoReflect.Descriptor instead.
func (*Character) Descriptor() ([]byte, []int) {
	return file_mortalkin_game_proto_rawDescGZIP(), []int{2}
}

func (x *Character) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Character) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Character) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

type CharacterPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharacterId int64     `protobuf:"varint,1,opt,name=character_id,json=characterId,proto3" json:"character_id,omitempty"`
	Position    *Position `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *CharacterPosition) Reset() {
	*x = CharacterPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mortalkin_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharacterPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterPosition) ProtoMessage() {}

func (x *CharacterPosition) ProtoReflect() protoreflect.Message {
	mi := &file_mortalkin_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterPosition.ProtoReflect.Descriptor instead.
func (*CharacterPosition) Descriptor() ([]byte, []int) {
	return file_mortalkin_game_proto_rawDescGZIP(), []int{3}
}

func (x *CharacterPosition) GetCharacterId() int64 {
	if x != nil {
		return x.CharacterId
	}
	return 0
}

func (x *CharacterPosition) GetPosition() *Position {
	if x != nil {
		return x.Position
	}
	return nil
}

type Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Position) Reset() {
	*x = Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mortalkin_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Position) ProtoMessage() {}

func (x *Position) ProtoReflect() protoreflect.Message {
	mi := &file_mortalkin_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Position.ProtoReflect.Descriptor instead.
func (*Position) Descriptor() ([]byte, []int) {
	return file_mortalkin_game_proto_rawDescGZIP(), []int{4}
}

func (x *Position) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Position) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_mortalkin_game_proto protoreflect.FileDescriptor

var file_mortalkin_game_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x2f, 0x67, 0x61, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x70, 0x75, 0x72, 0x73, 0x75, 0x69, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x22, 0x87, 0x01,
	0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70,
	0x75, 0x72, 0x73, 0x75, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x6b, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc3, 0x01, 0x0a, 0x09, 0x47, 0x61, 0x6d, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x12, 0x50, 0x0a, 0x13, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x75, 0x72, 0x73, 0x75, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x52, 0x11, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x4f,
	0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x12, 0x64, 0x0a, 0x19, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x75, 0x72,
	0x73, 0x75, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6b,
	0x69, 0x6e, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x22, 0x6c, 0x0a,
	0x09, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b,
	0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x70, 0x75, 0x72, 0x73, 0x75, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x11, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x75, 0x72, 0x73, 0x75, 0x69, 0x74, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x2e, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x26, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x32, 0x5e, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x56, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x79, 0x12, 0x26, 0x2e, 0x70, 0x75, 0x72, 0x73, 0x75,
	0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x1a, 0x20, 0x2e, 0x70, 0x75, 0x72, 0x73, 0x75, 0x69, 0x74, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x75, 0x72, 0x73, 0x75, 0x69, 0x74, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6b, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mortalkin_game_proto_rawDescOnce sync.Once
	file_mortalkin_game_proto_rawDescData = file_mortalkin_game_proto_rawDesc
)

func file_mortalkin_game_proto_rawDescGZIP() []byte {
	file_mortalkin_game_proto_rawDescOnce.Do(func() {
		file_mortalkin_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_mortalkin_game_proto_rawDescData)
	})
	return file_mortalkin_game_proto_rawDescData
}

var file_mortalkin_game_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_mortalkin_game_proto_goTypes = []interface{}{
	(*PlayGamePayload)(nil),   // 0: pursuit.api.mortalkin.PlayGamePayload
	(*GameNotif)(nil),         // 1: pursuit.api.mortalkin.GameNotif
	(*Character)(nil),         // 2: pursuit.api.mortalkin.Character
	(*CharacterPosition)(nil), // 3: pursuit.api.mortalkin.CharacterPosition
	(*Position)(nil),          // 4: pursuit.api.mortalkin.Position
}
var file_mortalkin_game_proto_depIdxs = []int32{
	4, // 0: pursuit.api.mortalkin.PlayGamePayload.position:type_name -> pursuit.api.mortalkin.Position
	2, // 1: pursuit.api.mortalkin.GameNotif.character_on_notifs:type_name -> pursuit.api.mortalkin.Character
	3, // 2: pursuit.api.mortalkin.GameNotif.character_position_notifs:type_name -> pursuit.api.mortalkin.CharacterPosition
	4, // 3: pursuit.api.mortalkin.Character.position:type_name -> pursuit.api.mortalkin.Position
	4, // 4: pursuit.api.mortalkin.CharacterPosition.position:type_name -> pursuit.api.mortalkin.Position
	0, // 5: pursuit.api.mortalkin.Game.Play:input_type -> pursuit.api.mortalkin.PlayGamePayload
	1, // 6: pursuit.api.mortalkin.Game.Play:output_type -> pursuit.api.mortalkin.GameNotif
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_mortalkin_game_proto_init() }
func file_mortalkin_game_proto_init() {
	if File_mortalkin_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mortalkin_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayGamePayload); i {
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
		file_mortalkin_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameNotif); i {
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
		file_mortalkin_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Character); i {
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
		file_mortalkin_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharacterPosition); i {
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
		file_mortalkin_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Position); i {
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
			RawDescriptor: file_mortalkin_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mortalkin_game_proto_goTypes,
		DependencyIndexes: file_mortalkin_game_proto_depIdxs,
		MessageInfos:      file_mortalkin_game_proto_msgTypes,
	}.Build()
	File_mortalkin_game_proto = out.File
	file_mortalkin_game_proto_rawDesc = nil
	file_mortalkin_game_proto_goTypes = nil
	file_mortalkin_game_proto_depIdxs = nil
}
