package logic

import (
	"block/block"
	"block/internal/model"
	"block/internal/svc"
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type AssistorBargainLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAssistorBargainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AssistorBargainLogic {
	return &AssistorBargainLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AssistorBargainLogic) AssistorBargain(in *block.AssistorBargainRequest) (*block.AssistorBargainResponse, error) {

	if in.BargainPrice != 0 {
		// 如果可以拿到 金额，说明缓存未失效

		BargainRecord, err := l.svcCtx.BargainModel.FindOne(l.ctx, in.BargainId)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		//
		unusedsplit := strings.Split(BargainRecord.UnusedData, ",")
		var isExist bool = false
		var isChange int
		for i, str := range unusedsplit {
			if str == fmt.Sprintf("%.2f", in.BargainPrice) {
				isExist = true
				isChange = i
				goto labelOut
			}
		}
	labelOut:
		if isExist == true {
			var newSplit []string
			newSplit = append(unusedsplit[:isChange], unusedsplit[isChange+1:]...)
			UnusedData := strings.Join(newSplit, ",")
			BargainRecord.UnusedData = UnusedData

			// 增加使用过数据的信息
			usedSplit := strings.Split(BargainRecord.UsedData, ",")
			if len(usedSplit) != 0 {
				usedSplit = append(usedSplit, fmt.Sprintf("%.2f", in.BargainPrice))
			}
			UsedData := strings.Join(usedSplit, ",")
			BargainRecord.UsedData = UsedData
		} else {
			return &block.AssistorBargainResponse{UnuserArr: "", RedisExist: true}, nil
		}

		// FriendAssistance增加单条好友砍价记录
		_, err = l.svcCtx.FriendAssistanceModel.Insert(l.ctx, &model.FriendAssistance{
			BargainId:    in.BargainId,
			UserId:       in.UserId,
			UserName:     in.UserName,
			TwitterName:  in.TwitterName,
			Avatar:       in.Avatar,
			BargainPrice: float64(in.BargainPrice),
		})
		if err != nil {
			logx.Error(err)
			return nil, err
		}

		// 更新bargain 剩余额度，已使用的金额集
		BargainRecord.RemainingPrice = BargainRecord.RemainingPrice - float64(in.BargainPrice)
		err = l.svcCtx.BargainModel.Update(l.ctx, BargainRecord)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		return &block.AssistorBargainResponse{UnuserArr: "", RedisExist: true}, nil

	} else {
		//缓存不存在，需要库里取到返回给redis数据

		BargainRecord, err := l.svcCtx.BargainModel.FindOne(l.ctx, in.BargainId)
		if err != nil {
			logx.Error(err)
			return nil, err
		}

		// 取走剩余数据中的第一个值，返回剩余数据做缓存
		unusedsplit := strings.Split(BargainRecord.UnusedData, ",")
		if len(unusedsplit) == 0 {
			return &block.AssistorBargainResponse{UnuserArr: "nilArr", RedisExist: false}, nil
		}
		usedSplit := strings.Split(BargainRecord.UsedData, ",")
		var theBargainPrice float64
		theBargainPrice, err = strconv.ParseFloat(unusedsplit[0], 64)
		if err != nil {
			logx.Error(err)
			return nil, err
		}

		if len(usedSplit) != 0 {
			usedSplit = append(usedSplit, unusedsplit[0])
		}
		UsedData := strings.Join(usedSplit, ",")
		BargainRecord.UsedData = UsedData
		unusedsplit = unusedsplit[1:]
		UnusedData := strings.Join(unusedsplit, ",")
		BargainRecord.UnusedData = UnusedData

		// FriendAssistance增加单条好友砍价记录
		_, err = l.svcCtx.FriendAssistanceModel.Insert(l.ctx, &model.FriendAssistance{
			BargainId:    in.BargainId,
			UserId:       in.UserId,
			UserName:     in.UserName,
			TwitterName:  in.TwitterName,
			Avatar:       in.Avatar,
			BargainPrice: theBargainPrice,
		})
		if err != nil {
			logx.Error(err)
			return nil, err
		}

		// 更新bargain 剩余额度，已使用的金额集
		BargainRecord.RemainingPrice = BargainRecord.RemainingPrice - theBargainPrice
		err = l.svcCtx.BargainModel.Update(l.ctx, BargainRecord)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		return &block.AssistorBargainResponse{UnuserArr: UnusedData, RedisExist: false}, nil
	}
}
