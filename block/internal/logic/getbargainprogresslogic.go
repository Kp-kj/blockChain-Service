package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainProgressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBargainProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainProgressLogic {
	return &GetBargainProgressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBargainProgressLogic) GetBargainProgress(in *block.GetBargainProgressRequest) (*block.GetBargainProgressResponse, error) {
	// todo: add your logic here and delete this line

	return &block.GetBargainProgressResponse{}, nil
}
