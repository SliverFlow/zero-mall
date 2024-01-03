/*
 Navicat Premium Data Transfer

 Source Server         : 本地_8
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : localhost:3306
 Source Schema         : zero-mall

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 02/01/2024 22:55:12
*/
create database if not exists `zero-mall`;
use `zero-mall`;


SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for business
-- ----------------------------
DROP TABLE IF EXISTS `business`;
CREATE TABLE `business`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `business_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户 id',
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '关联用户已存在的 user ',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商家名称',
  `detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商家详情',
  `score` bigint NOT NULL DEFAULT 0 COMMENT '商家评分',
  `image` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图片地址 json 格式',
  `status` bigint NOT NULL DEFAULT 0 COMMENT '商户状态.0-暂存 1-正常 2-禁用',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_business_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of business
-- ----------------------------
INSERT INTO `business` VALUES (1, '2023-09-06 22:07:17.862', '2023-10-04 17:14:19.304', NULL, '155831367513817088', '9da5c129-2cfe-41c6-add9-0efe9958578d', '脑子挖掉了专卖店', '专门售卖脑子的专卖店', 1, '[\"https://zheng-bbx-b.oss-cn-hangzhou.aliyuncs.com/rotation/6a632ae1-e24f-48dd-90b8-221074601969.jpg\"]', 1);
INSERT INTO `business` VALUES (5, '2024-01-02 15:22:28.652', '2024-01-02 15:33:28.599', NULL, '198491259315568640', '16a7470d-4489-4306-b448-1bd87b6baf52', '我是第二个商家', '第二个测试商家', 0, '[]', 1);

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `cart_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '购物车业务id',
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户业务id',
  `product_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品业务 id ',
  `quantity` bigint NOT NULL DEFAULT 0 COMMENT '数量',
  `checked` bigint NOT NULL DEFAULT 0 COMMENT '是否选择,1=已勾选,0=未勾选',
  `product_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '冗余商品名称',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cart_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cart
-- ----------------------------
INSERT INTO `cart` VALUES (1, '2023-10-27 21:34:28.878', '2023-10-27 22:11:15.942', NULL, '174304890041286656', '3bc22666-494f-4448-8785-8c6cec889911', '165907838173331456', 10, 0, 'oppo reno 10 ');
INSERT INTO `cart` VALUES (2, '2023-12-24 17:51:18.395', '2023-12-24 17:51:18.395', '2023-12-24 21:24:55.685', '195267222565175296', '3bc22666-494f-4448-8785-8c6cec889911', '164876320638386176', 1, 1, '小米 MIX FOLD 3');
INSERT INTO `cart` VALUES (3, '2023-12-26 16:34:31.166', '2023-12-26 16:34:31.166', '2023-12-26 16:35:11.038', '195972674160574464', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (4, '2023-12-26 16:34:32.349', '2023-12-26 16:34:32.349', '2023-12-26 16:35:12.972', '195972679160184832', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (5, '2023-12-26 16:34:33.229', '2023-12-26 16:34:33.229', '2023-12-26 16:35:14.703', '195972682851172352', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (6, '2023-12-26 16:34:33.858', '2023-12-26 16:34:33.858', '2023-12-26 16:35:16.262', '195972685489389568', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (7, '2023-12-26 16:34:34.479', '2023-12-26 16:34:34.479', '2023-12-26 16:35:17.924', '195972688094052352', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (8, '2023-12-26 16:34:35.074', '2023-12-26 16:34:35.074', '2023-12-26 16:35:20.564', '195972690589663232', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (9, '2023-12-26 16:34:35.665', '2023-12-26 16:34:35.665', '2023-12-26 16:35:22.365', '195972693068496896', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (10, '2023-12-26 16:34:36.192', '2023-12-26 16:34:36.192', '2023-12-26 16:35:24.445', '195972695278895104', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (11, '2023-12-26 16:34:36.760', '2023-12-26 16:34:36.760', '2023-12-26 16:35:26.209', '195972697661259776', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (12, '2023-12-26 16:34:37.304', '2023-12-26 16:34:37.304', '2023-12-26 16:34:59.513', '195972699942961152', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (13, '2023-12-26 16:34:37.871', '2023-12-26 16:34:37.871', '2023-12-26 16:35:06.171', '195972702321131520', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (14, '2023-12-26 16:34:38.321', '2023-12-27 15:12:53.847', NULL, '195972704212762624', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 3, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (15, '2023-12-26 16:34:38.924', '2023-12-26 16:34:38.924', '2023-12-26 16:35:03.985', '195972706741927936', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (16, '2023-12-26 16:34:39.369', '2023-12-26 16:34:39.369', '2023-12-26 16:34:57.311', '195972708608393216', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (17, '2023-12-26 16:34:39.955', '2023-12-26 16:34:39.955', '2023-12-26 16:35:01.937', '195972711062061056', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (18, '2023-12-26 16:34:40.393', '2023-12-26 16:34:40.393', '2023-12-26 16:34:55.085', '195972712899166208', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', 1, 1, '小米 MIX Alpna');
INSERT INTO `cart` VALUES (19, '2024-01-02 17:45:16.558', '2024-01-02 17:45:16.558', NULL, '198527195701198848', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '198519803806629888', 1, 1, '华为机通用官网GT3PRO智能手表');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `category_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品分类id',
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '父分类id 当id=0时说明是根节点,一级类别 ',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类名',
  `status` bigint NOT NULL DEFAULT 0 COMMENT '类别状态 0-暂存 1-正常,2-已废弃',
  `type` bigint NOT NULL DEFAULT 0 COMMENT '分类类别 0 系统类别 1 商家添加的类别',
  `business_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户 id 只有为商家类别的时候才拥有值',
  `sorted` bigint NOT NULL DEFAULT 0 COMMENT '排序标记',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_category_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 49 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category` VALUES (1, '2023-09-18 12:12:20.937', '2023-10-02 09:08:02.018', NULL, '160030298137313280', '0', '电子产品', 1, 0, '', 0);
