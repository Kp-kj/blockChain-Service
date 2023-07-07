package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartActivityLogic {
	return &StartActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StartActivityLogic) StartActivity(in *block.StartActivityRequest) (*block.IsSuccessResponse, error) {

	newActivity, err := l.svcCtx.ManageActivityModel.FindOne(l.ctx, in.ActivityId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	oldActivity, err := l.svcCtx.ManageActivityModel.FindOneByCryptominerTypeid(l.ctx, newActivity.CryptominerTypeid)
	if err == nil {
		oldActivity.IsActivation = 0
		err = l.svcCtx.ManageActivityModel.Update(l.ctx, oldActivity)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
	}
	newActivity.IsActivation = 1
	err = l.svcCtx.ManageActivityModel.Update(l.ctx, newActivity)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return &block.IsSuccessResponse{
		IsSuccess: true,
	}, nil
}
