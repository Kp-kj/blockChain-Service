package logic

import (
	"block/block"
	"block/internal/svc"
	"context"
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

	//createActivity := &model.ManageActivity{
	//	CryptominerTypeid: in.CryptominerTypeid,
	//	AdminuserId:       in.AdminuserId,
	//	UserAmount:        in.UserAmount,
	//	MinPrice:          float64(in.MinPrice),
	//	FirstBargainPer:   float64(in.FirstBargainPer),
	//	FriendBargainPer:  float64(in.FriendBargainPer),
	//	IsActivation:      in.IsActivation,
	//}
	//
	//_, err = l.svcCtx.ManageCryptominerModel.Insert(l.ctx, createCryptominer)
	//if err != nil {
	//	return &block.IsSuccessResponse{
	//		IsSuccess: false,
	//	}, nil
	//}
	//
	//return &block.IsSuccessResponse{
	//	IsSuccess: true,
	//}, nil

	return &block.IsSuccessResponse{}, nil
}
