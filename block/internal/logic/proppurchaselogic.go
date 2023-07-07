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

	// 查询道具，（此道具仅供提供基本属性用，不会被消费掉）
	Prop, err := l.svcCtx.PropModel.FindOneByPropId(l.ctx, in.PropId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

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
		GoodQuantity:     in.GoodQuantity,
		PurchaseTime:     Prop.PurchaseTime.Time,
		PurchasePrice:    sql.NullFloat64{Float64: float64(Prop.PropPrice) * float64(in.GoodQuantity), Valid: true},
		PaymentWay:       Prop.PaymentWay,
	}
	_, err = l.svcCtx.PurchaseRecordsModel.Insert(l.ctx, PurchaseRecordModel)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	var i int64
	var propList []*model.Prop
	for i = 0; i < in.GoodQuantity; i++ {
		propId, err := snowflake.NewNode(1)
		if err != nil {
			return nil, err
		}
		prop := &model.Prop{
			UserId:         in.UserId,
			PropId:         propId.Generate().Int64() + i, //同一微妙产生的雪花是一样的，加上数字以区分
			PropTypeid:     Prop.PropTypeid,
			PropName:       Prop.PropName,
			PropPicture:    Prop.PropPicture,
			PropDescribe:   Prop.PropDescribe,
			PropPrice:      Prop.PropPrice,
			PaymentWay:     Prop.PaymentWay,
			PurchaseTime:   sql.NullTime{Time: time.Now(), Valid: true},
			OptionalStatus: "1", // 状态 0：未购买 1：已购买
		}
		propList = append(propList, prop)
	}
	_, err = l.svcCtx.PropModel.InsertManyProp(l.ctx, propList)
	if err != nil {
		logx.Error(err)
		return &block.IsSuccessResponse{IsSuccess: false}, nil
	}
	// 暂时未添加用户道具信息（钱包区）

	return &block.IsSuccessResponse{IsSuccess: true}, nil
}
