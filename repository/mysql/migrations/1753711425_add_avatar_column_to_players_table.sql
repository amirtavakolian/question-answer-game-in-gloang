-- +migrate Up
ALTER TABLE players ADD COLUMN `avatar` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_persian_ci DEFAULT NULL;

-- +migrate Down
ALTER TABLE players DROP COLUMN `avatar`;