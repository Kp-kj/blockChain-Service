package logic

import (
	"context"
	"database/sql"
	"time"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivateCryptominerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivateCryptominerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivateCryptominerLogic {
	return &ActivateCryptominerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivateCryptominerLogic) ActivateCryptominer(in *block.ActivateCryptominerRequest) (*block.ActivateCryptominerResponse, error) {

	var isSuccess bool = true
	var propEnough bool = true

	cryptominer, err := l.svcCtx.CryptominerModel.FindOne(l.ctx, in.CryptominerId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	if cryptominer.CryptominerName == "黄金胡萝卜矿机" {
		golds, err := l.svcCtx.PropModel.FindUserPropForGold(l.ctx, in.UserId)
		if err != nil {
			propEnough = false
			isSuccess = false
		} else {
			if len(golds) == 2 {
				for _, gold := range golds {
					err := l.svcCtx.PropModel.Delete(l.ctx, gold.PropId)
					if err != nil {
						logx.Error(err)
						return nil, err
					}
				}
				cryptominer.OptionalStatus = "1"
				cryptominer.CryptominerStartTime = sql.NullTime{Time: time.Now(), Valid: true}
				cryptominer.CryptominerEndTime = sql.NullTime{Time: time.Now().Add(time.Hour * 24 * time.Duration(cryptominer.CryptominerDuration)), Valid: true}
				err := l.svcCtx.CryptominerModel.Update(l.ctx, cryptominer)
				if err != nil {
					logx.Error(err)
					isSuccess = false
				}
			} else if len(golds) < 2 {
				propEnough = false
				isSuccess = false
			}
		}
	} else { //否则是白银矿机
		sliver, err := l.svcCtx.PropModel.FindUserPropForSliver(l.ctx, in.UserId)
		if err != nil {
			propEnough = false
			isSuccess = false
		} else {
			err := l.svcCtx.PropModel.Delete(l.ctx, sliver.PropId)
			if err != nil {
				logx.Error(err)
				return nil, err
			}
			cryptominer.OptionalStatus = "1"
			cryptominer.CryptominerStartTime = sql.NullTime{Time: time.Now(), Valid: true}
			cryptominer.CryptominerEndTime = sql.NullTime{Time: time.Now().Add(time.Hour * 24 * time.Duration(cryptominer.CryptominerDuration)), Valid: true}
			err = l.svcCtx.CryptominerModel.Update(l.ctx, cryptominer)
			if err != nil {
				logx.Error(err)
				isSuccess = false
			}
		}

	}

	return &block.ActivateCryptominerResponse{IsSuccess: isSuccess, PropEnough: propEnough}, nil
}
