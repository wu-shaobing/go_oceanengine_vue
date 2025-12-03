# 数据库设计

## 数据库规范

### 命名规范
- 表名：小写，下划线分隔，前缀 `sys_`（系统表）、`ad_`（广告表）、`rpt_`（报告表）
- 字段名：小写，下划线分隔
- 索引名：`idx_表名_字段名`
- 唯一索引：`uk_表名_字段名`

### 通用字段
每个表都应包含以下字段：
- `id`: 主键，自增
- `created_at`: 创建时间
- `updated_at`: 更新时间
- `deleted_at`: 软删除时间（可选）
- `created_by`: 创建人ID（可选）
- `updated_by`: 更新人ID（可选）

## 表结构设计

### 1. 系统管理模块

#### sys_user - 系统用户表
```sql
CREATE TABLE `sys_user` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `username` VARCHAR(64) NOT NULL COMMENT '用户名',
    `password` VARCHAR(128) NOT NULL COMMENT '密码（加密）',
    `nickname` VARCHAR(128) DEFAULT '' COMMENT '昵称',
    `phone` VARCHAR(20) DEFAULT '' COMMENT '手机号',
    `email` VARCHAR(128) DEFAULT '' COMMENT '邮箱',
    `avatar` VARCHAR(255) DEFAULT '' COMMENT '头像URL',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：0-禁用，1-启用',
    `role_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '角色ID',
    `dept_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '部门ID',
    `last_login_at` DATETIME DEFAULT NULL COMMENT '最后登录时间',
    `last_login_ip` VARCHAR(50) DEFAULT '' COMMENT '最后登录IP',
    `remark` VARCHAR(500) DEFAULT '' COMMENT '备注',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
    `created_by` BIGINT UNSIGNED DEFAULT 0 COMMENT '创建人',
    `updated_by` BIGINT UNSIGNED DEFAULT 0 COMMENT '更新人',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_username` (`username`),
    KEY `idx_phone` (`phone`),
    KEY `idx_email` (`email`),
    KEY `idx_role_id` (`role_id`),
    KEY `idx_status` (`status`),
    KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统用户表';
```

#### sys_role - 角色表
```sql
CREATE TABLE `sys_role` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `name` VARCHAR(64) NOT NULL COMMENT '角色名称',
    `key` VARCHAR(64) NOT NULL COMMENT '角色标识',
    `sort` INT NOT NULL DEFAULT 0 COMMENT '排序',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：0-禁用，1-启用',
    `data_scope` TINYINT NOT NULL DEFAULT 1 COMMENT '数据范围：1-全部，2-自定义，3-本部门，4-本部门及以下，5-仅本人',
    `remark` VARCHAR(500) DEFAULT '' COMMENT '备注',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    `created_by` BIGINT UNSIGNED DEFAULT 0,
    `updated_by` BIGINT UNSIGNED DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_key` (`key`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';
```

#### sys_menu - 菜单表
```sql
CREATE TABLE `sys_menu` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `parent_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父菜单ID',
    `name` VARCHAR(64) NOT NULL COMMENT '菜单名称',
    `path` VARCHAR(255) DEFAULT '' COMMENT '路由路径',
    `component` VARCHAR(255) DEFAULT '' COMMENT '组件路径',
    `icon` VARCHAR(64) DEFAULT '' COMMENT '图标',
    `sort` INT NOT NULL DEFAULT 0 COMMENT '排序',
    `permission` VARCHAR(128) DEFAULT '' COMMENT '权限标识',
    `type` TINYINT NOT NULL DEFAULT 1 COMMENT '类型：1-目录，2-菜单，3-按钮',
    `visible` TINYINT NOT NULL DEFAULT 1 COMMENT '是否可见：0-隐藏，1-显示',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：0-禁用，1-启用',
    `is_frame` TINYINT NOT NULL DEFAULT 0 COMMENT '是否外链：0-否，1-是',
    `is_cache` TINYINT NOT NULL DEFAULT 0 COMMENT '是否缓存：0-否，1-是',
    `remark` VARCHAR(500) DEFAULT '' COMMENT '备注',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_parent_id` (`parent_id`),
    KEY `idx_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';
```

#### sys_role_menu - 角色菜单关联表
```sql
CREATE TABLE `sys_role_menu` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `role_id` BIGINT UNSIGNED NOT NULL COMMENT '角色ID',
    `menu_id` BIGINT UNSIGNED NOT NULL COMMENT '菜单ID',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_role_menu` (`role_id`, `menu_id`),
    KEY `idx_role_id` (`role_id`),
    KEY `idx_menu_id` (`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单关联表';
