// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: game.proto

package gamepb

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

type Begin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 房间号
	RoomCode int64 `protobuf:"varint,1,opt,name=roomCode,proto3" json:"roomCode,omitempty"`
	// map 选择的规则 key  匹配key  uint32 具体的值   根据匹配的规则定 当匹配到 需要bool 的时候 则0 为false  1 为true
	Rules map[string]uint32 `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// map 会话编号 玩家列表信息 key 为会话对象唯一标识
	Players map[int64]*Player `protobuf:"bytes,3,rep,name=players,proto3" json:"players,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 当前轮数
	RoundNum uint32 `protobuf:"varint,4,opt,name=roundNum,proto3" json:"roundNum,omitempty"`
	// 场次总人数
	PeopleNum uint32 `protobuf:"varint,5,opt,name=peopleNum,proto3" json:"peopleNum,omitempty"`
}

func (x *Begin) Reset() {
	*x = Begin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Begin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Begin) ProtoMessage() {}

func (x *Begin) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Begin.ProtoReflect.Descriptor instead.
func (*Begin) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *Begin) GetRoomCode() int64 {
	if x != nil {
		return x.RoomCode
	}
	return 0
}

func (x *Begin) GetRules() map[string]uint32 {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *Begin) GetPlayers() map[int64]*Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *Begin) GetRoundNum() uint32 {
	if x != nil {
		return x.RoundNum
	}
	return 0
}

func (x *Begin) GetPeopleNum() uint32 {
	if x != nil {
		return x.PeopleNum
	}
	return 0
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 玩家座位位置
	DeskPos uint32 `protobuf:"varint,1,opt,name=deskPos,proto3" json:"deskPos,omitempty"`
	// 玩家编号
	Uid int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	// 玩家昵称
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	// 玩家是否准备
	IsReady bool `protobuf:"varint,4,opt,name=isReady,proto3" json:"isReady,omitempty"`
	// 玩家性别
	Sex uint32 `protobuf:"varint,5,opt,name=sex,proto3" json:"sex,omitempty"`
	// 退出
	IsExit bool `protobuf:"varint,6,opt,name=isExit,proto3" json:"isExit,omitempty"`
	// 头像
	HeadUrl string `protobuf:"bytes,7,opt,name=headUrl,proto3" json:"headUrl,omitempty"`
	// 玩家余额
	Score float64 `protobuf:"fixed64,8,opt,name=score,proto3" json:"score,omitempty"`
	// 玩家ip地址
	Ip string `protobuf:"bytes,9,opt,name=ip,proto3" json:"ip,omitempty"`
	// 玩家是否离线
	Offline bool `protobuf:"varint,20,opt,name=offline,proto3" json:"offline,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *Player) GetDeskPos() uint32 {
	if x != nil {
		return x.DeskPos
	}
	return 0
}

func (x *Player) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *Player) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Player) GetIsReady() bool {
	if x != nil {
		return x.IsReady
	}
	return false
}

func (x *Player) GetSex() uint32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *Player) GetIsExit() bool {
	if x != nil {
		return x.IsExit
	}
	return false
}

func (x *Player) GetHeadUrl() string {
	if x != nil {
		return x.HeadUrl
	}
	return ""
}

func (x *Player) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Player) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Player) GetOffline() bool {
	if x != nil {
		return x.Offline
	}
	return false
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61,
	0x6d, 0x65, 0x22, 0xc3, 0x02, 0x0a, 0x05, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x72, 0x6f, 0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x42,
	0x65, 0x67, 0x69, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x6f,
	0x75, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x65, 0x6f, 0x70, 0x6c,
	0x65, 0x4e, 0x75, 0x6d, 0x1a, 0x38, 0x0a, 0x0a, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x48,
	0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xee, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x6b, 0x50, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x64, 0x65, 0x73, 0x6b, 0x50, 0x6f, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73,
	0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x45, 0x78, 0x69,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x45, 0x78, 0x69, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x70, 0x62, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_game_proto_goTypes = []interface{}{
	(*Begin)(nil),  // 0: game.Begin
	(*Player)(nil), // 1: game.Player
	nil,            // 2: game.Begin.RulesEntry
	nil,            // 3: game.Begin.PlayersEntry
}
var file_game_proto_depIdxs = []int32{
	2, // 0: game.Begin.rules:type_name -> game.Begin.RulesEntry
	3, // 1: game.Begin.players:type_name -> game.Begin.PlayersEntry
	1, // 2: game.Begin.PlayersEntry.value:type_name -> game.Player
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Begin); i {
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
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}