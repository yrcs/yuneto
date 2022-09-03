/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.23.111_3306
 Source Server Type    : MariaDB
 Source Server Version : 100608
 Source Host           : 192.168.23.111:3306
 Source Schema         : hospital

 Target Server Type    : MariaDB
 Target Server Version : 100608
 File Encoding         : 65001

 Date: 04/08/2022 22:58:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for HospitalSetting
-- ----------------------------
DROP TABLE IF EXISTS `HospitalSetting`;
CREATE TABLE `HospitalSetting`  (
  `Id` bigint(20) UNSIGNED NOT NULL,
  `CreatedAt` datetime NULL DEFAULT NULL,
  `UpdatedAt` datetime NULL DEFAULT NULL,
  `DeletedAt` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `Version` bigint(20) NULL DEFAULT 1 COMMENT '版本号（乐观锁专用）',
  `Name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `RegistrationNumber` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `ContactPerson` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `ContactMobile` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `Locked` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `ApiUrl` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `SignatureKey` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  UNIQUE INDEX `Name`(`Name`) USING BTREE,
  UNIQUE INDEX `RegistrationNumber`(`RegistrationNumber`) USING BTREE,
  INDEX `idx_HospitalSetting_DeletedAt`(`DeletedAt`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of HospitalSetting
-- ----------------------------
INSERT INTO `HospitalSetting` VALUES (1553670574587252736, '2022-07-31 17:15:05', NULL, NULL, 1, '中山大学附属第一医院', '350002440102150131', '肖海鹏', '16602000001', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1553671201707003904, '2022-07-31 17:17:34', NULL, NULL, 1, '南方医科大学南方医院', '73759012144011111A1001', '李文源', '16602000002', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554362468401156096, '2022-08-02 15:04:25', NULL, NULL, 1, '广东省人民医院', '350011440102210131', '余学清', '16602000003', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554363205801742336, '2022-08-02 15:07:21', NULL, NULL, 1, '中山大学孙逸仙纪念医院', '351361440104150131', '宋尔卫', '16602000004', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554363680773115904, '2022-08-02 15:09:14', NULL, NULL, 1, '广州医科大学附属第一医院', '45534420544010411A1001', '黄锦坤', '16602000005', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554364142545014784, '2022-08-02 15:11:04', NULL, NULL, 1, '中山大学附属第三医院', '352771440106150131', '戎利民', '16602000006', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554364584700153856, '2022-08-02 15:12:50', NULL, NULL, 1, '广东省中医院', '351383440104154031', '陈达灿', '16602000007', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554365595330285568, '2022-08-02 15:16:51', NULL, NULL, 1, '中山大学肿瘤防治中心', '350003440102151031', '徐瑞华', '16602000008', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554365972050087936, '2022-08-02 15:18:21', NULL, NULL, 1, '广州市妇女儿童医疗中心', '68329213644010411A5391', '韦建瑞', '16602000009', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554366517284442112, '2022-08-02 15:20:31', NULL, NULL, 1, '南方医科大学珠江医院', '72562750544010511A1001', '郭洪波', '16602000010', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554373472111562752, '2022-08-02 15:48:09', NULL, NULL, 1, '广东省人民医院珠海医院（珠海市金湾中心医院）', '31507713344040411A1001', '张忠林', '16602000011', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554374140826226688, '2022-08-02 15:50:48', NULL, NULL, 1, '广东省人民医院南海医院', '45607939444068211c2201', '罗建方', '16602000012', 0, 'http://127.0.0.1:8000', '');
INSERT INTO `HospitalSetting` VALUES (1554374781762015232, '2022-08-02 15:53:21', NULL, NULL, 1, '中山大学孙逸仙纪念医院深汕中心医院', 'B2D00110444150211A1001', '宋尔卫', '16602000004', 0, 'http://127.0.0.1:8000', '');

SET FOREIGN_KEY_CHECKS = 1;
