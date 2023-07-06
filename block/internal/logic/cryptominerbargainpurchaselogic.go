package logic

import (
	"block/block"
	"block/internal/model"
	"block/internal/svc"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"log"
	"math/rand"
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
	MinPrice         float64
	UserAmount       int
	FirstBargainPer  float64
	FriendBargainPer float64
}

func (l *CryptominerBargainPurchaseLogic) CryptominerBargainPurchase(in *block.CryptominerBargainRequest) (*block.CryptominerBargainResponse, error) {
	// 发起砍价
	l.ctx.Value("userId")
	Cryptominer, err := l.svcCtx.CryptominerModel.FindOneByCryptominerId(l.ctx, in.CryptominerId)
	if err == nil {
		logx.Error(err)
		return nil, err
	}
	Cryptominer.OptionalStatus = "3" // 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
	err = l.svcCtx.CryptominerModel.Update(l.ctx, Cryptominer)
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	// 创建 Redis 客户端连接
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // 根据实际情况修改 Redis 服务器地址和端口
		Password: "123456",         // 如果需要密码认证，可以在此处填写密码
		DB:       0,                // Redis 数据库编号
	})
	// 检查是否成功连接到 Redis
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis:", pong)

	// 砍价内部逻辑，生成砍价金额切片
	var arr []decimal.Decimal
	CryptominerBargain(Cryptominer, arr)

	// 存入redis
	for _, amount := range arr {
		err := client.RPush(context.Background(), "amounts", amount).Err()
		if err != nil {
			log.Printf("Failed to push amount to Redis queue: %v", err)
		}
	}

	return &block.CryptominerBargainResponse{}, nil
}

func CryptominerBargain(Cryptominer *model.Cryptominer, arr []decimal.Decimal) {
	CryptominerPer := CryptominerType{
		MinPrice:         1.0,
		UserAmount:       30,
		FirstBargainPer:  0.1,
		FriendBargainPer: 0.2,
	}
	var remainingPrice decimal.Decimal = decimal.NewFromFloat(float64(Cryptominer.CryptominerPrice)).Mul(decimal.NewFromFloat(1.0).Sub(decimal.NewFromFloat(CryptominerPer.FirstBargainPer)))
	var price decimal.Decimal = remainingPrice.Sub(decimal.NewFromFloat(CryptominerPer.FirstBargainPer).Mul(decimal.NewFromInt(int64(CryptominerPer.UserAmount))))

	for i := 1; i < CryptominerPer.UserAmount; i++ {
		amount := generateRandomAmount(price)
		price = price.Sub(amount)
		arr = append(arr, amount.Round(2).Add(decimal.NewFromFloat(1.0)))
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
				fmt.Println(arr[i])
			}
		}
	} else if lastRemainingPrice > 6 {
		remainingPrice = remainingPrice.Sub(decimal.NewFromFloat(5.0))
		for i := 0; i < 5; i++ {
			if len(arr) > i {
				arr[i] = arr[i].Add(decimal.NewFromFloat(1.0))
				fmt.Println(arr[i])
			}
		}
	}
	arr = append(arr, remainingPrice.Round(2))
	ShuffleDecimalSlice(arr)
	//var aaa decimal.Decimal
	//	for _, d := range arr {
	//		aaa = aaa.Add(d)
	//	} 测试最后结果
}

// 生成随机金额
func generateRandomAmount(price decimal.Decimal) decimal.Decimal {
	rand.Seed(time.Now().UnixNano())
	f, _ := price.Float64()
	randomFloat := rand.Float64() * f * 0.20
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
