package logic

import (
	"context"

	"block/block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FullPurchaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFullPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FullPurchaseLogic {
	return &FullPurchaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FullPurchaseLogic) FullPurchase(in *block.PurchaseRequest) (*block.FullPurchaseResponse, error) {
	// todo: add your logic here and delete this line

	return &block.FullPurchaseResponse{}, nil
}
