package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CryptominerBargainPurchaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCryptominerBargainPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CryptominerBargainPurchaseLogic {
	return &CryptominerBargainPurchaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CryptominerBargainPurchaseLogic) CryptominerBargainPurchase(in *block.CryptominerBargainRequest) (*block.CryptominerBargainResponse, error) {
	// todo: add your logic here and delete this line

	return &block.CryptominerBargainResponse{}, nil
}
