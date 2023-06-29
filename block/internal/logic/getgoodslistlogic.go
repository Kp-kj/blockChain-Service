package logic

import (
	"block/block"
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
			return nil, err
		}

		NewCryptominer := &block.Cryptominer{
			CryptominerId:     CryptominerId.Generate().Int64(),
			UserId:            in.UserId, // 用户ID,未购买时没有
			CryptominerTypeid: cryptominer.CryptominerTypeid,
			CryptominerName:   cryptominer.CryptominerName,

			CryptominerPicture:  cryptominer.CryptominerPicture.String,
			CryptominerDescribe: cryptominer.CryptominerDescribe.String,
			CryptominerPrice:    cryptominer.CryptominerPrice,
			CryptominerDuration: strconv.FormatInt(cryptominer.CryptominerDuration, 10),
			IsBargain:           true,
			OptionalStatus:      "0",
		}
		Cryptominers = append(Cryptominers, NewCryptominer)
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
			return nil, err
		}
		NewProp := &block.Prop{
			PropId:       PropId.Generate().Int64(),
			UserId:       in.UserId, // 用户ID,未购买时没有
			PropTypeid:   prop.PropTypeid,
			PropName:     prop.PropName,
			PropPicture:  prop.PropPicture.String,
			PropDescribe: prop.PropDescribe.String,
			PropPrice:    prop.PropPrice,
		}
		Props = append(Props, NewProp)
	}

	return &block.GetGoodsListResponse{Cryptominer: Cryptominers, Prop: Props}, nil
}
