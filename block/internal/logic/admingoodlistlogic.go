package logic

import (
	"context"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminGoodListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminGoodListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminGoodListLogic {
	return &AdminGoodListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminGoodListLogic) AdminGoodList(in *block.AdminGoodListRequest) (*block.AdminGoodListResponse, error) {
	var AdminGoodList []*block.AdminGood

	cryptominerList, _ := l.svcCtx.ManageCryptominerModel.SelectAllCryptominer(l.ctx)
	for _, cryptominer := range cryptominerList {
		NewAdminGood := &block.AdminGood{
			GoodTypeid:   cryptominer.CryptominerTypeid,
			GoodName:     cryptominer.CryptominerName,
			GoodDuration: cryptominer.CryptominerDuration,
			PaymentWay:   cryptominer.PaymentWay,
			PropPrice:    cryptominer.CryptominerPrice,
			GoodStatus:   cryptominer.GoodStatus,
			GoodType:     cryptominer.GoodType,
			GoodPower:    cryptominer.CryptominerPower,
		}
		AdminGoodList = append(AdminGoodList, NewAdminGood)
	}
	propList, _ := l.svcCtx.ManagePropModel.SelectAllProp(l.ctx)
	for _, prop := range propList {
		NewAdminGood := &block.AdminGood{
			GoodTypeid:   prop.PropTypeid,
			GoodName:     prop.PropName,
			GoodDuration: 0,
			PaymentWay:   prop.PaymentWay,
			PropPrice:    prop.PropPrice,
			GoodStatus:   prop.GoodStatus,
			GoodType:     prop.GoodType,
			GoodPower:    0,
		}
		AdminGoodList = append(AdminGoodList, NewAdminGood)
	}

	return &block.AdminGoodListResponse{
		AdminGood: AdminGoodList,
	}, nil

}
