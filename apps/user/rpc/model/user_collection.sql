CREATE TABLE `user_collection` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '收藏Id',
  `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
  `product_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '商品id',
  `is_delete` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '数据创建时间[禁止在代码中赋值]',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '数据更新时间[禁止在代码中赋值]',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UN_collection_uid_product_id` (`uid`,`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户收藏表';