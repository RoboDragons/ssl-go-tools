// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_match_stats.proto

package sslproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type TeamColor int32

const (
	TeamColor_TEAM_UNKNOWN TeamColor = 0
	TeamColor_TEAM_YELLOW  TeamColor = 1
	TeamColor_TEAM_BLUE    TeamColor = 2
)

var TeamColor_name = map[int32]string{
	0: "TEAM_UNKNOWN",
	1: "TEAM_YELLOW",
	2: "TEAM_BLUE",
}

var TeamColor_value = map[string]int32{
	"TEAM_UNKNOWN": 0,
	"TEAM_YELLOW":  1,
	"TEAM_BLUE":    2,
}

func (x TeamColor) String() string {
	return proto.EnumName(TeamColor_name, int32(x))
}

func (TeamColor) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{0}
}

type CommandType int32

const (
	CommandType_COMMAND_UNKNOWN         CommandType = 0
	CommandType_COMMAND_HALT            CommandType = 1
	CommandType_COMMAND_STOP            CommandType = 2
	CommandType_COMMAND_NORMAL_START    CommandType = 3
	CommandType_COMMAND_FORCE_START     CommandType = 4
	CommandType_COMMAND_PREPARE_KICKOFF CommandType = 5
	CommandType_COMMAND_PREPARE_PENALTY CommandType = 6
	CommandType_COMMAND_DIRECT_FREE     CommandType = 7
	CommandType_COMMAND_INDIRECT_FREE   CommandType = 8
	CommandType_COMMAND_TIMEOUT         CommandType = 9
	CommandType_COMMAND_BALL_PLACEMENT  CommandType = 10
)

var CommandType_name = map[int32]string{
	0:  "COMMAND_UNKNOWN",
	1:  "COMMAND_HALT",
	2:  "COMMAND_STOP",
	3:  "COMMAND_NORMAL_START",
	4:  "COMMAND_FORCE_START",
	5:  "COMMAND_PREPARE_KICKOFF",
	6:  "COMMAND_PREPARE_PENALTY",
	7:  "COMMAND_DIRECT_FREE",
	8:  "COMMAND_INDIRECT_FREE",
	9:  "COMMAND_TIMEOUT",
	10: "COMMAND_BALL_PLACEMENT",
}

var CommandType_value = map[string]int32{
	"COMMAND_UNKNOWN":         0,
	"COMMAND_HALT":            1,
	"COMMAND_STOP":            2,
	"COMMAND_NORMAL_START":    3,
	"COMMAND_FORCE_START":     4,
	"COMMAND_PREPARE_KICKOFF": 5,
	"COMMAND_PREPARE_PENALTY": 6,
	"COMMAND_DIRECT_FREE":     7,
	"COMMAND_INDIRECT_FREE":   8,
	"COMMAND_TIMEOUT":         9,
	"COMMAND_BALL_PLACEMENT":  10,
}

func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}

func (CommandType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{1}
}

type GamePhaseType int32

const (
	GamePhaseType_PHASE_UNKNOWN         GamePhaseType = 0
	GamePhaseType_PHASE_HALT            GamePhaseType = 1
	GamePhaseType_PHASE_STOP            GamePhaseType = 2
	GamePhaseType_PHASE_RUNNING         GamePhaseType = 3
	GamePhaseType_PHASE_TIMEOUT         GamePhaseType = 4
	GamePhaseType_PHASE_BREAK           GamePhaseType = 5
	GamePhaseType_PHASE_POST_GAME       GamePhaseType = 6
	GamePhaseType_PHASE_PREPARE_KICKOFF GamePhaseType = 7
	GamePhaseType_PHASE_PREPARE_PENALTY GamePhaseType = 8
	GamePhaseType_PHASE_DIRECT_FREE     GamePhaseType = 9
	GamePhaseType_PHASE_INDIRECT_FREE   GamePhaseType = 10
	GamePhaseType_PHASE_BALL_PLACEMENT  GamePhaseType = 11
)

