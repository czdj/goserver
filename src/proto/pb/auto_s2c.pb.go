// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auto_s2c.proto

package pb

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

type S2C struct {
	Error                             int32                              `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	Key                               string                             `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	GamerLoginS2C                     *GamerLoginS2C                     `protobuf:"bytes,3,opt,name=gamerLoginS2C,proto3" json:"gamerLoginS2C,omitempty"`
	GamerLoginGetDataS2C              *GamerLoginGetDataS2C              `protobuf:"bytes,4,opt,name=gamerLoginGetDataS2C,proto3" json:"gamerLoginGetDataS2C,omitempty"`
	ServerTimeS2C                     *ServerTimeS2C                     `protobuf:"bytes,5,opt,name=serverTimeS2C,proto3" json:"serverTimeS2C,omitempty"`
	GamerNotifyLoginOtherS2C          *GamerNotifyLoginOtherS2C          `protobuf:"bytes,6,opt,name=gamerNotifyLoginOtherS2C,proto3" json:"gamerNotifyLoginOtherS2C,omitempty"`
	GamerSubChatChannelS2C            *GamerSubChatChannelS2C            `protobuf:"bytes,7,opt,name=gamerSubChatChannelS2C,proto3" json:"gamerSubChatChannelS2C,omitempty"`
	GamerFriendChatS2C                *GamerFriendChatS2C                `protobuf:"bytes,8,opt,name=gamerFriendChatS2C,proto3" json:"gamerFriendChatS2C,omitempty"`
	GamerWorldChatS2C                 *GamerWorldChatS2C                 `protobuf:"bytes,9,opt,name=gamerWorldChatS2C,proto3" json:"gamerWorldChatS2C,omitempty"`
	GamerTestChatS2C                  *GamerTestChatS2C                  `protobuf:"bytes,10,opt,name=gamerTestChatS2C,proto3" json:"gamerTestChatS2C,omitempty"`
	GamerClubRequestS2C               *GamerClubRequestS2C               `protobuf:"bytes,11,opt,name=gamerClubRequestS2C,proto3" json:"gamerClubRequestS2C,omitempty"`
	GamerNotifyNewChatS2C             *GamerNotifyNewChatS2C             `protobuf:"bytes,12,opt,name=gamerNotifyNewChatS2C,proto3" json:"gamerNotifyNewChatS2C,omitempty"`
	GamerNewFriendReqS2C              *GamerNewFriendReqS2C              `protobuf:"bytes,13,opt,name=gamerNewFriendReqS2C,proto3" json:"gamerNewFriendReqS2C,omitempty"`
	GamerNotifyNewFriendReqS2C        *GamerNotifyNewFriendReqS2C        `protobuf:"bytes,14,opt,name=gamerNotifyNewFriendReqS2C,proto3" json:"gamerNotifyNewFriendReqS2C,omitempty"`
	GamerProcessFriendReqS2C          *GamerProcessFriendReqS2C          `protobuf:"bytes,15,opt,name=gamerProcessFriendReqS2C,proto3" json:"gamerProcessFriendReqS2C,omitempty"`
	GamerNotifyNewFriendS2C           *GamerNotifyNewFriendS2C           `protobuf:"bytes,16,opt,name=gamerNotifyNewFriendS2C,proto3" json:"gamerNotifyNewFriendS2C,omitempty"`
	GamerNotifyMailS2C                *GamerNotifyMailS2C                `protobuf:"bytes,17,opt,name=gamerNotifyMailS2C,proto3" json:"gamerNotifyMailS2C,omitempty"`
	GamerNotifyNewMailS2C             *GamerNotifyNewMailS2C             `protobuf:"bytes,18,opt,name=gamerNotifyNewMailS2C,proto3" json:"gamerNotifyNewMailS2C,omitempty"`
	GamerGetMailS2C                   *GamerGetMailS2C                   `protobuf:"bytes,19,opt,name=gamerGetMailS2C,proto3" json:"gamerGetMailS2C,omitempty"`
	GamerDelMailS2C                   *GamerDelMailS2C                   `protobuf:"bytes,20,opt,name=gamerDelMailS2C,proto3" json:"gamerDelMailS2C,omitempty"`
	GamerDelHaveReadMailS2C           *GamerDelHaveReadMailS2C           `protobuf:"bytes,21,opt,name=gamerDelHaveReadMailS2C,proto3" json:"gamerDelHaveReadMailS2C,omitempty"`
	GamerOneKeyRcvMailRewardS2C       *GamerOneKeyRcvMailRewardS2C       `protobuf:"bytes,22,opt,name=gamerOneKeyRcvMailRewardS2C,proto3" json:"gamerOneKeyRcvMailRewardS2C,omitempty"`
	GamerChangeMailStateS2C           *GamerChangeMailStateS2C           `protobuf:"bytes,23,opt,name=gamerChangeMailStateS2C,proto3" json:"gamerChangeMailStateS2C,omitempty"`
	GamerMatchS2C                     *GamerMatchS2C                     `protobuf:"bytes,24,opt,name=gamerMatchS2C,proto3" json:"gamerMatchS2C,omitempty"`
	GamerNotifyMatchInfoS2C           *GamerNotifyMatchInfoS2C           `protobuf:"bytes,25,opt,name=gamerNotifyMatchInfoS2C,proto3" json:"gamerNotifyMatchInfoS2C,omitempty"`
	GamerPVPSyncS2C                   *GamerPVPSyncS2C                   `protobuf:"bytes,26,opt,name=gamerPVPSyncS2C,proto3" json:"gamerPVPSyncS2C,omitempty"`
	GamerNotifyPVPSyncS2C             *GamerNotifyPVPSyncS2C             `protobuf:"bytes,27,opt,name=gamerNotifyPVPSyncS2C,proto3" json:"gamerNotifyPVPSyncS2C,omitempty"`
	GamerNotifyNewPVPResultS2C        *GamerNotifyNewPVPResultS2C        `protobuf:"bytes,28,opt,name=gamerNotifyNewPVPResultS2C,proto3" json:"gamerNotifyNewPVPResultS2C,omitempty"`
	GamerNotifyIconChangeS2C          *GamerNotifyIconChangeS2C          `protobuf:"bytes,29,opt,name=gamerNotifyIconChangeS2C,proto3" json:"gamerNotifyIconChangeS2C,omitempty"`
	GamerGetRealTimeRankS2C           *GamerGetRealTimeRankS2C           `protobuf:"bytes,30,opt,name=gamerGetRealTimeRankS2C,proto3" json:"gamerGetRealTimeRankS2C,omitempty"`
	GamerCheckPVPBattleS2C            *GamerCheckPVPBattleS2C            `protobuf:"bytes,31,opt,name=gamerCheckPVPBattleS2C,proto3" json:"gamerCheckPVPBattleS2C,omitempty"`
	GamerUploadWXInfoS2C              *GamerUploadWXInfoS2C              `protobuf:"bytes,32,opt,name=gamerUploadWXInfoS2C,proto3" json:"gamerUploadWXInfoS2C,omitempty"`
	CreateLeagueS2C                   *CreateLeagueS2C                   `protobuf:"bytes,33,opt,name=createLeagueS2C,proto3" json:"createLeagueS2C,omitempty"`
	GamerNotifyLeagueDataS2C          *GamerNotifyLeagueDataS2C          `protobuf:"bytes,34,opt,name=gamerNotifyLeagueDataS2C,proto3" json:"gamerNotifyLeagueDataS2C,omitempty"`
	GamerNotifyLeagueGamerOnlineS2C   *GamerNotifyLeagueGamerOnlineS2C   `protobuf:"bytes,35,opt,name=gamerNotifyLeagueGamerOnlineS2C,proto3" json:"gamerNotifyLeagueGamerOnlineS2C,omitempty"`
	GamerGetBackpackS2C               *GamerGetBackpackS2C               `protobuf:"bytes,36,opt,name=gamerGetBackpackS2C,proto3" json:"gamerGetBackpackS2C,omitempty"`
	GamerNotifyItemS2C                *GamerNotifyItemS2C                `protobuf:"bytes,37,opt,name=gamerNotifyItemS2C,proto3" json:"gamerNotifyItemS2C,omitempty"`
	GamerSellItemS2C                  *GamerSellItemS2C                  `protobuf:"bytes,38,opt,name=gamerSellItemS2C,proto3" json:"gamerSellItemS2C,omitempty"`
	GamerUseItemS2C                   *GamerUseItemS2C                   `protobuf:"bytes,39,opt,name=gamerUseItemS2C,proto3" json:"gamerUseItemS2C,omitempty"`
	GamerCompoundItemS2C              *GamerCompoundItemS2C              `protobuf:"bytes,40,opt,name=gamerCompoundItemS2C,proto3" json:"gamerCompoundItemS2C,omitempty"`
	GamerSplitItemS2C                 *GamerSplitItemS2C                 `protobuf:"bytes,41,opt,name=gamerSplitItemS2C,proto3" json:"gamerSplitItemS2C,omitempty"`
	GamerNotifyItemChangeS2C          *GamerNotifyItemChangeS2C          `protobuf:"bytes,42,opt,name=gamerNotifyItemChangeS2C,proto3" json:"gamerNotifyItemChangeS2C,omitempty"`
	GamerGetActorListS2C              *GamerGetActorListS2C              `protobuf:"bytes,43,opt,name=gamerGetActorListS2C,proto3" json:"gamerGetActorListS2C,omitempty"`
	GamerChangeActorNameS2C           *GamerChangeActorNameS2C           `protobuf:"bytes,44,opt,name=gamerChangeActorNameS2C,proto3" json:"gamerChangeActorNameS2C,omitempty"`
	GamerActorUpLevelS2C              *GamerActorUpLevelS2C              `protobuf:"bytes,45,opt,name=gamerActorUpLevelS2C,proto3" json:"gamerActorUpLevelS2C,omitempty"`
	GamerNotifyActorChangeS2C         *GamerNotifyActorChangeS2C         `protobuf:"bytes,46,opt,name=gamerNotifyActorChangeS2C,proto3" json:"gamerNotifyActorChangeS2C,omitempty"`
	GMS2C                             *GMS2C                             `protobuf:"bytes,47,opt,name=gMS2C,proto3" json:"gMS2C,omitempty"`
	GamerNotifyStoryListS2C           *GamerNotifyStoryListS2C           `protobuf:"bytes,48,opt,name=gamerNotifyStoryListS2C,proto3" json:"gamerNotifyStoryListS2C,omitempty"`
	GamerNotifyBuildInfoS2C           *GamerNotifyBuildInfoS2C           `protobuf:"bytes,49,opt,name=gamerNotifyBuildInfoS2C,proto3" json:"gamerNotifyBuildInfoS2C,omitempty"`
	BuildLevelUpS2C                   *BuildLevelUpS2C                   `protobuf:"bytes,50,opt,name=buildLevelUpS2C,proto3" json:"buildLevelUpS2C,omitempty"`
	BuildEndOfLevelUpS2C              *BuildEndOfLevelUpS2C              `protobuf:"bytes,51,opt,name=buildEndOfLevelUpS2C,proto3" json:"buildEndOfLevelUpS2C,omitempty"`
	BuildBreakUpS2C                   *BuildBreakUpS2C                   `protobuf:"bytes,52,opt,name=buildBreakUpS2C,proto3" json:"buildBreakUpS2C,omitempty"`
	BuildSpeedUpS2C                   *BuildSpeedUpS2C                   `protobuf:"bytes,53,opt,name=buildSpeedUpS2C,proto3" json:"buildSpeedUpS2C,omitempty"`
	BuildGetProductS2C                *BuildGetProductS2C                `protobuf:"bytes,54,opt,name=buildGetProductS2C,proto3" json:"buildGetProductS2C,omitempty"`
	ActorEnterBuildS2C                *ActorEnterBuildS2C                `protobuf:"bytes,55,opt,name=actorEnterBuildS2C,proto3" json:"actorEnterBuildS2C,omitempty"`
	ActorLeaveBuildS2C                *ActorLeaveBuildS2C                `protobuf:"bytes,56,opt,name=actorLeaveBuildS2C,proto3" json:"actorLeaveBuildS2C,omitempty"`
	GamerNotifyIntelligenceS2C        *GamerNotifyIntelligenceS2C        `protobuf:"bytes,57,opt,name=gamerNotifyIntelligenceS2C,proto3" json:"gamerNotifyIntelligenceS2C,omitempty"`
	GamerGetIntelligenceS2C           *GamerGetIntelligenceS2C           `protobuf:"bytes,58,opt,name=gamerGetIntelligenceS2C,proto3" json:"gamerGetIntelligenceS2C,omitempty"`
	GamerExtractIntelligenceS2C       *GamerExtractIntelligenceS2C       `protobuf:"bytes,59,opt,name=gamerExtractIntelligenceS2C,proto3" json:"gamerExtractIntelligenceS2C,omitempty"`
	GamerIntelligenceTreasureChestS2C *GamerIntelligenceTreasureChestS2C `protobuf:"bytes,60,opt,name=gamerIntelligenceTreasureChestS2C,proto3" json:"gamerIntelligenceTreasureChestS2C,omitempty"`
	GamerIntelligenceParseS2C         *GamerIntelligenceParseS2C         `protobuf:"bytes,61,opt,name=gamerIntelligenceParseS2C,proto3" json:"gamerIntelligenceParseS2C,omitempty"`
	StudioS2C                         *StudioS2C                         `protobuf:"bytes,62,opt,name=studioS2C,proto3" json:"studioS2C,omitempty"`
	StudioStoryListS2C                *StudioStoryListS2C                `protobuf:"bytes,63,opt,name=studioStoryListS2C,proto3" json:"studioStoryListS2C,omitempty"`
	StudioStorySelectS2C              *StudioStorySelectS2C              `protobuf:"bytes,64,opt,name=studioStorySelectS2C,proto3" json:"studioStorySelectS2C,omitempty"`
	StudioActorSelectS2C              *StudioActorSelectS2C              `protobuf:"bytes,65,opt,name=studioActorSelectS2C,proto3" json:"studioActorSelectS2C,omitempty"`
	StudioEventS2C                    *StudioEventS2C                    `protobuf:"bytes,66,opt,name=studioEventS2C,proto3" json:"studioEventS2C,omitempty"`
	StudioEventOptionSelectS2C        *StudioEventOptionSelectS2C        `protobuf:"bytes,67,opt,name=studioEventOptionSelectS2C,proto3" json:"studioEventOptionSelectS2C,omitempty"`
	StudioStarS2C                     *StudioStarS2C                     `protobuf:"bytes,68,opt,name=studioStarS2C,proto3" json:"studioStarS2C,omitempty"`
	StudioNextStepS2C                 *StudioNextStepS2C                 `protobuf:"bytes,69,opt,name=studioNextStepS2C,proto3" json:"studioNextStepS2C,omitempty"`
	StudioContinuedListS2C            *StudioContinuedListS2C            `protobuf:"bytes,70,opt,name=studioContinuedListS2C,proto3" json:"studioContinuedListS2C,omitempty"`
	StudioContinuedReceiveS2C         *StudioContinuedReceiveS2C         `protobuf:"bytes,71,opt,name=studioContinuedReceiveS2C,proto3" json:"studioContinuedReceiveS2C,omitempty"`
	StudioContinuedReceiveAllS2C      *StudioContinuedReceiveAllS2C      `protobuf:"bytes,72,opt,name=studioContinuedReceiveAllS2C,proto3" json:"studioContinuedReceiveAllS2C,omitempty"`
	XXX_NoUnkeyedLiteral              struct{}                           `json:"-"`
	XXX_unrecognized                  []byte                             `json:"-"`
	XXX_sizecache                     int32                              `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c717f7f57a1f7b1, []int{0}
}

func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (m *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(m, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *S2C) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *S2C) GetGamerLoginS2C() *GamerLoginS2C {
	if m != nil {
		return m.GamerLoginS2C
	}
	return nil
}

func (m *S2C) GetGamerLoginGetDataS2C() *GamerLoginGetDataS2C {
	if m != nil {
		return m.GamerLoginGetDataS2C
	}
	return nil
}

func (m *S2C) GetServerTimeS2C() *ServerTimeS2C {
	if m != nil {
		return m.ServerTimeS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyLoginOtherS2C() *GamerNotifyLoginOtherS2C {
	if m != nil {
		return m.GamerNotifyLoginOtherS2C
	}
	return nil
}

func (m *S2C) GetGamerSubChatChannelS2C() *GamerSubChatChannelS2C {
	if m != nil {
		return m.GamerSubChatChannelS2C
	}
	return nil
}

func (m *S2C) GetGamerFriendChatS2C() *GamerFriendChatS2C {
	if m != nil {
		return m.GamerFriendChatS2C
	}
	return nil
}

func (m *S2C) GetGamerWorldChatS2C() *GamerWorldChatS2C {
	if m != nil {
		return m.GamerWorldChatS2C
	}
	return nil
}

func (m *S2C) GetGamerTestChatS2C() *GamerTestChatS2C {
	if m != nil {
		return m.GamerTestChatS2C
	}
	return nil
}

func (m *S2C) GetGamerClubRequestS2C() *GamerClubRequestS2C {
	if m != nil {
		return m.GamerClubRequestS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyNewChatS2C() *GamerNotifyNewChatS2C {
	if m != nil {
		return m.GamerNotifyNewChatS2C
	}
	return nil
}

func (m *S2C) GetGamerNewFriendReqS2C() *GamerNewFriendReqS2C {
	if m != nil {
		return m.GamerNewFriendReqS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyNewFriendReqS2C() *GamerNotifyNewFriendReqS2C {
	if m != nil {
		return m.GamerNotifyNewFriendReqS2C
	}
	return nil
}

func (m *S2C) GetGamerProcessFriendReqS2C() *GamerProcessFriendReqS2C {
	if m != nil {
		return m.GamerProcessFriendReqS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyNewFriendS2C() *GamerNotifyNewFriendS2C {
	if m != nil {
		return m.GamerNotifyNewFriendS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyMailS2C() *GamerNotifyMailS2C {
	if m != nil {
		return m.GamerNotifyMailS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyNewMailS2C() *GamerNotifyNewMailS2C {
	if m != nil {
		return m.GamerNotifyNewMailS2C
	}
	return nil
}

func (m *S2C) GetGamerGetMailS2C() *GamerGetMailS2C {
	if m != nil {
		return m.GamerGetMailS2C
	}
	return nil
}

func (m *S2C) GetGamerDelMailS2C() *GamerDelMailS2C {
	if m != nil {
		return m.GamerDelMailS2C
	}
	return nil
}

func (m *S2C) GetGamerDelHaveReadMailS2C() *GamerDelHaveReadMailS2C {
	if m != nil {
		return m.GamerDelHaveReadMailS2C
	}
	return nil
}

func (m *S2C) GetGamerOneKeyRcvMailRewardS2C() *GamerOneKeyRcvMailRewardS2C {
	if m != nil {
		return m.GamerOneKeyRcvMailRewardS2C
	}
	return nil
}

func (m *S2C) GetGamerChangeMailStateS2C() *GamerChangeMailStateS2C {
	if m != nil {
		return m.GamerChangeMailStateS2C
	}
	return nil
}

func (m *S2C) GetGamerMatchS2C() *GamerMatchS2C {
	if m != nil {
		return m.GamerMatchS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyMatchInfoS2C() *GamerNotifyMatchInfoS2C {
	if m != nil {
		return m.GamerNotifyMatchInfoS2C
	}
	return nil
}

func (m *S2C) GetGamerPVPSyncS2C() *GamerPVPSyncS2C {
	if m != nil {
		return m.GamerPVPSyncS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyPVPSyncS2C() *GamerNotifyPVPSyncS2C {
	if m != nil {
		return m.GamerNotifyPVPSyncS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyNewPVPResultS2C() *GamerNotifyNewPVPResultS2C {
	if m != nil {
		return m.GamerNotifyNewPVPResultS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyIconChangeS2C() *GamerNotifyIconChangeS2C {
	if m != nil {
		return m.GamerNotifyIconChangeS2C
	}
	return nil
}

func (m *S2C) GetGamerGetRealTimeRankS2C() *GamerGetRealTimeRankS2C {
	if m != nil {
		return m.GamerGetRealTimeRankS2C
	}
	return nil
}

func (m *S2C) GetGamerCheckPVPBattleS2C() *GamerCheckPVPBattleS2C {
	if m != nil {
		return m.GamerCheckPVPBattleS2C
	}
	return nil
}

func (m *S2C) GetGamerUploadWXInfoS2C() *GamerUploadWXInfoS2C {
	if m != nil {
		return m.GamerUploadWXInfoS2C
	}
	return nil
}

func (m *S2C) GetCreateLeagueS2C() *CreateLeagueS2C {
	if m != nil {
		return m.CreateLeagueS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyLeagueDataS2C() *GamerNotifyLeagueDataS2C {
	if m != nil {
		return m.GamerNotifyLeagueDataS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyLeagueGamerOnlineS2C() *GamerNotifyLeagueGamerOnlineS2C {
	if m != nil {
		return m.GamerNotifyLeagueGamerOnlineS2C
	}
	return nil
}

func (m *S2C) GetGamerGetBackpackS2C() *GamerGetBackpackS2C {
	if m != nil {
		return m.GamerGetBackpackS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyItemS2C() *GamerNotifyItemS2C {
	if m != nil {
		return m.GamerNotifyItemS2C
	}
	return nil
}

func (m *S2C) GetGamerSellItemS2C() *GamerSellItemS2C {
	if m != nil {
		return m.GamerSellItemS2C
	}
	return nil
}

func (m *S2C) GetGamerUseItemS2C() *GamerUseItemS2C {
	if m != nil {
		return m.GamerUseItemS2C
	}
	return nil
}

func (m *S2C) GetGamerCompoundItemS2C() *GamerCompoundItemS2C {
	if m != nil {
		return m.GamerCompoundItemS2C
	}
	return nil
}

func (m *S2C) GetGamerSplitItemS2C() *GamerSplitItemS2C {
	if m != nil {
		return m.GamerSplitItemS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyItemChangeS2C() *GamerNotifyItemChangeS2C {
	if m != nil {
		return m.GamerNotifyItemChangeS2C
	}
	return nil
}

func (m *S2C) GetGamerGetActorListS2C() *GamerGetActorListS2C {
	if m != nil {
		return m.GamerGetActorListS2C
	}
	return nil
}

func (m *S2C) GetGamerChangeActorNameS2C() *GamerChangeActorNameS2C {
	if m != nil {
		return m.GamerChangeActorNameS2C
	}
	return nil
}

func (m *S2C) GetGamerActorUpLevelS2C() *GamerActorUpLevelS2C {
	if m != nil {
		return m.GamerActorUpLevelS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyActorChangeS2C() *GamerNotifyActorChangeS2C {
	if m != nil {
		return m.GamerNotifyActorChangeS2C
	}
	return nil
}

func (m *S2C) GetGMS2C() *GMS2C {
	if m != nil {
		return m.GMS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyStoryListS2C() *GamerNotifyStoryListS2C {
	if m != nil {
		return m.GamerNotifyStoryListS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyBuildInfoS2C() *GamerNotifyBuildInfoS2C {
	if m != nil {
		return m.GamerNotifyBuildInfoS2C
	}
	return nil
}

func (m *S2C) GetBuildLevelUpS2C() *BuildLevelUpS2C {
	if m != nil {
		return m.BuildLevelUpS2C
	}
	return nil
}

func (m *S2C) GetBuildEndOfLevelUpS2C() *BuildEndOfLevelUpS2C {
	if m != nil {
		return m.BuildEndOfLevelUpS2C
	}
	return nil
}

func (m *S2C) GetBuildBreakUpS2C() *BuildBreakUpS2C {
	if m != nil {
		return m.BuildBreakUpS2C
	}
	return nil
}

func (m *S2C) GetBuildSpeedUpS2C() *BuildSpeedUpS2C {
	if m != nil {
		return m.BuildSpeedUpS2C
	}
	return nil
}

func (m *S2C) GetBuildGetProductS2C() *BuildGetProductS2C {
	if m != nil {
		return m.BuildGetProductS2C
	}
	return nil
}

func (m *S2C) GetActorEnterBuildS2C() *ActorEnterBuildS2C {
	if m != nil {
		return m.ActorEnterBuildS2C
	}
	return nil
}

func (m *S2C) GetActorLeaveBuildS2C() *ActorLeaveBuildS2C {
	if m != nil {
		return m.ActorLeaveBuildS2C
	}
	return nil
}

func (m *S2C) GetGamerNotifyIntelligenceS2C() *GamerNotifyIntelligenceS2C {
	if m != nil {
		return m.GamerNotifyIntelligenceS2C
	}
	return nil
}

func (m *S2C) GetGamerGetIntelligenceS2C() *GamerGetIntelligenceS2C {
	if m != nil {
		return m.GamerGetIntelligenceS2C
	}
	return nil
}

func (m *S2C) GetGamerExtractIntelligenceS2C() *GamerExtractIntelligenceS2C {
	if m != nil {
		return m.GamerExtractIntelligenceS2C
	}
	return nil
}

func (m *S2C) GetGamerIntelligenceTreasureChestS2C() *GamerIntelligenceTreasureChestS2C {
	if m != nil {
		return m.GamerIntelligenceTreasureChestS2C
	}
	return nil
}

func (m *S2C) GetGamerIntelligenceParseS2C() *GamerIntelligenceParseS2C {
	if m != nil {
		return m.GamerIntelligenceParseS2C
	}
	return nil
}

func (m *S2C) GetStudioS2C() *StudioS2C {
	if m != nil {
		return m.StudioS2C
	}
	return nil
}

func (m *S2C) GetStudioStoryListS2C() *StudioStoryListS2C {
	if m != nil {
		return m.StudioStoryListS2C
	}
	return nil
}

func (m *S2C) GetStudioStorySelectS2C() *StudioStorySelectS2C {
	if m != nil {
		return m.StudioStorySelectS2C
	}
	return nil
}

func (m *S2C) GetStudioActorSelectS2C() *StudioActorSelectS2C {
	if m != nil {
		return m.StudioActorSelectS2C
	}
	return nil
}

func (m *S2C) GetStudioEventS2C() *StudioEventS2C {
	if m != nil {
		return m.StudioEventS2C
	}
	return nil
}

func (m *S2C) GetStudioEventOptionSelectS2C() *StudioEventOptionSelectS2C {
	if m != nil {
		return m.StudioEventOptionSelectS2C
	}
	return nil
}

func (m *S2C) GetStudioStarS2C() *StudioStarS2C {
	if m != nil {
		return m.StudioStarS2C
	}
	return nil
}

func (m *S2C) GetStudioNextStepS2C() *StudioNextStepS2C {
	if m != nil {
		return m.StudioNextStepS2C
	}
	return nil
}

func (m *S2C) GetStudioContinuedListS2C() *StudioContinuedListS2C {
	if m != nil {
		return m.StudioContinuedListS2C
	}
	return nil
}

func (m *S2C) GetStudioContinuedReceiveS2C() *StudioContinuedReceiveS2C {
	if m != nil {
		return m.StudioContinuedReceiveS2C
	}
	return nil
}

func (m *S2C) GetStudioContinuedReceiveAllS2C() *StudioContinuedReceiveAllS2C {
	if m != nil {
		return m.StudioContinuedReceiveAllS2C
	}
	return nil
}

func init() {
	proto.RegisterType((*S2C)(nil), "S2C")
}

func init() { proto.RegisterFile("auto_s2c.proto", fileDescriptor_8c717f7f57a1f7b1) }

var fileDescriptor_8c717f7f57a1f7b1 = []byte{
	// 1413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x98, 0x7b, 0x53, 0x14, 0x47,
	0x14, 0xc5, 0x8b, 0x18, 0x4c, 0x1c, 0x15, 0x71, 0x44, 0x5d, 0x11, 0x23, 0x9a, 0x17, 0xe6, 0x41,
	0x12, 0x34, 0x31, 0x31, 0x31, 0x91, 0x1d, 0x70, 0xa5, 0x82, 0xb0, 0x75, 0x47, 0xd4, 0xaa, 0x54,
	0x25, 0xd5, 0xcc, 0x5e, 0x61, 0xc2, 0x30, 0xb3, 0xce, 0xf6, 0x2e, 0xf2, 0x09, 0xf3, 0xb5, 0x52,
	0x7d, 0x7b, 0x1e, 0x77, 0xba, 0x7b, 0xd6, 0xff, 0xd8, 0x7b, 0xce, 0xf9, 0x4d, 0xef, 0xf4, 0xed,
	0x07, 0xeb, 0xcd, 0x89, 0xb1, 0xcc, 0xfe, 0x19, 0xad, 0x45, 0xab, 0xc3, 0x3c, 0x93, 0xd9, 0xe2,
	0x85, 0x28, 0x89, 0x31, 0x95, 0xfa, 0xd3, 0xdd, 0xff, 0xee, 0x79, 0x67, 0xc2, 0xb5, 0xc0, 0x5f,
	0xf0, 0x66, 0x31, 0xcf, 0xb3, 0xbc, 0x33, 0xb3, 0x3c, 0xb3, 0x32, 0x0b, 0xfa, 0x83, 0x3f, 0xef,
	0x9d, 0x39, 0xc2, 0xd3, 0xce, 0x07, 0xcb, 0x33, 0x2b, 0xe7, 0x40, 0xfd, 0xe9, 0x3f, 0xf0, 0x2e,
	0x1e, 0x88, 0x63, 0xcc, 0xb7, 0xb3, 0x83, 0x38, 0x0d, 0xd7, 0x82, 0xce, 0x99, 0xe5, 0x99, 0x95,
	0xf3, 0x6b, 0x73, 0xab, 0x3d, 0x5e, 0x85, 0xa6, 0xc9, 0xdf, 0xf2, 0x16, 0xea, 0x42, 0x0f, 0xe5,
	0x86, 0x90, 0x42, 0x85, 0x3f, 0xa4, 0xf0, 0x55, 0x16, 0xae, 0x45, 0x70, 0x46, 0xd4, 0x00, 0x46,
	0x98, 0x4f, 0x30, 0x7f, 0x11, 0x1f, 0xa3, 0x62, 0xcc, 0x16, 0x03, 0x08, 0x79, 0x15, 0x9a, 0x26,
	0x7f, 0xcf, 0xeb, 0x10, 0x6d, 0x27, 0x93, 0xf1, 0x9b, 0x53, 0x62, 0xee, 0xca, 0x43, 0xcc, 0x15,
	0xe0, 0x2c, 0x01, 0x6e, 0xe8, 0x41, 0x38, 0x0c, 0xd0, 0x1a, 0xf5, 0x77, 0xbd, 0x6b, 0xa4, 0x85,
	0xe3, 0xfd, 0xe0, 0x50, 0xc8, 0xe0, 0x50, 0xa4, 0x29, 0x26, 0x0a, 0xfa, 0x11, 0x41, 0xaf, 0x6b,
	0xa8, 0x25, 0x43, 0x4b, 0xcc, 0x0f, 0x3c, 0x9f, 0x94, 0xa7, 0x79, 0x8c, 0xe9, 0x40, 0x89, 0x0a,
	0xf6, 0x31, 0xc1, 0xae, 0x68, 0x58, 0x43, 0x02, 0x87, 0xdd, 0x7f, 0xe2, 0x5d, 0xa6, 0xea, 0xab,
	0x2c, 0x4f, 0x2a, 0xc6, 0x39, 0x62, 0xf8, 0x9a, 0xc1, 0x15, 0xb0, 0xcd, 0xfe, 0x63, 0x6f, 0x9e,
	0x8a, 0x2f, 0x70, 0x24, 0x4b, 0x80, 0x47, 0x80, 0xcb, 0x1a, 0xc0, 0x04, 0xb0, 0xac, 0xfe, 0x53,
	0xef, 0x0a, 0xd5, 0x82, 0x64, 0xbc, 0x0f, 0xf8, 0x76, 0x8c, 0x23, 0x22, 0x9c, 0x27, 0xc2, 0x82,
	0x26, 0x34, 0x35, 0x70, 0x05, 0xfc, 0x6d, 0xef, 0x2a, 0x7b, 0xf5, 0x3b, 0x78, 0x52, 0x8e, 0xe5,
	0x02, 0x91, 0xae, 0xf1, 0x29, 0xab, 0x55, 0x70, 0x87, 0xaa, 0x26, 0xdc, 0xc1, 0x13, 0xfd, 0xbe,
	0x00, 0xdf, 0x2a, 0xd8, 0x45, 0xde, 0x84, 0x86, 0x08, 0xce, 0x88, 0xff, 0x97, 0xb7, 0xd8, 0x7c,
	0x46, 0x03, 0x38, 0x47, 0xc0, 0x9b, 0xc6, 0xe8, 0x1a, 0xd8, 0x29, 0xf1, 0xaa, 0x57, 0xfb, 0x79,
	0x16, 0xe1, 0x68, 0xd4, 0x40, 0x5f, 0xe2, 0xbd, 0xea, 0x30, 0x40, 0x6b, 0xd4, 0x07, 0xef, 0xba,
	0xeb, 0xa1, 0x8a, 0x3a, 0x4f, 0xd4, 0x8e, 0x73, 0xc0, 0x0a, 0xda, 0x16, 0xac, 0xda, 0x55, 0x4b,
	0xcf, 0x45, 0x4c, 0xbd, 0x7f, 0x99, 0xb7, 0x6b, 0x43, 0x02, 0x87, 0xdd, 0x9e, 0xe5, 0x92, 0xe3,
	0x3b, 0x67, 0xb9, 0x44, 0xb9, 0x43, 0xfe, 0x23, 0xef, 0x12, 0x09, 0x3d, 0x94, 0x25, 0xe7, 0x0a,
	0x71, 0xe6, 0x35, 0xa7, 0xae, 0x83, 0x69, 0xac, 0xb2, 0x1b, 0x98, 0x94, 0xd9, 0x05, 0x9e, 0xad,
	0xeb, 0x60, 0x1a, 0xab, 0xd7, 0xbb, 0x81, 0xc9, 0x33, 0x31, 0x41, 0x40, 0x31, 0x28, 0x19, 0x57,
	0xf9, 0xeb, 0xb5, 0x75, 0x68, 0x0b, 0xfa, 0x7f, 0x7b, 0x37, 0x49, 0xda, 0x4d, 0xf1, 0x4f, 0x3c,
	0x85, 0x68, 0xa2, 0x04, 0xc0, 0x13, 0x91, 0xd3, 0xb4, 0x5d, 0x23, 0xee, 0x92, 0xe6, 0xba, 0x3d,
	0x30, 0x0d, 0x50, 0x8d, 0x59, 0x6d, 0x40, 0x07, 0x48, 0x4f, 0x95, 0x42, 0xd2, 0xae, 0x7a, 0x9d,
	0x8f, 0xd9, 0xd6, 0xa1, 0x2d, 0x58, 0x1d, 0x10, 0xcf, 0x85, 0x8c, 0x0e, 0x15, 0xa9, 0xc3, 0x0f,
	0x88, 0xb2, 0x0a, 0x4d, 0x93, 0xd1, 0x9c, 0x54, 0xde, 0x4a, 0xdf, 0x64, 0x2a, 0x7f, 0xc3, 0x6e,
	0x4e, 0xae, 0x43, 0x5b, 0xb0, 0x9a, 0xcd, 0xfe, 0xcb, 0x7e, 0x78, 0x9a, 0x46, 0x8a, 0xb5, 0xc8,
	0x67, 0xb3, 0xae, 0x83, 0x69, 0x34, 0x7a, 0x92, 0x11, 0x6e, 0xda, 0x3d, 0xc9, 0x38, 0xee, 0x90,
	0xbd, 0x5d, 0xf4, 0x5f, 0xf6, 0x01, 0x47, 0xe3, 0x84, 0x36, 0xb3, 0x25, 0xe7, 0x76, 0xc1, 0x2d,
	0x30, 0x25, 0x6e, 0x1c, 0x6d, 0x5b, 0x51, 0x96, 0xea, 0x59, 0x51, 0xe8, 0x5b, 0xf6, 0xd1, 0xd6,
	0x30, 0x40, 0x6b, 0xb4, 0x9a, 0x91, 0x1e, 0x4a, 0x40, 0x91, 0xa8, 0x83, 0x14, 0x44, 0x7a, 0xa4,
	0xa8, 0x9f, 0xf0, 0x19, 0xb1, 0x75, 0x68, 0x0b, 0x56, 0xc7, 0x65, 0x70, 0x88, 0xd1, 0x51, 0xff,
	0x65, 0xbf, 0x2b, 0xa4, 0x4c, 0x68, 0xa0, 0xb7, 0xf9, 0x71, 0x69, 0xc9, 0xd0, 0x12, 0xab, 0xb6,
	0xf4, 0xbd, 0x61, 0x92, 0x89, 0xc1, 0xab, 0xd7, 0x65, 0xcf, 0x2c, 0xf3, 0x2d, 0xdd, 0x10, 0xc1,
	0x19, 0x51, 0xdd, 0x12, 0xe5, 0x28, 0x24, 0x6e, 0xa3, 0x38, 0x18, 0xd3, 0xa0, 0xee, 0x14, 0xdd,
	0x12, 0x34, 0xeb, 0x60, 0x1a, 0xcd, 0xdb, 0x05, 0xd5, 0xcb, 0x2b, 0xce, 0x5d, 0xc7, 0xed, 0x82,
	0x1b, 0xa0, 0x35, 0xea, 0xff, 0xeb, 0xdd, 0xb6, 0xb4, 0x62, 0xad, 0x27, 0x71, 0x4a, 0x43, 0xfc,
	0x94, 0xe8, 0xcb, 0x36, 0xbd, 0xe9, 0x83, 0xf7, 0x81, 0xaa, 0x23, 0xbb, 0x87, 0xb2, 0x2b, 0xa2,
	0xa3, 0xa1, 0x88, 0x68, 0xaa, 0x3f, 0xe3, 0x47, 0x76, 0x53, 0x03, 0x57, 0xc0, 0x38, 0x11, 0xb6,
	0x24, 0x1e, 0x2b, 0xcc, 0xe7, 0xf6, 0x89, 0x50, 0x48, 0xe0, 0xb0, 0x57, 0xd7, 0x8f, 0x10, 0x93,
	0xa4, 0x44, 0x7c, 0xc1, 0xaf, 0x1f, 0x4c, 0x00, 0xcb, 0x5a, 0x2d, 0xfc, 0xbd, 0x11, 0x96, 0xe9,
	0x2f, 0xf9, 0xc2, 0xaf, 0xeb, 0x60, 0x1a, 0xab, 0x8e, 0x0a, 0xb2, 0xe3, 0x61, 0x36, 0x4e, 0x07,
	0x25, 0x60, 0x85, 0x77, 0x94, 0x21, 0x82, 0x33, 0x52, 0x5d, 0xc3, 0xc2, 0x61, 0x12, 0xcb, 0x92,
	0x73, 0x8f, 0x5f, 0xc3, 0xb8, 0x02, 0xb6, 0xd9, 0x5c, 0xda, 0x12, 0x8f, 0xeb, 0xa5, 0xfd, 0x95,
	0x63, 0x69, 0x73, 0x03, 0xb4, 0x46, 0xab, 0xef, 0xd8, 0x43, 0xb9, 0x1e, 0xc9, 0x2c, 0xdf, 0x8e,
	0xf5, 0xfd, 0xec, 0x6b, 0xfe, 0x1d, 0x0d, 0x11, 0x9c, 0x11, 0xe3, 0x04, 0x21, 0x69, 0x47, 0xe8,
	0x7b, 0xf9, 0x37, 0xf6, 0x09, 0xc2, 0x75, 0x68, 0x0b, 0x56, 0xc3, 0xa3, 0xe2, 0xde, 0x70, 0x1b,
	0x27, 0xfa, 0x4a, 0xfd, 0x2d, 0x1f, 0x9e, 0x21, 0x82, 0x33, 0xe2, 0xbf, 0xf6, 0x6e, 0xb0, 0xb7,
	0x40, 0x6a, 0xfd, 0x06, 0x57, 0x89, 0xb7, 0xc8, 0xdf, 0x60, 0xd3, 0x01, 0xed, 0x61, 0x7f, 0xc9,
	0x9b, 0x3d, 0x78, 0xae, 0x28, 0xdf, 0x11, 0xe5, 0xec, 0x6a, 0x4f, 0x7d, 0x02, 0x5d, 0x34, 0x8e,
	0xb3, 0x50, 0x66, 0xf9, 0x69, 0xf9, 0x92, 0xbf, 0xb7, 0x8f, 0x33, 0xae, 0x43, 0x5b, 0xd0, 0x60,
	0x76, 0xc7, 0x71, 0x32, 0x28, 0xb7, 0xbb, 0x1f, 0x6c, 0x26, 0xd7, 0xa1, 0x2d, 0xa8, 0x56, 0xca,
	0xbe, 0xfa, 0x4c, 0x2f, 0x6c, 0x6f, 0xa8, 0x58, 0x6b, 0xc5, 0x4a, 0xe9, 0x36, 0xeb, 0x60, 0x1a,
	0xd5, 0x34, 0x51, 0x69, 0x33, 0x1d, 0xec, 0xbe, 0x61, 0x80, 0xfb, 0xc5, 0x34, 0x75, 0x1d, 0x22,
	0x38, 0x23, 0xd5, 0x30, 0xba, 0x39, 0x8a, 0x23, 0x4d, 0x79, 0xc0, 0x87, 0x51, 0xd7, 0xc1, 0x34,
	0x56, 0xd9, 0x70, 0x88, 0x38, 0xd0, 0xd9, 0x1f, 0x79, 0xb6, 0xae, 0x83, 0x69, 0x54, 0x9b, 0x15,
	0x95, 0x7a, 0x28, 0xfb, 0x79, 0x36, 0x18, 0x47, 0x34, 0x43, 0x3f, 0x15, 0x9b, 0x55, 0xd7, 0x92,
	0xc0, 0x61, 0x57, 0x10, 0xa1, 0x7a, 0x63, 0x33, 0x95, 0x98, 0xeb, 0x47, 0xae, 0x05, 0x9d, 0x87,
	0x05, 0x64, 0xdd, 0x92, 0xc0, 0x61, 0xaf, 0x20, 0xdb, 0x28, 0x26, 0x58, 0x41, 0x7e, 0xe6, 0x90,
	0x86, 0x04, 0x0e, 0xbb, 0x71, 0xcd, 0xd8, 0x4a, 0x25, 0x26, 0x49, 0x7c, 0x80, 0x69, 0x44, 0xed,
	0xfe, 0x8b, 0x7d, 0xcd, 0x30, 0x2c, 0x30, 0x25, 0xce, 0xef, 0x03, 0x26, 0xf9, 0x91, 0x71, 0x1f,
	0x30, 0xb1, 0x6d, 0xc1, 0xea, 0x7e, 0xbb, 0xf9, 0x4e, 0xe6, 0x22, 0xb2, 0xb8, 0xbf, 0xf2, 0xfb,
	0xad, 0xdb, 0x03, 0xd3, 0x00, 0xfe, 0xd0, 0xbb, 0x43, 0x32, 0xaf, 0xbf, 0xc8, 0x51, 0x8c, 0xc6,
	0x39, 0x06, 0x87, 0xc5, 0x7f, 0xa5, 0xbf, 0xd1, 0x53, 0xee, 0xea, 0xa7, 0x4c, 0x73, 0xc2, 0xfb,
	0x61, 0xd5, 0x86, 0xc3, 0x4d, 0x7d, 0x91, 0x8f, 0xe8, 0xfb, 0x3c, 0xe6, 0x1b, 0x8e, 0xcb, 0x01,
	0xed, 0x61, 0x7f, 0xc5, 0x3b, 0x37, 0x92, 0xe3, 0x41, 0x4c, 0x0b, 0xfe, 0x77, 0x22, 0x79, 0xab,
	0x61, 0x59, 0x81, 0x5a, 0x54, 0xbd, 0x54, 0x7c, 0xe0, 0xfb, 0xce, 0x1f, 0x45, 0x2f, 0x85, 0x96,
	0x04, 0x0e, 0xbb, 0x5a, 0xdd, 0xac, 0x1a, 0x62, 0x82, 0x7a, 0x71, 0x3c, 0x29, 0x56, 0x77, 0xe8,
	0x10, 0xc1, 0x19, 0xa9, 0x51, 0xd4, 0xc6, 0x35, 0x6a, 0xbd, 0x81, 0x6a, 0x8a, 0xe0, 0x8c, 0xf8,
	0x0f, 0xbd, 0x39, 0x5d, 0xdf, 0x9c, 0x60, 0x4a, 0x90, 0x2e, 0x41, 0x2e, 0x15, 0x90, 0xb2, 0x0c,
	0x86, 0x4d, 0x2d, 0x0d, 0x56, 0xd9, 0x1d, 0xca, 0x38, 0x4b, 0xeb, 0x91, 0x04, 0xc5, 0xd2, 0x08,
	0x5b, 0x2d, 0x30, 0x25, 0x4e, 0x3f, 0x49, 0x15, 0x5f, 0x5c, 0xd0, 0x2f, 0x4a, 0x1b, 0xe5, 0x4f,
	0x52, 0xbc, 0x0a, 0x4d, 0x93, 0xba, 0x1e, 0xe8, 0xc2, 0x0e, 0xbe, 0x93, 0xa1, 0x44, 0xda, 0xba,
	0x36, 0x8b, 0xeb, 0x41, 0x68, 0x2a, 0x60, 0x9b, 0xd5, 0x75, 0x5a, 0x17, 0x83, 0x2c, 0x95, 0x71,
	0x3a, 0xc6, 0x41, 0x39, 0xd9, 0x4f, 0x8b, 0xeb, 0x74, 0xe8, 0x94, 0xa1, 0x25, 0xa6, 0xba, 0xd7,
	0x50, 0x00, 0x23, 0x8c, 0x27, 0xd4, 0xbd, 0xbd, 0xa2, 0x7b, 0xc3, 0x36, 0x07, 0xb4, 0x87, 0x7d,
	0xe1, 0x2d, 0xb9, 0xc5, 0xf5, 0x84, 0xce, 0xf6, 0x67, 0x04, 0xbf, 0xd5, 0x02, 0xd7, 0x26, 0x98,
	0x8a, 0xd8, 0x3f, 0x4b, 0x3f, 0x68, 0xde, 0xff, 0x3f, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x04, 0x27,
	0xba, 0xf0, 0x14, 0x00, 0x00,
}