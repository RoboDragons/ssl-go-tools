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
	TeamColor_TEAM_NONE    TeamColor = 3
)

var TeamColor_name = map[int32]string{
	0: "TEAM_UNKNOWN",
	1: "TEAM_YELLOW",
	2: "TEAM_BLUE",
	3: "TEAM_NONE",
}

var TeamColor_value = map[string]int32{
	"TEAM_UNKNOWN": 0,
	"TEAM_YELLOW":  1,
	"TEAM_BLUE":    2,
	"TEAM_NONE":    3,
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
	CommandType_COMMAND_BALL_PLACEMENT  CommandType = 3
	CommandType_COMMAND_NORMAL_START    CommandType = 4
	CommandType_COMMAND_FORCE_START     CommandType = 5
	CommandType_COMMAND_DIRECT_FREE     CommandType = 6
	CommandType_COMMAND_INDIRECT_FREE   CommandType = 7
	CommandType_COMMAND_PREPARE_KICKOFF CommandType = 8
	CommandType_COMMAND_PREPARE_PENALTY CommandType = 9
	CommandType_COMMAND_TIMEOUT         CommandType = 10
	CommandType_COMMAND_GOAL            CommandType = 11
)

var CommandType_name = map[int32]string{
	0:  "COMMAND_UNKNOWN",
	1:  "COMMAND_HALT",
	2:  "COMMAND_STOP",
	3:  "COMMAND_BALL_PLACEMENT",
	4:  "COMMAND_NORMAL_START",
	5:  "COMMAND_FORCE_START",
	6:  "COMMAND_DIRECT_FREE",
	7:  "COMMAND_INDIRECT_FREE",
	8:  "COMMAND_PREPARE_KICKOFF",
	9:  "COMMAND_PREPARE_PENALTY",
	10: "COMMAND_TIMEOUT",
	11: "COMMAND_GOAL",
}

var CommandType_value = map[string]int32{
	"COMMAND_UNKNOWN":         0,
	"COMMAND_HALT":            1,
	"COMMAND_STOP":            2,
	"COMMAND_BALL_PLACEMENT":  3,
	"COMMAND_NORMAL_START":    4,
	"COMMAND_FORCE_START":     5,
	"COMMAND_DIRECT_FREE":     6,
	"COMMAND_INDIRECT_FREE":   7,
	"COMMAND_PREPARE_KICKOFF": 8,
	"COMMAND_PREPARE_PENALTY": 9,
	"COMMAND_TIMEOUT":         10,
	"COMMAND_GOAL":            11,
}

func (x CommandType) String() string {
	return proto.EnumName(CommandType_name, int32(x))
}

func (CommandType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{1}
}

type StageType int32

const (
	StageType_STAGE_UNKNOWN                StageType = 0
	StageType_STAGE_NORMAL_FIRST_HALF_PRE  StageType = 1
	StageType_STAGE_NORMAL_FIRST_HALF      StageType = 2
	StageType_STAGE_NORMAL_HALF_TIME       StageType = 3
	StageType_STAGE_NORMAL_SECOND_HALF_PRE StageType = 4
	StageType_STAGE_NORMAL_SECOND_HALF     StageType = 5
	StageType_STAGE_EXTRA_TIME_BREAK       StageType = 6
	StageType_STAGE_EXTRA_FIRST_HALF_PRE   StageType = 7
	StageType_STAGE_EXTRA_FIRST_HALF       StageType = 8
	StageType_STAGE_EXTRA_HALF_TIME        StageType = 9
	StageType_STAGE_EXTRA_SECOND_HALF_PRE  StageType = 10
	StageType_STAGE_EXTRA_SECOND_HALF      StageType = 11
	StageType_STAGE_PENALTY_SHOOTOUT_BREAK StageType = 12
	StageType_STAGE_PENALTY_SHOOTOUT       StageType = 13
	StageType_STAGE_POST_GAME              StageType = 14
)

