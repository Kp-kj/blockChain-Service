package logic

import (
	"block/block"
	"block/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeGoodsPurchaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeGoodsPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeGoodsPurchaseLogic {
	return &JudgeGoodsPurchaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeGoodsPurchaseLogic) JudgeGoodsPurchase(in *block.JudgeGoodsPurchaseRequest) (*block.JudgeGoodsPurchaseResponse, error) {

	// 查询购买记录，是否有今日商品记录
	_, err := l.svcCtx.PurchaseRecordsModel.JudgeGoodsPurchase(l.ctx, in.UserId, in.GoodName)
	if err == nil {
		// 查询到了数据，有今日商品记录
		return &block.JudgeGoodsPurchaseResponse{
			IsPurchase: true,
		}, nil
	}
	return &block.JudgeGoodsPurchaseResponse{
		IsPurchase: false,
	}, nil

}
