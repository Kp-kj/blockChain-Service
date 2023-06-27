package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainGoodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBargainGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainGoodLogic {
	return &GetBargainGoodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBargainGoodLogic) GetBargainGood(in *block.GetBargainGoodRequest) (*block.GetBargainGoodResponse, error) {
	// todo: add your logic here and delete this line

	return &block.GetBargainGoodResponse{}, nil
}