INSERT INTO `category` VALUES (2, '2023-09-18 12:13:10.221', '2023-10-04 10:36:39.147', NULL, '160030504853585920', '0', '电视 智慧屏', 1, 0, '', 6);
INSERT INTO `category` VALUES (3, '2023-09-18 15:44:20.465', '2023-10-01 21:03:20.842', NULL, '160083647704481792', '160030298137313280', '折叠屏', 1, 0, '', 1);
INSERT INTO `category` VALUES (4, '2023-09-18 20:48:30.781', '2023-10-01 21:03:31.564', NULL, '160160195069493248', '0', '生活 百货', 1, 0, '', 3);
INSERT INTO `category` VALUES (5, '2023-09-18 20:53:41.570', '2023-09-18 20:53:41.570', NULL, '160161498629816320', '0', '出行 穿戴', 1, 0, '', 0);
INSERT INTO `category` VALUES (6, '2023-09-18 20:55:07.514', '2023-10-01 21:03:33.417', NULL, '160161859100884992', '0', '健康 儿童', 1, 0, '', 4);
INSERT INTO `category` VALUES (7, '2023-09-18 21:00:37.762', '2023-10-01 21:03:29.549', NULL, '160163244261392384', '0', '电源 配件', 1, 0, '', 2);
INSERT INTO `category` VALUES (8, '2023-09-19 09:53:11.686', '2023-09-19 09:53:11.686', '2023-09-21 16:07:13.188', '160357666710241280', '160030504853585920', '11111', 0, 0, '', 0);
INSERT INTO `category` VALUES (9, '2023-09-19 09:54:25.735', '2023-09-19 09:54:25.735', '2023-09-21 16:07:13.188', '160357977294258176', '160030504853585920', '78778', 0, 0, '', 0);
INSERT INTO `category` VALUES (10, '2023-09-19 10:11:21.088', '2023-09-19 10:11:21.088', '2023-09-21 16:07:13.188', '160362235989213184', '160030504853585920', '0909090', 0, 0, '', 0);
INSERT INTO `category` VALUES (11, '2023-09-19 10:11:31.364', '2023-09-19 10:11:31.364', '2023-09-21 16:07:13.188', '160362279094075392', '160030504853585920', '90877', 0, 0, '', 0);
INSERT INTO `category` VALUES (12, '2023-09-19 10:12:47.494', '2023-09-19 10:12:47.494', '2023-09-21 16:07:13.188', '160362598410633216', '160030504853585920', '111', 1, 0, '', 0);
INSERT INTO `category` VALUES (13, '2023-09-19 10:14:17.631', '2023-09-19 10:14:17.631', '2023-09-21 16:07:13.188', '160362976460029952', '160030504853585920', '88980', 1, 0, '', 4);
INSERT INTO `category` VALUES (14, '2023-09-19 20:54:53.748', '2023-10-01 21:03:58.413', NULL, '160524189227696128', '160160195069493248', '洗漱用品', 1, 0, '', 1);
INSERT INTO `category` VALUES (15, '2023-09-19 20:55:24.929', '2023-10-01 21:03:59.789', NULL, '160524320010289152', '160160195069493248', '夏天衬衣', 1, 0, '', 2);
INSERT INTO `category` VALUES (16, '2023-09-19 20:55:39.561', '2023-10-01 21:04:01.267', NULL, '160524381381345280', '160160195069493248', '洗漱', 1, 0, '', 3);
INSERT INTO `category` VALUES (17, '2023-09-19 20:59:23.596', '2023-09-19 20:59:23.596', NULL, '160525321052241920', '160030298137313280', '小米', 1, 0, '', 1);
INSERT INTO `category` VALUES (18, '2023-09-19 20:59:33.239', '2023-09-19 20:59:33.239', NULL, '160525361497915392', '160030298137313280', '荣耀', 1, 0, '', 2);
INSERT INTO `category` VALUES (19, '2023-09-19 20:59:44.507', '2023-09-19 20:59:44.507', NULL, '160525408759332864', '160030298137313280', '华为', 1, 0, '', 0);
INSERT INTO `category` VALUES (20, '2023-09-19 21:02:28.111', '2023-09-19 21:02:28.111', '2023-09-21 16:07:00.269', '160526094964244480', '0', 'test 撇量删除', 0, 0, '', 0);
INSERT INTO `category` VALUES (21, '2023-09-19 21:02:33.658', '2023-09-19 21:02:33.658', '2023-09-21 16:07:00.269', '160526118230048768', '160526094964244480', '12132', 0, 0, '', 0);
INSERT INTO `category` VALUES (22, '2023-09-19 21:02:37.752', '2023-09-19 21:02:37.752', '2023-09-21 16:07:00.269', '160526135401529344', '160526094964244480', '人防2方法2', 0, 0, '', 0);
INSERT INTO `category` VALUES (23, '2023-09-19 21:21:26.801', '2023-09-19 21:21:26.801', NULL, '160530870938517504', '160030504853585920', '海信', 1, 0, '', 0);
INSERT INTO `category` VALUES (24, '2023-09-19 22:21:28.517', '2023-09-19 22:23:01.566', '2023-09-21 16:07:23.736', '160545977651314688', '0', '热斯特', 1, 0, '', 0);
INSERT INTO `category` VALUES (25, '2023-09-19 22:21:33.110', '2023-09-19 22:22:59.672', '2023-09-21 16:07:23.736', '160545996932530176', '160545977651314688', '23日日3月2 ', 0, 0, '', 0);
INSERT INTO `category` VALUES (26, '2023-09-19 22:21:43.244', '2023-09-19 22:22:59.672', '2023-09-21 16:07:23.736', '160546039437606912', '160545977651314688', '它3天3弹钢琴4个', 0, 0, '', 0);
INSERT INTO `category` VALUES (27, '2023-09-19 22:21:46.922', '2023-09-19 22:22:59.672', '2023-09-21 16:07:23.736', '160546054864257024', '160545977651314688', '切换过热害人', 0, 0, '', 0);
INSERT INTO `category` VALUES (28, '2023-09-19 22:22:53.067', '2023-09-19 22:22:59.672', '2023-09-21 16:07:23.736', '160546332296495104', '160545977651314688', '八旗二', 0, 0, '', 3);
INSERT INTO `category` VALUES (29, '2023-09-21 10:03:39.738', '2023-09-21 10:03:39.738', NULL, '161085076678262784', '160030298137313280', 'oppo', 1, 0, '', 3);
INSERT INTO `category` VALUES (30, '2023-09-21 10:03:48.760', '2023-09-21 10:03:48.760', NULL, '161085114523467776', '160030298137313280', 'vivo', 1, 0, '', 4);
INSERT INTO `category` VALUES (31, '2023-09-21 10:04:19.879', '2023-10-01 21:03:46.062', NULL, '161085245046013952', '160161859100884992', '益智文具', 1, 0, '', 1);
INSERT INTO `category` VALUES (32, '2023-09-21 10:04:44.943', '2023-10-01 21:03:47.612', NULL, '161085350172049408', '160161859100884992', '图书', 1, 0, '', 2);
INSERT INTO `category` VALUES (33, '2023-09-21 10:59:27.224', '2023-09-21 10:59:27.224', '2023-09-21 16:07:00.269', '161099117043793920', '0', '13535', 0, 0, '', 0);
INSERT INTO `category` VALUES (34, '2023-09-21 10:59:32.882', '2023-09-21 10:59:32.882', '2023-09-21 16:07:00.269', '161099140787748864', '161099117043793920', '3题前天33 ', 0, 0, '', 0);
INSERT INTO `category` VALUES (35, '2023-09-21 10:59:41.343', '2023-09-21 10:59:41.343', '2023-09-21 16:07:00.269', '161099176275755008', '161099117043793920', '34条3有4去', 0, 0, '', 0);
INSERT INTO `category` VALUES (36, '2023-09-21 10:59:45.265', '2023-09-21 10:59:45.265', '2023-09-21 16:07:00.269', '161099192725815296', '161099117043793920', '34永泰4有', 0, 0, '', 0);
INSERT INTO `category` VALUES (37, '2023-09-21 14:57:49.047', '2023-09-21 14:58:02.939', NULL, '161159103228821504', '160161498629816320', '自行车', 1, 0, '', 1);
INSERT INTO `category` VALUES (38, '2023-09-21 14:58:00.643', '2023-09-21 14:58:00.643', NULL, '161159151886942208', '160161498629816320', '代步', 1, 0, '', 2);
INSERT INTO `category` VALUES (39, '2023-09-21 14:58:21.458', '2023-09-21 14:58:21.458', NULL, '161159239195574272', '160161498629816320', '智能手表', 1, 0, '', 3);
INSERT INTO `category` VALUES (40, '2023-09-21 14:58:38.796', '2023-09-21 14:58:38.796', NULL, '161159311912222720', '160161498629816320', '卫星电话', 1, 0, '', 4);
INSERT INTO `category` VALUES (41, '2023-09-21 15:52:34.066', '2023-10-01 21:03:49.488', NULL, '161172881588764672', '160161859100884992', '生命源泉', 1, 0, '', 2);
INSERT INTO `category` VALUES (42, '2023-09-21 15:52:56.348', '2023-10-01 21:03:50.899', NULL, '161172975075606528', '160161859100884992', '玩具', 1, 0, '', 4);
INSERT INTO `category` VALUES (43, '2023-09-21 15:53:15.191', '2023-10-01 21:03:52.673', NULL, '161173054108876800', '160161859100884992', '迪士尼乐园', 1, 0, '', 5);
INSERT INTO `category` VALUES (44, '2023-10-01 21:13:12.935', '2023-10-01 21:13:12.935', NULL, '164877453939326976', '160161498629816320', '无线耳机', 1, 0, '', 1);
INSERT INTO `category` VALUES (45, '2023-10-04 10:36:15.711', '2023-10-04 10:36:15.711', NULL, '165804322859663360', '0', '办公 桌椅', 1, 0, '', 4);
INSERT INTO `category` VALUES (46, '2023-10-04 10:37:11.047', '2023-10-04 10:37:11.047', NULL, '165804554976641024', '0', '零食 坚果', 1, 0, '', 4);
INSERT INTO `category` VALUES (47, '2023-10-04 10:38:17.395', '2023-10-04 10:38:17.395', NULL, '165804833260322816', '0', '图书 教育', 1, 0, '', 4);
INSERT INTO `category` VALUES (48, '2023-10-04 10:38:42.785', '2023-10-04 10:38:42.785', NULL, '165804939753701376', '0', '二手', 1, 0, '', 7);

