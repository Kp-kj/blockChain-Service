package other

//
//import "time"
//
//// Wallet 钱包结构体
//type Wallet struct {
//	gorm.Model
//	WalletID           int64   `gorm:"primaryKey;type:bigint;not null" json:"walletId"`      // 钱包ID
//	UserID             int64   `gorm:"type:bigint;not null;index:idx_user_id" json:"userId"` // 用户ID
//	HoldingQuantity    float32 `gorm:"type:float;not null" json:"holdingQuantity"`           // 持仓数量
//	AvailableQuantity  float32 `gorm:"type:float;not null" json:"availableQuantity"`         // 可用数量
//	FrozenQuantity     float32 `gorm:"type:float;not null" json:"frozenQuantity"`            // 冻结数量
//	AccumulatedReward  float32 `gorm:"type:float;not null" json:"accumulatedReward"`         // 累计收益
//	AccumulatedExpense float32 `gorm:"type:float;not null" json:"accumulatedExpense"`        // 累计支出
//}
//
//// BillDetail 账单明细版块
//type BillDetail struct {
//	gorm.Model
//	WalletID    int64     `gorm:"primaryKey;type:bigint;not null" json:"walletId"` // 钱包ID
//	GetType     string    `gorm:"type:varchar(256);not null" json:"getType"`       // 获取类型
//	GetTime     time.Time `gorm:"type:datetime;not null" json:"getTime"`           // 获取时间
//	GetQuantity float32   `gorm:"type:float;not null" json:"getQuantity"`          // 获取数量
//}
//
//// YesterdayReward 昨日收益版块
//type YesterdayReward struct {
//	gorm.Model
//	WalletID                 int64   `gorm:"primaryKey;type:bigint;not null" json:"walletId"`     // 钱包ID
//	YesterdayRewardQuantity  float32 `gorm:"type:float;not null" json:"yesterdayRewardQuantity"`  // 昨日收益数量
//	YesterdayExpenseQuantity float32 `gorm:"type:float;not null" json:"yesterdayExpenseQuantity"` // 昨日支出数量
//}
//
//// YesterdayRewardDetail 昨日收益明细版块
//type YesterdayRewardDetail struct {
//	gorm.Model
//	WalletID     int64     `gorm:"primaryKey;type:bigint;not null" json:"walletId"` // 钱包ID
//	RewardType   string    `gorm:"type:varchar(256);not null" json:"rewardType"`    // 收益类型
//	RewardAmount float32   `gorm:"type:float;not null" json:"rewardAmount"`         // 收益金额
//	RewardTime   time.Time `gorm:"type:datetime;not null" json:"rewardTime"`        // 收益时间
//}
//
//// YesterdayExpenseDetail 昨日支出明细版块
//type YesterdayExpenseDetail struct {
//	gorm.Model
//	WalletID      int64     `gorm:"primaryKey;type:bigint;not null" json:"walletId"` // 钱包ID
//	ExpenseName   string    `gorm:"type:varchar(256);not null" json:"expenseName"`   // 支出名称
//	ItemQuantity  int       `gorm:"type:int;not null" json:"itemQuantity"`           // 购买道具个数
//	ExpenseAmount float32   `gorm:"type:float;not null" json:"expenseAmount"`        // 支出金额
//	ExpenseTime   time.Time `gorm:"type:datetime;no
//t null" json:"expenseTime"` // 支出时间
//}
