package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCryptominerStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCryptominerStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCryptominerStatusLogic {
	return &GetUserCryptominerStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCryptominerStatusLogic) GetUserCryptominerStatus(in *block.GetUserCryptominerStatusRequest) (*block.GetUserCryptominerStatusResponse, error) {

	var UserCryptominers []*block.UserCryptominer
	cryptominers, err := l.svcCtx.CryptominerModel.FindUserCryptominers(l.ctx, in.UserId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	for _, cryptominer := range cryptominers {
		UserCryptominer := &block.UserCryptominer{
			CryptominerId:      cryptominer.CryptominerId,
			CryptominerName:    cryptominer.CryptominerName,
			CryptominerPicture: cryptominer.CryptominerPicture.String,
			OptionalStatus:     cryptominer.OptionalStatus,
			CryptominerEndTime: cryptominer.CryptominerEndTime.Time.String(),
		}

		UserCryptominers = append(UserCryptominers, UserCryptominer)
	}

	return &block.GetUserCryptominerStatusResponse{UserCryptominer: UserCryptominers}, nil
}
