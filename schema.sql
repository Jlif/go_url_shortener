CREATE TABLE `shortener_url`
(
    `id`          int unsigned  NOT NULL AUTO_INCREMENT,
    `url`         varchar(2048) NOT NULL COMMENT '原始url',
    `short_url`   varchar(8)    NOT NULL COMMENT '短网址字符串',
    `create_time` datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `short_url` (`short_url`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;