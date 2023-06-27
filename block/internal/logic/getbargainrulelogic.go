package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBargainRuleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBargainRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBargainRuleLogic {
	return &GetBargainRuleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBargainRuleLogic) GetBargainRule(in *block.GetBargainRuleRequest) (*block.GetBargainRuleResponse, error) {
	// todo: add your logic here and delete this line

	return &block.GetBargainRuleResponse{}, nil
}
