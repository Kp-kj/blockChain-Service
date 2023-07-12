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

type BargainPayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBargainPayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BargainPayLogic {
	return &BargainPayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BargainPayLogic) BargainPay(in *block.BargainPayRequest) (*block.IsSuccessResponse, error) {

	// 添加矿机属性
	Cryptominer, err := l.svcCtx.CryptominerModel.FindOneByCryptominerId(l.ctx, in.CryptominerId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	Cryptominer.PurchaseWay = sql.NullString{String: "1", Valid: true} // 购买方式 0：全额购买 1：限时砍价
	Cryptominer.PurchaseTime = sql.NullTime{Time: time.Now(), Valid: true}
	Cryptominer.OptionalStatus = "1" // 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
	Cryptominer.CryptominerStartTime = Cryptominer.PurchaseTime
	Cryptominer.CryptominerEndTime = sql.NullTime{Time: time.Now().Add(time.Hour * 24 * time.Duration(Cryptominer.CryptominerDuration)), Valid: true}

	err = l.svcCtx.CryptominerModel.Update(l.ctx, Cryptominer)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	BargainRecord, err := l.svcCtx.BargainModel.FindBargainByCryptominerId(l.ctx, in.UserId, in.CryptominerId)
	// 添加购买记录
	recordId, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	PurchaseRecordModel := &model.PurchaseRecords{
		UserId:           in.UserId,
		PurchaseRecordId: recordId.Generate().Int64(),
		GoodName:         Cryptominer.CryptominerName,
		GoodPicture:      Cryptominer.CryptominerPicture,
		PurchaseWay:      "1", // 购买方式 0：全额购买 1：限时砍价
		GoodQuantity:     1,
		PurchaseTime:     Cryptominer.PurchaseTime.Time,
		PurchasePrice:    sql.NullFloat64{Float64: float64(BargainRecord.RemainingPrice), Valid: true},
		PaymentWay:       Cryptominer.PaymentWay,
	}

	_, err = l.svcCtx.PurchaseRecordsModel.Insert(l.ctx, PurchaseRecordModel)
	if err != nil {
		logx.Error(err)
		return &block.IsSuccessResponse{IsSuccess: false}, nil
	}

	return &block.IsSuccessResponse{IsSuccess: true}, nil
}
