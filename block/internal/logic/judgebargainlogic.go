package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeBargainLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeBargainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeBargainLogic {
	return &JudgeBargainLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeBargainLogic) JudgeBargain(in *block.JudgeBargainRequest) (*block.JudgeBargainResponse, error) {

	// 查询购买记录，是否有今日砍价账单
	_, err := l.svcCtx.PurchaseRecordsModel.FindBargainByUserId(l.ctx, in.UserId, in.CryptominerTypeid)
	if err == nil {
		// 查询到了数据，此种矿机今日不可砍
		return &block.JudgeBargainResponse{
			IsBargain: false,
		}, nil
	}
	return &block.JudgeBargainResponse{
		IsBargain: true,
	}, nil

}
