CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `name` char(20) DEFAULT NULL COMMENT '姓名',
    `email` char(20) DEFAULT NULL COMMENT '邮箱',
    `dept_name` char(20) DEFAULT NULL COMMENT '部门',
    `create_time` char(30) NOT NULL COMMENT '创建时间',
    `update_time` char(30) DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;