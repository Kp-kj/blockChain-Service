package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGoodsListLogic {
	return &CreateGoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGoodsListLogic) CreateGoodsList(in *block.CreateGoodsListRequest) (*block.CreateGoodsListResponse, error) {
	// todo: add your logic here and delete this line

	return &block.CreateGoodsListResponse{}, nil
}
