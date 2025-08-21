-- +migrate Up
CREATE TABLE `profiles` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `player_id` BIGINT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  CONSTRAINT `fk_profiles_players`
    FOREIGN KEY (`player_id`) REFERENCES `players`(`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_persian_ci ROW_FORMAT=Dynamic;

-- +migrate Down
DROP TABLE IF EXISTS `profiles`;