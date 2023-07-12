package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPropStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPropStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPropStatusLogic {
	return &GetUserPropStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPropStatusLogic) GetUserPropStatus(in *block.GetUserPropStatusRequest) (*block.GetUserPropStatusResponse, error) {

	var UserProps []*block.UserProp
	props, err := l.svcCtx.PropModel.FindUserProps(l.ctx, in.UserId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	for _, prop := range props {
		UserProp := &block.UserProp{
			PropId:      prop.PropId,
			PropName:    prop.PropName,
			PropPicture: prop.PropPicture.String,
		}

		UserProps = append(UserProps, UserProp)
	}

	return &block.GetUserPropStatusResponse{UserProp: UserProps}, nil
}
