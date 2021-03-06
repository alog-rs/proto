// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: rs3.proto

package rs3

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

type Skill int32

const (
	Skill_ATTACK        Skill = 0
	Skill_DEFENCE       Skill = 1
	Skill_STRENGTH      Skill = 2
	Skill_CONSTITUTION  Skill = 3
	Skill_RANGED        Skill = 4
	Skill_PRAYER        Skill = 5
	Skill_MAGIC         Skill = 6
	Skill_COOKING       Skill = 7
	Skill_WOODCUTTING   Skill = 8
	Skill_FLETCHING     Skill = 9
	Skill_FISHING       Skill = 10
	Skill_FIREMAKING    Skill = 11
	Skill_CRAFTING      Skill = 12
	Skill_SMITHING      Skill = 13
	Skill_MINING        Skill = 14
	Skill_HERBLORE      Skill = 15
	Skill_AGILITY       Skill = 16
	Skill_THIEVING      Skill = 17
	Skill_SLAYER        Skill = 18
	Skill_FARMING       Skill = 19
	Skill_RUNECRAFTING  Skill = 20
	Skill_HUNTER        Skill = 21
	Skill_CONSTRUCTION  Skill = 22
	Skill_SUMMONING     Skill = 23
	Skill_DUNGEONEERING Skill = 24
	Skill_DIVINATION    Skill = 25
	Skill_INVENTION     Skill = 26
	Skill_ARCHAEOLOGY   Skill = 27
)

// Enum value maps for Skill.
var (
	Skill_name = map[int32]string{
		0:  "ATTACK",
		1:  "DEFENCE",
		2:  "STRENGTH",
		3:  "CONSTITUTION",
		4:  "RANGED",
		5:  "PRAYER",
		6:  "MAGIC",
		7:  "COOKING",
		8:  "WOODCUTTING",
		9:  "FLETCHING",
		10: "FISHING",
		11: "FIREMAKING",
		12: "CRAFTING",
		13: "SMITHING",
		14: "MINING",
		15: "HERBLORE",
		16: "AGILITY",
		17: "THIEVING",
		18: "SLAYER",
		19: "FARMING",
		20: "RUNECRAFTING",
		21: "HUNTER",
		22: "CONSTRUCTION",
		23: "SUMMONING",
		24: "DUNGEONEERING",
		25: "DIVINATION",
		26: "INVENTION",
		27: "ARCHAEOLOGY",
	}
	Skill_value = map[string]int32{
		"ATTACK":        0,
		"DEFENCE":       1,
		"STRENGTH":      2,
		"CONSTITUTION":  3,
		"RANGED":        4,
		"PRAYER":        5,
		"MAGIC":         6,
		"COOKING":       7,
		"WOODCUTTING":   8,
		"FLETCHING":     9,
		"FISHING":       10,
		"FIREMAKING":    11,
		"CRAFTING":      12,
		"SMITHING":      13,
		"MINING":        14,
		"HERBLORE":      15,
		"AGILITY":       16,
		"THIEVING":      17,
		"SLAYER":        18,
		"FARMING":       19,
		"RUNECRAFTING":  20,
		"HUNTER":        21,
		"CONSTRUCTION":  22,
		"SUMMONING":     23,
		"DUNGEONEERING": 24,
		"DIVINATION":    25,
		"INVENTION":     26,
		"ARCHAEOLOGY":   27,
	}
)

func (x Skill) Enum() *Skill {
	p := new(Skill)
	*p = x
	return p
}

func (x Skill) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Skill) Descriptor() protoreflect.EnumDescriptor {
	return file_rs3_proto_enumTypes[0].Descriptor()
}

func (Skill) Type() protoreflect.EnumType {
	return &file_rs3_proto_enumTypes[0]
}

func (x Skill) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Skill.Descriptor instead.
func (Skill) EnumDescriptor() ([]byte, []int) {
	return file_rs3_proto_rawDescGZIP(), []int{0}
}

type SkillData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skill        Skill `protobuf:"varint,1,opt,name=skill,proto3,enum=rs3.Skill" json:"skill,omitempty"`
	Rank         int32 `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Level        int32 `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	VirtualLevel int32 `protobuf:"varint,4,opt,name=virtual_level,json=virtualLevel,proto3" json:"virtual_level,omitempty"`
	Xp           int64 `protobuf:"varint,5,opt,name=xp,proto3" json:"xp,omitempty"`
}

func (x *SkillData) Reset() {
	*x = SkillData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rs3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkillData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkillData) ProtoMessage() {}

func (x *SkillData) ProtoReflect() protoreflect.Message {
	mi := &file_rs3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkillData.ProtoReflect.Descriptor instead.
func (*SkillData) Descriptor() ([]byte, []int) {
	return file_rs3_proto_rawDescGZIP(), []int{0}
}

func (x *SkillData) GetSkill() Skill {
	if x != nil {
		return x.Skill
	}
	return Skill_ATTACK
}

func (x *SkillData) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *SkillData) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *SkillData) GetVirtualLevel() int32 {
	if x != nil {
		return x.VirtualLevel
	}
	return 0
}

