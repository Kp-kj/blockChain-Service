package logic

import (
	"block/block"
	"block/internal/model"
	"block/internal/svc"
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
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
	ManageCryptominers, err := l.svcCtx.ManageCryptominerModel.SelectAllCryptominer(l.ctx)
	if err == nil {
		logx.Error(err)
		return nil, err
	}
	var Cryptominers []*block.Cryptominer

	for _, cryptominer := range ManageCryptominers {
		CryptominerId, err := snowflake.NewNode(1)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		// 提供给用户展示
		NewCryptominer := &block.Cryptominer{
			CryptominerId:       CryptominerId.Generate().Int64(),
			CryptominerTypeid:   cryptominer.CryptominerTypeid,
			CryptominerName:     cryptominer.CryptominerName,
			CryptominerPicture:  cryptominer.CryptominerPicture.String,
			CryptominerDescribe: cryptominer.CryptominerDescribe.String,
			CryptominerPrice:    cryptominer.CryptominerPrice,
			CryptominerDuration: strconv.FormatInt(cryptominer.CryptominerDuration, 10),
			OptionalStatus:      "0", // 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
			IsBargain:           true,
		}
		Cryptominers = append(Cryptominers, NewCryptominer)
		// 存数据库
		CryptominerModel := &model.Cryptominer{
			CryptominerId:       CryptominerId.Generate().Int64(),
			CryptominerTypeid:   cryptominer.CryptominerTypeid,
			CryptominerName:     cryptominer.CryptominerName,
			CryptominerPicture:  cryptominer.CryptominerPicture,
			CryptominerDescribe: cryptominer.CryptominerDescribe,
			CryptominerPrice:    cryptominer.CryptominerPrice,
			CryptominerDuration: cryptominer.CryptominerDuration,
			OptionalStatus:      "0", // 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
			IsBargain:           1,
		}
		_, err = l.svcCtx.CryptominerModel.Insert(l.ctx, CryptominerModel)
	}

	// 创建实体道具
	ManageProps, err := l.svcCtx.ManagePropModel.SelectAllProp(l.ctx)
	if err == nil {
		logx.Error(err)
		return nil, err
	}
	var Props []*block.Prop
	for _, prop := range ManageProps {
		PropId, err := snowflake.NewNode(1)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		NewProp := &block.Prop{
			PropId:       PropId.Generate().Int64(),
			PropTypeid:   prop.PropTypeid,
			PropName:     prop.PropName,
			PropPicture:  prop.PropPicture.String,
			PropDescribe: prop.PropDescribe.String,
			PropPrice:    prop.PropPrice,
		}
		Props = append(Props, NewProp)

		PropModel := &model.Prop{
			PropId:       PropId.Generate().Int64(),
			PropTypeid:   prop.PropTypeid,
			PropName:     prop.PropName,
			PropPicture:  prop.PropPicture,
			PropDescribe: prop.PropDescribe,
			PropPrice:    prop.PropPrice,
		}
		_, err = l.svcCtx.PropModel.Insert(l.ctx, PropModel)
	}

	return &block.GetGoodsListResponse{Cryptominer: Cryptominers, Prop: Props}, nil
}
