
-- 创建 ManageGoods 管理系统控制商品表
CREATE TABLE `manage_goods` (
     `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
     `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
     `updated_at` datetime DEFAULT NULL,
     `deleted_at` datetime DEFAULT NULL,
     `user_id` bigint(20) NOT NULL,
     `good_name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
     `good_picture` VARCHAR(256) NULL DEFAULT NULL ,
     `good_describe` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
     `good_price` int NULL DEFAULT NULL,
     `currency_type` varchar(256) NOT NULL,     -- 支付货币种类 0：ADF 1：U
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