func (x *SkillData) GetXp() int64 {
	if x != nil {
		return x.Xp
	}
	return 0
}

type QuestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Completed  int32 `protobuf:"varint,1,opt,name=completed,proto3" json:"completed,omitempty"`
	Started    int32 `protobuf:"varint,2,opt,name=started,proto3" json:"started,omitempty"`
	NotStarted int32 `protobuf:"varint,3,opt,name=not_started,json=notStarted,proto3" json:"not_started,omitempty"`
}

func (x *QuestData) Reset() {
	*x = QuestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rs3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestData) ProtoMessage() {}

func (x *QuestData) ProtoReflect() protoreflect.Message {
	mi := &file_rs3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestData.ProtoReflect.Descriptor instead.
func (*QuestData) Descriptor() ([]byte, []int) {
	return file_rs3_proto_rawDescGZIP(), []int{1}
}

func (x *QuestData) GetCompleted() int32 {
	if x != nil {
		return x.Completed
	}
	return 0
}

func (x *QuestData) GetStarted() int32 {
	if x != nil {
		return x.Started
	}
	return 0
}

func (x *QuestData) GetNotStarted() int32 {
	if x != nil {
		return x.NotStarted
	}
	return 0
}

type PlayerActivityItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Long      string `protobuf:"bytes,2,opt,name=long,proto3" json:"long,omitempty"`
	Short     string `protobuf:"bytes,3,opt,name=short,proto3" json:"short,omitempty"`
}

func (x *PlayerActivityItem) Reset() {
	*x = PlayerActivityItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rs3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerActivityItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerActivityItem) ProtoMessage() {}

func (x *PlayerActivityItem) ProtoReflect() protoreflect.Message {
	mi := &file_rs3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerActivityItem.ProtoReflect.Descriptor instead.
func (*PlayerActivityItem) Descriptor() ([]byte, []int) {
	return file_rs3_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerActivityItem) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PlayerActivityItem) GetLong() string {
	if x != nil {
		return x.Long
	}
	return ""
}

func (x *PlayerActivityItem) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

type PlayerProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rank        int32                 `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	TotalLevel  int32                 `protobuf:"varint,3,opt,name=total_level,json=totalLevel,proto3" json:"total_level,omitempty"`
	TotalXp     int64                 `protobuf:"varint,4,opt,name=total_xp,json=totalXp,proto3" json:"total_xp,omitempty"`
	CombatLevel *int32                `protobuf:"varint,5,opt,name=combat_level,json=combatLevel,proto3,oneof" json:"combat_level,omitempty"` // Only available form RuneMetrics data source
	QuestInfo   *QuestData            `protobuf:"bytes,6,opt,name=quest_info,json=questInfo,proto3,oneof" json:"quest_info,omitempty"`        // Only available from RuneMetrics data source
	Skills      []*SkillData          `protobuf:"bytes,7,rep,name=skills,proto3" json:"skills,omitempty"`
	Activity    []*PlayerActivityItem `protobuf:"bytes,8,rep,name=activity,proto3" json:"activity,omitempty"` // Only available from RuneMetrics data source
}

func (x *PlayerProfile) Reset() {
	*x = PlayerProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rs3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerProfile) ProtoMessage() {}

func (x *PlayerProfile) ProtoReflect() protoreflect.Message {
	mi := &file_rs3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerProfile.ProtoReflect.Descriptor instead.
func (*PlayerProfile) Descriptor() ([]byte, []int) {
	return file_rs3_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerProfile) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *PlayerProfile) GetTotalLevel() int32 {
	if x != nil {
		return x.TotalLevel
	}
	return 0
}

func (x *PlayerProfile) GetTotalXp() int64 {
	if x != nil {
		return x.TotalXp
	}
	return 0
}

func (x *PlayerProfile) GetCombatLevel() int32 {
	if x != nil && x.CombatLevel != nil {
		return *x.CombatLevel
	}
	return 0
}

func (x *PlayerProfile) GetQuestInfo() *QuestData {
	if x != nil {
		return x.QuestInfo
	}
	return nil
}

func (x *PlayerProfile) GetSkills() []*SkillData {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *PlayerProfile) GetActivity() []*PlayerActivityItem {
	if x != nil {
		return x.Activity
	}
	return nil
}

type GetPlayerProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ActivityCount int32  `protobuf:"varint,2,opt,name=activity_count,json=activityCount,proto3" json:"activity_count,omitempty"`
}

func (x *GetPlayerProfileRequest) Reset() {
	*x = GetPlayerProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rs3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerProfileRequest) ProtoMessage() {}

func (x *GetPlayerProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rs3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerProfileRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerProfileRequest) Descriptor() ([]byte, []int) {
	return file_rs3_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlayerProfileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPlayerProfileRequest) GetActivityCount() int32 {
	if x != nil {
		return x.ActivityCount
	}
	return 0
}

var File_rs3_proto protoreflect.FileDescriptor

var file_rs3_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x73, 0x33,
	0x22, 0x8c, 0x01, 0x0a, 0x09, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20,
	0x0a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e,
	0x72, 0x73, 0x33, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x78, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x78, 0x70, 0x22,
	0x64, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x6f, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x22, 0x5c, 0x0a, 0x12, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x22, 0xcc, 0x02, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x19,
	0x0a, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x78, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x58, 0x70, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x62, 0x61, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48,
	0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01,
	0x01, 0x12, 0x32, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x73, 0x33, 0x2e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x73, 0x33, 0x2e, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x33, 0x0a,
	0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x72, 0x73, 0x33, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x54, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x99, 0x03, 0x0a, 0x05, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x54, 0x54, 0x41, 0x43, 0x4b, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x45, 0x46, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x54, 0x52, 0x45, 0x4e, 0x47, 0x54, 0x48, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e,
	0x53, 0x54, 0x49, 0x54, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x52, 0x41, 0x59, 0x45,
	0x52, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x47, 0x49, 0x43, 0x10, 0x06, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x57,
	0x4f, 0x4f, 0x44, 0x43, 0x55, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09,
	0x46, 0x4c, 0x45, 0x54, 0x43, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x46,
	0x49, 0x53, 0x48, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x46, 0x49, 0x52, 0x45,
	0x4d, 0x41, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x41, 0x46,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x4d, 0x49, 0x54, 0x48, 0x49,
	0x4e, 0x47, 0x10, 0x0d, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0e,
	0x12, 0x0c, 0x0a, 0x08, 0x48, 0x45, 0x52, 0x42, 0x4c, 0x4f, 0x52, 0x45, 0x10, 0x0f, 0x12, 0x0b,
	0x0a, 0x07, 0x41, 0x47, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x54,
	0x48, 0x49, 0x45, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x11, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4c, 0x41,
	0x59, 0x45, 0x52, 0x10, 0x12, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x41, 0x52, 0x4d, 0x49, 0x4e, 0x47,
	0x10, 0x13, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x55, 0x4e, 0x45, 0x43, 0x52, 0x41, 0x46, 0x54, 0x49,
	0x4e, 0x47, 0x10, 0x14, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x15,
	0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f, 0x4e, 0x53, 0x54, 0x52, 0x55, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x16, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x4d, 0x4d, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x17, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x55, 0x4e, 0x47, 0x45, 0x4f, 0x4e, 0x45, 0x45, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x18, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x56, 0x49, 0x4e, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x19, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x56, 0x45, 0x4e, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x1a, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x52, 0x43, 0x48, 0x41, 0x45, 0x4f, 0x4c, 0x4f,
	0x47, 0x59, 0x10, 0x1b, 0x32, 0x53, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x65, 0x73, 0x63, 0x61, 0x70,
	0x65, 0x12, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x72, 0x73, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x72, 0x73, 0x33, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x6f, 0x67, 0x2d, 0x72, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x73, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rs3_proto_rawDescOnce sync.Once
	file_rs3_proto_rawDescData = file_rs3_proto_rawDesc
)

func file_rs3_proto_rawDescGZIP() []byte {
	file_rs3_proto_rawDescOnce.Do(func() {
		file_rs3_proto_rawDescData = protoimpl.X.CompressGZIP(file_rs3_proto_rawDescData)
	})
	return file_rs3_proto_rawDescData
}

var file_rs3_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rs3_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rs3_proto_goTypes = []interface{}{
	(Skill)(0),                      // 0: rs3.Skill
	(*SkillData)(nil),               // 1: rs3.SkillData
	(*QuestData)(nil),               // 2: rs3.QuestData
	(*PlayerActivityItem)(nil),      // 3: rs3.PlayerActivityItem
	(*PlayerProfile)(nil),           // 4: rs3.PlayerProfile
	(*GetPlayerProfileRequest)(nil), // 5: rs3.GetPlayerProfileRequest
}
var file_rs3_proto_depIdxs = []int32{
	0, // 0: rs3.SkillData.skill:type_name -> rs3.Skill
	2, // 1: rs3.PlayerProfile.quest_info:type_name -> rs3.QuestData
	1, // 2: rs3.PlayerProfile.skills:type_name -> rs3.SkillData
	3, // 3: rs3.PlayerProfile.activity:type_name -> rs3.PlayerActivityItem
	5, // 4: rs3.Runescape.GetPlayerProfile:input_type -> rs3.GetPlayerProfileRequest
	4, // 5: rs3.Runescape.GetPlayerProfile:output_type -> rs3.PlayerProfile
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rs3_proto_init() }
func file_rs3_proto_init() {
	if File_rs3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rs3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkillData); i {
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
		file_rs3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuestData); i {
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
		file_rs3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerActivityItem); i {
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
		file_rs3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerProfile); i {
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
		file_rs3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerProfileRequest); i {
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
	file_rs3_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rs3_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rs3_proto_goTypes,
		DependencyIndexes: file_rs3_proto_depIdxs,
		EnumInfos:         file_rs3_proto_enumTypes,
		MessageInfos:      file_rs3_proto_msgTypes,
	}.Build()
	File_rs3_proto = out.File
	file_rs3_proto_rawDesc = nil
	file_rs3_proto_goTypes = nil
	file_rs3_proto_depIdxs = nil
}
