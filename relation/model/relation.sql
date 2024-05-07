CREATE TABLE `relation`
(
    `id`          bigint unsigned     NOT NULL AUTO_INCREMENT,
    `owner_id`   bigint unsigned NOT NULL COMMENT '谁的关系信息',
    `target_id`  bigint unsigned NOT NULL COMMENT '对应的谁',
    `type`       int          NOT NULL DEFAULT 1 COMMENT '关系类型： 1表示好友关系 2表示群关系',
    `desc`        varchar(255)        NOT NULL DEFAULT ''  COMMENT '描述',
    `create_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_owner_id` (`owner_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
