package other

//
//import (
//	"gorm.io/gorm"
//	"time"
//)
//
//// Mall 商城表
//type Mall struct {
//	gorm.Model
//	// 商城ID
//	MallId int64 `gorm:"primaryKey;type:bigint;not null" json:"mallId"`
//	// 用户ID
//	UserId int64 `gorm:"type:bigint;not null;index:idx_user_id" json:"userId"`
//}
//
//// Good 道具表
//type Good struct {
//	gorm.Model
//	// GoodId
//	GoodId int64 `gorm:"primaryKey;type:bigint;not null" json:"goodId"`
//	// 商城ID
//	MallId int64 `gorm:"type:bigint;not null;index:idx_mall_id" json:"mallId"`
//	// 道具名称
//	GoodName string `gorm:"type:varchar(256);not null;" json:"goodName"`
//	// 道具图片
//	GoodPicture string `gorm:"type:varchar(256);not null;" json:"goodPicture"`
//	// 道具价格
//	GoodPrice int `gorm:"type:int;not null;index:idx_good_price" json:"goodPrice"`
//	// 道具描述
//	GoodDescribe string `gorm:"type:varchar(256);not null;" json:"goodDescribe"`
//	// 道具种类 0：黄金胡萝卜矿机 1：白银胡萝卜矿机 2：能量水
//	GoodType string `gorm:"type:varchar(256);not null;index:idx_good_type" json:"goodType"`
//	// 支付货币种类 0：ADF 1：U
//	CurrencyType string `gorm:"type:varchar(256);not null;index:idx_currency_type" json:"currencyType"`
//	// 是否可砍，矿机可砍，能量水不可砍
//	IsBargain bool `gorm:"type:bool;not null;index:idx_is_bargain" json:"isBargain"`
//	// 购买方式 0：全额购买 1：限时砍价
//	PurchaseWay string `gorm:"type:varchar(256);not null;index:idx_purchase_way" json:"purchaseWay"`
//	// 购买时间
//	PurchaseTime time.Time `gorm:"type:datetime;index:idx_purchase_time" json:"purchaseTime"`
//	// 道具状态 0：未购买 1：砍价中 2：待支付 3：工作中 4：已失效 5：已使用(能量水)
//	OptionalStatus string `gorm:"type:varchar(256);not null;index:idx_optional_status" json:"optionalStatus"`
//	// 道具开始运作时间 -- 如果失效后激活，按照激活时间开始
//	GoodStartTime time.Time `gorm:"type:datetime" json:"goodStartTime"`
//}
//
//// PurchaseRecord 购买记录
//type PurchaseRecord struct {
//	gorm.Model
//	// 购买记录ID
//	PurchaseRecordID int `gorm:"primaryKey;not null" json:"purchaseRecordId"`
//	// 商城ID
//	MallId int64 `gorm:"type:bigint;not null;index:idx_mall_id" json:"mallId"`
//	// 道具名称
//	GoodName string `gorm:"type:varchar(256);not null;index:idx_good_name" json:"goodName"`
//	// 道具图片
//	GoodPicture string `gorm:"type:varchar(256);not null;index:idx_good_picture" json:"goodPicture"`
//	// 购买方式 0：全额购买 1：限时砍价
//	PurchaseWay string `gorm:"type:varchar(256);not null;index:idx_purchase_way" json:"purchaseWay"`
//	// 道具数量
//	GoodQuantity int `gorm:"type:int;not null;" json:"goodQuantity"`
//	// 购买时间
//	PurchaseTime time.Time `gorm:"type:datetime;not null" json:"purchaseTime"`
//	// 购买最终价格
//	PurchasePrice int `gorm:"type:int;not null" json:"purchasePrice"`
//	// 支付货币种类 0：ADF 1：U
//	CurrencyType string `gorm:"type:varchar(256);not null;index:idx_currency_type" json:"currencyType"`
//}
