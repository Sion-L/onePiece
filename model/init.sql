CREATE TABLE `user` (
    `id` bigint unsigned NOT NULL DEFAULT '1' COMMENT '用户id',
    `cn` char(20) UNIQUE NOT NULL COMMENT '中文名',
    `en` char(20) UNIQUE NOT NULL COMMENT '英文名',
    `email` char(20) NOT NULL COMMENT '邮箱',
    `dept_name` char(20) NOT NULL DEFAULT 'employee' COMMENT '部门',
    `role` char(20) DEFAULT NULL COMMENT '角色',
    `business` char(20) DEFAULT NULL COMMENT '业务线',
    `create_time` char(30) DEFAULT NULL COMMENT '创建时间',
    `update_time` char(30) DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;