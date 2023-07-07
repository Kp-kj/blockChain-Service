package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartGoodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartGoodLogic {
	return &StartGoodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StartGoodLogic) StartGood(in *block.StartGoodRequest) (*block.IsSuccessResponse, error) {

	// 先找矿机（一共种类就几个），找到了就改为上架状态
	cryptominer, err := l.svcCtx.ManageCryptominerModel.FindOne(l.ctx, in.GoodTypeid)
	if err != nil {
		prop, err := l.svcCtx.ManagePropModel.FindOne(l.ctx, in.GoodTypeid)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		prop.GoodStatus = "1"
		err = l.svcCtx.ManagePropModel.Update(l.ctx, prop)
		if err != nil {
			return &block.IsSuccessResponse{
				IsSuccess: false,
			}, nil
		}
	} else {
		cryptominer.GoodStatus = "1"
		err := l.svcCtx.ManageCryptominerModel.Update(l.ctx, cryptominer)
		if err != nil {
			return &block.IsSuccessResponse{
				IsSuccess: false,
			}, nil
		}
	}

	return &block.IsSuccessResponse{
		IsSuccess: true,
	}, nil
}
