CREATE DATABASE IF NOT EXISTS `hello` DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;

USE hello;
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255),
  `age` INT,
  `stats_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY `pk_id`(`id`)
) ENGINE = InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO `users` (`name`, `age`) VALUES
("a君", NULL),
("b君", NULL),
("c君", 18),
("d君", 20),
("e君", 23),
("f君", 2);
