package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BargainPurchaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBargainPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BargainPurchaseLogic {
	return &BargainPurchaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BargainPurchaseLogic) BargainPurchase(in *block.PurchaseRequest) (*block.BargainPurchaseResponse, error) {
	// todo: add your logic here and delete this line

	return &block.BargainPurchaseResponse{}, nil
}
