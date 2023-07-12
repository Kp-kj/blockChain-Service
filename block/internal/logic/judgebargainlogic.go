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

	// 在点击限时砍价时触发,查询购买记录里有无砍价记录
	theRule, err := l.svcCtx.BargainRuleModel.FindTheRule(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var isfirst bool = false
	_, err = l.svcCtx.BargainModel.FindBargainForRule(l.ctx, in.UserId)
	if err != nil {
		logx.Error(err)
		isfirst = true
	}

	// 查询购买记录，是否有今日砍价账单
	_, err = l.svcCtx.BargainModel.FindBargainByUserId(l.ctx, in.UserId, in.CryptominerTypeid)
	if err == nil {
		// 查询到了数据，此种矿机今日不可砍
		return &block.JudgeBargainResponse{
			IsBargain:           false,
			BargainRuleDescribe: "",
			IsFirst:             false,
		}, nil
	}
	return &block.JudgeBargainResponse{
		IsBargain:           true,
		BargainRuleDescribe: theRule.BargainRule,
		IsFirst:             isfirst,
	}, nil

}
