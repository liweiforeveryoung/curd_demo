CREATE TABLE IF NOT EXISTS `users`
(
    `id`         BIGINT(20)   NOT NULL AUTO_INCREMENT,
    `user_id`    BIGINT(20)   NOT NULL DEFAULT 0 COMMENT '用户 id',
    `name`       VARCHAR(191) NOT NULL DEFAULT '' COMMENT '用户昵称',
    `age`        INT          NOT NULL DEFAULT 0 COMMENT '用户年龄',
    `deleted_at` BIGINT(20)   NOT NULL DEFAULT 0,
    `created_at` BIGINT(20)   NOT NULL DEFAULT 0,
    `updated_at` BIGINT(20)   NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    UNIQUE INDEX `udx_user_id` (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT '用户信息表';