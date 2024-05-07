CREATE TABLE `group`
(
    `id`          bigint unsigned     NOT NULL AUTO_INCREMENT,
    `owner_id`   bigint unsigned NOT NULL COMMENT '群拥有者',
    `name`  varchar(255) NOT NULL  COMMENT '群名称',
    `type`       int          NOT NULL DEFAULT 0 COMMENT '群类型',
    `image`  varchar(255) NOT NULL NULL DEFAULT '' COMMENT '头像',
    `desc`        varchar(255)        NOT NULL DEFAULT ''  COMMENT '描述',
    `create_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_owner_id` (`owner_id`),
    UNIQUE KEY `idx_name_unique` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;