
-- 创建 ManageCryptominer 管理系统矿机表
CREATE TABLE `manage_cryptominer` (
      `cryptominer_typeid` bigint(20) NOT NULL,        -- 矿机种类 根据管理系统生成的金银矿机种类（雪花算法）分类
      `id` int NOT NULL AUTO_INCREMENT UNIQUE ,
      `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
      `updated_at` datetime DEFAULT NULL,
      `deleted_at` datetime DEFAULT NULL,
      `adminuser_id` bigint(20) NOT NULL,
      `cryptominer_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
      `cryptominer_picture` VARCHAR(256) NULL  ,
      `cryptominer_describe` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL ,
      `cryptominer_price` int NOT NULL,
      `cryptominer_duration` bigint(20) NOT NULL,       -- 矿机使用时效（天）
      `cryptominer_power` bigint(20) NOT NULL,          -- 矿机算力
      `payment_way` varchar(256) NOT NULL,              -- 支付方式 0：U 1：ADF
      `good_status` varchar(256) NOT NULL,              -- 商品状态 0：待上架  1：已上架 2：未上架
      `good_type` varchar(256) NOT NULL ,               -- 道具类型 0：胡萝卜 1：能量水
      PRIMARY KEY (`cryptominer_typeid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 ManageProp 管理系统道具表
CREATE TABLE `manage_prop` (
       `prop_typeid` bigint(20) NOT NULL,         -- 道具种类 根据管理系统生成的道具种类（雪花算法）分类，目前只有能量水
       `id` int NOT NULL AUTO_INCREMENT UNIQUE ,
       `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
       `updated_at` datetime DEFAULT NULL,
       `deleted_at` datetime DEFAULT NULL,
       `adminuser_id` bigint(20) NOT NULL,
       `prop_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
       `prop_picture` VARCHAR(256) NULL  ,
       `prop_describe` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL ,
       `prop_price` int NOT NULL,
       `payment_way` varchar(256) NOT NULL ,             -- 支付方式 0：U 1：ADF
       `good_status` varchar(256) NOT NULL,              -- 商品状态 0：待上架  1：已上架 2：未上架
       `good_type` varchar(256) NOT NULL ,               -- 道具类型 0：胡萝卜 1：能量水
       PRIMARY KEY (`prop_typeid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 ManageActivity 管理系统活动表(关联矿机)
CREATE TABLE `manage_activity` (
       `activity_id` bigint(20) NOT NULL,
       `id` int NOT NULL AUTO_INCREMENT UNIQUE ,
       `cryptominer_typeid` bigint(20) NOT NULL,         -- 道具种类 根据管理系统生成的道具种类（雪花算法）分类，目前只有能量水
       `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
       `updated_at` datetime DEFAULT NULL,
       `deleted_at` datetime DEFAULT NULL,
       `adminuser_id` bigint(20) NOT NULL,
       `user_amount` bigint(20) NOT NULL,              -- 砍价人数
       `min_price` decimal(10,2)  NOT NULL,                     -- 最小砍价金额
       `first_bargain_per` decimal(10,2)  NOT NULL,             -- 首次砍价比率
       `friend_bargain_per` decimal(10,2)  NOT NULL,            -- 好友砍价比率
       `is_activation` tinyint(1) NOT NULL,            -- 是否开启活动 0：未开启  1：开启
       PRIMARY KEY (`activity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 Cryptominer 矿机表
CREATE TABLE `cryptominer` (
       `cryptominer_id` bigint(20) NOT NULL,
       `id` int NOT NULL AUTO_INCREMENT UNIQUE ,
       `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
       `updated_at` datetime DEFAULT NULL,
       `deleted_at` datetime DEFAULT NULL,
       `user_id` bigint(20) NOT NULL ,
       `cryptominer_typeid` bigint(20) NOT NULL,         -- 矿机种类 根据管理系统生成的金银矿机种类（雪花算法）分类
       `cryptominer_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
       `cryptominer_picture` VARCHAR(256) NULL ,
       `cryptominer_price` int NOT NULL,
       `payment_way` varchar(256) NOT NULL ,             -- 支付方式 0：U 1：ADF
       `cryptominer_describe` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL ,
       `is_bargain` tinyint(1) NOT NULL,                 -- 是否可砍，根据当天是否砍价决定 0：不可砍 1：可砍
       `purchase_way` varchar(256) NULL ,                -- 购买方式 0：全额购买 1：限时砍价
       `purchase_time` datetime NULL,                    -- 购买时间
       `optional_status` varchar(256) NOT NULL,          -- 矿机状态 0：未购买 1：工作中 2：已失效 3：砍价中
       `cryptominer_start_time` datetime ,               -- 矿机开始运作时间 -- 如果失效后激活，按照激活时间开始
       `cryptominer_end_time` datetime NULL,             -- 矿机失效时间
       `cryptominer_duration` bigint(20) NOT NULL,       -- 矿机使用时效（天）
       `cryptominer_power` bigint(20) NOT NULL,          -- 矿机算力
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
        `id` int NOT NULL AUTO_INCREMENT UNIQUE ,
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

-- 创建 PurchaseRecord 商品购买记录表
CREATE TABLE `purchase_records` (
        `user_id` bigint(20) NOT NULL,
        `id` int NOT NULL AUTO_INCREMENT UNIQUE ,
        `purchase_record_id` bigint(20) NOT NULL,          -- 购买记录id
        `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
        `updated_at` datetime DEFAULT NULL,
        `deleted_at` datetime DEFAULT NULL,
        `good_name` varchar(256) NOT NULL,
        `good_picture` varchar(256) NULL ,
        `purchase_way` varchar(256) NOT NULL,       -- 购买方式 0：全额购买 1：限时砍价
        `good_quantity` int NOT NULL,
        `purchase_time` datetime NOT NULL,
        `purchase_price` decimal(10,2) NULL ,              -- 购买最终价格
        `payment_way` varchar(256) NOT NULL ,                 -- 支付方式 0：U 1：ADF
        PRIMARY KEY (`purchase_record_id`),
        INDEX `idx_good_name` (`good_name`),
        INDEX `idx_good_picture` (`good_picture`),
        INDEX `idx_purchase_way` (`purchase_way`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- 支付账号.-- 支付单id
-- `chain_purchase_account` varchar(256) NOT NULL,
-- `chain_purchase_id` varchar(256) NOT NULL,

-- 创建合约 config
CREATE TABLE `chain_config` (
        `id` int NOT NULL AUTO_INCREMENT,
        `name` varchar(255) NOT NULL,
        `chain_id` int NOT NULL,
        `contract_ad_machine` varchar(255) NOT NULL,
        `chain_block` varchar(255) NOT NULL,
        `from_block` varchar(255) NOT NULL,
        `update_date` datetime NOT NULL,
        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 bargain 砍价表
CREATE TABLE `bargain` (
       `bargain_id` bigint(20) NOT NULL,
       `id` int NOT NULL AUTO_INCREMENT UNIQUE ,
       `user_id` bigint(20) NOT NULL,
       `cryptominer_id` bigint(20) NOT NULL,
       `cryptominer_typeid` bigint(20) NOT NULL,          -- 矿机种类 根据管理系统生成的金银矿机种类（雪花算法）分类
       `cryptominer_price` int NOT NULL,
       `remaining_price` decimal(10,2) NOT NULL,        -- 剩余金额
       `activity_start_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
       `activity_end_time` datetime NULL,
       `unused_data` varchar(512) NOT NULL,
       `used_data` varchar(512) NOT NULL,
       PRIMARY KEY (`bargain_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 bargain 砍价规则表
CREATE TABLE `bargain_rule` (
       `id` bigint(20) NOT NULL,
       `bargain_rule` varchar(256) NOT NULL,
       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建 friend_assistance 好友助力表
CREATE TABLE `friend_assistance` (
       `id` int NOT NULL AUTO_INCREMENT UNIQUE ,
       `bargain_id` bigint(20) NOT NULL,
       `user_id` bigint(20) NOT NULL,
       `user_name` varchar(256) NOT NULL,
       `twitter_name` varchar(256) NOT NULL,
       `avatar` varchar(256) NOT NULL,
       `bargain_price` decimal(10,2) NOT NULL,
       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