```

#### sys_dict_type - 字典类型表
```sql
CREATE TABLE `sys_dict_type` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL COMMENT '字典名称',
    `type` VARCHAR(64) NOT NULL COMMENT '字典类型',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态',
    `remark` VARCHAR(500) DEFAULT '',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典类型表';
```

#### sys_dict_data - 字典数据表
```sql
CREATE TABLE `sys_dict_data` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `dict_type` VARCHAR(64) NOT NULL COMMENT '字典类型',
    `label` VARCHAR(128) NOT NULL COMMENT '字典标签',
    `value` VARCHAR(128) NOT NULL COMMENT '字典值',
    `sort` INT NOT NULL DEFAULT 0 COMMENT '排序',
    `status` TINYINT NOT NULL DEFAULT 1,
    `css_class` VARCHAR(128) DEFAULT '' COMMENT 'CSS类名',
    `list_class` VARCHAR(128) DEFAULT '' COMMENT '列表类名',
    `is_default` TINYINT NOT NULL DEFAULT 0 COMMENT '是否默认',
    `remark` VARCHAR(500) DEFAULT '',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_dict_type` (`dict_type`),
    KEY `idx_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典数据表';
```

#### sys_operation_log - 操作日志表
```sql
CREATE TABLE `sys_operation_log` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作用户ID',
    `username` VARCHAR(64) DEFAULT '' COMMENT '操作用户名',
    `module` VARCHAR(64) DEFAULT '' COMMENT '模块名称',
    `action` VARCHAR(64) DEFAULT '' COMMENT '操作类型',
    `method` VARCHAR(16) DEFAULT '' COMMENT '请求方法',
    `path` VARCHAR(255) DEFAULT '' COMMENT '请求路径',
    `query` TEXT COMMENT '请求参数',
    `body` TEXT COMMENT '请求体',
    `response` TEXT COMMENT '响应内容',
    `ip` VARCHAR(50) DEFAULT '' COMMENT '操作IP',
    `user_agent` VARCHAR(500) DEFAULT '' COMMENT 'User-Agent',
    `status` INT DEFAULT 0 COMMENT '状态码',
    `latency` INT DEFAULT 0 COMMENT '响应时间（毫秒）',
    `error_msg` TEXT COMMENT '错误信息',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_module` (`module`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='操作日志表';
```

### 2. 广告主管理模块

