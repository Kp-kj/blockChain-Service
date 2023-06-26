package other

import (
	"gorm.io/gorm"
	"time"
)

type Bargain struct {
	gorm.Model
	// BargainId
	BargainId int64 `gorm:"primary_key;type:bigint;not null" json:"bargainId"`
	// 砍价人信息，如设定最大数量50人助力，将有51个，第一个是自己
	UsersProfile []Profile `gorm:"foreignKey:ProfileId" json:"usersProfile"`
	// 道具 包含（道具名称，道具图片，GoodId）
	Good Good `gorm:"foreignKey:GoodId" json:"good"`
	// 道具价格
	GoodPrice int `gorm:"type:int;not null;index:idx_good_price" json:"goodPrice"`
	// 剩余金额 如设定50人助力，金额有51份，第一刀是自己
	BargainAmounts []BargainAmount `gorm:"foreignKey:BargainId" json:"bargainAmounts"`
	//剩余金额 通过砍价金额赋值给
	RemainingPrice float32 `gorm:"type:float;not null;index:idx_bargain_percentage" json:"remainingPrice"`
	// 活动开始时间
	ActivityStartTime time.Time `gorm:"type:datetime;not null" json:"activityStartTime"`
	// 活动结束时间
	ActivityEndTime time.Time `gorm:"type:datetime;not null" json:"activityEndTime"`
}

// BargainAmount 金额计算
type BargainAmount struct {
	gorm.Model
	// Bargain 获取 BargainAmount.Bargain.GoodPrice
	Bargain Bargain `gorm:"foreignKey:BargainId" json:"bargain"`
	// 初次砍价百分比
	FirstBargainPercentage float32 `gorm:"type:float;not null;index:first_bargain_percentage" json:"firstBargainPercentage"`
	// 砍价最小额度
	BargainMinPrice float32 `gorm:"type:float;not null;column:bargain_min_price" json:"bargainMinPrice"`
	// 本次砍价额度
	BargainPrice float32 `gorm:"type:float;not null;column:bargain_price" json:"bargainPrice"`
}

// Profile 暂时从用户服务拿到
type Profile struct {
	gorm.Model
	//使用雪花id 方便迁移
	ProfileId int64 `gorm:"primary_key;not null" json:"ProfileId"`
	//用户id
	UserId int64 `gorm:"type:bigint;not null;index:idx_user_id" json:"userId"`
	//用户名
	UserName string `gorm:"type:varchar(256);not null;index:idx_user_name" json:"userName"`
	//用户头像
	Avatar string `gorm:"type:varchar(256);not null" json:"avatar"`
}