-- ----------------------------
-- Table structure for category_product
-- ----------------------------
DROP TABLE IF EXISTS `category_product`;
CREATE TABLE `category_product`  (
  `product_id` bigint UNSIGNED NOT NULL,
  `category_id` bigint UNSIGNED NOT NULL,
  PRIMARY KEY (`product_id`, `category_id`) USING BTREE,
  INDEX `fk_category_product_category`(`category_id`) USING BTREE,
  CONSTRAINT `fk_category_product_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_category_product_product` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of category_product
-- ----------------------------
INSERT INTO `category_product` VALUES (6, 1);
INSERT INTO `category_product` VALUES (7, 1);
INSERT INTO `category_product` VALUES (8, 1);
INSERT INTO `category_product` VALUES (9, 1);
INSERT INTO `category_product` VALUES (11, 1);
INSERT INTO `category_product` VALUES (12, 1);
INSERT INTO `category_product` VALUES (13, 1);
INSERT INTO `category_product` VALUES (14, 1);
INSERT INTO `category_product` VALUES (18, 1);
INSERT INTO `category_product` VALUES (10, 5);
INSERT INTO `category_product` VALUES (15, 5);
INSERT INTO `category_product` VALUES (16, 5);
INSERT INTO `category_product` VALUES (17, 5);
INSERT INTO `category_product` VALUES (19, 5);
INSERT INTO `category_product` VALUES (20, 5);
INSERT INTO `category_product` VALUES (21, 5);
INSERT INTO `category_product` VALUES (22, 5);
INSERT INTO `category_product` VALUES (23, 5);
INSERT INTO `category_product` VALUES (7, 17);
INSERT INTO `category_product` VALUES (8, 17);
INSERT INTO `category_product` VALUES (9, 17);
INSERT INTO `category_product` VALUES (11, 17);
INSERT INTO `category_product` VALUES (6, 19);
INSERT INTO `category_product` VALUES (12, 19);
INSERT INTO `category_product` VALUES (13, 19);
INSERT INTO `category_product` VALUES (14, 19);
INSERT INTO `category_product` VALUES (18, 29);
INSERT INTO `category_product` VALUES (15, 39);
INSERT INTO `category_product` VALUES (16, 39);
INSERT INTO `category_product` VALUES (19, 39);
INSERT INTO `category_product` VALUES (20, 39);
INSERT INTO `category_product` VALUES (21, 39);
INSERT INTO `category_product` VALUES (22, 39);
INSERT INTO `category_product` VALUES (23, 39);
INSERT INTO `category_product` VALUES (10, 44);
INSERT INTO `category_product` VALUES (17, 44);

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `order_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单业务id',
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户业务id',
  `shopping_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收获信息业务id',
  `payment` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '实际付款金额',
  `payment_type` bigint NOT NULL DEFAULT 1 COMMENT '支付类型,1-在线支付',
  `postage` bigint NOT NULL DEFAULT 0 COMMENT '运费 单位是 元',
  `status` bigint NOT NULL DEFAULT 10 COMMENT '订单状态:70-已取消-10-未付款，20-已付款，30-待发货 40-待收货，50-交易成功，60-交易关闭 80-订单完成',
  `payment_time` bigint NOT NULL DEFAULT 0 COMMENT '支付时间',
  `send_time` bigint NOT NULL DEFAULT 0 COMMENT '发货时间',
  `end_time` bigint NOT NULL DEFAULT 0 COMMENT '交易完成时间',
  `close_time` bigint NOT NULL DEFAULT 0 COMMENT '交易关闭时间',
  `business_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '冗余字段 商品的商家信息',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_order_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES (1, '2023-10-11 20:42:39.702', '2023-10-11 20:42:39.702', '2023-11-02 10:41:28.873', '168493643433656320', '3bc22666-494f-4448-8785-8c6cec889911', '', 0.00, 1, 0, 70, 0, 0, 0, 0, '155831367513817088');
INSERT INTO `order` VALUES (4, '2023-10-18 10:17:10.645', '2023-10-18 10:17:10.645', NULL, '170872950244130816', '3bc22666-494f-4448-8785-8c6cec889911', '', 10999.00, 1, 0, 20, 0, 0, 0, 0, '155831367513817088');
INSERT INTO `order` VALUES (5, '2023-10-18 10:19:24.300', '2023-10-18 10:19:24.300', NULL, '170873510833831936', '3bc22666-494f-4448-8785-8c6cec889911', '', 0.00, 1, 0, 10, 0, 0, 0, 0, '155831367513817088');
INSERT INTO `order` VALUES (6, '2023-12-27 15:21:28.386', '2023-12-27 15:51:12.771', '2023-12-27 16:57:27.451', '196316679142522880', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '', 0.00, 1, 0, 70, 0, 0, 0, 0, '155831367513817088');
INSERT INTO `order` VALUES (7, '2024-01-02 17:18:39.663', '2024-01-02 17:18:39.663', NULL, '198520497846501376', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '', 0.00, 1, 0, 10, 0, 0, 0, 0, '198491259315568640');

-- ----------------------------
-- Table structure for order_item
-- ----------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `order_item_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单子表业务id',
  `order_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '订单业务id',
  `user_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户业务id',
  `product_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品业务 id',
  `prod_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
  `prod_image` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品图片地址',
  `currentunit_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '生成订单时的商品单价，单位是元,保留两位小数',
  `quantity` bigint NOT NULL DEFAULT 0 COMMENT '商品数量',
  `total_price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品总价,单位是元,保留两位小数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_order_item_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order_item
-- ----------------------------
INSERT INTO `order_item` VALUES (1, '2023-10-11 20:42:39.721', '2023-10-11 20:42:39.721', '2023-11-02 10:41:28.877', '168493643534319617', '168493643433656320', '3bc22666-494f-4448-8785-8c6cec889911', '164876320638386176', '小米 MIX FOLD 3', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/20a171a9-1c02-bb31-fc61-d3dcb9a179be.jpg\"}]', 10999.00, 9, 98991.00);
INSERT INTO `order_item` VALUES (4, '2023-10-18 10:17:10.647', '2023-10-18 10:17:10.647', NULL, '170872950252519425', '170872950244130816', '3bc22666-494f-4448-8785-8c6cec889911', '164876320638386176', '小米 MIX FOLD 3', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/20a171a9-1c02-bb31-fc61-d3dcb9a179be.jpg\"}]', 10999.00, 1, 10999.00);
INSERT INTO `order_item` VALUES (5, '2023-10-18 10:19:24.303', '2023-10-18 10:19:24.303', NULL, '170873510846414849', '170873510833831936', '3bc22666-494f-4448-8785-8c6cec889911', '164876320638386176', '小米 MIX FOLD 3', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/20a171a9-1c02-bb31-fc61-d3dcb9a179be.jpg\"}]', 10999.00, 1, 10999.00);
INSERT INTO `order_item` VALUES (6, '2023-12-27 15:21:28.418', '2023-12-27 15:21:28.418', NULL, '196316679486455809', '196316679142522880', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '164875718655098880', '小米 MIX Alpna', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/d65e316f-b19d-88de-d857-e91899739545.jpg\"}]', 19999.00, 1, 19999.00);
INSERT INTO `order_item` VALUES (7, '2024-01-02 17:18:39.682', '2024-01-02 17:18:39.682', NULL, '198520497934581761', '198520497846501376', 'ba479a90-f156-4c2e-a219-f0635bb49f74', '198517763139321856', '果坊华强北S9Ultra智能运动手表', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/16a7470d-4489-4306-b448-1bd87b6baf52/a2e8bc98-e15c-ae34-9699-d6bd1e783f6e.avif\"}]', 368.00, 1, 368.00);

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `product_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品业务 id',
  `business_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属商户 id ',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名',
  `subtitle` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品副标题',
  `image` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '图片地址 json 格式',
  `detail` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商品详情',
  `stock` bigint NOT NULL DEFAULT 0 COMMENT '商品数量',
  `status` bigint NOT NULL DEFAULT 0 COMMENT '商品状态.0-暂存 1-在售 2-下架 3-删除',
  `price` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '商品价格',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_product_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES (2, '2023-09-26 11:23:51.666', '2023-10-01 13:26:51.380', '2023-10-01 20:18:17.034', '162917198703509504', '155831367513817088', 'test', '2r3r23r', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/e99d136c-1f08-1f52-f758-b9ffef2e15a6.jpg\"}]', '2343423', 14, 0, 4.00);
INSERT INTO `product` VALUES (3, '2023-09-26 11:28:50.090', '2023-10-01 13:26:45.099', '2023-10-01 20:18:17.716', '162918450384486400', '155831367513817088', 'test', '2r3r23r', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/093658ea-89a5-656f-366c-3d9601dc48b7.jpg\"}]', '2343423', 10, 0, 1.00);
INSERT INTO `product` VALUES (4, '2023-09-26 11:37:25.620', '2023-10-01 13:26:38.575', '2023-10-01 20:18:18.414', '162920612669833216', '155831367513817088', 'test', '2r3r23r', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/d4ce88cb-fb89-2444-0782-7be8edf442f5.jpg\"}]', '2343423', 10, 0, 1000.90);
INSERT INTO `product` VALUES (5, '2023-09-30 13:25:35.426', '2023-10-01 13:26:30.166', '2023-10-01 20:18:19.086', '164397384343109632', '155831367513817088', '华为mate60 pro ', '我是华为的一个手机', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/b87a01b9-e7c6-2ca2-27e3-35d188301065.jpg\"}]', '12人如果其二  97879', 5, 0, 1006.78);
INSERT INTO `product` VALUES (6, '2023-10-01 21:05:11.990', '2023-12-10 16:51:57.653', NULL, '164875436705595392', '155831367513817088', '小米 MIX FOLD 2', '11231', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/f5727884-bfc6-915e-c267-5d083e366787.png\"}]', '31231', 10, 1, 9997.00);
INSERT INTO `product` VALUES (7, '2023-10-01 21:06:19.211', '2024-01-02 16:45:01.547', NULL, '164875718655098880', '155831367513817088', '小米 MIX Alpna', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/d65e316f-b19d-88de-d857-e91899739545.jpg\"}]', '', 9, 2, 19999.00);
INSERT INTO `product` VALUES (8, '2023-10-01 21:08:42.735', '2023-11-02 10:41:28.858', NULL, '164876320638386176', '155831367513817088', '小米 MIX FOLD 3', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/20a171a9-1c02-bb31-fc61-d3dcb9a179be.jpg\"}]', '', 9, 1, 10999.00);
INSERT INTO `product` VALUES (9, '2023-10-01 21:10:40.691', '2023-10-04 17:14:26.000', NULL, '164876815385903104', '155831367513817088', 'XIAOMI Cicv 3', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/924b6bc8-da06-b5ab-0a93-2db89c91d17d.png\"}]', '', 10, 1, 2999.00);
INSERT INTO `product` VALUES (10, '2023-10-01 21:14:00.175', '2023-12-10 16:24:09.366', NULL, '164877652082442240', '155831367513817088', 'Redmi Buds5 ', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/faf3c8e4-9a69-29d0-0588-fab98fb8df0e.webp\"}]', '', 10, 1, 199.00);
INSERT INTO `product` VALUES (11, '2023-10-02 09:07:19.059', '2023-10-04 17:25:38.532', NULL, '165057163600216064', '155831367513817088', 'Xiao Mi Pad 6 Pro ', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/3e99fdcd-d3d4-e883-a5f1-cf68e3058263.png\"}]', '', 10, 1, 2999.00);
INSERT INTO `product` VALUES (12, '2023-10-02 09:34:15.370', '2023-10-04 17:14:35.582', NULL, '165063942908297216', '155831367513817088', '华为 mate 60 pro+', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/a87951cf-1845-ec1d-c2da-6b4c60751e82.jpg\"}]', '', 10, 1, 14999.00);
INSERT INTO `product` VALUES (13, '2023-10-02 09:35:28.431', '2023-10-04 17:14:34.557', NULL, '165064249352536064', '155831367513817088', '华为 mate 50 pro+ ', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/e8e59c24-7800-134e-6f44-d3f737e45adb.jpg\"}]', '', 10, 1, 6298.00);
INSERT INTO `product` VALUES (14, '2023-10-04 15:03:03.191', '2023-10-04 17:14:32.616', NULL, '165871463097057280', '155831367513817088', '三星 SAMSUNG Galaxy S23 Ultra ', '智能手写 | 性能强劲', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/974d733c-5ea0-b6c8-a2b4-8a2fd71438f7.avif\"}]', '', 10, 1, 7999.00);
INSERT INTO `product` VALUES (15, '2023-10-04 15:05:56.127', '2023-10-04 17:25:34.510', NULL, '165872188464185344', '155831367513817088', '新款GT3智能手表watch3', 'NFC | 智能', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/79db29da-8b42-870c-a026-d55c853bf973.avif\"}]', '', 1, 1, 165.00);
INSERT INTO `product` VALUES (16, '2023-10-04 15:08:06.632', '2023-10-04 15:08:06.632', NULL, '165872735841828864', '155831367513817088', '华强北GT3pro顶配', '性能强劲 | 智能门禁', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/8edaa123-990a-1c0a-665f-a1c0e8d36141.avif\"}]', '', 10, 1, 175.00);
INSERT INTO `product` VALUES (17, '2023-10-04 15:37:29.333', '2023-10-04 15:37:29.333', NULL, '165880129141489664', '155831367513817088', 'ReadMi Buds4 ', 'AI 智能降噪', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/6572873e-ff53-9632-1314-5e9a3c012b5f.webp\"}]', '', 10, 0, 179.00);
INSERT INTO `product` VALUES (18, '2023-10-04 17:27:35.680', '2023-10-04 17:27:41.867', NULL, '165907838173331456', '155831367513817088', 'oppo reno 10 ', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/9da5c129-2cfe-41c6-add9-0efe9958578d/2a513353-6b51-c000-a474-3c577c3435dd.avif\"}]', '', 10, 1, 2999.00);
INSERT INTO `product` VALUES (19, '2024-01-02 16:59:04.862', '2024-01-02 16:59:36.966', NULL, '198515570378162176', '198491259315568640', 'OPPO Watch 4 pro', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/16a7470d-4489-4306-b448-1bd87b6baf52/ea1d57a9-7114-21d5-0079-1957a44bc699.avif\"}]', '', 10, 1, 2199.00);
INSERT INTO `product` VALUES (20, '2024-01-02 17:04:52.153', '2024-01-02 17:04:53.962', NULL, '198517027030581248', '198491259315568640', '华强北智能手表 pro ', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/16a7470d-4489-4306-b448-1bd87b6baf52/da241d1d-b330-1462-c12e-c53dec249753.avif\"}]', '', 10, 1, 298.00);
INSERT INTO `product` VALUES (21, '2024-01-02 17:07:47.655', '2024-01-02 17:18:39.804', NULL, '198517763139321856', '198491259315568640', '果坊华强北S9Ultra智能运动手表', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/16a7470d-4489-4306-b448-1bd87b6baf52/a2e8bc98-e15c-ae34-9699-d6bd1e783f6e.avif\"}]', '', 9, 1, 368.00);
INSERT INTO `product` VALUES (22, '2024-01-02 17:15:03.632', '2024-01-02 17:15:04.867', NULL, '198519591759396864', '198491259315568640', '阿玛丁【顶配版GT3】华强北手表watch3', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/16a7470d-4489-4306-b448-1bd87b6baf52/34888708-cda4-f658-ab8a-dee0e678db1f.avif\"}]', '', 10, 1, 208.00);
INSERT INTO `product` VALUES (23, '2024-01-02 17:15:54.188', '2024-01-02 17:15:57.741', NULL, '198519803806629888', '198491259315568640', '华为机通用官网GT3PRO智能手表', '', '[{\"url\":\"http://zjxzero.oss-cn-hangzhou.aliyuncs.com/product/16a7470d-4489-4306-b448-1bd87b6baf52/90ea8b2d-4b50-4b8e-b0eb-f3b6fa5bce33.avif\"}]', '', 10, 1, 189.00);

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `parent_id` bigint UNSIGNED NOT NULL DEFAULT 0 COMMENT '父菜单 id',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '路由name',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '附加属性',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '路由path',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '组件位置对应的前端路径',
  `sorted` bigint NOT NULL DEFAULT 0 COMMENT '排序标记',
  `role` bigint NOT NULL DEFAULT 0 COMMENT '角色 1 普通用户 2 系统商家 3 系统管理员',
  `status` bigint NOT NULL DEFAULT 0 COMMENT '菜单状态',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_menu_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '2023-08-02 17:47:49.095', '2023-08-03 22:26:04.493', NULL, 0, 'Dashboard', '仪表盘', '/layout/dashboard', 'Odometer', 'dashboard/index.vue', 0, 3, 1);
INSERT INTO `sys_menu` VALUES (2, '2023-08-02 17:49:19.179', '2023-12-11 13:54:28.723', NULL, 0, 'System', '超级管理员', '/layout/system', 'User', 'system/index.vue', 10, 3, 1);
INSERT INTO `sys_menu` VALUES (3, '2023-08-02 17:50:04.675', '2023-08-03 19:55:40.457', NULL, 2, 'SystemMenu', '菜单管理', '/layout/system/menu', 'Memo', 'system/menu/index.vue', 0, 3, 1);
INSERT INTO `sys_menu` VALUES (4, '2023-08-06 22:19:15.951', '2023-08-30 17:17:16.067', NULL, 0, 'Dashboard', '仪表盘', '/layout/dashboard', 'data-board', 'dashboard/index.vue', 1, 2, 1);
INSERT INTO `sys_menu` VALUES (5, '2023-08-06 23:08:34.991', '2023-09-07 10:24:28.983', NULL, 0, 'Log', '操作日志', '/layout/log', 'notification', 'log/index.vue', 30, 3, 1);
INSERT INTO `sys_menu` VALUES (6, '2023-08-06 23:19:57.733', '2023-09-21 20:59:37.038', NULL, 0, 'About', '关于我', '/layout/about', 'paperclip', 'about/index.vue', 43, 3, 1);
INSERT INTO `sys_menu` VALUES (8, '2023-08-06 23:35:01.847', '2023-09-21 18:02:25.452', NULL, 0, 'Tool', '系统工具', '/layout/tool', 'timer', 'tool/index.vue', 15, 3, 1);
INSERT INTO `sys_menu` VALUES (9, '2023-08-10 19:23:49.892', '2023-09-07 10:53:13.702', NULL, 0, 'Business', '商户管理', '/layout/business', 'calendar', 'business/index.vue', 2, 2, 1);
INSERT INTO `sys_menu` VALUES (10, '2023-08-10 19:28:31.555', '2023-09-05 19:52:38.766', NULL, 9, 'BusinessInfo', '信息设置', '/layout/business/info', 'info-filled', 'business/info/index.vue', 1, 2, 1);
INSERT INTO `sys_menu` VALUES (11, '2023-08-10 19:37:01.090', '2023-09-25 12:24:54.599', NULL, 9, 'Businesscategory', '分类管理', '/layout/business/category', 'notification', 'business/category/index.vue', 6, 2, 0);
INSERT INTO `sys_menu` VALUES (12, '2023-08-10 19:39:00.475', '2023-09-25 21:24:30.735', NULL, 9, 'BusinessProduct', '商品管理', '/layout/business/product', 'pear', 'business/product/index.vue', 2, 2, 1);
INSERT INTO `sys_menu` VALUES (13, '2023-08-10 19:56:11.823', '2023-09-21 18:02:25.452', NULL, 8, 'ToolPackage', '自动化', '/layout/tool/package', 'bottom', 'tool/package/index.vue', 2, 3, 1);
INSERT INTO `sys_menu` VALUES (14, '2023-08-30 17:11:29.724', '2023-09-25 21:24:20.914', NULL, 9, 'Panic', '秒杀商品管理', '/layout/business/panic', 'data-line', 'business/panic/index.vue', 3, 2, 1);
INSERT INTO `sys_menu` VALUES (16, '2023-09-07 10:00:48.256', '2023-09-07 10:02:11.441', NULL, 2, 'User', '用户管理', '/layout/system/user', 'user-filled', 'system/user/index.vue', 0, 3, 1);
INSERT INTO `sys_menu` VALUES (17, '2023-09-07 10:07:07.017', '2023-12-11 13:54:28.723', NULL, 2, 'Role', '角色管理', '/layout/system/role', 'cherry', 'system/role/index.vue', 0, 3, 1);
INSERT INTO `sys_menu` VALUES (18, '2023-09-07 10:58:08.452', '2023-09-07 10:58:14.128', NULL, 2, 'SystemBusiness', '店铺管理', '/layout/system/business', 'collection', 'system/business/index.vue', 0, 3, 1);
INSERT INTO `sys_menu` VALUES (19, '2023-09-11 14:18:48.247', '2023-09-11 14:19:36.490', NULL, 0, 'State', '信息统计', '/layout/business/state', 'histogram', 'business/state/index.vue', 4, 2, 1);
INSERT INTO `sys_menu` VALUES (20, '2023-09-11 14:21:04.390', '2023-09-11 14:21:04.390', NULL, 0, 'Order', '订单管理', '/layout/business/order', 'coffee-cup', 'business/order/index.vue', 3, 2, 1);
INSERT INTO `sys_menu` VALUES (21, '2023-09-18 09:19:59.668', '2023-09-18 09:19:59.668', NULL, 2, 'Category', '商品分类管理', '/layout/category', 'coffee', 'system/category/index.vue', 0, 3, 1);
INSERT INTO `sys_menu` VALUES (22, '2023-09-21 20:58:18.031', '2023-09-21 20:58:25.395', '2023-09-22 09:36:11.699', 0, '', '11111', '', 'aim', '', 0, 3, 0);
INSERT INTO `sys_menu` VALUES (23, '2023-09-21 21:05:50.317', '2023-09-21 21:06:47.272', '2023-09-22 09:36:11.699', 22, '1241424', '12412', '14214241', 'add-location', '142412', 8, 3, 0);
INSERT INTO `sys_menu` VALUES (24, '2023-09-25 12:10:40.417', '2023-09-25 12:10:40.417', '2023-09-25 12:10:55.470', 0, '', '1231412', '', 'aim', '', 0, 0, 0);
INSERT INTO `sys_menu` VALUES (25, '2023-12-11 13:55:46.647', '2024-01-02 16:10:59.853', NULL, 0, 'SystemOrder', '订单管理', '/layout/system/order', 'dessert', 'system/order/index.vue', 11, 3, 1);
INSERT INTO `sys_menu` VALUES (26, '2024-01-02 16:12:11.664', '2024-01-02 16:12:11.664', NULL, 0, 'SystemProduct', '商品管理', '/layout/product', 'aim', 'system/product/index.vue', 13, 3, 1);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `status` bigint NOT NULL DEFAULT 0 COMMENT '角色状态',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '2023-09-04 20:03:17.556', '2023-09-04 20:03:17.556', NULL, '普通用户', 0);
INSERT INTO `sys_role` VALUES (2, '2023-09-04 20:03:43.875', '2023-09-04 20:03:43.875', NULL, '系统商家', 0);
INSERT INTO `sys_role` VALUES (3, '2023-09-04 20:03:58.621', '2023-09-04 20:03:58.621', NULL, '系统管理员', 0);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户uuid',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录名',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `nickname` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户显示昵称',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录密码',
  `avatar` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `role` bigint NOT NULL DEFAULT 1 COMMENT '用户类型 1 普通用户 2 商家 3 系统管理员',
  `status` bigint NOT NULL DEFAULT 1 COMMENT '用户状态 1 开启 0 禁用 ',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号码',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2023-09-06 20:01:28.000', '2023-12-21 15:22:54.519', NULL, '7b369a36-31f8-4242-88f8-e025d390dfb9', 'admin', '202606540@qq.com', ' 笑脸带眼红', '$2a$10$PUMkbq.s8Z.tjSDomo92qu71WJVTW/LiC4eH6tOwO2OE9iZteQgpu', 'https://zpshop.oss-cn-hangzhou.aliyuncs.com/2608045f-a589-45a8-b8c8-a8201021a2c2.jpg', 3, 1, '15593336497');
INSERT INTO `user` VALUES (2, '2023-09-06 22:07:17.843', '2023-09-25 10:25:32.259', NULL, '9da5c129-2cfe-41c6-add9-0efe9958578d', 'zhengjunxing', '202606540@qq.com', '脑子挖掉了专卖店', '$2a$10$PUMkbq.s8Z.tjSDomo92qu71WJVTW/LiC4eH6tOwO2OE9iZteQgpu', 'https://zpshop.oss-cn-hangzhou.aliyuncs.com/2608045f-a589-45a8-b8c8-a8201021a2c2.jpg', 2, 1, '15593336497');
INSERT INTO `user` VALUES (3, '2023-09-23 22:06:08.223', '2023-09-23 22:41:38.335', NULL, '3bc22666-494f-4448-8785-8c6cec889911', 'zjx', '202606540@qq.com', 'xiaolian', '$2a$10$EfkU4idgsfr22d8aqeCjfOXXpVt9l6ga5WfaFrIRtTImTbmjvwza2', '', 1, 1, '15593336410');
INSERT INTO `user` VALUES (4, '2023-09-23 22:12:21.021', '2023-12-11 14:10:42.727', NULL, 'bcf254b4-a9fa-4b57-802c-80a5472d9a8f', 'zjx1', '202606540@qq.com', 'xiaolian', '$2a$10$HJ7UkB5w7dfN9EsMc9b6f.i6MaeqCFZxBzg4..PS7QvJqzVSOb22m', '', 1, 1, '');
INSERT INTO `user` VALUES (5, '2023-09-23 22:12:23.655', '2023-09-23 22:12:23.655', '2023-09-23 22:40:12.910', 'aacebad0-a93d-4464-9ab8-10c366b492c5', 'zjx21', '202606540@qq.com', 'xiaolian', '$2a$10$KatCT3tdSFCEDiozaNdsPuSy6oK/MqFhgwBF/xgSCBjXOCavTDLYW', '', 1, 1, '');
INSERT INTO `user` VALUES (6, '2023-09-23 22:12:28.046', '2023-09-23 22:12:28.046', '2023-09-23 22:40:17.086', '12d5dbf5-f544-4f0d-98e7-d10aae5f7d96', 'zjx221', '202606540@qq.com', 'xiaolian', '$2a$10$CrsbBXsLbyWFP4Ox28igBeboa67iE1mpM32b.T12ap/2TRk8KmQf.', '', 1, 1, '');
INSERT INTO `user` VALUES (7, '2023-09-23 22:12:30.158', '2023-09-23 22:12:30.158', '2023-09-23 22:40:20.514', '84e0319f-c604-4609-967d-a1d5bd2facc6', 'zjx3221', '202606540@qq.com', 'xiaolian', '$2a$10$NMudAWiotI97j9blR3nSi.uAZVaZT86A9kEmtZubOnG3eO7/ycpfm', '', 1, 1, '');
INSERT INTO `user` VALUES (8, '2023-09-23 22:12:31.815', '2023-09-23 22:41:38.895', '2023-09-25 12:10:18.045', 'fd7388c4-6446-4a45-8769-16e4f6d382fc', 'zj4x3221', '202606540@qq.com', 'xiaolian', '$2a$10$19FFZmyQEXs9C9hn5jr9auJfBNYYEOIEG1pR3wjMyhkJA0Qgd.lA2', '', 1, 1, '');
INSERT INTO `user` VALUES (9, '2023-09-23 22:12:41.915', '2023-09-25 18:54:19.652', NULL, 'fa5fea4b-b352-4593-a895-1b5e90b7be0b', 'zj4x2323221', '202606540@qq.com', 'xiaolian', '$2a$10$eW6O7GvZAGYWe52QpnWQ5.rraAhbfyYbweMpTrFKqBo8Rdny5wuCG', '', 1, 0, '');
INSERT INTO `user` VALUES (10, '2023-09-23 22:12:43.687', '2023-09-23 22:12:43.687', '2023-09-23 22:35:43.606', '2f390ff4-ea9a-4e0e-95be-25db8b16835a', 'zj4x24323221', '202606540@qq.com', 'xiaolian', '$2a$10$B1NGbELDJJ2AoxUNQcs9aeUBF871TwbARDuQITJEcmHRbmI5FXg8e', '', 1, 1, '');
INSERT INTO `user` VALUES (11, '2023-09-23 22:12:47.788', '2023-09-25 18:54:20.963', NULL, 'f0275324-2628-454d-9540-20f33f100449', 'zj4x2223221', '202606540@qq.com', 'xiaolian', '$2a$10$x4U2PbwOBZKEhw7ajtpaZO//VpG15qvKD9YLBhDSxOyOs4y9.yQ7C', '', 1, 0, '');
INSERT INTO `user` VALUES (12, '2023-09-23 22:12:49.488', '2023-12-11 14:10:44.268', NULL, 'de1653ea-eac0-40d5-aafc-250ceb53dea6', 'zj4x223221', '202606540@qq.com', 'xiaolian', '$2a$10$VLhnvwl1tlHBESoBGKgBbOLAOF2pOinQrq3LBkDYxZgNUQD7hkXIW', '', 1, 1, '');
INSERT INTO `user` VALUES (13, '2023-09-23 22:12:51.107', '2023-09-25 18:54:18.428', NULL, '4d8ac4e5-5fa8-41d0-9f7b-929b08e5488d', 'zj4x22321', '202606540@qq.com', 'xiaolian', '$2a$10$KxlCmVe6GsMy7zF97JScHONUbHBSVYagBrFbnYnFTBpn6ysiOrpJC', '', 1, 0, '');
INSERT INTO `user` VALUES (14, '2023-09-23 22:12:52.949', '2023-09-25 18:54:21.695', NULL, 'f887ce2c-fc7d-4878-8a46-b45470bbfeba', 'zj4x2231', '202606540@qq.com', 'xiaolian', '$2a$10$Wax.OEKfiqA3rEjSXGf4aur6Uo5LxnbGraYq7XyDpJ8qQAl5oCYsC', '', 1, 0, '');
INSERT INTO `user` VALUES (15, '2023-09-23 22:12:54.947', '2023-12-11 14:10:45.463', NULL, '0bf493da-791b-4274-a410-0dda54fc3cc7', 'z4x2231', '202606540@qq.com', 'xiaolian', '$2a$10$4vR3DYT02hPeTsYIZYrmV.Zgr06/P3Nm1jRKRJ97lHQfo.MXLibXS', '', 1, 1, '');
INSERT INTO `user` VALUES (16, '2023-10-09 15:12:11.550', '2023-10-09 15:12:11.550', NULL, 'ba479a90-f156-4c2e-a219-f0635bb49f74', 'test', '', '测试用户', '$2a$10$PUMkbq.s8Z.tjSDomo92qu71WJVTW/LiC4eH6tOwO2OE9iZteQgpu', '', 1, 1, '15593336499');
INSERT INTO `user` VALUES (20, '2024-01-02 15:22:28.650', '2024-01-02 15:22:28.650', NULL, '16a7470d-4489-4306-b448-1bd87b6baf52', 'zjx', '202606540@qq.com', 'xiaolian', '$2a$10$EfkU4idgsfr22d8aqeCjfOXXpVt9l6ga5WfaFrIRtTImTbmjvwza2', '', 2, 1, '15593336410');

SET FOREIGN_KEY_CHECKS = 1;
