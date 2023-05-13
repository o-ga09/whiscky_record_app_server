CREATE TABLE `user`
(
    `user_id` VARCHAR(80) NOT NULL COMMENT 'ユーザー名のid',
    `created_at` DATETIME(6) NOT NULL COMMENT 'ユーザー登録日時',
    `modified_at` DATETIME(6) NOT NULL COMMENT 'ユーザーログイン日時',
    PRIMARY KEY (`user_id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ユーザー';

CREATE TABLE `whicky_record` (
  `user_id` VARCHAR(80) NOT NULL,
  `whisky_name` VARCHAR(255) NOT NULL,
  `drankAt` DATE NOT NULL,
  `taste` VARCHAR(255),
  `smell` VARCHAR(255),
  `evaluate` VARCHAR(10),
  `imageUrl` VARCHAR(255),
  CONSTRAINT `fk_user_id`
    FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
      ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;