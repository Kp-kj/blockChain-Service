package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainCryptominerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBargainCryptominerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainCryptominerLogic {
	return &GetBargainCryptominerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBargainCryptominerLogic) GetBargainCryptominer(in *block.GetBargainCryptominerRequest) (*block.GetBargainCryptominerResponse, error) {
	// todo: add your logic here and delete this line

	return &block.GetBargainCryptominerResponse{}, nil
}
