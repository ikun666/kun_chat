CREATE TABLE `messages`
(
    `id`          bigint unsigned     NOT NULL AUTO_INCREMENT,
    `form_id`        bigint unsigned        NOT NULL  COMMENT '信息发送者',
    `target_id`      bigint unsigned NOT NULL  COMMENT '信息接收者',
    `type`       int          NOT NULL DEFAULT 1 COMMENT '聊天类型:私聊 1 群聊2  广播3',
    `media`       int          NOT NULL DEFAULT 1 COMMENT '信息类型:文字1 图片2 音频3',
    `content`    longtext        NOT NULL  COMMENT '消息内容',
    `pic`    longtext        NOT NULL  COMMENT '图片相关',
    `url`    longtext        NOT NULL  COMMENT '文件相关',
    `desc`    longtext        NOT NULL  COMMENT '文件描述',
    `amount`     bigint NOT NULL DEFAULT 0 COMMENT '其他数据大小',
    `create_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp           NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_form_id` (`form_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;