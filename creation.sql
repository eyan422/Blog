create database assignment;

use assignment;
create table `assignment`.`article` (
    id bigint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    title varchar(256) NOT NULL DEFAULT '' COMMENT 'title',
    content varchar(1024) NOT NULL DEFAULT '' COMMENT 'content',
    author varchar(256) NOT NULL DEFAULT '' COMMENT 'author',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'article table';

-- desc `assignment`.`article`;

-- show create table `assignment`.`article`;