var StageType_name = map[int32]string{
	0:  "STAGE_UNKNOWN",
	1:  "STAGE_NORMAL_FIRST_HALF_PRE",
	2:  "STAGE_NORMAL_FIRST_HALF",
	3:  "STAGE_NORMAL_HALF_TIME",
	4:  "STAGE_NORMAL_SECOND_HALF_PRE",
	5:  "STAGE_NORMAL_SECOND_HALF",
	6:  "STAGE_EXTRA_TIME_BREAK",
	7:  "STAGE_EXTRA_FIRST_HALF_PRE",
	8:  "STAGE_EXTRA_FIRST_HALF",
	9:  "STAGE_EXTRA_HALF_TIME",
	10: "STAGE_EXTRA_SECOND_HALF_PRE",
	11: "STAGE_EXTRA_SECOND_HALF",
	12: "STAGE_PENALTY_SHOOTOUT_BREAK",
	13: "STAGE_PENALTY_SHOOTOUT",
	14: "STAGE_POST_GAME",
}

var StageType_value = map[string]int32{
	"STAGE_UNKNOWN":                0,
	"STAGE_NORMAL_FIRST_HALF_PRE":  1,
	"STAGE_NORMAL_FIRST_HALF":      2,
	"STAGE_NORMAL_HALF_TIME":       3,
	"STAGE_NORMAL_SECOND_HALF_PRE": 4,
	"STAGE_NORMAL_SECOND_HALF":     5,
	"STAGE_EXTRA_TIME_BREAK":       6,
	"STAGE_EXTRA_FIRST_HALF_PRE":   7,
	"STAGE_EXTRA_FIRST_HALF":       8,
	"STAGE_EXTRA_HALF_TIME":        9,
	"STAGE_EXTRA_SECOND_HALF_PRE":  10,
	"STAGE_EXTRA_SECOND_HALF":      11,
	"STAGE_PENALTY_SHOOTOUT_BREAK": 12,
	"STAGE_PENALTY_SHOOTOUT":       13,
	"STAGE_POST_GAME":              14,
}

func (x StageType) String() string {
	return proto.EnumName(StageType_name, int32(x))
}

func (StageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{2}
}

type GamePhaseType int32

const (
	GamePhaseType_PHASE_UNKNOWN         GamePhaseType = 0
	GamePhaseType_PHASE_RUNNING         GamePhaseType = 1
	GamePhaseType_PHASE_DIRECT_FREE     GamePhaseType = 2
	GamePhaseType_PHASE_INDIRECT_FREE   GamePhaseType = 3
	GamePhaseType_PHASE_PREPARE_KICKOFF GamePhaseType = 4
	GamePhaseType_PHASE_PREPARE_PENALTY GamePhaseType = 5
	GamePhaseType_PHASE_STOP            GamePhaseType = 6
	GamePhaseType_PHASE_BALL_PLACEMENT  GamePhaseType = 7
	GamePhaseType_PHASE_TIMEOUT         GamePhaseType = 8
	GamePhaseType_PHASE_BREAK           GamePhaseType = 9
	GamePhaseType_PHASE_HALT            GamePhaseType = 10
	GamePhaseType_PHASE_POST_GAME       GamePhaseType = 11
)

var GamePhaseType_name = map[int32]string{
	0:  "PHASE_UNKNOWN",
	1:  "PHASE_RUNNING",
	2:  "PHASE_DIRECT_FREE",
	3:  "PHASE_INDIRECT_FREE",
	4:  "PHASE_PREPARE_KICKOFF",
	5:  "PHASE_PREPARE_PENALTY",
	6:  "PHASE_STOP",
	7:  "PHASE_BALL_PLACEMENT",
	8:  "PHASE_TIMEOUT",
	9:  "PHASE_BREAK",
	10: "PHASE_HALT",
	11: "PHASE_POST_GAME",
}

var GamePhaseType_value = map[string]int32{
	"PHASE_UNKNOWN":         0,
	"PHASE_RUNNING":         1,
	"PHASE_DIRECT_FREE":     2,
	"PHASE_INDIRECT_FREE":   3,
	"PHASE_PREPARE_KICKOFF": 4,
	"PHASE_PREPARE_PENALTY": 5,
	"PHASE_STOP":            6,
	"PHASE_BALL_PLACEMENT":  7,
	"PHASE_TIMEOUT":         8,
	"PHASE_BREAK":           9,
	"PHASE_HALT":            10,
	"PHASE_POST_GAME":       11,
}

func (x GamePhaseType) String() string {
	return proto.EnumName(GamePhaseType_name, int32(x))
}

