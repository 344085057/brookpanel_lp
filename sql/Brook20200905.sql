/*
 Navicat Premium Data Transfer

 Source Server         : brook
 Source Server Type    : MySQL
 Source Server Version : 50727
 Source Host           : 122.112.204.106:3306
 Source Schema         : MyBrookData

 Target Server Type    : MySQL
 Target Server Version : 50727
 File Encoding         : 65001

 Date: 05/09/2020 14:55:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for lp_brook_commodity
-- ----------------------------
DROP TABLE IF EXISTS `lp_brook_commodity`;
CREATE TABLE `lp_brook_commodity`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '商品',
  `sort` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '商品类别',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `describe` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '商品描述',
  `money` int(200) NOT NULL DEFAULT 0 COMMENT '商品价格 100 = 1元',
  `time` int(10) NOT NULL DEFAULT 0 COMMENT '时长（天数）',
  `cover` int(1) NOT NULL DEFAULT -1 COMMENT '-1:覆盖/1:叠加/ 默认覆盖',
  `state` int(1) NOT NULL DEFAULT 1 COMMENT '-1:禁用/1:启用 默认启用',
  `ll` decimal(40, 5) NOT NULL DEFAULT 0.00000 COMMENT '流量  mb',
  `sx` int(255) NOT NULL DEFAULT 0 COMMENT '顺序',
  `table_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '直接修改表的日期',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lp_brook_commodity
-- ----------------------------
INSERT INTO `lp_brook_commodity` VALUES (1, '活动套餐', '1天体验', '体验', 1, 1, -1, 1, 1000.00000, 0, '2020-08-20 22:07:21');
INSERT INTO `lp_brook_commodity` VALUES (2, '活动套餐', '3天体验', '体验', 3, 0, -1, 1, 1000.00000, 0, '2020-08-20 22:07:20');
INSERT INTO `lp_brook_commodity` VALUES (4, '包月', '中杯包月', '包月', 15, 30, 1, 1, 50000.00000, 0, '2020-08-20 22:07:24');
INSERT INTO `lp_brook_commodity` VALUES (5, '包月', '大杯包月', '包月', 20, 30, 1, 1, 100000.00000, 0, '2020-08-20 22:07:24');
INSERT INTO `lp_brook_commodity` VALUES (15, 'test', 'test', '1', 1, 1, -1, 1, 1000.00000, 0, '2020-09-02 12:21:20');

-- ----------------------------
-- Table structure for lp_brook_gg
-- ----------------------------
DROP TABLE IF EXISTS `lp_brook_gg`;
CREATE TABLE `lp_brook_gg`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `text` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `state` int(255) NOT NULL DEFAULT 1 COMMENT '-1:禁用/1:启用 默认启用',
  `g_type` int(10) NOT NULL DEFAULT 1 COMMENT '1:公告 2使用教程',
  `sx` int(255) NOT NULL DEFAULT 0 COMMENT '顺序',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lp_brook_gg
-- ----------------------------
INSERT INTO `lp_brook_gg` VALUES (1, '请使用最新版本的brook V20200901', '<h1>\r\n	<strong>请使用最新版本的brook V20200901</strong>\r\n</h1>', 1, 1, 0, '2020-08-20 17:03:27');
INSERT INTO `lp_brook_gg` VALUES (2, '购买套餐后请等待30秒左右', '<h1>\r\n	购买套餐后请等待30秒左右\r\n</h1>', 1, 1, 1, '2020-08-20 17:10:06');
INSERT INTO `lp_brook_gg` VALUES (3, '连接服务器的密码', '<h2>\r\n	节点连接的密码是您注册时的密码\r\n</h2>\r\n<p>\r\n</p>\r\n<h2>\r\n	因此你需要保护好你的密码\r\n</h2>', 1, 2, 0, '2020-08-20 17:21:50');

-- ----------------------------
-- Table structure for lp_brook_moneycdk
-- ----------------------------
DROP TABLE IF EXISTS `lp_brook_moneycdk`;
CREATE TABLE `lp_brook_moneycdk`  (
  `id` int(255) NOT NULL AUTO_INCREMENT COMMENT '充值码',
  `cdk` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT 'cdk',
  `money` int(255) NOT NULL DEFAULT 0 COMMENT '金额',
  `create_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `use_time` datetime(0) NULL DEFAULT NULL COMMENT '使用时间',
  `use_uid` int(11) NOT NULL DEFAULT 0 COMMENT '使用者id',
  `lp_brook_user_id` int(11) NOT NULL,
  `use_uid2` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for lp_brook_server
-- ----------------------------
DROP TABLE IF EXISTS `lp_brook_server`;
CREATE TABLE `lp_brook_server`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT 'ip地址',
  `flow_ratio` double(255, 2) NOT NULL DEFAULT 1.00 COMMENT '流量比例',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '服务标题',
  `domain` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '域名',
  `type` int(255) NOT NULL DEFAULT 1 COMMENT '服务器类型 1为Brook/2为socks5 /3为ws/4为wss/-1关闭',
  `delay` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '服务器延迟',
  `dk` int(255) NOT NULL DEFAULT 0 COMMENT '带宽 Mbps',
  `peed` int(255) NOT NULL DEFAULT 0 COMMENT '限速 s/Mb',
  `describe` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '描述',
  `state` int(255) NOT NULL DEFAULT 1 COMMENT '状态 -1:停用 1启用',
  `sx` int(255) NOT NULL DEFAULT 0 COMMENT '顺序',
  `table_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '直接修改表的日期',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lp_brook_server
-- ----------------------------
INSERT INTO `lp_brook_server` VALUES (1, '127.0.0.1', 1.00, 'test', '', 1, '', 100, 20, '', 1, 0, '2020-09-05 14:43:07');

-- ----------------------------
-- Table structure for lp_brook_user
-- ----------------------------
DROP TABLE IF EXISTS `lp_brook_user`;
CREATE TABLE `lp_brook_user`  (
  `u_id` int(11) NOT NULL AUTO_INCREMENT,
  `u_email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '邮箱',
  `u_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `u_passwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `u_port` int(255) NOT NULL COMMENT '端口',
  `u_flow` decimal(40, 5) NOT NULL DEFAULT 99999999.00000 COMMENT '剩余流量 mb',
  `u_is_admin` int(1) NOT NULL DEFAULT 0 COMMENT '是否是管理员 0普通用户/1管理员/-1停用',
  `expire_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'vip到期时间',
  `u_flow_total` decimal(40, 5) NOT NULL DEFAULT 0.00000 COMMENT '总使用流量',
  `u_money` int(20) NOT NULL DEFAULT 0 COMMENT '金币 100 = 1元',
  `table_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '直接修改表的日期',
  `create_time` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
  `update_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '更新日期',
  `lp_brook_moneycdk_id` int(11) NOT NULL,
  PRIMARY KEY (`u_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 40 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lp_brook_user
-- ----------------------------
INSERT INTO `lp_brook_user` VALUES (1, 'admin@gmail.com', 'admin', '123456', 12345, 1000.00000, 1, '2020-09-04 19:23:58', 0.00000, 999, '2020-09-05 14:55:13', '2020-08-20 12:03:38', '2020-09-05 14:55:13', 0);

-- ----------------------------
-- Table structure for lp_brook_user_commodity_log
-- ----------------------------
DROP TABLE IF EXISTS `lp_brook_user_commodity_log`;
CREATE TABLE `lp_brook_user_commodity_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT 0 COMMENT '用户id',
  `cid` int(11) NOT NULL DEFAULT 0 COMMENT '商品id',
  `sendingtime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发生时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Fixed;

-- ----------------------------
-- Table structure for lp_brook_user_login_log
-- ----------------------------
DROP TABLE IF EXISTS `lp_brook_user_login_log`;
CREATE TABLE `lp_brook_user_login_log`  (
  `l_id` int(255) NOT NULL AUTO_INCREMENT COMMENT '登录日志',
  `u_id` int(11) NOT NULL COMMENT '用户id',
  `login_time` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登录时间',
  `login_ip` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录ip',
  `login_ip_address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '登录归属地',
  PRIMARY KEY (`l_id`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 41 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lp_brook_user_login_log
-- ----------------------------
INSERT INTO `lp_brook_user_login_log` VALUES (1, 1, '2020-09-02 13:10:26', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (2, 1, '2020-09-02 13:11:33', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (3, 1, '2020-09-02 18:15:59', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (4, 1, '2020-09-02 18:21:20', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (5, 1, '2020-09-02 18:26:44', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (6, 1, '2020-09-03 14:39:43', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (7, 1, '2020-09-03 14:40:49', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (8, 1, '2020-09-03 14:43:31', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (9, 1, '2020-09-03 15:13:16', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (10, 1, '2020-09-03 15:16:35', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (11, 1, '2020-09-03 15:17:44', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (12, 1, '2020-09-03 16:41:26', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (13, 1, '2020-09-03 16:53:01', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (14, 1, '2020-09-03 16:54:24', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (15, 1, '2020-09-03 17:23:44', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (16, 1, '2020-09-03 17:24:54', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (17, 1, '2020-09-03 17:26:58', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (18, 1, '2020-09-03 17:35:52', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (19, 1, '2020-09-03 17:38:23', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (20, 1, '2020-09-03 17:38:56', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (21, 1, '2020-09-03 17:43:46', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (22, 1, '2020-09-03 17:59:33', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (23, 1, '2020-09-03 18:15:22', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (24, 1, '2020-09-03 18:16:53', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (25, 1, '2020-09-03 18:20:39', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (26, 1, '2020-09-03 18:27:27', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (27, 1, '2020-09-03 18:28:25', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (28, 1, '2020-09-03 18:29:53', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (29, 1, '2020-09-03 18:31:02', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (30, 1, '2020-09-03 18:33:32', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (31, 1, '2020-09-03 18:40:24', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (32, 1, '2020-09-03 18:44:34', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (33, 1, '2020-09-03 18:48:01', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (34, 1, '2020-09-03 18:49:41', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (35, 1, '2020-09-03 18:58:27', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (36, 1, '2020-09-03 18:59:27', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (37, 1, '2020-09-03 19:05:59', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (38, 1, '2020-09-03 19:23:44', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (39, 1, '2020-09-03 20:02:13', '127.0.0.1', '');
INSERT INTO `lp_brook_user_login_log` VALUES (40, 1, '2020-09-05 14:55:15', '127.0.0.1', '');

-- ----------------------------
-- Table structure for lp_brook_user_money_log
-- ----------------------------
DROP TABLE IF EXISTS `lp_brook_user_money_log`;
CREATE TABLE `lp_brook_user_money_log`  (
  `id` int(11) NOT NULL,
  `uid` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户id',
  `money` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '金额',
  `sendingtime` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发生时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for lp_sys
-- ----------------------------
DROP TABLE IF EXISTS `lp_sys`;
CREATE TABLE `lp_sys`  (
  `s_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `s_value` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  PRIMARY KEY (`s_name`) USING BTREE
) ENGINE = MyISAM CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lp_sys
-- ----------------------------
INSERT INTO `lp_sys` VALUES ('remote_u', 'qweqwesadasdf');
INSERT INTO `lp_sys` VALUES ('remote_p', 'dfgfhbfgjtyjyjytj');

SET FOREIGN_KEY_CHECKS = 1;
