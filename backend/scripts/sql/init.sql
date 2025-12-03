-- OceanEngine 后台管理系统数据库初始化脚本
-- 字符集: utf8mb4
-- 排序规则: utf8mb4_unicode_ci

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- 系统用户表
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(64) NOT NULL COMMENT '用户名',
  `password` varchar(128) NOT NULL COMMENT '密码',
  `nickname` varchar(64) DEFAULT NULL COMMENT '昵称',
  `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机号',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像URL',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 0-禁用, 1-启用',
  `role_id` bigint unsigned DEFAULT NULL COMMENT '角色ID',
  `last_login_at` datetime DEFAULT NULL COMMENT '最后登录时间',
  `last_login_ip` varchar(45) DEFAULT NULL COMMENT '最后登录IP',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_role_id` (`role_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统用户表';

-- ----------------------------
-- 系统角色表
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` varchar(64) NOT NULL COMMENT '角色名称',
  `code` varchar(64) NOT NULL COMMENT '角色编码',
  `description` varchar(255) DEFAULT NULL COMMENT '角色描述',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态: 0-禁用, 1-启用',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统角色表';

-- ----------------------------
-- 系统菜单表
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父级ID',
  `name` varchar(64) NOT NULL COMMENT '菜单名称',
  `path` varchar(255) DEFAULT NULL COMMENT '路由路径',
  `component` varchar(255) DEFAULT NULL COMMENT '组件路径',
  `icon` varchar(64) DEFAULT NULL COMMENT '图标',
  `type` tinyint NOT NULL DEFAULT '1' COMMENT '类型: 1-目录, 2-菜单, 3-按钮',
  `permission` varchar(128) DEFAULT NULL COMMENT '权限标识',
  `visible` tinyint NOT NULL DEFAULT '1' COMMENT '是否可见: 0-隐藏, 1-显示',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统菜单表';

-- ----------------------------
-- 角色菜单关联表
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `menu_id` bigint unsigned NOT NULL COMMENT '菜单ID',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_role_menu` (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色菜单关联表';

-- ----------------------------
-- 操作日志表
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_log`;
CREATE TABLE `sys_operation_log` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint unsigned DEFAULT NULL COMMENT '用户ID',
  `username` varchar(64) DEFAULT NULL COMMENT '用户名',
  `module` varchar(64) DEFAULT NULL COMMENT '模块',
  `action` varchar(64) DEFAULT NULL COMMENT '操作',
  `method` varchar(16) DEFAULT NULL COMMENT '请求方法',
  `path` varchar(255) DEFAULT NULL COMMENT '请求路径',
  `ip` varchar(45) DEFAULT NULL COMMENT 'IP地址',
  `user_agent` varchar(512) DEFAULT NULL COMMENT 'User-Agent',
  `request_body` text COMMENT '请求体',
  `response_body` text COMMENT '响应体',
  `status_code` int DEFAULT NULL COMMENT '状态码',
  `latency` bigint DEFAULT NULL COMMENT '耗时(毫秒)',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志表';

-- ----------------------------
-- 广告主表
-- ----------------------------
DROP TABLE IF EXISTS `ad_advertiser`;
CREATE TABLE `ad_advertiser` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `advertiser_id` bigint unsigned NOT NULL COMMENT '巨量广告主ID',
  `name` varchar(128) NOT NULL COMMENT '广告主名称',
  `company` varchar(255) DEFAULT NULL COMMENT '公司名称',
  `role` varchar(32) DEFAULT NULL COMMENT '角色',
  `status` varchar(32) DEFAULT NULL COMMENT '状态',
  `address` varchar(255) DEFAULT NULL COMMENT '地址',
  `license_url` varchar(512) DEFAULT NULL COMMENT '执照URL',
  `license_no` varchar(128) DEFAULT NULL COMMENT '执照号',
  `access_token` varchar(512) DEFAULT NULL COMMENT 'Access Token',
  `refresh_token` varchar(512) DEFAULT NULL COMMENT 'Refresh Token',
  `token_expires_at` datetime DEFAULT NULL COMMENT 'Token过期时间',
  `balance` decimal(18,2) DEFAULT '0.00' COMMENT '账户余额',
  `total_cost` decimal(18,2) DEFAULT '0.00' COMMENT '总消耗',
  `sync_status` tinyint NOT NULL DEFAULT '0' COMMENT '同步状态: 0-未同步, 1-同步中, 2-已同步, 3-同步失败',
  `last_sync_at` datetime DEFAULT NULL COMMENT '最后同步时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_advertiser_id` (`advertiser_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告主表';

-- ----------------------------
-- 广告主资金流水表
-- ----------------------------
DROP TABLE IF EXISTS `ad_advertiser_fund`;
CREATE TABLE `ad_advertiser_fund` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `advertiser_id` bigint unsigned NOT NULL COMMENT '广告主ID',
  `transaction_type` varchar(32) DEFAULT NULL COMMENT '交易类型',
  `amount` decimal(18,2) DEFAULT NULL COMMENT '金额',
  `balance_after` decimal(18,2) DEFAULT NULL COMMENT '交易后余额',
  `transaction_time` datetime DEFAULT NULL COMMENT '交易时间',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_advertiser_id` (`advertiser_id`),
  KEY `idx_transaction_time` (`transaction_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告主资金流水表';

-- ----------------------------
-- 广告系列表
-- ----------------------------
DROP TABLE IF EXISTS `ad_campaign`;
CREATE TABLE `ad_campaign` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `campaign_id` bigint unsigned NOT NULL COMMENT '巨量广告系列ID',
  `advertiser_id` bigint unsigned NOT NULL COMMENT '广告主ID',
  `name` varchar(128) NOT NULL COMMENT '广告系列名称',
  `budget` decimal(18,2) DEFAULT NULL COMMENT '预算',
  `budget_mode` varchar(32) DEFAULT NULL COMMENT '预算类型',
  `landing_type` varchar(32) DEFAULT NULL COMMENT '推广目标',
  `status` varchar(32) DEFAULT NULL COMMENT '状态',
  `opt_status` varchar(32) DEFAULT NULL COMMENT '操作状态',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_campaign_id` (`campaign_id`),
  KEY `idx_advertiser_id` (`advertiser_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告系列表';

-- ----------------------------
-- 广告组表
-- ----------------------------
DROP TABLE IF EXISTS `ad_ad`;
CREATE TABLE `ad_ad` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `ad_id` bigint unsigned NOT NULL COMMENT '巨量广告ID',
  `advertiser_id` bigint unsigned NOT NULL COMMENT '广告主ID',
  `campaign_id` bigint unsigned NOT NULL COMMENT '广告系列ID',
  `name` varchar(128) NOT NULL COMMENT '广告名称',
  `budget` decimal(18,2) DEFAULT NULL COMMENT '预算',
  `budget_mode` varchar(32) DEFAULT NULL COMMENT '预算类型',
  `bid` decimal(18,4) DEFAULT NULL COMMENT '出价',
  `pricing` varchar(32) DEFAULT NULL COMMENT '计费方式',
  `status` varchar(32) DEFAULT NULL COMMENT '状态',
  `opt_status` varchar(32) DEFAULT NULL COMMENT '操作状态',
  `delivery_range` varchar(32) DEFAULT NULL COMMENT '投放范围',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_ad_id` (`ad_id`),
  KEY `idx_advertiser_id` (`advertiser_id`),
  KEY `idx_campaign_id` (`campaign_id`),
  KEY `idx_status` (`status`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告组表';

-- ----------------------------
-- 创意表
-- ----------------------------
DROP TABLE IF EXISTS `ad_creative`;
CREATE TABLE `ad_creative` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `creative_id` bigint unsigned NOT NULL COMMENT '巨量创意ID',
  `advertiser_id` bigint unsigned NOT NULL COMMENT '广告主ID',
  `ad_id` bigint unsigned NOT NULL COMMENT '广告ID',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `image_ids` json DEFAULT NULL COMMENT '图片ID列表',
  `video_id` varchar(128) DEFAULT NULL COMMENT '视频ID',
  `status` varchar(32) DEFAULT NULL COMMENT '状态',
  `opt_status` varchar(32) DEFAULT NULL COMMENT '操作状态',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_creative_id` (`creative_id`),
  KEY `idx_advertiser_id` (`advertiser_id`),
  KEY `idx_ad_id` (`ad_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='创意表';

-- ----------------------------
-- 广告主报告表
-- ----------------------------
DROP TABLE IF EXISTS `rpt_advertiser`;
CREATE TABLE `rpt_advertiser` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `advertiser_id` bigint unsigned NOT NULL COMMENT '广告主ID',
  `stat_date` date NOT NULL COMMENT '统计日期',
  `cost` decimal(18,2) DEFAULT '0.00' COMMENT '消耗',
  `show_cnt` bigint DEFAULT '0' COMMENT '展示数',
  `click_cnt` bigint DEFAULT '0' COMMENT '点击数',
  `convert_cnt` bigint DEFAULT '0' COMMENT '转化数',
  `ctr` decimal(10,4) DEFAULT '0.0000' COMMENT '点击率',
  `cpc` decimal(10,4) DEFAULT '0.0000' COMMENT '平均点击单价',
  `cpm` decimal(10,4) DEFAULT '0.0000' COMMENT '千次展示费用',
  `convert_rate` decimal(10,4) DEFAULT '0.0000' COMMENT '转化率',
  `convert_cost` decimal(10,4) DEFAULT '0.0000' COMMENT '转化成本',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_advertiser_date` (`advertiser_id`,`stat_date`),
  KEY `idx_stat_date` (`stat_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告主报告表';

-- ----------------------------
-- 广告系列报告表
-- ----------------------------
DROP TABLE IF EXISTS `rpt_campaign`;
CREATE TABLE `rpt_campaign` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `campaign_id` bigint unsigned NOT NULL COMMENT '广告系列ID',
  `advertiser_id` bigint unsigned NOT NULL COMMENT '广告主ID',
  `stat_date` date NOT NULL COMMENT '统计日期',
  `cost` decimal(18,2) DEFAULT '0.00' COMMENT '消耗',
  `show_cnt` bigint DEFAULT '0' COMMENT '展示数',
  `click_cnt` bigint DEFAULT '0' COMMENT '点击数',
  `convert_cnt` bigint DEFAULT '0' COMMENT '转化数',
  `ctr` decimal(10,4) DEFAULT '0.0000' COMMENT '点击率',
  `cpc` decimal(10,4) DEFAULT '0.0000' COMMENT '平均点击单价',
  `cpm` decimal(10,4) DEFAULT '0.0000' COMMENT '千次展示费用',
  `convert_rate` decimal(10,4) DEFAULT '0.0000' COMMENT '转化率',
  `convert_cost` decimal(10,4) DEFAULT '0.0000' COMMENT '转化成本',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_campaign_date` (`campaign_id`,`stat_date`),
  KEY `idx_advertiser_id` (`advertiser_id`),
  KEY `idx_stat_date` (`stat_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告系列报告表';

-- ----------------------------
-- 广告报告表
-- ----------------------------
DROP TABLE IF EXISTS `rpt_ad`;
CREATE TABLE `rpt_ad` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `ad_id` bigint unsigned NOT NULL COMMENT '广告ID',
  `campaign_id` bigint unsigned NOT NULL COMMENT '广告系列ID',
  `advertiser_id` bigint unsigned NOT NULL COMMENT '广告主ID',
  `stat_date` date NOT NULL COMMENT '统计日期',
  `cost` decimal(18,2) DEFAULT '0.00' COMMENT '消耗',
  `show_cnt` bigint DEFAULT '0' COMMENT '展示数',
  `click_cnt` bigint DEFAULT '0' COMMENT '点击数',
  `convert_cnt` bigint DEFAULT '0' COMMENT '转化数',
  `ctr` decimal(10,4) DEFAULT '0.0000' COMMENT '点击率',
  `cpc` decimal(10,4) DEFAULT '0.0000' COMMENT '平均点击单价',
  `cpm` decimal(10,4) DEFAULT '0.0000' COMMENT '千次展示费用',
  `convert_rate` decimal(10,4) DEFAULT '0.0000' COMMENT '转化率',
  `convert_cost` decimal(10,4) DEFAULT '0.0000' COMMENT '转化成本',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_ad_date` (`ad_id`,`stat_date`),
  KEY `idx_campaign_id` (`campaign_id`),
  KEY `idx_advertiser_id` (`advertiser_id`),
  KEY `idx_stat_date` (`stat_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告报告表';

-- ----------------------------
-- 报告导出任务表
-- ----------------------------
DROP TABLE IF EXISTS `rpt_export_task`;
CREATE TABLE `rpt_export_task` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `type` varchar(32) NOT NULL COMMENT '报告类型',
  `params` json DEFAULT NULL COMMENT '查询参数',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态: 0-待处理, 1-处理中, 2-已完成, 3-失败',
  `file_url` varchar(512) DEFAULT NULL COMMENT '文件URL',
  `error_msg` varchar(512) DEFAULT NULL COMMENT '错误信息',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='报告导出任务表';

-- ----------------------------
-- 初始化数据
-- ----------------------------

-- 初始化超级管理员角色
INSERT INTO `sys_role` (`id`, `name`, `code`, `description`, `status`, `sort`) VALUES
(1, '超级管理员', 'admin', '系统超级管理员，拥有所有权限', 1, 0);

-- 初始化管理员用户 (密码: admin123)
-- 密码使用 bcrypt 加密
INSERT INTO `sys_user` (`id`, `username`, `password`, `nickname`, `status`, `role_id`) VALUES
(1, 'admin', '$2a$10$q0hiBoTiIKFdQhSlIbE4luHBRFV2n17GMDyN1BJnaTMIto.YEKAX6', '管理员', 1, 1);

-- 初始化菜单
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `path`, `component`, `icon`, `type`, `permission`, `visible`, `sort`) VALUES
(1, 0, '系统管理', '/system', NULL, 'setting', 1, NULL, 1, 100),
(2, 1, '用户管理', '/system/users', 'system/user/index', 'user', 2, 'system:user:list', 1, 1),
(3, 1, '角色管理', '/system/roles', 'system/role/index', 'peoples', 2, 'system:role:list', 1, 2),
(4, 1, '菜单管理', '/system/menus', 'system/menu/index', 'tree-table', 2, 'system:menu:list', 1, 3),
(5, 0, '广告管理', '/ad', NULL, 'ad', 1, NULL, 1, 10),
(6, 5, '广告主管理', '/ad/advertisers', 'ad/advertiser/index', 'company', 2, 'ad:advertiser:list', 1, 1),
(7, 5, '广告系列', '/ad/campaigns', 'ad/campaign/index', 'list', 2, 'ad:campaign:list', 1, 2),
(8, 5, '广告组', '/ad/ads', 'ad/ad/index', 'component', 2, 'ad:ad:list', 1, 3),
(9, 5, '创意管理', '/ad/creatives', 'ad/creative/index', 'star', 2, 'ad:creative:list', 1, 4),
(10, 0, '数据报表', '/report', NULL, 'chart', 1, NULL, 1, 20),
(11, 10, '广告主报表', '/report/advertiser', 'report/advertiser/index', 'chart', 2, 'report:advertiser:list', 1, 1),
(12, 10, '广告系列报表', '/report/campaign', 'report/campaign/index', 'chart', 2, 'report:campaign:list', 1, 2),
(13, 10, '广告报表', '/report/ad', 'report/ad/index', 'chart', 2, 'report:ad:list', 1, 3),
(14, 10, '实时数据', '/report/realtime', 'report/realtime/index', 'monitor', 2, 'report:realtime:view', 1, 4);

-- 初始化角色菜单关联 (超级管理员拥有所有菜单权限)
INSERT INTO `sys_role_menu` (`role_id`, `menu_id`) VALUES
(1, 1), (1, 2), (1, 3), (1, 4), (1, 5), (1, 6), (1, 7),
(1, 8), (1, 9), (1, 10), (1, 11), (1, 12), (1, 13), (1, 14);

SET FOREIGN_KEY_CHECKS = 1;