#### ad_advertiser - 广告主表
```sql
CREATE TABLE `ad_advertiser` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT 'Ocean Engine 广告主ID',
    `name` VARCHAR(255) NOT NULL COMMENT '广告主名称',
    `company` VARCHAR(255) DEFAULT '' COMMENT '公司名称',
    `status` VARCHAR(50) DEFAULT '' COMMENT '状态',
    `role` VARCHAR(50) DEFAULT '' COMMENT '角色类型',
    `balance` DECIMAL(15,2) DEFAULT 0.00 COMMENT '账户余额',
    `valid_balance` DECIMAL(15,2) DEFAULT 0.00 COMMENT '可用余额',
    `cash_balance` DECIMAL(15,2) DEFAULT 0.00 COMMENT '现金余额',
    `industry` VARCHAR(100) DEFAULT '' COMMENT '行业',
    `license_url` VARCHAR(500) DEFAULT '' COMMENT '营业执照URL',
    `license_no` VARCHAR(100) DEFAULT '' COMMENT '营业执照号',
    `contact_name` VARCHAR(64) DEFAULT '' COMMENT '联系人',
    `contact_phone` VARCHAR(20) DEFAULT '' COMMENT '联系电话',
    `contact_email` VARCHAR(128) DEFAULT '' COMMENT '联系邮箱',
    `address` VARCHAR(500) DEFAULT '' COMMENT '联系地址',
    `access_token` VARCHAR(500) DEFAULT '' COMMENT '访问令牌',
    `refresh_token` VARCHAR(500) DEFAULT '' COMMENT '刷新令牌',
    `token_expire_at` DATETIME DEFAULT NULL COMMENT 'Token过期时间',
    `last_sync_at` DATETIME DEFAULT NULL COMMENT '最后同步时间',
    `remark` VARCHAR(500) DEFAULT '' COMMENT '备注',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    `created_by` BIGINT UNSIGNED DEFAULT 0,
    `updated_by` BIGINT UNSIGNED DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_advertiser_id` (`advertiser_id`),
    KEY `idx_name` (`name`),
    KEY `idx_status` (`status`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告主表';
```

#### ad_advertiser_fund - 资金流水表
```sql
CREATE TABLE `ad_advertiser_fund` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT '广告主ID',
    `transaction_type` VARCHAR(50) NOT NULL COMMENT '交易类型：recharge-充值，consume-消费，refund-退款',
    `amount` DECIMAL(15,2) NOT NULL COMMENT '交易金额',
    `balance_before` DECIMAL(15,2) DEFAULT 0.00 COMMENT '交易前余额',
    `balance_after` DECIMAL(15,2) DEFAULT 0.00 COMMENT '交易后余额',
    `transaction_seq` VARCHAR(100) DEFAULT '' COMMENT '交易流水号',
    `transaction_time` DATETIME DEFAULT NULL COMMENT '交易时间',
    `remark` VARCHAR(500) DEFAULT '',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_advertiser_id` (`advertiser_id`),
    KEY `idx_transaction_type` (`transaction_type`),
    KEY `idx_transaction_time` (`transaction_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资金流水表';
```

### 3. 广告投放模块

#### ad_campaign - 广告系列表
```sql
CREATE TABLE `ad_campaign` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `campaign_id` BIGINT UNSIGNED NOT NULL COMMENT 'Ocean Engine 广告系列ID',
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT '广告主ID',
    `name` VARCHAR(255) NOT NULL COMMENT '系列名称',
    `budget_mode` VARCHAR(50) DEFAULT '' COMMENT '预算类型：BUDGET_MODE_DAY-日预算，BUDGET_MODE_TOTAL-总预算',
    `budget` DECIMAL(15,2) DEFAULT 0.00 COMMENT '预算金额',
    `landing_type` VARCHAR(50) DEFAULT '' COMMENT '推广目的',
    `marketing_goal` VARCHAR(50) DEFAULT '' COMMENT '营销目标',
    `delivery_related_num` VARCHAR(50) DEFAULT '' COMMENT '广告组数量限制',
    `status` VARCHAR(50) DEFAULT '' COMMENT '状态',
    `opt_status` VARCHAR(50) DEFAULT '' COMMENT '操作状态',
    `modify_time` DATETIME DEFAULT NULL COMMENT '修改时间',
    `create_time` DATETIME DEFAULT NULL COMMENT 'Ocean端创建时间',
    `last_sync_at` DATETIME DEFAULT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_campaign_id` (`campaign_id`),
    KEY `idx_advertiser_id` (`advertiser_id`),
    KEY `idx_status` (`status`),
    KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告系列表';
```

#### ad_ad - 广告组表
```sql
CREATE TABLE `ad_ad` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `ad_id` BIGINT UNSIGNED NOT NULL COMMENT 'Ocean Engine 广告组ID',
    `campaign_id` BIGINT UNSIGNED NOT NULL COMMENT '广告系列ID',
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT '广告主ID',
    `name` VARCHAR(255) NOT NULL COMMENT '广告组名称',
    `status` VARCHAR(50) DEFAULT '' COMMENT '状态',
    `opt_status` VARCHAR(50) DEFAULT '' COMMENT '操作状态',
    `delivery_range` VARCHAR(50) DEFAULT '' COMMENT '投放范围',
    `budget_mode` VARCHAR(50) DEFAULT '' COMMENT '预算类型',
    `budget` DECIMAL(15,2) DEFAULT 0.00 COMMENT '预算金额',
    `bid_type` VARCHAR(50) DEFAULT '' COMMENT '出价类型',
    `bid` DECIMAL(10,2) DEFAULT 0.00 COMMENT '出价金额',
    `cpa_bid` DECIMAL(10,2) DEFAULT 0.00 COMMENT 'CPA出价',
    `deep_bid_type` VARCHAR(50) DEFAULT '' COMMENT '深度出价类型',
    `deep_cpabid` DECIMAL(10,2) DEFAULT 0.00 COMMENT '深度CPA出价',
    `pricing` VARCHAR(50) DEFAULT '' COMMENT '计费方式',
    `start_time` DATETIME DEFAULT NULL COMMENT '投放开始时间',
    `end_time` DATETIME DEFAULT NULL COMMENT '投放结束时间',
    `schedule_time` VARCHAR(1000) DEFAULT '' COMMENT '投放时段',
    `audience_package_id` BIGINT UNSIGNED DEFAULT 0 COMMENT '定向包ID',
    `modify_time` DATETIME DEFAULT NULL,
    `create_time` DATETIME DEFAULT NULL,
    `last_sync_at` DATETIME DEFAULT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_ad_id` (`ad_id`),
    KEY `idx_campaign_id` (`campaign_id`),
    KEY `idx_advertiser_id` (`advertiser_id`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告组表';
```

#### ad_creative - 创意表
```sql
CREATE TABLE `ad_creative` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `creative_id` BIGINT UNSIGNED NOT NULL COMMENT 'Ocean Engine 创意ID',
    `ad_id` BIGINT UNSIGNED NOT NULL COMMENT '广告组ID',
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT '广告主ID',
    `name` VARCHAR(255) DEFAULT '' COMMENT '创意名称',
    `status` VARCHAR(50) DEFAULT '' COMMENT '状态',
    `opt_status` VARCHAR(50) DEFAULT '' COMMENT '操作状态',
    `creative_material_mode` VARCHAR(50) DEFAULT '' COMMENT '创意类型',
    `image_mode` VARCHAR(50) DEFAULT '' COMMENT '素材类型',
    `title` VARCHAR(500) DEFAULT '' COMMENT '标题',
    `source` VARCHAR(255) DEFAULT '' COMMENT '来源',
    `image_ids` TEXT COMMENT '图片ID列表（JSON）',
    `video_id` VARCHAR(100) DEFAULT '' COMMENT '视频ID',
    `third_party_id` VARCHAR(255) DEFAULT '' COMMENT '第三方监测ID',
    `modify_time` DATETIME DEFAULT NULL,
    `create_time` DATETIME DEFAULT NULL,
    `last_sync_at` DATETIME DEFAULT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_creative_id` (`creative_id`),
    KEY `idx_ad_id` (`ad_id`),
    KEY `idx_advertiser_id` (`advertiser_id`),
    KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='创意表';
```

### 4. 素材管理模块

#### ad_material_image - 图片素材表
```sql
CREATE TABLE `ad_material_image` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `image_id` VARCHAR(100) NOT NULL COMMENT 'Ocean Engine 图片ID',
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT '广告主ID',
    `filename` VARCHAR(255) DEFAULT '' COMMENT '文件名',
    `size` INT DEFAULT 0 COMMENT '文件大小（字节）',
    `width` INT DEFAULT 0 COMMENT '宽度',
    `height` INT DEFAULT 0 COMMENT '高度',
    `format` VARCHAR(20) DEFAULT '' COMMENT '格式',
    `url` VARCHAR(500) DEFAULT '' COMMENT '图片URL',
    `material_id` VARCHAR(100) DEFAULT '' COMMENT '素材ID',
    `signature` VARCHAR(100) DEFAULT '' COMMENT 'MD5签名',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_image_id` (`image_id`),
    KEY `idx_advertiser_id` (`advertiser_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='图片素材表';
```

#### ad_material_video - 视频素材表
```sql
CREATE TABLE `ad_material_video` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `video_id` VARCHAR(100) NOT NULL COMMENT 'Ocean Engine 视频ID',
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT '广告主ID',
    `filename` VARCHAR(255) DEFAULT '' COMMENT '文件名',
    `size` INT DEFAULT 0 COMMENT '文件大小（字节）',
    `width` INT DEFAULT 0 COMMENT '宽度',
    `height` INT DEFAULT 0 COMMENT '高度',
    `duration` DECIMAL(10,2) DEFAULT 0.00 COMMENT '时长（秒）',
    `format` VARCHAR(20) DEFAULT '' COMMENT '格式',
    `url` VARCHAR(500) DEFAULT '' COMMENT '视频URL',
    `poster_url` VARCHAR(500) DEFAULT '' COMMENT '封面URL',
    `material_id` VARCHAR(100) DEFAULT '' COMMENT '素材ID',
    `signature` VARCHAR(100) DEFAULT '' COMMENT 'MD5签名',
    `bit_rate` INT DEFAULT 0 COMMENT '比特率',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_video_id` (`video_id`),
    KEY `idx_advertiser_id` (`advertiser_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频素材表';
```

### 5. 人群定向模块

#### ad_audience_package - 定向包表
```sql
CREATE TABLE `ad_audience_package` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `package_id` BIGINT UNSIGNED NOT NULL COMMENT 'Ocean Engine 定向包ID',
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT '广告主ID',
    `name` VARCHAR(255) NOT NULL COMMENT '定向包名称',
    `description` VARCHAR(500) DEFAULT '' COMMENT '描述',
    `status` VARCHAR(50) DEFAULT '' COMMENT '状态',
    `landing_type` VARCHAR(50) DEFAULT '' COMMENT '推广类型',
    `audience` TEXT COMMENT '定向设置（JSON）',
    `delivery_range` VARCHAR(50) DEFAULT '' COMMENT '投放范围',
    `modify_time` DATETIME DEFAULT NULL,
    `create_time` DATETIME DEFAULT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_package_id` (`package_id`),
    KEY `idx_advertiser_id` (`advertiser_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='定向包表';
```

#### ad_custom_audience - 自定义人群表
```sql
CREATE TABLE `ad_custom_audience` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `custom_audience_id` BIGINT UNSIGNED NOT NULL COMMENT 'Ocean Engine 人群包ID',
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT '广告主ID',
    `name` VARCHAR(255) NOT NULL COMMENT '人群包名称',
    `source` VARCHAR(50) DEFAULT '' COMMENT '来源类型',
    `status` INT DEFAULT 0 COMMENT '状态',
    `cover_num` BIGINT DEFAULT 0 COMMENT '覆盖人数',
    `tag` VARCHAR(100) DEFAULT '' COMMENT '标签',
    `push_status` INT DEFAULT 0 COMMENT '推送状态',
    `modify_time` DATETIME DEFAULT NULL,
    `create_time` DATETIME DEFAULT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_custom_audience_id` (`custom_audience_id`),
    KEY `idx_advertiser_id` (`advertiser_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='自定义人群表';
```

### 6. 数据报告模块

#### rpt_advertiser_daily - 广告主日报表
```sql
CREATE TABLE `rpt_advertiser_daily` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `advertiser_id` BIGINT UNSIGNED NOT NULL COMMENT '广告主ID',
    `stat_date` DATE NOT NULL COMMENT '统计日期',
    `cost` DECIMAL(15,2) DEFAULT 0.00 COMMENT '消耗',
    `show_cnt` BIGINT DEFAULT 0 COMMENT '展示次数',
    `click_cnt` BIGINT DEFAULT 0 COMMENT '点击次数',
    `ctr` DECIMAL(10,4) DEFAULT 0.0000 COMMENT '点击率',
    `avg_click_cost` DECIMAL(10,2) DEFAULT 0.00 COMMENT '平均点击成本',
    `avg_show_cost` DECIMAL(10,4) DEFAULT 0.0000 COMMENT '千次展示成本',
    `convert_cnt` BIGINT DEFAULT 0 COMMENT '转化数',
    `convert_cost` DECIMAL(10,2) DEFAULT 0.00 COMMENT '转化成本',
    `convert_rate` DECIMAL(10,4) DEFAULT 0.0000 COMMENT '转化率',
    `deep_convert_cnt` BIGINT DEFAULT 0 COMMENT '深度转化数',
    `deep_convert_cost` DECIMAL(10,2) DEFAULT 0.00 COMMENT '深度转化成本',
    `deep_convert_rate` DECIMAL(10,4) DEFAULT 0.0000 COMMENT '深度转化率',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_advertiser_date` (`advertiser_id`, `stat_date`),
    KEY `idx_stat_date` (`stat_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告主日报表';
```

#### rpt_campaign_daily - 广告系列日报表
```sql
CREATE TABLE `rpt_campaign_daily` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `advertiser_id` BIGINT UNSIGNED NOT NULL,
    `campaign_id` BIGINT UNSIGNED NOT NULL COMMENT '广告系列ID',
    `stat_date` DATE NOT NULL COMMENT '统计日期',
    `cost` DECIMAL(15,2) DEFAULT 0.00,
    `show_cnt` BIGINT DEFAULT 0,
    `click_cnt` BIGINT DEFAULT 0,
    `ctr` DECIMAL(10,4) DEFAULT 0.0000,
    `avg_click_cost` DECIMAL(10,2) DEFAULT 0.00,
    `avg_show_cost` DECIMAL(10,4) DEFAULT 0.0000,
    `convert_cnt` BIGINT DEFAULT 0,
    `convert_cost` DECIMAL(10,2) DEFAULT 0.00,
    `convert_rate` DECIMAL(10,4) DEFAULT 0.0000,
    `deep_convert_cnt` BIGINT DEFAULT 0,
    `deep_convert_cost` DECIMAL(10,2) DEFAULT 0.00,
    `deep_convert_rate` DECIMAL(10,4) DEFAULT 0.0000,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_campaign_date` (`campaign_id`, `stat_date`),
    KEY `idx_advertiser_id` (`advertiser_id`),
    KEY `idx_stat_date` (`stat_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告系列日报表';
```

#### rpt_ad_daily - 广告组日报表
```sql
CREATE TABLE `rpt_ad_daily` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `advertiser_id` BIGINT UNSIGNED NOT NULL,
    `campaign_id` BIGINT UNSIGNED NOT NULL,
    `ad_id` BIGINT UNSIGNED NOT NULL COMMENT '广告组ID',
    `stat_date` DATE NOT NULL,
    `cost` DECIMAL(15,2) DEFAULT 0.00,
    `show_cnt` BIGINT DEFAULT 0,
    `click_cnt` BIGINT DEFAULT 0,
    `ctr` DECIMAL(10,4) DEFAULT 0.0000,
    `avg_click_cost` DECIMAL(10,2) DEFAULT 0.00,
    `avg_show_cost` DECIMAL(10,4) DEFAULT 0.0000,
    `convert_cnt` BIGINT DEFAULT 0,
    `convert_cost` DECIMAL(10,2) DEFAULT 0.00,
    `convert_rate` DECIMAL(10,4) DEFAULT 0.0000,
    `deep_convert_cnt` BIGINT DEFAULT 0,
    `deep_convert_cost` DECIMAL(10,2) DEFAULT 0.00,
    `deep_convert_rate` DECIMAL(10,4) DEFAULT 0.0000,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_ad_date` (`ad_id`, `stat_date`),
    KEY `idx_advertiser_id` (`advertiser_id`),
    KEY `idx_campaign_id` (`campaign_id`),
    KEY `idx_stat_date` (`stat_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='广告组日报表';
```

## ER 关系图

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│    sys_user     │     │    sys_role     │     │    sys_menu     │
├─────────────────┤     ├─────────────────┤     ├─────────────────┤
│ id              │     │ id              │     │ id              │
│ role_id ────────┼────>│ name            │<────┼─ role_id        │
│ username        │     │ key             │     │ name            │
│ ...             │     │ ...             │     │ ...             │
└─────────────────┘     └─────────────────┘     └─────────────────┘
                              │
                              │ sys_role_menu
                              ▼
                        ┌─────────────────┐
                        │ role_id         │
                        │ menu_id         │
                        └─────────────────┘

┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│  ad_advertiser  │────>│   ad_campaign   │────>│     ad_ad       │
├─────────────────┤     ├─────────────────┤     ├─────────────────┤
│ id              │     │ id              │     │ id              │
│ advertiser_id   │     │ campaign_id     │     │ ad_id           │
│ name            │     │ advertiser_id   │     │ campaign_id     │
│ balance         │     │ name            │     │ name            │
│ ...             │     │ budget          │     │ budget          │
└─────────────────┘     └─────────────────┘     └─────────────────┘
        │                                               │
        │                                               │
        ▼                                               ▼
┌─────────────────┐                           ┌─────────────────┐
│ rpt_adv_daily   │                           │   ad_creative   │
├─────────────────┤                           ├─────────────────┤
│ advertiser_id   │                           │ creative_id     │
│ stat_date       │                           │ ad_id           │
│ cost            │                           │ title           │
│ show_cnt        │                           │ image_mode      │
│ ...             │                           │ ...             │
└─────────────────┘                           └─────────────────┘
```

## 数据同步策略

### 同步方式
1. **手动同步**：用户点击同步按钮，实时从 Ocean Engine API 拉取数据
2. **定时同步**：定时任务定期同步数据（建议频率：广告数据 5分钟，报告数据 1小时）
3. **Webhook**：（如 Ocean Engine 支持）实时接收状态变更通知

### 同步字段
- `last_sync_at`：记录最后同步时间
- 本地新增字段不影响 Ocean Engine 原始数据

### 数据一致性
- 以 Ocean Engine 数据为准
- 本地缓存用于快速查询和减少 API 调用
- 冲突时以远程数据覆盖本地
