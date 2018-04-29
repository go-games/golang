/*
 Navicat Premium Data Transfer

 Source Server         : 10.211.55.4
 Source Server Type    : MySQL
 Source Server Version : 50639
 Source Host           : 10.211.55.4:3306
 Source Schema         : game

 Target Server Type    : MySQL
 Target Server Version : 50639
 File Encoding         : 65001

 Date: 29/04/2018 21:57:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `key` varchar(50) NOT NULL DEFAULT '0',
  `value` varchar(50) NOT NULL DEFAULT '0',
  `parent` int(11) NOT NULL DEFAULT '0',
  `struct` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of config
-- ----------------------------
BEGIN;
INSERT INTO `config` VALUES (1, 'Version', '0.4', 0, 0);
INSERT INTO `config` VALUES (2, 'LogPath', '0', 0, 0);
INSERT INTO `config` VALUES (3, 'TCPAddr', '127.0.0.1:3563', 0, 0);
INSERT INTO `config` VALUES (4, 'WSAddr', '127.0.0.1:3653', 0, 0);
INSERT INTO `config` VALUES (5, 'ConsolePort', '7771', 0, 0);
INSERT INTO `config` VALUES (6, 'MaxConnNum', '20000', 0, 0);
INSERT INTO `config` VALUES (7, 'Mysql', '0', 0, 1);
INSERT INTO `config` VALUES (8, 'DBname', 'game', 7, 0);
INSERT INTO `config` VALUES (9, 'DBaddr', '10.211.55.4', 7, 0);
INSERT INTO `config` VALUES (10, 'DBport', '3306', 7, 0);
INSERT INTO `config` VALUES (11, 'DBuser', 'root', 7, 0);
INSERT INTO `config` VALUES (12, 'DBpasswd', '123456', 7, 0);
INSERT INTO `config` VALUES (14, 'LogLevel', 'LogLevel', 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` varchar(30) DEFAULT '0',
  `passwd` varchar(50) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user` (`user`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'Golang.Ltd', '123456');
INSERT INTO `user` VALUES (2, 'Gola1ng.Ltd', '123456');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