func (GamePhaseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{3}
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
	Stage                StageType     `protobuf:"varint,11,opt,name=stage,proto3,enum=StageType" json:"stage,omitempty"`
	StageTimeLeftEntry   int32         `protobuf:"zigzag32,12,opt,name=stage_time_left_entry,json=stageTimeLeftEntry,proto3" json:"stage_time_left_entry,omitempty"`
	StageTimeLeftExit    int32         `protobuf:"zigzag32,13,opt,name=stage_time_left_exit,json=stageTimeLeftExit,proto3" json:"stage_time_left_exit,omitempty"`
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

func (m *GamePhase) GetStage() StageType {
	if m != nil {
		return m.Stage
	}
	return StageType_STAGE_UNKNOWN
}

func (m *GamePhase) GetStageTimeLeftEntry() int32 {
	if m != nil {
		return m.StageTimeLeftEntry
	}
	return 0
}

func (m *GamePhase) GetStageTimeLeftExit() int32 {
	if m != nil {
		return m.StageTimeLeftExit
	}
	return 0
}

type TeamStats struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Goals                uint32   `protobuf:"varint,2,opt,name=goals,proto3" json:"goals,omitempty"`
	Fouls                uint32   `protobuf:"varint,3,opt,name=fouls,proto3" json:"fouls,omitempty"`
	YellowCards          uint32   `protobuf:"varint,4,opt,name=yellow_cards,json=yellowCards,proto3" json:"yellow_cards,omitempty"`
	RedCards             uint32   `protobuf:"varint,5,opt,name=red_cards,json=redCards,proto3" json:"red_cards,omitempty"`
	TimeoutTime          uint32   `protobuf:"varint,6,opt,name=timeout_time,json=timeoutTime,proto3" json:"timeout_time,omitempty"`
	Timeouts             uint32   `protobuf:"varint,7,opt,name=timeouts,proto3" json:"timeouts,omitempty"`
	PenaltyShotsTotal    uint32   `protobuf:"varint,8,opt,name=penalty_shots_total,json=penaltyShotsTotal,proto3" json:"penalty_shots_total,omitempty"`
	PenaltyShotsSuccess  uint32   `protobuf:"varint,9,opt,name=penalty_shots_success,json=penaltyShotsSuccess,proto3" json:"penalty_shots_success,omitempty"`
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

func (m *TeamStats) GetTimeoutTime() uint32 {
	if m != nil {
		return m.TimeoutTime
	}
	return 0
}

func (m *TeamStats) GetTimeouts() uint32 {
	if m != nil {
		return m.Timeouts
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

type MatchStatsCollection struct {
	MatchStats           []*MatchStats `protobuf:"bytes,1,rep,name=match_stats,json=matchStats,proto3" json:"match_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MatchStatsCollection) Reset()         { *m = MatchStatsCollection{} }
func (m *MatchStatsCollection) String() string { return proto.CompactTextString(m) }
func (*MatchStatsCollection) ProtoMessage()    {}
func (*MatchStatsCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b017991eb90546, []int{4}
}

func (m *MatchStatsCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchStatsCollection.Unmarshal(m, b)
}
func (m *MatchStatsCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchStatsCollection.Marshal(b, m, deterministic)
}
func (m *MatchStatsCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchStatsCollection.Merge(m, src)
}
func (m *MatchStatsCollection) XXX_Size() int {
	return xxx_messageInfo_MatchStatsCollection.Size(m)
}
func (m *MatchStatsCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchStatsCollection.DiscardUnknown(m)
}

var xxx_messageInfo_MatchStatsCollection proto.InternalMessageInfo

func (m *MatchStatsCollection) GetMatchStats() []*MatchStats {
	if m != nil {
		return m.MatchStats
	}
	return nil
}

func init() {
	proto.RegisterEnum("TeamColor", TeamColor_name, TeamColor_value)
	proto.RegisterEnum("CommandType", CommandType_name, CommandType_value)
	proto.RegisterEnum("StageType", StageType_name, StageType_value)
	proto.RegisterEnum("GamePhaseType", GamePhaseType_name, GamePhaseType_value)
	proto.RegisterType((*Command)(nil), "Command")
	proto.RegisterType((*GamePhase)(nil), "GamePhase")
	proto.RegisterType((*TeamStats)(nil), "TeamStats")
	proto.RegisterType((*MatchStats)(nil), "MatchStats")
	proto.RegisterMapType((map[string]float32)(nil), "MatchStats.RelTimePerGamePhaseEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "MatchStats.TimePerGamePhaseEntry")
	proto.RegisterType((*MatchStatsCollection)(nil), "MatchStatsCollection")
}

func init() { proto.RegisterFile("ssl_match_stats.proto", fileDescriptor_21b017991eb90546) }

var fileDescriptor_21b017991eb90546 = []byte{
	// 1223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0xdd, 0x6e, 0xe3, 0x54,
	0x10, 0x26, 0xff, 0xf1, 0x38, 0x69, 0x4f, 0x4f, 0x1a, 0xd6, 0xdb, 0xdd, 0x85, 0x6c, 0xc5, 0x4a,
	0x55, 0x17, 0x02, 0x5b, 0x10, 0x02, 0xc4, 0x8d, 0x9b, 0x3a, 0xdd, 0xaa, 0x89, 0x1d, 0x1d, 0xbb,
	0x5a, 0x96, 0x1b, 0xcb, 0xdb, 0x9c, 0xfe, 0x08, 0x27, 0x8e, 0xec, 0x93, 0xa5, 0x79, 0x09, 0x1e,
	0x84, 0x0b, 0x1e, 0x00, 0xf1, 0x1c, 0x3c, 0x0f, 0x3a, 0x3f, 0xfe, 0x49, 0xda, 0x8a, 0xbb, 0x33,
	0xf3, 0xcd, 0x4c, 0xbe, 0xf9, 0x66, 0xc6, 0x2d, 0x74, 0x93, 0x24, 0xf4, 0x67, 0x01, 0xbb, 0xbc,
	0xf1, 0x13, 0x16, 0xb0, 0xa4, 0xbf, 0x88, 0x23, 0x16, 0xed, 0x3d, 0xe5, 0xee, 0xeb, 0x60, 0x46,
	0x7d, 0xfa, 0x91, 0xce, 0x99, 0x7f, 0xf4, 0xcd, 0x9b, 0x1f, 0x25, 0xb4, 0x4f, 0xa0, 0x31, 0x88,
	0x66, 0xb3, 0x60, 0x3e, 0xc5, 0x3d, 0xa8, 0xb2, 0xd5, 0x82, 0x1a, 0xa5, 0x5e, 0xe9, 0x60, 0xeb,
	0xa8, 0xd5, 0x57, 0x7e, 0x6f, 0xb5, 0xa0, 0x44, 0x20, 0xf8, 0x15, 0x34, 0xaf, 0xa2, 0xd8, 0x67,
	0x34, 0x98, 0x19, 0x65, 0x11, 0x05, 0x7d, 0x8f, 0x06, 0xb3, 0x41, 0x14, 0x46, 0x31, 0x69, 0x5c,
	0x45, 0x31, 0xb7, 0xf6, 0xff, 0xae, 0x82, 0x76, 0x1a, 0xcc, 0xe8, 0xe4, 0x26, 0x48, 0x28, 0x7e,
	0x01, 0x90, 0xb0, 0x20, 0x66, 0x3e, 0xbb, 0x9d, 0xc9, 0xe2, 0x55, 0xa2, 0x09, 0x8f, 0x77, 0x3b,
	0xa3, 0xf8, 0x29, 0x34, 0xe9, 0x7c, 0x2a, 0xc1, 0xb2, 0x00, 0x1b, 0x74, 0x3e, 0x15, 0xd0, 0x1e,
	0x34, 0xa7, 0xcb, 0x38, 0x60, 0xb7, 0xd1, 0xdc, 0xa8, 0xf4, 0x4a, 0x07, 0x6d, 0x92, 0xd9, 0x78,
	0x5f, 0x91, 0xad, 0x0a, 0x1a, 0x5b, 0xfd, 0xec, 0xf7, 0x0a, 0x74, 0xbf, 0x82, 0xf6, 0xa5, 0xec,
	0xc1, 0xa7, 0x73, 0x16, 0xaf, 0x8c, 0x5a, 0xaf, 0x74, 0xa0, 0x1f, 0x35, 0xd3, 0xce, 0x48, 0x4b,
	0xc1, 0x16, 0x47, 0xf1, 0x6b, 0x68, 0x65, 0xe1, 0x77, 0xb7, 0xcc, 0xa8, 0x6f, 0x44, 0xeb, 0x69,
	0xf4, 0xdd, 0x2d, 0x5b, 0x93, 0xa2, 0xf1, 0xa8, 0x14, 0xf8, 0x7b, 0xd8, 0xc9, 0x75, 0x4f, 0x14,
	0x8d, 0x66, 0xaf, 0x72, 0xa0, 0x1f, 0x81, 0xe0, 0x6c, 0x71, 0x80, 0x6c, 0x5f, 0xa7, 0xcf, 0x44,
	0x72, 0xf9, 0x0e, 0xd0, 0x5a, 0x1e, 0xe7, 0xa3, 0xdd, 0x4b, 0xdb, 0x2a, 0xa4, 0x71, 0x52, 0x3f,
	0x43, 0x77, 0x4e, 0xef, 0x98, 0x9f, 0xb6, 0xb1, 0x88, 0xa3, 0x45, 0x94, 0xd0, 0xa9, 0x01, 0x1b,
	0xad, 0x74, 0x78, 0x98, 0x32, 0x26, 0x2a, 0x08, 0xf7, 0xa0, 0x96, 0xb0, 0xe0, 0x9a, 0x1a, 0xba,
	0xea, 0xc7, 0xe5, 0x96, 0xd0, 0x53, 0x02, 0xf8, 0x0d, 0x74, 0xc5, 0x43, 0x4c, 0xcb, 0x0f, 0xe9,
	0x15, 0x53, 0x1d, 0xb5, 0x7a, 0xa5, 0x83, 0x1d, 0x82, 0x05, 0xc8, 0x47, 0x37, 0xa2, 0x57, 0x4c,
	0x36, 0xf2, 0x35, 0xec, 0xde, 0x4b, 0xe1, 0xcd, 0xb4, 0x45, 0xc6, 0xce, 0x7a, 0xc6, 0xdd, 0x2d,
	0xdb, 0xff, 0xab, 0x0c, 0x1a, 0x97, 0xce, 0xe5, 0xfb, 0x8b, 0x31, 0x54, 0xe7, 0x81, 0x5a, 0x1b,
	0x8d, 0x88, 0x37, 0xde, 0x85, 0xda, 0x75, 0x14, 0x84, 0x89, 0x58, 0x97, 0x36, 0x91, 0x06, 0xf7,
	0x5e, 0x45, 0xcb, 0x30, 0x51, 0x9b, 0x22, 0x0d, 0xfc, 0x12, 0x5a, 0x2b, 0x1a, 0x86, 0xd1, 0xef,
	0xfe, 0x65, 0x10, 0x4f, 0x13, 0xb1, 0x2e, 0x6d, 0xa2, 0x4b, 0xdf, 0x80, 0xbb, 0xf0, 0x33, 0xd0,
	0x62, 0x3a, 0x55, 0x78, 0x4d, 0xae, 0x59, 0x4c, 0xa7, 0x12, 0x7c, 0x09, 0x2d, 0x4e, 0x3c, 0x5a,
	0xaa, 0xf5, 0xad, 0xcb, 0x7c, 0xe5, 0x4b, 0xb7, 0x54, 0x99, 0x89, 0xd8, 0x84, 0x36, 0xc9, 0x6c,
	0xdc, 0x87, 0xce, 0x82, 0xce, 0x83, 0x90, 0xad, 0xfc, 0xe4, 0x26, 0x62, 0x89, 0xcf, 0x22, 0x16,
	0x84, 0x46, 0x53, 0x84, 0xed, 0x28, 0xc8, 0xe5, 0x88, 0xc7, 0x01, 0x7c, 0x04, 0xdd, 0xf5, 0xf8,
	0x64, 0x79, 0x79, 0x49, 0x93, 0xc4, 0xd0, 0x44, 0x46, 0xa7, 0x98, 0xe1, 0x4a, 0x68, 0xff, 0xdf,
	0x2a, 0xc0, 0x98, 0x9f, 0xfc, 0xe3, 0x8a, 0xbd, 0x06, 0x5d, 0x6c, 0xd3, 0x82, 0x1f, 0x08, 0xd7,
	0x2d, 0x5f, 0x24, 0x71, 0x33, 0x04, 0xae, 0xd3, 0x67, 0x82, 0x5f, 0xc1, 0x96, 0xfc, 0x82, 0x6c,
	0xdc, 0x5e, 0x5b, 0x78, 0x4f, 0xd2, 0x03, 0x7c, 0x01, 0x40, 0xef, 0x58, 0x1c, 0x48, 0x5d, 0xb8,
	0xae, 0x4d, 0xa2, 0x09, 0x4f, 0xaa, 0x4a, 0x72, 0x13, 0x45, 0x2c, 0x5a, 0x32, 0x21, 0x6a, 0x93,
	0x64, 0x36, 0x9e, 0x40, 0x47, 0x6c, 0xc3, 0x82, 0xc6, 0x7e, 0xce, 0xcb, 0xa8, 0x0b, 0x5a, 0x2f,
	0xfb, 0x79, 0x33, 0x7d, 0x5e, 0x6a, 0x42, 0xe3, 0x8c, 0xa8, 0xd8, 0x29, 0x82, 0xd8, 0x86, 0x1b,
	0xff, 0x0a, 0x4f, 0x62, 0x1a, 0xfa, 0x0f, 0x55, 0x6d, 0x88, 0xaa, 0x5f, 0x14, 0xab, 0x12, 0x1a,
	0x3e, 0x5c, 0xb8, 0x13, 0xdf, 0x47, 0xf8, 0x09, 0xf3, 0x2b, 0x97, 0x1f, 0x54, 0x5f, 0x6e, 0x8e,
	0x98, 0xa0, 0xae, 0x4e, 0x5e, 0x14, 0x25, 0xdb, 0x2c, 0x7d, 0xbe, 0x17, 0x21, 0xf8, 0x08, 0xb6,
	0x0b, 0x79, 0x1f, 0xc2, 0x25, 0x15, 0x53, 0x5c, 0xcf, 0x6a, 0x67, 0x59, 0xc7, 0xe1, 0x92, 0xee,
	0x0d, 0xa0, 0xfb, 0x20, 0x33, 0x8c, 0xa0, 0xf2, 0x1b, 0x5d, 0xa9, 0xa1, 0xf2, 0x27, 0xdf, 0xf7,
	0x8f, 0x01, 0x2f, 0xaa, 0xae, 0x40, 0x18, 0x3f, 0x95, 0x7f, 0x28, 0xed, 0x0d, 0xc1, 0x78, 0xac,
	0xc3, 0xff, 0xab, 0x53, 0x2e, 0xd4, 0xd9, 0x3f, 0x81, 0xdd, 0x5c, 0xb4, 0x41, 0x14, 0x86, 0xf4,
	0x52, 0x4c, 0xfe, 0x4b, 0xd0, 0x0b, 0x7f, 0x62, 0x8c, 0x92, 0x10, 0x58, 0x2f, 0x08, 0x4c, 0x60,
	0x96, 0xbd, 0x0f, 0x47, 0xf2, 0x9c, 0xc5, 0x77, 0x11, 0x23, 0x68, 0x79, 0x96, 0x39, 0xf6, 0x2f,
	0xec, 0x73, 0xdb, 0x79, 0x67, 0xa3, 0x4f, 0xf0, 0x36, 0xe8, 0xc2, 0xf3, 0xde, 0x1a, 0x8d, 0x9c,
	0x77, 0xa8, 0x84, 0xdb, 0xa0, 0x09, 0xc7, 0xf1, 0xe8, 0xc2, 0x42, 0xe5, 0xcc, 0xb4, 0x1d, 0xdb,
	0x42, 0x95, 0xc3, 0x3f, 0xcb, 0xa0, 0x17, 0xfe, 0x2e, 0xe1, 0x0e, 0x6c, 0x0f, 0x9c, 0xf1, 0xd8,
	0xb4, 0x4f, 0x0a, 0x35, 0x11, 0xb4, 0x52, 0xe7, 0x5b, 0x73, 0xe4, 0xa1, 0x52, 0xd1, 0xe3, 0x7a,
	0xce, 0x04, 0x95, 0xf1, 0x1e, 0x7c, 0x9a, 0x7a, 0x8e, 0xcd, 0xd1, 0xc8, 0x9f, 0x8c, 0xcc, 0x81,
	0x35, 0xb6, 0x6c, 0x0f, 0x55, 0xb0, 0x01, 0xbb, 0x29, 0x66, 0x3b, 0x64, 0x6c, 0x8e, 0x7c, 0xd7,
	0x33, 0x89, 0x87, 0xaa, 0xf8, 0x09, 0x74, 0x52, 0x64, 0xe8, 0x90, 0x81, 0xa5, 0x80, 0x5a, 0x11,
	0x38, 0x39, 0x23, 0xd6, 0xc0, 0xf3, 0x87, 0xc4, 0xb2, 0x50, 0x1d, 0x3f, 0x85, 0x6e, 0x0a, 0x9c,
	0xd9, 0x45, 0xa8, 0x81, 0x9f, 0xc1, 0x93, 0x14, 0x9a, 0x10, 0x6b, 0x62, 0x12, 0xcb, 0x3f, 0x3f,
	0x1b, 0x9c, 0x3b, 0xc3, 0x21, 0x6a, 0x3e, 0x04, 0x4e, 0x2c, 0xdb, 0x1c, 0x79, 0xef, 0x91, 0x56,
	0xec, 0xda, 0x3b, 0x1b, 0x5b, 0xce, 0x85, 0x87, 0xa0, 0xd8, 0xe3, 0xa9, 0x63, 0x8e, 0x90, 0x7e,
	0xf8, 0x4f, 0x05, 0xb4, 0xec, 0x1b, 0x8e, 0x77, 0xa0, 0xed, 0x7a, 0xe6, 0xa9, 0x55, 0x10, 0xea,
	0x73, 0x78, 0x26, 0x5d, 0xaa, 0xcd, 0xe1, 0x19, 0x71, 0x3d, 0xae, 0xd9, 0x90, 0xff, 0x28, 0x2a,
	0x71, 0x16, 0x8f, 0x04, 0x48, 0x09, 0xd7, 0x40, 0x91, 0xc7, 0xf9, 0xa0, 0x0a, 0xee, 0xc1, 0xf3,
	0x35, 0xcc, 0xb5, 0x06, 0x8e, 0x1c, 0x87, 0x2c, 0x5d, 0xc5, 0xcf, 0xc1, 0x78, 0x2c, 0x02, 0xd5,
	0xf2, 0xda, 0xd6, 0x2f, 0x1e, 0x31, 0x45, 0x55, 0xff, 0x98, 0x58, 0xe6, 0x39, 0xaa, 0xe3, 0xcf,
	0x60, 0xaf, 0x88, 0x6d, 0x90, 0x6e, 0x6c, 0xe6, 0x16, 0x38, 0x37, 0xf9, 0x38, 0x8a, 0x58, 0x4e,
	0x59, 0xcb, 0xc5, 0x90, 0xd0, 0x26, 0x63, 0xc8, 0xc5, 0xb8, 0x17, 0x80, 0xf4, 0xbc, 0x61, 0x35,
	0x25, 0xdf, 0x7d, 0xeb, 0x38, 0x9e, 0x73, 0xe1, 0x29, 0xda, 0xad, 0x9c, 0xd6, 0x66, 0x04, 0x6a,
	0xf3, 0x81, 0x2a, 0xcc, 0x71, 0x3d, 0xff, 0xd4, 0x1c, 0x5b, 0x68, 0xeb, 0xf0, 0x8f, 0x32, 0xb4,
	0xd7, 0xfe, 0xad, 0xe1, 0x23, 0x9c, 0xbc, 0x35, 0xdd, 0xe2, 0x08, 0x33, 0x17, 0xb9, 0xb0, 0xed,
	0x33, 0xfb, 0x14, 0x95, 0x70, 0x17, 0x76, 0xa4, 0xab, 0xb8, 0x6e, 0x65, 0xbe, 0xa2, 0xd2, 0xbd,
	0xbe, 0x87, 0x15, 0xae, 0x89, 0x04, 0x36, 0xb7, 0xb0, 0x7a, 0x1f, 0x4a, 0x77, 0xb0, 0x86, 0xb7,
	0x00, 0x24, 0x24, 0x0e, 0xaa, 0xce, 0x8f, 0x46, 0xda, 0x1b, 0xe7, 0xd4, 0xc8, 0x29, 0xa6, 0xbb,
	0xda, 0xe4, 0x57, 0xaf, 0x82, 0x85, 0x38, 0x5a, 0x5e, 0x4d, 0x1c, 0x2c, 0x70, 0x41, 0xd4, 0x0f,
	0x67, 0x82, 0xe8, 0x1f, 0xea, 0xe2, 0x5f, 0xd6, 0x6f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc3,
	0x28, 0x07, 0x84, 0xe6, 0x0a, 0x00, 0x00,
}
