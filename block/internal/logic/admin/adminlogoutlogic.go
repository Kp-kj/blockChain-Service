package admin

import (
	"context"

	"block/internal/svc"
	"block/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogoutLogic {
	return &AdminLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogoutLogic) AdminLogout() (resp *types.AdminLoginReply, err error) {
	// todo: add your logic here and delete this line

	return
}
