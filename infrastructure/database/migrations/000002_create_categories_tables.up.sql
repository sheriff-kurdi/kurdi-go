DROP TABLE IF EXISTS "categories";
CREATE TABLE "public"."categories"
(
    "id"              text NOT NULL,
    "is_parent"       boolean,
    "parent_category" text,
    "created_at"      timestamptz,
    "updated_at"      timestamptz,
    "deleted_at"      timestamptz,
    CONSTRAINT "categories_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

CREATE INDEX "idx_categories_deleted_at" ON "public"."categories" USING btree ("deleted_at");

INSERT INTO "categories" ("id", "is_parent", "parent_category", "created_at", "updated_at", "deleted_at")
VALUES ('MEN', '1', NULL, '2022-11-26 17:32:29.932498+00', NULL, NULL),
       ('Women', '1', NULL, '2022-11-26 17:32:41.653415+00', NULL, NULL);
--Categories
ALTER TABLE ONLY "public"."categories"
    ADD CONSTRAINT "fk_categories_parent" FOREIGN KEY (parent_category) REFERENCES categories (id) NOT DEFERRABLE;

DROP TABLE IF EXISTS "category_details";
CREATE TABLE "public"."category_details"
(
    "id"            text NOT NULL,
    "name"          text NOT NULL,
    "language_code" text NOT NULL,
    "created_at"    timestamptz,
    "updated_at"    timestamptz,
    "deleted_at"    timestamptz,
    CONSTRAINT "category_details_pkey" PRIMARY KEY ("id", "name", "language_code")
) WITH (oids = false);

CREATE INDEX "idx_category_details_deleted_at" ON "public"."category_details" USING btree ("deleted_at");
--Categories Details
INSERT INTO "category_details" ("id", "name", "language_code", "created_at", "updated_at", "deleted_at")
VALUES ('MEN', 'Men', 'en', '2022-11-26 17:33:24.446545+00', NULL, NULL),
       ('MEN', 'رجالي', 'ar', '2022-11-26 17:33:46.25756+00', NULL, NULL),
       ('WOMEN', 'Women', 'en', '2022-11-26 17:33:24.446545+00', NULL, NULL),
       ('WOMEN', 'حريمي', 'ar', '2022-11-26 17:33:24.446545+00', NULL, NULL);

ALTER TABLE ONLY "public"."category_details"
    ADD CONSTRAINT "fk_category_details_language" FOREIGN KEY (language_code) REFERENCES languages (code) NOT DEFERRABLE;