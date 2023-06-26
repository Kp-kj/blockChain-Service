package user

import (
	"context"

	"block/internal/svc"
	"block/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelpCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHelpCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelpCategoryListLogic {
	return &GetHelpCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHelpCategoryListLogic) GetHelpCategoryList() (resp *types.HelpCategoryListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
