package logic

import (
	"block/internal/model"
	"context"
	"database/sql"
	"github.com/bwmarrin/snowflake"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePropLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePropLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePropLogic {
	return &CreatePropLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePropLogic) CreateProp(in *block.CreatePropRequest) (*block.IsSuccessResponse, error) {
	// 查询是否有道具
	_, err := l.svcCtx.ManagePropModel.FindOneByPropName(l.ctx, in.PropName)
	if err == nil {
		logx.Error(err)
		return nil, err
	}
	typeId, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	createProp := &model.ManageProp{
		PropTypeid:   typeId.Generate().Int64(),
		AdminuserId:  in.AdminuserId,
		PropName:     in.PropName,
		PropPicture:  sql.NullString{String: in.PropPicture},
		PropDescribe: sql.NullString{String: in.PropDescribe},
		PropPrice:    in.PropPrice,
	}

	_, err = l.svcCtx.ManagePropModel.Insert(l.ctx, createProp)
	if err != nil {
		return &block.IsSuccessResponse{
			IsSuccess: false,
		}, nil
	}
	return &block.IsSuccessResponse{
		IsSuccess: true,
	}, nil
}
