CREATE TABLE `user`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `user_id` VARCHAR(80) NOT NULL COMMENT 'ユーザー名のuid',
    `created_at` DATETIME(6) NOT NULL COMMENT 'ユーザー登録日時',
    `modified_at` DATETIME(6) NOT NULL COMMENT 'ユーザーログイン日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';
