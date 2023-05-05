CREATE TABLE `user`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーの識別子',
    `user_id` VARCHAR(80) NOT NULL COMMENT 'ユーザー名のuid',
    `created_at` DATETIME(6) NOT NULL COMMENT 'ユーザー登録日時',
    `modified_at` DATETIME(6) NOT NULL COMMENT 'ユーザーログイン日時',
    PRIMARY KEY (`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `whisky_record` (
  `user_id` BIGINT UNSIGNED NOT NULL,
  `whisky_name` VARCHAR(255) NOT NULL,
  `drankAt` DATE NOT NULL,
  `imageUrl` VARCHAR(255),
  CONSTRAINT `fk_user_id`
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
      ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;