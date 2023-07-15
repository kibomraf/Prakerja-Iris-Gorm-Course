-- Active: 1689349311789@@127.0.0.1@5432@course
CREATE TABLE "students" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" varchar  NOT NULL,
  "address" varchar NOT NULL,
  "no_handphone" VARCHAR NOT NULL,
  "email" varchar  NOT NULL,
  "password_hash" varchar NOT NULL,
  "avatar_file_name" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);
