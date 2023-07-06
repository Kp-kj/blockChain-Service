package logic

import (
	"block/block"
	"block/internal/svc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPurchaseRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPurchaseRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPurchaseRecordLogic {
	return &GetPurchaseRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPurchaseRecordLogic) GetPurchaseRecord(in *block.GetPurchaseRecordRequest) (*block.GetPurchaseRecordResponse, error) {

	record, err := l.svcCtx.PurchaseRecordsModel.FindAllPurchaseRecord(l.ctx, in.UserId)
	if err == nil {
		logx.Error(err)
		return nil, err
	}

	var PurchaseRecordBlock []*block.PurchaseRecord

	for _, onePurchaseRecord := range record {
		purchaseRecord := &block.PurchaseRecord{
			UserId:           onePurchaseRecord.UserId,
			PurchaseRecordId: onePurchaseRecord.PurchaseRecordId,
			GoodName:         onePurchaseRecord.GoodName,
			GoodPicture:      onePurchaseRecord.GoodPicture.String,
			PurchaseWay:      onePurchaseRecord.PurchaseWay, // 购买方式 0：全额购买 1：限时砍价
			GoodQuantity:     onePurchaseRecord.GoodQuantity,
			PurchaseTime:     onePurchaseRecord.PurchaseTime.String(),
			PurchasePrice:    onePurchaseRecord.PurchasePrice.Int64,
		}
		PurchaseRecordBlock = append(PurchaseRecordBlock, purchaseRecord)
	}

	return &block.GetPurchaseRecordResponse{PurchaseRecord: PurchaseRecordBlock}, nil
}
