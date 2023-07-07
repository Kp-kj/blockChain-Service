package logic

import (
	"block/block"
	"block/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminActivityListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminActivityListLogic {
	return &AdminActivityListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminActivityListLogic) AdminActivityList(in *block.AdminActivityListRequest) (*block.AdminActivityListResponse, error) {

	list, err := l.svcCtx.ManageActivityModel.SelectActivityList(l.ctx)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var Activitylist []*block.Activity
	for _, activity := range list {
		Cryptominer, err := l.svcCtx.ManageCryptominerModel.FindOne(l.ctx, activity.CryptominerTypeid)
		if err != nil {
			logx.Error(err)
			return nil, err
		}
		NewActivity := &block.Activity{
			ActivityId:       activity.ActivityId,
			CryptominerName:  Cryptominer.CryptominerName,
			UserAmount:       activity.UserAmount,
			MinPrice:         float32(activity.MinPrice),
			FirstBargainPer:  float32(activity.FirstBargainPer),
			FriendBargainPer: float32(activity.FriendBargainPer),
			IsActivation:     activity.IsActivation, // 是否开启活动 0：未开启  1：开启
		}
		Activitylist = append(Activitylist, NewActivity)
	}

	return &block.AdminActivityListResponse{
		Activity: Activitylist,
	}, nil
}
