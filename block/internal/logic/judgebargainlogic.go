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
	// todo: add your logic here and delete this line

	return &block.JudgeBargainResponse{}, nil
}
