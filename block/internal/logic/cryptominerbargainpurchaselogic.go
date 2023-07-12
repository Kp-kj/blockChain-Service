package logic

import (
	"block/block"
	"block/internal/model"
	"block/internal/svc"
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/shopspring/decimal"
	"math/rand"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CryptominerBargainPurchaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCryptominerBargainPurchaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CryptominerBargainPurchaseLogic {
	return &CryptominerBargainPurchaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type CryptominerType struct {
	MinPrice         decimal.Decimal
	UserAmount       int64
	FirstBargainPer  decimal.Decimal
	FriendBargainPer decimal.Decimal
}

func (l *CryptominerBargainPurchaseLogic) CryptominerBargainPurchase(in *block.CryptominerBargainRequest) (*block.CryptominerBargainResponse, error) {
	// 发起砍价
	Cryptominer, err := l.svcCtx.CryptominerModel.FindOneByCryptominerId(l.ctx, in.CryptominerId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	// 1.查找砍价活动，获取砍价数据
	activityData, err := l.svcCtx.ManageActivityModel.FindOneByCryptominerTypeid(l.ctx, Cryptominer.CryptominerTypeid)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	// 2. 砍价内部逻辑，生成砍价金额切片
	priceArr := CryptominerBargain(Cryptominer, activityData)
	//改变方案
	values := make([]string, len(*priceArr))
	for i, decimal := range *priceArr {
		values[i] = decimal.String()
	}
	UnusedData := strings.Join(values, ",")

	// 3.改变矿机状态
	Cryptominer.OptionalStatus = "3" // 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
	err = l.svcCtx.CryptominerModel.Update(l.ctx, Cryptominer)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	// 6.添加砍价记录,后续变动
	bargainid, err := snowflake.NewNode(1)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	per := float64(Cryptominer.CryptominerPrice) * (1.0 - activityData.FirstBargainPer)
	BargainRecord := &model.Bargain{
		BargainId:         bargainid.Generate().Int64(),
		UserId:            in.UserId,
		CryptominerId:     Cryptominer.CryptominerId,
		CryptominerTypeid: Cryptominer.CryptominerTypeid,
		CryptominerPrice:  Cryptominer.CryptominerPrice,
		RemainingPrice:    per,
		ActivityStartTime: time.Now(),
		UnusedData:        UnusedData,
		UsedData:          "",
	}
	_, err = l.svcCtx.BargainModel.Insert(l.ctx, BargainRecord)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &block.CryptominerBargainResponse{BargainId: BargainRecord.BargainId}, nil
}

func CryptominerBargain(Cryptominer *model.Cryptominer, activityData *model.ManageActivity) *[]decimal.Decimal {
	var arr []decimal.Decimal
	CryptominerPer := CryptominerType{
		MinPrice:         decimal.NewFromFloat(activityData.MinPrice),         // 最小砍价金额
		UserAmount:       activityData.UserAmount,                             // 帮砍价人数
		FirstBargainPer:  decimal.NewFromFloat(activityData.FirstBargainPer),  // 第一次砍价金额
		FriendBargainPer: decimal.NewFromFloat(activityData.FriendBargainPer), // 朋友帮砍权重（权重0-1之间，越大砍价的峰值越高，建议 0.2-0.3 之间）
	}

	var remainingPrice decimal.Decimal = decimal.NewFromFloat(float64(Cryptominer.CryptominerPrice)).Mul(decimal.NewFromFloat(1.0).Sub(CryptominerPer.FirstBargainPer))
	var price decimal.Decimal = remainingPrice.Sub(CryptominerPer.MinPrice.Mul(decimal.NewFromInt(CryptominerPer.UserAmount)))

	var i int64
	for i = 1; i < CryptominerPer.UserAmount; i++ {
		amount := generateRandomAmount(price, activityData.FriendBargainPer)
		price = price.Sub(amount)
		arr = append(arr, amount.Round(2).Add(CryptominerPer.MinPrice))
	}
	for _, num := range arr {
		remainingPrice = remainingPrice.Sub(num)
	}
	lastRemainingPrice, _ := remainingPrice.Float64()
	// 判断最后一次生成数额较大时，平衡
	if lastRemainingPrice > 4 && lastRemainingPrice <= 6 {
		remainingPrice = remainingPrice.Sub(decimal.NewFromFloat(3.0))
		for i := 0; i < 3; i++ {
			if len(arr) > i {
				arr[i] = arr[i].Add(decimal.NewFromFloat(1.0))
			}
		}
	} else if lastRemainingPrice > 6 {
		remainingPrice = remainingPrice.Sub(decimal.NewFromFloat(5.0))
		for i := 0; i < 5; i++ {
			if len(arr) > i {
				arr[i] = arr[i].Add(decimal.NewFromFloat(1.0))
			}
		}
	}

	arr = append(arr, remainingPrice.Round(2))
	ShuffleDecimalSlice(arr)
	return &arr
}

// 生成随机金额
func generateRandomAmount(price decimal.Decimal, friendBargainPer float64) decimal.Decimal {
	rand.Seed(time.Now().UnixNano())
	f, _ := price.Float64()
	randomFloat := rand.Float64() * f * friendBargainPer
	return decimal.NewFromFloat(randomFloat)
}

func ShuffleDecimalSlice(arr []decimal.Decimal) {
	rand.Seed(time.Now().UnixNano())
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}
