/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : product

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 13/09/2022 15:41:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` smallint unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `parentid` smallint NOT NULL DEFAULT '0' COMMENT '父类别id当id=0时说明是根节点,一级类别',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '类别名称',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '类别状态1-正常,2-已废弃',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品类别表';

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `cateid` smallint unsigned NOT NULL DEFAULT '0' COMMENT '类别Id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '商品名称',
  `subtitle` varchar(200) DEFAULT '' COMMENT '商品副标题',
  `images` text COMMENT '图片地址,json格式,扩展用',
  `detail` text COMMENT '商品详情',
  `price` decimal(20,2) NOT NULL DEFAULT '0.00' COMMENT '价格,单位-元保留两位小数',
  `stock` int NOT NULL DEFAULT '0' COMMENT '库存数量',
  `status` int NOT NULL DEFAULT '1' COMMENT '商品状态.1-在售 2-下架 3-删除',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `ix_cateid` (`cateid`),
  KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='商品表';

SET FOREIGN_KEY_CHECKS = 1;
