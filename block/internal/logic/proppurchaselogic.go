package logic

import (
	"block/internal/model"
	"context"
	"database/sql"
	"github.com/bwmarrin/snowflake"
	"time"

	"block/block"
	"block/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PropPurchaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPropPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PropPurchaseLogic {
	return &PropPurchaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PropPurchaseLogic) PropPurchase(in *block.PropPurchaseRequest) (*block.IsSuccessResponse, error) {
	// 添加矿机属性
	Prop, err := l.svcCtx.PropModel.FindOneByPropId(l.ctx, in.PropId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	Prop.UserId = in.UserId
	Prop.PurchaseTime = sql.NullTime{Time: time.Now(), Valid: true}
	Prop.OptionalStatus = "1" // 状态 0：未购买 1：已购买

	l.svcCtx.PropModel.Update(l.ctx, Prop)

	// 添加购买记录
	recordId, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	PurchaseRecordModel := &model.PurchaseRecords{
		UserId:           in.UserId,
		PurchaseRecordId: recordId.Generate().Int64(),
		GoodName:         Prop.PropName,
		GoodPicture:      Prop.PropPicture,
		PurchaseWay:      "0", // 购买方式 0：全额购买 1：限时砍价
		GoodQuantity:     1,
		PurchaseTime:     Prop.PurchaseTime.Time,
		PurchasePrice:    sql.NullInt64{Int64: int64(Prop.PropPrice), Valid: true},
		PaymentWay:       Prop.PaymentWay,
	}

	_, err = l.svcCtx.PurchaseRecordsModel.Insert(l.ctx, PurchaseRecordModel)
	if err != nil {
		logx.Error(err)
		return &block.IsSuccessResponse{IsSuccess: false}, nil
	}
	// 暂时未添加用户道具信息（钱包区）

	return &block.IsSuccessResponse{IsSuccess: true}, nil
}
