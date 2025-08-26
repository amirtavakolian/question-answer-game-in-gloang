-- +migrate Up
CREATE TABLE `permission_role` (
  `role_id` BIGINT UNSIGNED NOT NULL,
  `permission_id` INT NOT NULL,
  PRIMARY KEY (`role_id`, `permission_id`),
  CONSTRAINT `fk_permission_role_role`
    FOREIGN KEY (`role_id`) REFERENCES `roles`(`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_permission_role_permission`
    FOREIGN KEY (`permission_id`) REFERENCES `permissions`(`id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_persian_ci;

-- +migrate Down
DROP TABLE IF EXISTS `permission_role`;
