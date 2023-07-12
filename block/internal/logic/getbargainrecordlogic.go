package logic

import (
	"context"
	"time"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBargainRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainRecordLogic {
	return &GetBargainRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBargainRecordLogic) GetBargainRecord(in *block.GetBargainRecordRequest) (*block.GetBargainRecordResponse, error) {
	var loginStatus string = "0"
	var inActivity bool = true
	var bargainMax bool = false
	var isBargain bool = false

	_, err := l.svcCtx.BargainModel.FindBargainRecord(l.ctx, in.UserId, in.BargainId)
	if err != nil {
		loginStatus = "1"
	}
	record, err := l.svcCtx.BargainModel.FindOne(l.ctx, in.BargainId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	Cryptominer, _ := l.svcCtx.CryptominerModel.FindOne(l.ctx, record.CryptominerId)
	assistanceRecord, _ := l.svcCtx.FriendAssistanceModel.FindOneAssistanceRecord(l.ctx, in.BargainId)

	_, err = l.svcCtx.PurchaseRecordsModel.JudgeGoodsPurchaseForBargain(l.ctx, Cryptominer.UserId, Cryptominer.CryptominerName)
	if err == nil {
		//今日已砍价购买过本种类矿机，活动结束
		inActivity = false
	} else if time.Now().Unix() > record.ActivityStartTime.Add(time.Hour*6).Unix() {
		//活动已超时，活动结束
		inActivity = false
	}
	if record.RemainingPrice == 0 {
		bargainMax = true
	}

	_, err = l.svcCtx.FriendAssistanceModel.FindUserAssistanceRecord(l.ctx, in.BargainId, in.UserId)
	if err != nil {
		isBargain = true
	}

	var SupportUserList []*block.SupportUser
	for _, assistance := range assistanceRecord {
		supportUser := &block.SupportUser{
			AssistorId:   assistance.UserId,
			AssistorName: assistance.UserName,
			Avatar:       assistance.Avatar,
			TwitterName:  assistance.TwitterName,
			BargainPrice: float32(assistance.BargainPrice),
		}
		SupportUserList = append(SupportUserList, supportUser)
	}
	Record := &block.BargainRecord{
		BargainId:          record.BargainId,
		CryptominerName:    Cryptominer.CryptominerName,
		CryptominerPicture: Cryptominer.CryptominerPicture.String,
		CryptominerPrice:   Cryptominer.CryptominerPrice,
		ActivityStartTime:  record.ActivityStartTime.Unix(),
		RemainingPrice:     float32(record.RemainingPrice),
		SupportUser:        SupportUserList,
	}

	return &block.GetBargainRecordResponse{BargainRecord: Record, LoginStatus: loginStatus, InActivity: inActivity, BargainMax: bargainMax, IsBargain: isBargain}, nil
}
