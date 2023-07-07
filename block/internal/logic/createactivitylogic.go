package logic

import (
	"block/block"
	"block/internal/model"
	"block/internal/svc"
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateActivityLogic {
	return &CreateActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateActivityLogic) CreateActivity(in *block.CreateActivityRequest) (*block.IsSuccessResponse, error) {

	ActivityId, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	createActivity := &model.ManageActivity{
		ActivityId:        ActivityId.Generate().Int64(),
		CryptominerTypeid: in.CryptominerTypeid,
		AdminuserId:       in.AdminuserId,
		UserAmount:        in.UserAmount,
		MinPrice:          float64(in.MinPrice),
		FirstBargainPer:   float64(in.FirstBargainPer),
		FriendBargainPer:  float64(in.FriendBargainPer),
		IsActivation:      0, //默认不开启活动
	}

	_, err = l.svcCtx.ManageActivityModel.Insert(l.ctx, createActivity)
	if err != nil {
		return &block.IsSuccessResponse{
			IsSuccess: false,
		}, nil
	}

	return &block.IsSuccessResponse{
		IsSuccess: true,
	}, nil

}
