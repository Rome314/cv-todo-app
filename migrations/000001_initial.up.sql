-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS id_seq;

-- Table Definition
CREATE TABLE "public"."users"
(
    "id"           int4 NOT NULL DEFAULT nextval('id_seq'::regclass),
    "name"         text,
    "mail"         text,
    "phone"        text,
    "created"      timestamp     DEFAULT now(),
    "last_updated" timestamp     DEFAULT now(),
    PRIMARY KEY ("id")
);