var GamePhaseType_name = map[int32]string{
	0:  "PHASE_UNKNOWN",
	1:  "PHASE_HALT",
	2:  "PHASE_STOP",
	3:  "PHASE_RUNNING",
	4:  "PHASE_TIMEOUT",
	5:  "PHASE_BREAK",
	6:  "PHASE_POST_GAME",
	7:  "PHASE_PREPARE_KICKOFF",
	8:  "PHASE_PREPARE_PENALTY",
	9:  "PHASE_DIRECT_FREE",
	10: "PHASE_INDIRECT_FREE",
	11: "PHASE_BALL_PLACEMENT",
}

var GamePhaseType_value = map[string]int32{
	"PHASE_UNKNOWN":         0,
	"PHASE_HALT":            1,
	"PHASE_STOP":            2,
	"PHASE_RUNNING":         3,
	"PHASE_TIMEOUT":         4,
	"PHASE_BREAK":           5,
	"PHASE_POST_GAME":       6,
	"PHASE_PREPARE_KICKOFF": 7,
	"PHASE_PREPARE_PENALTY": 8,
	"PHASE_DIRECT_FREE":     9,
	"PHASE_INDIRECT_FREE":   10,
	"PHASE_BALL_PLACEMENT":  11,
}

func (x GamePhaseType) String() string {
	return proto.EnumName(GamePhaseType_name, int32(x))
}

func (GamePhaseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{2}
}

