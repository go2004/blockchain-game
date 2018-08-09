/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : 127.0.0.1:3306
 Source Schema         : blockchain

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : 65001

 Date: 03/08/2018 14:32:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for accounts
-- ----------------------------
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts`  (
  `id` int(32) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `sdk_account` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `sdk_client_id` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `sdk_access_token` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `channel_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `imei` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `game_ver` int(10) UNSIGNED NULL DEFAULT NULL,
  `extra` varchar(255) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `phone` varchar(32) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `createtime` bigint(64) NULL DEFAULT 0,
  `name` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL,
  `style` int(10) UNSIGNED NULL DEFAULT 0,
  `app_coins` double NULL DEFAULT 0,
  `compute_power` double NULL DEFAULT 0,
  `logintime` bigint(20) NULL DEFAULT NULL,
  `offflinetime` bigint(20) NULL DEFAULT NULL,
  `btc` double NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `account`(`account`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 268435456 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for block_machines
-- ----------------------------
DROP TABLE IF EXISTS `block_machines`;
CREATE TABLE `block_machines`  (
  `index` int(11) NULL DEFAULT NULL,
  `timestamp` bigint(64) NULL DEFAULT NULL,
  `player_id` int(32) UNSIGNED NULL DEFAULT NULL,
  `data_id` int(11) NULL DEFAULT NULL,
  `price` double(64, 0) NULL DEFAULT NULL,
  `count` int(11) NULL DEFAULT NULL,
  `hash` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL,
  `prev_hash` varchar(64) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for mine_machines
-- ----------------------------
DROP TABLE IF EXISTS `mine_machines`;
CREATE TABLE `mine_machines`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_id` int(11) NULL DEFAULT NULL,
  `player_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `location` int(11) NULL DEFAULT NULL,
  `start_time` bigint(64) NULL DEFAULT NULL,
  `end_time` bigint(64) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
