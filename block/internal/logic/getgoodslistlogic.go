package logic

import (
	"block/block"
	"block/internal/model"
	"block/internal/svc"
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListLogic {
	return &GetGoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsListLogic) GetGoodsList(in *block.GetGoodsListRequest) (*block.GetGoodsListResponse, error) {

	// 创建实体矿机
	ManageCryptominers, _ := l.svcCtx.ManageCryptominerModel.SelectAllCryptominer(l.ctx)

	var Cryptominers []*block.Cryptominer
	for _, cryptominer := range ManageCryptominers {

		// 查询用户矿机表中是否有未购买的矿机存在
		userCryptominers, err := l.svcCtx.CryptominerModel.FindCryptominerInUser(l.ctx, in.UserId, cryptominer.CryptominerTypeid)
		if err != nil {
			// 未查到，创建实体矿机
			CryptominerId, err := snowflake.NewNode(1)
			if err != nil {
				logx.Error(err)
				return nil, err
			}
			// 存数据库
			CryptominerModel := &model.Cryptominer{
				UserId:              in.UserId,
				CryptominerId:       CryptominerId.Generate().Int64(),
				CryptominerTypeid:   cryptominer.CryptominerTypeid,
				CryptominerName:     cryptominer.CryptominerName,
				CryptominerPicture:  cryptominer.CryptominerPicture,
				CryptominerDescribe: cryptominer.CryptominerDescribe,
				CryptominerPrice:    cryptominer.CryptominerPrice,
				CryptominerDuration: cryptominer.CryptominerDuration,
				CryptominerPower:    cryptominer.CryptominerPower,
				PaymentWay:          cryptominer.PaymentWay,
				OptionalStatus:      "0", // 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
				IsBargain:           1,
			}
			_, err = l.svcCtx.CryptominerModel.Insert(l.ctx, CryptominerModel)
			if err != nil {
				logx.Error(err)
				return nil, err
			}
			// 提供给用户展示
			NewCryptominer := &block.Cryptominer{
				UserId:              in.UserId,
				CryptominerId:       CryptominerId.Generate().Int64(),
				CryptominerTypeid:   cryptominer.CryptominerTypeid,
				CryptominerName:     cryptominer.CryptominerName,
				CryptominerPicture:  cryptominer.CryptominerPicture.String,
				CryptominerDescribe: cryptominer.CryptominerDescribe.String,
				CryptominerPrice:    cryptominer.CryptominerPrice,
				CryptominerDuration: cryptominer.CryptominerDuration,
				CryptominerPower:    cryptominer.CryptominerPower,
				PaymentWay:          cryptominer.PaymentWay,
				OptionalStatus:      "0", // 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
				IsBargain:           true,
			}
			Cryptominers = append(Cryptominers, NewCryptominer)
		} else {
			var judgeTheBargain bool = false
			if userCryptominers.IsBargain == 0 {
				judgeTheBargain = false
			} else {
				judgeTheBargain = true
			}
			// 提供给用户展示
			NewCryptominer := &block.Cryptominer{
				UserId:              userCryptominers.UserId,
				CryptominerId:       userCryptominers.CryptominerId,
				CryptominerTypeid:   userCryptominers.CryptominerTypeid,
				CryptominerName:     userCryptominers.CryptominerName,
				CryptominerPicture:  userCryptominers.CryptominerPicture.String,
				CryptominerDescribe: userCryptominers.CryptominerDescribe.String,
				CryptominerPrice:    userCryptominers.CryptominerPrice,
				CryptominerDuration: cryptominer.CryptominerDuration,
				CryptominerPower:    cryptominer.CryptominerPower,
				PaymentWay:          userCryptominers.PaymentWay,
				OptionalStatus:      userCryptominers.OptionalStatus, // 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
				IsBargain:           judgeTheBargain,                 //  是否可砍，根据当天是否砍价决定 0：不可砍 1：可砍
			}
			Cryptominers = append(Cryptominers, NewCryptominer)
		}
	}

	// 搜索用户道具库中是否存在
	ManageProps, _ := l.svcCtx.ManagePropModel.SelectAllProp(l.ctx)
	var Props []*block.Prop

	for _, prop := range ManageProps {
		userPorps, err := l.svcCtx.PropModel.FindPropInUser(l.ctx, in.UserId, prop.PropTypeid)
		if err != nil {
			// 未查到，创建实体道具
			PropId, err := snowflake.NewNode(1)
			if err != nil {
				logx.Error(err)
				return nil, err
			}
			// 存数据库
			PropModel := &model.Prop{
				UserId:         in.UserId,
				PropId:         PropId.Generate().Int64(),
				PropTypeid:     prop.PropTypeid,
				PropName:       prop.PropName,
				PropPicture:    prop.PropPicture,
				PropPrice:      prop.PropPrice,
				PropDescribe:   prop.PropDescribe,
				PaymentWay:     prop.PaymentWay,
				OptionalStatus: "0", // 道具状态 0：未购买 1：已购买
			}
			_, err = l.svcCtx.PropModel.Insert(l.ctx, PropModel)
			if err != nil {
				logx.Error(err)
				return nil, err
			}
			// 提供给用户展示
			NewProp := &block.Prop{
				UserId:       in.UserId,
				PropId:       PropId.Generate().Int64(),
				PropTypeid:   prop.PropTypeid,
				PropName:     prop.PropName,
				PropPicture:  prop.PropPicture.String,
				PropDescribe: prop.PropDescribe.String,
				PropPrice:    prop.PropPrice,
				PaymentWay:   prop.PaymentWay,
			}
			Props = append(Props, NewProp)
		} else {
			// 提供给用户展示
			NewProp := &block.Prop{
				UserId:       userPorps.UserId,
				PropId:       userPorps.PropId,
				PropTypeid:   userPorps.PropTypeid,
				PropName:     userPorps.PropName,
				PropPicture:  userPorps.PropPicture.String,
				PropDescribe: userPorps.PropDescribe.String,
				PropPrice:    userPorps.PropPrice,
				PaymentWay:   userPorps.PaymentWay,
			}
			Props = append(Props, NewProp)
		}

	}
	return &block.GetGoodsListResponse{Cryptominer: Cryptominers, Prop: Props}, nil
}