type Command struct {
	Type                 CommandType `protobuf:"varint,1,opt,name=type,proto3,enum=CommandType" json:"type,omitempty"`
	ForTeam              TeamColor   `protobuf:"varint,2,opt,name=for_team,json=forTeam,proto3,enum=TeamColor" json:"for_team,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{0}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetType() CommandType {
	if m != nil {
		return m.Type
	}
	return CommandType_COMMAND_UNKNOWN
}

func (m *Command) GetForTeam() TeamColor {
	if m != nil {
		return m.ForTeam
	}
	return TeamColor_TEAM_UNKNOWN
}

type GamePhase struct {
	StartTime            uint64        `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              uint64        `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Duration             uint32        `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Type                 GamePhaseType `protobuf:"varint,4,opt,name=type,proto3,enum=GamePhaseType" json:"type,omitempty"`
	CommandEntry         *Command      `protobuf:"bytes,5,opt,name=command_entry,json=commandEntry,proto3" json:"command_entry,omitempty"`
	CommandExit          *Command      `protobuf:"bytes,6,opt,name=command_exit,json=commandExit,proto3" json:"command_exit,omitempty"`
	ForTeam              TeamColor     `protobuf:"varint,7,opt,name=for_team,json=forTeam,proto3,enum=TeamColor" json:"for_team,omitempty"`
	GameEventsEntry      []*GameEvent  `protobuf:"bytes,8,rep,name=game_events_entry,json=gameEventsEntry,proto3" json:"game_events_entry,omitempty"`
	GameEventsExit       []*GameEvent  `protobuf:"bytes,9,rep,name=game_events_exit,json=gameEventsExit,proto3" json:"game_events_exit,omitempty"`
	NextCommandProposed  *Command      `protobuf:"bytes,10,opt,name=next_command_proposed,json=nextCommandProposed,proto3" json:"next_command_proposed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GamePhase) Reset()         { *m = GamePhase{} }
func (m *GamePhase) String() string { return proto.CompactTextString(m) }
func (*GamePhase) ProtoMessage()    {}
func (*GamePhase) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{1}
}

func (m *GamePhase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GamePhase.Unmarshal(m, b)
}
func (m *GamePhase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GamePhase.Marshal(b, m, deterministic)
}
func (m *GamePhase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GamePhase.Merge(m, src)
}
func (m *GamePhase) XXX_Size() int {
	return xxx_messageInfo_GamePhase.Size(m)
}
func (m *GamePhase) XXX_DiscardUnknown() {
	xxx_messageInfo_GamePhase.DiscardUnknown(m)
}

var xxx_messageInfo_GamePhase proto.InternalMessageInfo

func (m *GamePhase) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *GamePhase) GetEndTime() uint64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *GamePhase) GetDuration() uint32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *GamePhase) GetType() GamePhaseType {
	if m != nil {
		return m.Type
	}
	return GamePhaseType_PHASE_UNKNOWN
}

func (m *GamePhase) GetCommandEntry() *Command {
	if m != nil {
		return m.CommandEntry
	}
	return nil
}

func (m *GamePhase) GetCommandExit() *Command {
	if m != nil {
		return m.CommandExit
	}
	return nil
}

func (m *GamePhase) GetForTeam() TeamColor {
	if m != nil {
		return m.ForTeam
	}
	return TeamColor_TEAM_UNKNOWN
}

func (m *GamePhase) GetGameEventsEntry() []*GameEvent {
	if m != nil {
		return m.GameEventsEntry
	}
	return nil
}

func (m *GamePhase) GetGameEventsExit() []*GameEvent {
	if m != nil {
		return m.GameEventsExit
	}
	return nil
}

func (m *GamePhase) GetNextCommandProposed() *Command {
	if m != nil {
		return m.NextCommandProposed
	}
	return nil
}

type TeamStats struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Goals                uint32   `protobuf:"varint,2,opt,name=goals,proto3" json:"goals,omitempty"`
	Fouls                uint32   `protobuf:"varint,3,opt,name=fouls,proto3" json:"fouls,omitempty"`
	YellowCards          uint32   `protobuf:"varint,4,opt,name=yellow_cards,json=yellowCards,proto3" json:"yellow_cards,omitempty"`
	RedCards             uint32   `protobuf:"varint,5,opt,name=red_cards,json=redCards,proto3" json:"red_cards,omitempty"`
	PenaltyShotsTotal    uint32   `protobuf:"varint,6,opt,name=penalty_shots_total,json=penaltyShotsTotal,proto3" json:"penalty_shots_total,omitempty"`
	PenaltyShotsSuccess  uint32   `protobuf:"varint,7,opt,name=penalty_shots_success,json=penaltyShotsSuccess,proto3" json:"penalty_shots_success,omitempty"`
	TimeoutTime          uint32   `protobuf:"varint,9,opt,name=timeout_time,json=timeoutTime,proto3" json:"timeout_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamStats) Reset()         { *m = TeamStats{} }
func (m *TeamStats) String() string { return proto.CompactTextString(m) }
func (*TeamStats) ProtoMessage()    {}
func (*TeamStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{2}
}

func (m *TeamStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamStats.Unmarshal(m, b)
}
func (m *TeamStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamStats.Marshal(b, m, deterministic)
}
func (m *TeamStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamStats.Merge(m, src)
}
func (m *TeamStats) XXX_Size() int {
	return xxx_messageInfo_TeamStats.Size(m)
}
func (m *TeamStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamStats.DiscardUnknown(m)
}

var xxx_messageInfo_TeamStats proto.InternalMessageInfo

func (m *TeamStats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TeamStats) GetGoals() uint32 {
	if m != nil {
		return m.Goals
	}
	return 0
}

func (m *TeamStats) GetFouls() uint32 {
	if m != nil {
		return m.Fouls
	}
	return 0
}

func (m *TeamStats) GetYellowCards() uint32 {
	if m != nil {
		return m.YellowCards
	}
	return 0
}

func (m *TeamStats) GetRedCards() uint32 {
	if m != nil {
		return m.RedCards
	}
	return 0
}

func (m *TeamStats) GetPenaltyShotsTotal() uint32 {
	if m != nil {
		return m.PenaltyShotsTotal
	}
	return 0
}

func (m *TeamStats) GetPenaltyShotsSuccess() uint32 {
	if m != nil {
		return m.PenaltyShotsSuccess
	}
	return 0
}

func (m *TeamStats) GetTimeoutTime() uint32 {
	if m != nil {
		return m.TimeoutTime
	}
	return 0
}

type MatchStats struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	GamePhases           []*GamePhase       `protobuf:"bytes,2,rep,name=game_phases,json=gamePhases,proto3" json:"game_phases,omitempty"`
	MatchDuration        uint32             `protobuf:"varint,3,opt,name=match_duration,json=matchDuration,proto3" json:"match_duration,omitempty"`
	ExtraTime            bool               `protobuf:"varint,4,opt,name=extra_time,json=extraTime,proto3" json:"extra_time,omitempty"`
	Shootout             bool               `protobuf:"varint,5,opt,name=shootout,proto3" json:"shootout,omitempty"`
	TimePerGamePhase     map[string]uint32  `protobuf:"bytes,6,rep,name=time_per_game_phase,json=timePerGamePhase,proto3" json:"time_per_game_phase,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	RelTimePerGamePhase  map[string]float32 `protobuf:"bytes,7,rep,name=rel_time_per_game_phase,json=relTimePerGamePhase,proto3" json:"rel_time_per_game_phase,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	TeamStatsYellow      *TeamStats         `protobuf:"bytes,8,opt,name=team_stats_yellow,json=teamStatsYellow,proto3" json:"team_stats_yellow,omitempty"`
	TeamStatsBlue        *TeamStats         `protobuf:"bytes,9,opt,name=team_stats_blue,json=teamStatsBlue,proto3" json:"team_stats_blue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MatchStats) Reset()         { *m = MatchStats{} }
func (m *MatchStats) String() string { return proto.CompactTextString(m) }
func (*MatchStats) ProtoMessage()    {}
func (*MatchStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{3}
}

func (m *MatchStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchStats.Unmarshal(m, b)
}
func (m *MatchStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchStats.Marshal(b, m, deterministic)
}
func (m *MatchStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchStats.Merge(m, src)
}
func (m *MatchStats) XXX_Size() int {
	return xxx_messageInfo_MatchStats.Size(m)
}
func (m *MatchStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchStats.DiscardUnknown(m)
}

var xxx_messageInfo_MatchStats proto.InternalMessageInfo

func (m *MatchStats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MatchStats) GetGamePhases() []*GamePhase {
	if m != nil {
		return m.GamePhases
	}
	return nil
}

func (m *MatchStats) GetMatchDuration() uint32 {
	if m != nil {
		return m.MatchDuration
	}
	return 0
}

func (m *MatchStats) GetExtraTime() bool {
	if m != nil {
		return m.ExtraTime
	}
	return false
}

func (m *MatchStats) GetShootout() bool {
	if m != nil {
		return m.Shootout
	}
	return false
}

func (m *MatchStats) GetTimePerGamePhase() map[string]uint32 {
	if m != nil {
		return m.TimePerGamePhase
	}
	return nil
}

func (m *MatchStats) GetRelTimePerGamePhase() map[string]float32 {
	if m != nil {
		return m.RelTimePerGamePhase
	}
	return nil
}

func (m *MatchStats) GetTeamStatsYellow() *TeamStats {
	if m != nil {
		return m.TeamStatsYellow
	}
	return nil
}

func (m *MatchStats) GetTeamStatsBlue() *TeamStats {
	if m != nil {
		return m.TeamStatsBlue
	}
	return nil
}

func init() {
	proto.RegisterEnum("TeamColor", TeamColor_name, TeamColor_value)
	proto.RegisterEnum("CommandType", CommandType_name, CommandType_value)
	proto.RegisterEnum("GamePhaseType", GamePhaseType_name, GamePhaseType_value)
	proto.RegisterType((*Command)(nil), "Command")
	proto.RegisterType((*GamePhase)(nil), "GamePhase")
	proto.RegisterType((*TeamStats)(nil), "TeamStats")
	proto.RegisterType((*MatchStats)(nil), "MatchStats")
	proto.RegisterMapType((map[string]float32)(nil), "MatchStats.RelTimePerGamePhaseEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "MatchStats.TimePerGamePhaseEntry")
}

func init() { proto.RegisterFile("ssl_match_stats.proto", fileDescriptor_21b017991eb90546) }

var fileDescriptor_21b017991eb90546 = []byte{
	// 977 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x2e, 0x3f, 0x09, 0xf6, 0x01, 0x07, 0x67, 0x58, 0x1a, 0x87, 0xd5, 0x4a, 0x2c, 0xea, 0x4a,
	0x51, 0x56, 0x45, 0x2d, 0xad, 0xaa, 0xb6, 0x6a, 0x2f, 0x1c, 0x62, 0xb2, 0x28, 0x60, 0xac, 0xc1,
	0xd1, 0x2a, 0xbd, 0x19, 0x79, 0x61, 0x92, 0xa0, 0x1a, 0x8c, 0xec, 0x61, 0x1b, 0x5e, 0xa2, 0x17,
	0xbd, 0xea, 0x83, 0xf4, 0x59, 0xfa, 0x3c, 0xd5, 0xfc, 0x18, 0x0c, 0x49, 0xd4, 0xbb, 0x99, 0xef,
	0x3b, 0xe7, 0xf0, 0x9d, 0xef, 0x9c, 0x31, 0x50, 0x4f, 0x92, 0x90, 0xcc, 0x03, 0x36, 0x79, 0x20,
	0x09, 0x0b, 0x58, 0xd2, 0x5e, 0xc6, 0x11, 0x8b, 0x1a, 0xa7, 0x1c, 0xbe, 0x0f, 0xe6, 0x94, 0xd0,
	0xcf, 0x74, 0xc1, 0x48, 0xe7, 0x9b, 0x6f, 0x7f, 0x92, 0x54, 0x0b, 0x43, 0xa9, 0x1b, 0xcd, 0xe7,
	0xc1, 0x62, 0x8a, 0x9a, 0x50, 0x64, 0xeb, 0x25, 0xb5, 0x72, 0xcd, 0xdc, 0xd9, 0x51, 0xa7, 0xd2,
	0x56, 0xb8, 0xbf, 0x5e, 0x52, 0x2c, 0x18, 0xf4, 0x0e, 0xb4, 0xbb, 0x28, 0x26, 0x8c, 0x06, 0x73,
	0x2b, 0x2f, 0xa2, 0xa0, 0xed, 0xd3, 0x60, 0xde, 0x8d, 0xc2, 0x28, 0xc6, 0xa5, 0xbb, 0x28, 0xe6,
	0xb7, 0xd6, 0x3f, 0x05, 0xd0, 0xaf, 0x82, 0x39, 0xf5, 0x1e, 0x82, 0x84, 0xa2, 0x37, 0x00, 0x09,
	0x0b, 0x62, 0x46, 0xd8, 0x6c, 0x2e, 0x8b, 0x17, 0xb1, 0x2e, 0x10, 0x7f, 0x36, 0xa7, 0xe8, 0x14,
	0x34, 0xba, 0x98, 0x4a, 0x32, 0x2f, 0xc8, 0x12, 0x5d, 0x4c, 0x05, 0xd5, 0x00, 0x6d, 0xba, 0x8a,
	0x03, 0x36, 0x8b, 0x16, 0x56, 0xa1, 0x99, 0x3b, 0x33, 0xf0, 0xe6, 0x8e, 0x5a, 0x4a, 0x6c, 0x51,
	0xc8, 0x38, 0x6a, 0x6f, 0x7e, 0x2f, 0x23, 0xf7, 0x6b, 0x30, 0x26, 0xb2, 0x07, 0x42, 0x17, 0x2c,
	0x5e, 0x5b, 0x07, 0xcd, 0xdc, 0x59, 0xb9, 0xa3, 0xa5, 0x9d, 0xe1, 0x8a, 0xa2, 0x1d, 0xce, 0xa2,
	0xf7, 0x50, 0xd9, 0x84, 0x3f, 0xce, 0x98, 0x75, 0xb8, 0x17, 0x5d, 0x4e, 0xa3, 0x1f, 0x67, 0x6c,
	0xc7, 0x8a, 0xd2, 0x8b, 0x56, 0xa0, 0x1f, 0xe0, 0x78, 0xeb, 0x7b, 0xa2, 0x64, 0x68, 0xcd, 0xc2,
	0x59, 0xb9, 0x03, 0x42, 0xb3, 0xc3, 0x09, 0x5c, 0xbd, 0x4f, 0x8f, 0x89, 0xd4, 0xf2, 0x3d, 0x98,
	0x3b, 0x79, 0x5c, 0x8f, 0xfe, 0x24, 0xed, 0x28, 0x93, 0xc6, 0x45, 0xfd, 0x02, 0xf5, 0x05, 0x7d,
	0x64, 0x24, 0x6d, 0x63, 0x19, 0x47, 0xcb, 0x28, 0xa1, 0x53, 0x0b, 0xf6, 0x5a, 0xa9, 0xf1, 0x30,
	0x75, 0xf1, 0x54, 0x50, 0xeb, 0xaf, 0x3c, 0xe8, 0x5c, 0xf4, 0x98, 0x6f, 0x0e, 0x42, 0x50, 0x5c,
	0x04, 0x6a, 0x60, 0x3a, 0x16, 0x67, 0xf4, 0x0a, 0x0e, 0xee, 0xa3, 0x20, 0x4c, 0xc4, 0xa0, 0x0c,
	0x2c, 0x2f, 0x1c, 0xbd, 0x8b, 0x56, 0x61, 0xa2, 0x66, 0x24, 0x2f, 0xe8, 0x2d, 0x54, 0xd6, 0x34,
	0x0c, 0xa3, 0x3f, 0xc8, 0x24, 0x88, 0xa7, 0x89, 0x18, 0x94, 0x81, 0xcb, 0x12, 0xeb, 0x72, 0x08,
	0xbd, 0x06, 0x3d, 0xa6, 0x53, 0xc5, 0x1f, 0xc8, 0x01, 0xc7, 0x74, 0x2a, 0xc9, 0x36, 0xd4, 0x96,
	0x74, 0x11, 0x84, 0x6c, 0x4d, 0x92, 0x87, 0x88, 0x25, 0x84, 0x45, 0x2c, 0x08, 0xc5, 0x50, 0x0c,
	0x7c, 0xac, 0xa8, 0x31, 0x67, 0x7c, 0x4e, 0xa0, 0x0e, 0xd4, 0x77, 0xe3, 0x93, 0xd5, 0x64, 0x42,
	0x93, 0x44, 0x4c, 0xc7, 0xc0, 0xb5, 0x6c, 0xc6, 0x58, 0x52, 0x5c, 0x23, 0xdf, 0xbb, 0x68, 0xa5,
	0x96, 0x53, 0x97, 0x1a, 0x15, 0xc6, 0x77, 0xb0, 0xf5, 0x6f, 0x11, 0x60, 0xc8, 0x1f, 0xd4, 0xcb,
	0xae, 0xbc, 0x87, 0xb2, 0x98, 0xd5, 0x92, 0xaf, 0x1f, 0xf7, 0x66, 0x3b, 0x26, 0xb1, 0x91, 0x18,
	0xee, 0xd3, 0x63, 0x82, 0xde, 0xc1, 0x91, 0x7c, 0x9f, 0x7b, 0x9b, 0x6d, 0x08, 0xf4, 0x32, 0x5d,
	0xef, 0x37, 0x00, 0xf4, 0x91, 0xc5, 0x81, 0xd4, 0xc5, 0xbd, 0xd3, 0xb0, 0x2e, 0x90, 0xf4, 0x65,
	0x24, 0x0f, 0x51, 0xc4, 0xa2, 0x15, 0x13, 0xc6, 0x69, 0x78, 0x73, 0x47, 0x1e, 0xd4, 0x78, 0x12,
	0x59, 0xd2, 0x98, 0x6c, 0x75, 0x59, 0x87, 0x42, 0xd6, 0xdb, 0xf6, 0xb6, 0x99, 0x36, 0x2f, 0xe5,
	0xd1, 0x78, 0x23, 0x54, 0xac, 0x1e, 0x36, 0xd9, 0x1e, 0x8c, 0x7e, 0x83, 0x93, 0x98, 0x86, 0xe4,
	0xb9, 0xaa, 0x25, 0x51, 0xf5, 0xab, 0x6c, 0x55, 0x4c, 0xc3, 0xe7, 0x0b, 0xd7, 0xe2, 0xa7, 0x0c,
	0x7f, 0x20, 0xfc, 0x0d, 0xc9, 0xcf, 0x15, 0x91, 0xdb, 0x61, 0x69, 0x62, 0x5d, 0xe5, 0x83, 0x12,
	0x45, 0x71, 0x95, 0xa5, 0xc7, 0x5b, 0x11, 0x82, 0x3a, 0x50, 0xcd, 0xe4, 0x7d, 0x0a, 0x57, 0x72,
	0x7a, 0xbb, 0x59, 0xc6, 0x26, 0xeb, 0x22, 0x5c, 0xd1, 0x46, 0x17, 0xea, 0xcf, 0x2a, 0x43, 0x26,
	0x14, 0x7e, 0xa7, 0x6b, 0x35, 0x54, 0x7e, 0xe4, 0x3b, 0xfd, 0x39, 0xe0, 0x45, 0xd5, 0xa6, 0x8b,
	0xcb, 0xcf, 0xf9, 0x1f, 0x73, 0x8d, 0x1e, 0x58, 0x2f, 0x75, 0xf8, 0x7f, 0x75, 0xf2, 0x99, 0x3a,
	0xe7, 0xbf, 0xca, 0xc7, 0x26, 0xbe, 0x17, 0xc8, 0x84, 0x8a, 0xef, 0xd8, 0x43, 0x72, 0xe3, 0x5e,
	0xbb, 0xa3, 0x8f, 0xae, 0xf9, 0x05, 0xaa, 0x42, 0x59, 0x20, 0xb7, 0xce, 0x60, 0x30, 0xfa, 0x68,
	0xe6, 0x90, 0x01, 0xba, 0x00, 0x2e, 0x06, 0x37, 0x8e, 0x99, 0x3f, 0xff, 0x3b, 0x0f, 0xe5, 0xcc,
	0x07, 0x1a, 0xd5, 0xa0, 0xda, 0x1d, 0x0d, 0x87, 0xb6, 0x7b, 0x99, 0x29, 0x62, 0x42, 0x25, 0x05,
	0x3f, 0xd8, 0x03, 0xdf, 0xcc, 0x65, 0x91, 0xb1, 0x3f, 0xf2, 0xcc, 0x3c, 0xb2, 0xe0, 0x55, 0x8a,
	0xb8, 0x23, 0x3c, 0xb4, 0x07, 0x64, 0xec, 0xdb, 0xd8, 0x37, 0x0b, 0xe8, 0x04, 0x6a, 0x29, 0xd3,
	0x1b, 0xe1, 0xae, 0xa3, 0x88, 0x22, 0x7a, 0x0d, 0x27, 0x29, 0xe1, 0x61, 0xc7, 0xb3, 0xb1, 0x43,
	0xae, 0xfb, 0xdd, 0xeb, 0x51, 0xaf, 0x67, 0x1e, 0x3c, 0x47, 0x7a, 0x8e, 0x6b, 0x0f, 0xfc, 0x5b,
	0xf3, 0x30, 0x5b, 0xf2, 0xb2, 0x8f, 0x9d, 0xae, 0x4f, 0x7a, 0xd8, 0x71, 0xcc, 0x12, 0x3a, 0x85,
	0x7a, 0x4a, 0xf4, 0xdd, 0x2c, 0xa5, 0x65, 0x3b, 0xf3, 0xfb, 0x43, 0x67, 0x74, 0xe3, 0x9b, 0x3a,
	0x6a, 0xc0, 0x97, 0x29, 0x78, 0x61, 0x0f, 0x06, 0xc4, 0x1b, 0xd8, 0x5d, 0x67, 0xe8, 0xb8, 0xbe,
	0x09, 0xe7, 0x7f, 0xe6, 0xc1, 0xd8, 0xf9, 0x3b, 0x40, 0xc7, 0x60, 0x78, 0x1f, 0xec, 0xb1, 0x93,
	0xb1, 0xe6, 0x08, 0x40, 0x42, 0xca, 0x98, 0xcd, 0x5d, 0xd9, 0xb2, 0x49, 0xc1, 0x37, 0xae, 0xdb,
	0x77, 0xaf, 0xcc, 0xc2, 0x16, 0x4a, 0x65, 0x14, 0xf9, 0x94, 0x24, 0x74, 0x81, 0x1d, 0xfb, 0xda,
	0x3c, 0xe0, 0x62, 0x25, 0xe0, 0x8d, 0xc6, 0x3e, 0xb9, 0xb2, 0x87, 0x8e, 0x79, 0xc8, 0x9b, 0x53,
	0xe0, 0x9e, 0x5b, 0xa5, 0xa7, 0x54, 0xea, 0x95, 0x86, 0xea, 0x70, 0x2c, 0xa9, 0xac, 0x1d, 0x3a,
	0xb7, 0x50, 0xc2, 0xbb, 0x3e, 0x01, 0x1f, 0xa4, 0xd2, 0xb2, 0x6b, 0x48, 0xf9, 0xd3, 0xa1, 0xf8,
	0xab, 0xff, 0xee, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x2c, 0x77, 0x0f, 0x1e, 0x08, 0x00,
	0x00,
}
