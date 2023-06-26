-- 创建 Mall 商城表
CREATE TABLE `malls` (
     `mall_id` bigint(20) NOT NULL,
     `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL,
     `user_id` bigint(20) NOT NULL,
     PRIMARY KEY (`mall_id`),
     UNIQUE KEY `idx_mall_id` (`mall_id`),
     INDEX `idx_user_id` (`user_id`)
     -- CONSTRAINT `fk_malls_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 Good 商品表
CREATE TABLE `goods` (
     `good_id` bigint(20) NOT NULL,
     `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL,
     `mall_id` bigint(20) NOT NULL,
     `good_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `good_picture` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `good_price` int NULL DEFAULT NULL,
     `good_describe` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `good_type` varchar(256) NOT NULL,         -- 道具种类 0：黄金胡萝卜矿机 1：白银胡萝卜矿机 2：能量水
     `currency_type` varchar(256) NOT NULL,     -- 支付货币种类 0：ADF 1：U
     `is_bargain` tinyint(1) NOT NULL,          -- 是否可砍，矿机可砍(true)，能量水不可砍(false)
     `purchase_way` varchar(256) NOT NULL,      -- 购买方式 0：全额购买 1：限时砍价
     `purchase_time` datetime DEFAULT NULL,
     `optional_status` varchar(256) NOT NULL,   -- 道具状态 0：未购买 1：砍价中 2：待支付 3：工作中 4：已失效 5：已使用(能量水)
     `good_start_time` datetime DEFAULT NULL,   -- 道具开始运作时间 -- 如果失效后激活，按照激活时间开始
     PRIMARY KEY (`good_id`),
     UNIQUE KEY `idx_goods_good_id` (`good_id`),
     INDEX `idx_mall_id` (`mall_id`),
     INDEX `idx_good_price` (`good_price`),
     INDEX `idx_good_type` (`good_type`),
     INDEX `idx_currency_type` (`currency_type`),
     INDEX `idx_purchase_way` (`purchase_way`),
     INDEX `idx_optional_status` (`optional_status`)
--      CONSTRAINT `fk_goods_mall_id` FOREIGN KEY (`mall_id`) REFERENCES `malls` (`mall_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 PurchaseRecord 购买记录表
CREATE TABLE `purchase_records` (
    `purchase_record_id` int NOT NULL,
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT NULL,
    `deleted_at` datetime DEFAULT NULL,
    `mall_id` bigint(20) NOT NULL,
    `good_name` varchar(256) NOT NULL,
    `good_picture` varchar(256) NOT NULL,
    `purchase_way` varchar(256) NOT NULL,       -- 购买方式 0：全额购买 1：限时砍价
    `good_quantity` int NOT NULL,
    `purchase_time` datetime NOT NULL,
    `purchase_price` int NOT NULL,              -- 购买最终价格
    `currency_type` varchar(256) NOT NULL,      -- 支付货币种类 0：ADF 1：U
    PRIMARY KEY (`purchase_record_id`),
    UNIQUE KEY `idx_purchase_record_id` (`purchase_record_id`),
    INDEX `idx_mall_id` (`mall_id`),
    INDEX `idx_good_name` (`good_name`),
    INDEX `idx_good_picture` (`good_picture`),
    INDEX `idx_purchase_way` (`purchase_way`),
    INDEX `idx_currency_type` (`currency_type`)
--     CONSTRAINT `fk_purchase_records_mall_id` FOREIGN KEY (`mall_id`) REFERENCES `malls` (`mall_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 profile 个人信息表
CREATE TABLE `profile` (
   `profile_id` BIGINT(20) NOT NULL,
   `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
   `updated_at` DATETIME DEFAULT NULL,
   `deleted_at` DATETIME DEFAULT NULL,
   `user_id` BIGINT(20) NOT NULL,
   `twitter_name` VARCHAR(256) NOT NULL,
   `user_name` VARCHAR(256) NOT NULL,
   `avatar` VARCHAR(256) NOT NULL,
   `is_new` TINYINT(1) NOT NULL DEFAULT 1,
   PRIMARY KEY (`profile_id`),
   INDEX `idx_user_id` (`user_id`),
   INDEX `idx_user_name` (`user_name`),
   INDEX `idx_twitter_name` (`twitter_name`),
   INDEX `idx_is_new` (`is_new`),
   UNIQUE (`profile_id`)
);

-- 创建 bargain 砍价表
CREATE TABLE `bargain` (
   `bargain_id` bigint(20) NOT NULL,
   `good_id` bigint(20) NOT NULL,
   `good_price` int NOT NULL,
   `remaining_price` float NOT NULL,        -- 剩余金额
   `activity_start_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
   `activity_end_time` datetime NOT NULL,
   PRIMARY KEY (`bargain_id`),
   INDEX `idx_good_price` (`good_price`),
   INDEX `idx_bargain_percentage` (`remaining_price`)
--    FOREIGN KEY (`good_id`) REFERENCES `good` (`good_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 BargainAmount 砍价额度表
CREATE TABLE `bargain_amount` (
      `bargain_id` bigint(20) NOT NULL,
      `user_id` bigint(20) NOT NULL,
      `first_bargain_percentage` float NOT NULL,
      `bargain_min_price` float NOT NULL,
      `bargain_price` float NOT NULL,
      PRIMARY KEY (`bargain_id`),
      INDEX `first_bargain_percentage` (`first_bargain_percentage`)
--       FOREIGN KEY (`bargain_id`) REFERENCES `bargain` (`bargain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 关联表：bargain_users_profile
-- CREATE TABLE `bargain_users_profile` (
--      `user_id` bigint(20) NOT NULL,
--      `bargain_id` bigint(20) NOT NULL
-- --      FOREIGN KEY (`user_id`) REFERENCES `profile` (`user_id`),
-- --      FOREIGN KEY (`bargain_id`) REFERENCES `bargain` (`bargain_id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 wallet 钱包表
CREATE TABLE `wallets` (
       `walletId` bigint(20) NOT NULL,
       `userId` bigint(20) NOT NULL,
       `holdingQuantity` float NOT NULL,
       `availableQuantity` float NOT NULL,
       `frozenQuantity` float NOT NULL,
       `accumulatedReward` float NOT NULL,
       `accumulatedExpense` float NOT NULL,
       PRIMARY KEY (`walletId`),
       KEY `idx_user_id` (`userId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 bill_detail 账单明细表
CREATE TABLE `bill_details` (
        `walletId` bigint(20) NOT NULL,
        `getType` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
        `getTime` datetime NOT NULL,
        `getQuantity` float NOT NULL,
        PRIMARY KEY (`walletId`)
--         CONSTRAINT `fk_bill_details_wallets` FOREIGN KEY (`walletId`) REFERENCES `wallets` (`walletId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 yesterday_bill 昨日账单表
CREATE TABLE `yesterday_bill` (
         `walletId` bigint(20) NOT NULL,
         `yesterdayRewardQuantity` float NOT NULL,
         `yesterdayExpenseQuantity` float NOT NULL,
         PRIMARY KEY (`walletId`)
--          CONSTRAINT `fk_yesterday_rewards_wallets` FOREIGN KEY (`walletId`) REFERENCES `wallets` (`walletId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 yesterday_reward_details 昨日收益明细表
CREATE TABLE `yesterday_reward_details` (
        `walletId` bigint(20) NOT NULL,
        `rewardType` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
        `rewardAmount` float NOT NULL,
        `rewardTime` datetime NOT NULL,
        PRIMARY KEY (`walletId`)
--         CONSTRAINT `fk_yesterday_reward_details_wallets` FOREIGN KEY (`walletId`) REFERENCES `wallets` (`walletId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 yesterday_expense_details 昨日支出明细
CREATE TABLE `yesterday_expense_details` (
     `walletId` bigint(20) NOT NULL,
     `expenseName` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
     `itemQuantity` int NOT NULL,
     `expenseAmount` float NOT NULL,
     `expenseTime` datetime NOT NULL,
     PRIMARY KEY (`walletId`)
--      CONSTRAINT `fk_yesterday_expense_details_wallets` FOREIGN KEY (`walletId`) REFERENCES `wallets` (`walletId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

