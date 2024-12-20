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
    "type"            account_type NOT NULL,
    "name"            varchar(45)  NOT NULL,
    "user_name"       varchar(45)  NOT NULL,
    "password"        varchar(225) NOT NULL,
    "is_verified"     bool         NOT NULL DEFAULT false,
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
    "id"                SERIAL    NOT NULL,
    "account_id"        int       NOT NULL,
    "total_swipe_a_day" int       NOT NULL,
    "total_swipe"       int       NOT NULL,
    "last_updated_at"   timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- create unique index on swipe_count
CREATE UNIQUE INDEX ON "swipe_count_account_id_unique" USING BTREE ("account_id");

-- create foreign key
ALTER TABLE "swipe_count"
    ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

-- create function to set timestamp
CREATE
OR REPLACE FUNCTION trigger_set_timestamp_last_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
  NEW.last_updated_at
= NOW();
RETURN NEW;
END;
$$;

-- CREATE TRIGGER updated_at
CREATE TRIGGER set_timestamp_last_update
    BEFORE UPDATE
    ON swipe_count
    FOR EACH ROW EXECUTE FUNCTION public.trigger_set_timestamp_last_update();


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

-- func trigger for update swipe counter
CREATE
OR REPLACE FUNCTION update_swipe_count()
RETURNS TRIGGER AS $$
DECLARE
current_day_start DATE;
    today_swipe_count
INT;
BEGIN
    -- Define the start of the current day in UTC+7
    current_day_start
:= (CURRENT_TIMESTAMP)::DATE;

    -- Count today's swipes for the current swiper_id
SELECT COUNT(*)
INTO today_swipe_count
FROM user_swipe_log
WHERE swiper_id = NEW.swiper_id
  AND (created_at)::DATE = current_day_start;

-- Upsert the swipe count record
INSERT INTO swipe_count (account_id, total_swipe_a_day, total_swipe)
VALUES (NEW.swiper_id, today_swipe_count, 1) ON CONFLICT (account_id)
    DO
UPDATE SET
    total_swipe_a_day = today_swipe_count,
    total_swipe = swipe_count.total_swipe + 1;


RETURN NEW;
END;
$$
LANGUAGE plpgsql;

-- Create the trigger
CREATE TRIGGER trigger_update_swipe_count
    AFTER INSERT
    ON user_swipe_log
    FOR EACH ROW
    EXECUTE FUNCTION update_swipe_count();

CREATE TABLE "premium_package"
(
    "id"          SERIAL       NOT NULL,
    "package_uid" uuid UNIQUE  NOT NULL DEFAULT (uuid_generate_v4()),
    "title"       varchar      NOT NULL,
    "description" text         NOT NULL,
    "price"       float        NOT NULL,
    "is_active"   bool         NOT NULL,
    "created_at"  timestamp    NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "created_by"  varchar(225) NOT NULL,
    "updated_at"  timestamp    NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    "updated_by"  varchar(225)
);

CREATE TABLE "premium_package_user"
(
    "id"                 SERIAL    NOT NULL,
    "premium_package_id" int       NOT NULL,
    "account_id"         int       NOT NULL,
    "purchased_date"      timestamp NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

ALTER TABLE "premium_package_user"
    ADD CONSTRAINT "fk_premium_package_id" FOREIGN KEY ("premium_package_id") REFERENCES "premium_package" ("id");
ALTER TABLE "premium_package_user"
    ADD CONSTRAINT "fk_premium_package_user_account_id" FOREIGN KEY ("account_id") REFERENCES "account" ("id");

CREATE UNIQUE INDEX IF NOT EXISTS premium_package_user_account_id_premium_package_id_unique_idx ON premium_package_user (premium_package_id,account_id);