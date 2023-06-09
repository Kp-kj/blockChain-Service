package logic

import (
	"block/block"
	"block/internal/model"
	"block/internal/svc"
	"context"
	"database/sql"
	"github.com/bwmarrin/snowflake"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCryptominerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCryptominerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCryptominerLogic {
	return &CreateCryptominerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCryptominerLogic) CreateCryptominer(in *block.CreateCryptominerRequest) (*block.IsSuccessResponse, error) {

	//查询是否有矿机
	_, err := l.svcCtx.ManageCryptominerModel.FindOneByCryptominerName(l.ctx, in.CryptominerName)
	if err == nil {
		logx.Error(err)
		return nil, err
	}
	typeId, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	createCryptominer := &model.ManageCryptominer{
		CryptominerTypeid:   typeId.Generate().Int64(),
		AdminuserId:         in.AdminuserId,
		CryptominerName:     in.CryptominerName,
		CryptominerPicture:  sql.NullString{String: in.CryptominerPicture, Valid: true},
		CryptominerDescribe: sql.NullString{String: in.CryptominerDescribe, Valid: true},
		CryptominerPrice:    in.CryptominerPrice,
		CryptominerDuration: in.CryptominerDuration,
		CryptominerPower:    in.CryptominerPower,
		PaymentWay:          in.PaymentWay,
		GoodStatus:          "0", // 商品状态  0：待上架  1：已上架 2：未上架
		GoodType:            "0", // 类型 0：矿机 1：能量水
	}

	_, err = l.svcCtx.ManageCryptominerModel.Insert(l.ctx, createCryptominer)
	if err != nil {
		return &block.IsSuccessResponse{
			IsSuccess: false,
		}, nil
	}

	return &block.IsSuccessResponse{
		IsSuccess: true,
	}, nil
}
