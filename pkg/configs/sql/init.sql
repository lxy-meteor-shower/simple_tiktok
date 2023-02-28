CREATE TABLE `user`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `name`              varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`          varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `follow_count`      bigint unsigned NOT NULL DEFAULT 0 COMMENT 'FollowCount',
    `follower_count`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'FollowerCount',
    `avatar`            varchar(128) NOT NULL DEFAULT '' COMMENT 'UserAvatar',
    `background_image`  varchar(128) NOT NULL DEFAULT '' COMMENT 'BackgroundImage',
    `signature`         varchar(128) NOT NULL DEFAULT '' COMMENT 'Signature',
    `total_favorited`   bigint unsigned NOT NULL DEFAULT 0 COMMENT 'FavoritedCount',
    `work_count`        bigint unsigned NOT NULL DEFAULT 0 COMMENT 'WorkCount',
    `favorite_count`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'FavoriteCount',
    PRIMARY KEY (`id`),
    KEY         `idx_name` (`name`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User table';

CREATE TABLE `video`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `author`            bigint unsigned NOT NULL DEFAULT 0 COMMENT 'AuthorID',
    `play_url`          varchar(128) NOT NULL DEFAULT '' COMMENT 'PlayURL',
    `cover_url`         varchar(128) NOT NULL DEFAULT '' COMMENT 'CoverURL',
    `favorite_count`    bigint unsigned NOT NULL DEFAULT 0 COMMENT 'FavoriteCount',
    `comment_count`     bigint unsigned NOT NULL DEFAULT 0 COMMENT 'CommentCount',
    `title`             varchar(128) NOT NULL DEFAULT '' COMMENT 'Title',
    `create_at`         timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Video create time',
    PRIMARY KEY (`id`),
    KEY         `id_user_id_title` (`author`, `title`) COMMENT 'AuthorID Title index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video table';

CREATE TABLE `favorite`
(
    `user_id`       bigint unsigned NOT NULL DEFAULT 0 COMMENT 'UserID',
    `video_id`      bigint unsigned NOT NULL DEFAULT 0 COMMENT 'VideoID',
    `create_at`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Favorite create time',
    PRIMARY KEY (`user_id`, `video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favorite table';

CREATE TABLE `comment`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`       bigint unsigned NOT NULL DEFAULT 0 COMMENT 'UserID',
    `video_id`      bigint unsigned NOT NULL DEFAULT 0 COMMENT 'VideoID',
    `content`       varchar(128) NOT NULL DEFAULT "" COMMENT 'Content',
    `create_at`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Comment create time',
    PRIMARY KEY (`id`),
    KEY         `video_id_index` (`video_id`) COMMENT 'VideoID Index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment table';
