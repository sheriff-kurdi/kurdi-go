DROP TABLE IF EXISTS "languages";
CREATE TABLE "public"."languages"
(
    "code"       text NOT NULL,
    "name"       text,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    CONSTRAINT "languages_pkey" PRIMARY KEY ("code")
) WITH (oids = false);

CREATE INDEX "idx_languages_deleted_at" ON "public"."languages" USING btree ("deleted_at");

INSERT INTO "languages" ("code", "name", "created_at", "updated_at", "deleted_at")
VALUES ('en', 'English', '2022-11-26 17:26:57.001534+00', NULL, NULL),
       ('ar', 'العربيه', '2022-11-26 17:30:33.828364+00', NULL, NULL);