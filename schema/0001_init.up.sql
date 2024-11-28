-- create enum account_type
CREATE TYPE "account_type" AS ENUM (
  'PREMIUM',
  'FREE'
);

-- create enum swipe_type
CREATE TYPE "swipe_type" AS ENUM (
  'LIKE',
  'PASS'
);

-- create function to generate uuid
CREATE
OR REPLACE FUNCTION uuid_generate_v4()
 RETURNS uuid
 LANGUAGE c
 PARALLEL SAFE STRICT
AS '$libdir/uuid-ossp', $function$uuid_generate_v4$function$;

-- create table account
CREATE TABLE "account"
(
    "id"              SERIAL       NOT NULL,
    "account_mask_id" uuid UNIQUE  NOT NULL DEFAULT (uuid_generate_v4()),
    "type"            account_type          DEFAULT null,
    "name"            varchar(45)           DEFAULT null,
    "user_name"       varchar(45)           DEFAULT null,
    "password"        varchar(225)          DEFAULT null,
    "created_at"      timestamp    NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "created_by"      varchar(225) NOT NULL,
    "updated_at"      timestamp    NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_by"      varchar(225),
    PRIMARY KEY ("id")
);
-- create unique index on account
CREATE UNIQUE INDEX ON "account" USING BTREE ("user_name");

-- create table swipe_count
CREATE TABLE "swipe_count"
(
    "id"                SERIAL NOT NULL,
    "account_id"        int    NOT NULL,
    "total_swipe_a_day" int    NOT NULL,
    "total_swipe"       int    NOT NULL
);

-- create unique index on swipe_count
CREATE UNIQUE INDEX ON "swipe_count_account_id_unique" USING BTREE ("account_id");

-- create foreign key
ALTER TABLE "swipe_count"
    ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

-- create table user_swipe_log
CREATE TABLE "user_swipe_log"
(
    "id"         SERIAL     NOT NULL,
    "swiper_id"  int        NOT NULL,
    "swipee_id"  int        NOT NULL,
    "swipe_type" swipe_type NOT NULL,
    "created_at" timestamp  NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE "user_swipe_log"
    ADD CONSTRAINT "fk_user_swipe_log_swiper_id" FOREIGN KEY ("swiper_id") REFERENCES "account" ("id");

ALTER TABLE "user_swipe_log"
    ADD CONSTRAINT "fk_user_swipe_log_swipee_id" FOREIGN KEY ("swipee_id") REFERENCES "account" ("id");
