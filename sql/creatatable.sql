CREATE TABLE `shop_product1` (
  `id`          INT UNSIGNED AUTO_INCREMENT,
  `created_at`  TIMESTAMP    NULL,
  `updated_at`  TIMESTAMP    NULL,
  `deleted_at`  TIMESTAMP    NULL,
  `create_by`   VARCHAR(255),
  `name`        VARCHAR(255) NOT NULL,
  `total`       INT UNSIGNED DEFAULT 0,
  `left`        INT UNSIGNED DEFAULT 0,
  `state`       INT UNSIGNED
  COMMENT '状态 1为上架 2为下架 3为仓库中 ',
  `des`         VARCHAR(255),
  `category_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`id`),
  INDEX (category_id)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COMMENT ='商品';

CREATE TABLE `shop_category` (
  `id`          INT UNSIGNED AUTO_INCREMENT,
  `created_at`  TIMESTAMP    NULL,
  `updated_at`  TIMESTAMP    NULL,
  `deleted_at`  TIMESTAMP    NULL,
  `create_by`   VARCHAR(255),
  `name`        VARCHAR(255) NOT NULL,
  `state`       INT UNSIGNED
  COMMENT '状态 0为激活 1为未激活',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COMMENT ='商品类别';


