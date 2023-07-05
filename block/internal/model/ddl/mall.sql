
-- 创建 ManageCryptominer 管理系统矿机表
CREATE TABLE `manage_cryptominer` (
      `cryptominer_typeid` bigint(20) NOT NULL,        -- 矿机种类 根据管理系统生成的金银矿机种类（雪花算法）分类
      `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
      `updated_at` datetime DEFAULT NULL,
      `deleted_at` datetime DEFAULT NULL,
      `adminuser_id` bigint(20) NOT NULL,
      `cryptominer_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
      `cryptominer_picture` VARCHAR(256) NULL  ,
      `cryptominer_describe` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL ,
      `cryptominer_price` int NOT NULL,
      `cryptominer_duration` bigint(20) NOT NULL,       -- 矿机使用时效（天）
      `payment_way` varchar(256) NOT NULL,                 -- 支付方式 0：U 1：ADF
      PRIMARY KEY (`cryptominer_typeid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 ManageProp 管理系统道具表
CREATE TABLE `manage_prop` (
       `prop_typeid` bigint(20) NOT NULL,         -- 道具种类 根据管理系统生成的道具种类（雪花算法）分类，目前只有能量水
       `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
       `updated_at` datetime DEFAULT NULL,
       `deleted_at` datetime DEFAULT NULL,
       `adminuser_id` bigint(20) NOT NULL,
       `prop_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
       `prop_picture` VARCHAR(256) NULL  ,
       `prop_describe` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL ,
       `prop_price` int NOT NULL,
       `payment_way` varchar(256) NOT NULL ,                 -- 支付方式 0：U 1：ADF
       PRIMARY KEY (`prop_typeid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 ManageActivity 管理系统活动表(关联矿机)
CREATE TABLE `manage_activity` (
       `cryptominer_typeid` bigint(20) NOT NULL,         -- 道具种类 根据管理系统生成的道具种类（雪花算法）分类，目前只有能量水
       `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
       `updated_at` datetime DEFAULT NULL,
       `deleted_at` datetime DEFAULT NULL,
       `adminuser_id` bigint(20) NOT NULL,
       `user_amount` bigint(20) NOT NULL,              -- 砍价人数
       `min_price` float NOT NULL,                     -- 最小砍价金额
       `first_bargain_per` float NOT NULL,             -- 首次砍价比率
       `friend_bargain_per` float NOT NULL,            -- 好友砍价比率
       `is_activation` tinyint(1) NOT NULL,            -- 是否开启活动 0：未开启  1：开启
       PRIMARY KEY (`cryptominer_typeid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 Cryptominer 矿机表
CREATE TABLE `cryptominer` (
       `cryptominer_id` bigint(20) NOT NULL,
       `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
       `updated_at` datetime DEFAULT NULL,
       `deleted_at` datetime DEFAULT NULL,
       `user_id` bigint(20) NOT NULL ,
       `cryptominer_typeid` bigint(20) NOT NULL,          -- 矿机种类 根据管理系统生成的金银矿机种类（雪花算法）分类
       `cryptominer_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
       `cryptominer_picture` VARCHAR(256) NULL ,
       `cryptominer_price` int NOT NULL,
       `payment_way` varchar(256) NOT NULL ,                 -- 支付方式 0：U 1：ADF
       `cryptominer_describe` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL ,
       `is_bargain` tinyint(1) NOT NULL,                  -- 是否可砍，根据当天是否砍价决定 0：不可砍 1：可砍
       `purchase_way` varchar(256) NULL ,     -- 购买方式 0：全额购买 1：限时砍价
       `purchase_time` datetime NULL,             -- 购买时间
       `optional_status` varchar(256) NOT NULL,           -- 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
       `cryptominer_start_time` datetime ,    -- 矿机开始运作时间 -- 如果失效后激活，按照激活时间开始
       `cryptominer_end_time` datetime NULL,      -- 矿机失效时间
       `cryptominer_duration` bigint(20) NOT NULL,        -- 矿机使用时效（天）
       PRIMARY KEY (`cryptominer_id`),
       INDEX `idx_user_id` (`user_id`),
       INDEX `idx_cryptominer_price` (`cryptominer_price`),
       INDEX `idx_cryptominer_typeid` (`cryptominer_typeid`),
       INDEX `idx_purchase_way` (`purchase_way`),
       INDEX `idx_optional_status` (`optional_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 Prop 道具表
CREATE TABLE `prop` (
        `prop_id` bigint(20) NOT NULL,
        `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_at` datetime DEFAULT NULL,
        `deleted_at` datetime DEFAULT NULL,
        `user_id` bigint(20) NOT NULL ,
        `prop_typeid` bigint(20) NOT NULL,                 -- 道具种类 根据管理系统生成的道具种类（雪花算法）分类，目前只有能量水
        `prop_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
        `prop_picture` VARCHAR(256) NULL ,
        `prop_price` int NOT NULL,
        `payment_way` varchar(256) NOT NULL ,                 -- 支付方式 0：U 1：ADF
        `prop_describe` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL ,
        `purchase_time` datetime NULL,
        `optional_status` varchar(256) NOT NULL,            -- 道具状态 0：未购买 1：已购买
        PRIMARY KEY (`prop_id`),
        INDEX `idx_user_id` (`user_id`),
        INDEX `idx_prop_typeid` (`prop_typeid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 UserCryptominer 用户矿机表
CREATE TABLE `user_cryptominer` (
        `user_id` bigint(20) ,
        `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_at` datetime DEFAULT NULL,
        `deleted_at` datetime DEFAULT NULL,
        `cryptominer_amount` int NOT NULL,
        `cryptominer_typeid` bigint(20) NOT NULL,   -- 矿机种类 根据管理系统生成的金银矿机种类（雪花算法）分类
        PRIMARY KEY (`user_id`),
        UNIQUE KEY `idx_user_id` (`user_id`),
        INDEX `idx_cryptominer_typeid` (`cryptominer_typeid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 UserProp 用户道具表
CREATE TABLE `user_prop` (
         `user_id` bigint(20) ,
         `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
         `updated_at` datetime DEFAULT NULL,
         `deleted_at` datetime DEFAULT NULL,
         `prop_amount` int NOT NULL,
         `prop_typeid` bigint(20) NOT NULL,         -- 道具种类 根据管理系统生成的道具种类（雪花算法）分类，目前只有能量水
         PRIMARY KEY (`user_id`),
         UNIQUE KEY `idx_user_id` (`user_id`),
         INDEX `idx_prop_typeid` (`prop_typeid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 PurchaseRecord 商品购买记录表
CREATE TABLE `purchase_records` (
        `user_id` bigint(20) NOT NULL,
        `purchase_record_id` bigint(20) NOT NULL,          -- 购买记录id
        `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_at` datetime DEFAULT NULL,
        `deleted_at` datetime DEFAULT NULL,
        `good_name` varchar(256) NOT NULL,
        `good_picture` varchar(256) NULL ,
        `purchase_way` varchar(256) NOT NULL,       -- 购买方式 0：全额购买 1：限时砍价
        `good_quantity` int NOT NULL,
        `purchase_time` datetime NOT NULL,
        `purchase_price` int NULL ,              -- 购买最终价格
        `payment_way` varchar(256) NOT NULL ,                 -- 支付方式 0：U 1：ADF
        PRIMARY KEY (`purchase_record_id`),
        INDEX `idx_good_name` (`good_name`),
        INDEX `idx_good_picture` (`good_picture`),
        INDEX `idx_purchase_way` (`purchase_way`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 bargain 砍价表
CREATE TABLE `bargain` (
       `bargain_id` bigint(20) NOT NULL,
       `cryptominer_id` bigint(20) NOT NULL,
       `cryptominer_typeid` bigint(20) NOT NULL,          -- 矿机种类 根据管理系统生成的金银矿机种类（雪花算法）分类
       `cryptominer_price` int NOT NULL,
       `remaining_price` float NOT NULL,        -- 剩余金额
       `activity_start_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
       `activity_end_time` datetime NULL,
       PRIMARY KEY (`bargain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 friend_assistance 好友助力表
CREATE TABLE `friend_assistance` (
       `bargain_id` bigint(20) NOT NULL,
       `user_id` bigint(20) NOT NULL,
       `user_name` bigint(20) NOT NULL,
       `avatar` int NOT NULL,
       `bargain_price` float NOT NULL,
       PRIMARY KEY (`bargain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 BargainAmount 砍价额度表
CREATE TABLE `bargain_amount` (
          `bargain_id` bigint(20) NOT NULL,
          `first_bargain_percentage` float NOT NULL,
          `bargain_min_price` float NOT NULL,
          `bargain_price` float NOT NULL,
          PRIMARY KEY (`bargain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 wallet 钱包表
CREATE TABLE `wallets` (
       `user_id` bigint(20) NOT NULL,
       `holding_quantity` float NOT NULL,
       `available_quantity` float NOT NULL,
       `frozen_quantity` float NOT NULL,
       `accumulated_reward` float NOT NULL,
       `accumulated_expense` float NOT NULL,
       PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 bill_detail 账单明细表
CREATE TABLE `bill_details` (
        `user_id` bigint(20) NOT NULL,
        `getType` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
        `getTime` datetime NOT NULL,
        `getQuantity` float NOT NULL,
        PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 yesterday_bill 昨日账单表
CREATE TABLE `yesterday_bill` (
          `user_id` bigint(20) NOT NULL,
          `yesterdayRewardQuantity` float NOT NULL,
          `yesterdayExpenseQuantity` float NOT NULL,
          PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 yesterday_reward_details 昨日收益明细表
CREATE TABLE `yesterday_reward_details` (
        `user_id` bigint(20) NOT NULL,
        `rewardType` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
        `rewardAmount` float NOT NULL,
        `rewardTime` datetime NOT NULL,
        PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 yesterday_expense_details 昨日支出明细
CREATE TABLE `yesterday_expense_details` (
         `user_id` bigint(20) NOT NULL,
         `expenseName` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
         `itemQuantity` int NOT NULL,
         `expenseAmount` float NOT NULL,
         `expenseTime` datetime NOT NULL,
         PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
