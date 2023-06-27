package logic

import (
	"context"

	"block/block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBargainRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainRecordLogic {
	return &GetBargainRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBargainRecordLogic) GetBargainRecord(in *block.GetBargainRecordRequest) (*block.GetBargainRecordResponse, error) {
	// todo: add your logic here and delete this line

	return &block.GetBargainRecordResponse{}, nil
}
