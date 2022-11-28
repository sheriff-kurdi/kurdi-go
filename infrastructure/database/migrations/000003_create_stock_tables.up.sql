--Stock
DROP TABLE IF EXISTS "stock_items";
CREATE TABLE "public"."stock_items"
(
    "sku"             text NOT NULL,
    "cost_price"      numeric,
    "selling_price"   numeric,
    "is_discounted"   boolean,
    "discount"        numeric,
    "total_stock"     bigint,
    "available_stock" bigint,
    "reserved_stock"  bigint,
    "created_at"      timestamptz,
    "updated_at"      timestamptz,
    "deleted_at"      timestamptz,
    CONSTRAINT "stock_items_pkey" PRIMARY KEY ("sku")
) WITH (oids = false);

CREATE INDEX "idx_stock_items_deleted_at" ON "public"."stock_items" USING btree ("deleted_at");

INSERT INTO "stock_items" ("sku", "cost_price", "selling_price", "is_discounted", "discount", "total_stock",
                           "available_stock", "reserved_stock", "created_at", "updated_at", "deleted_at")
VALUES ('122', 25, 60, '1', 55, 6, 5, 1, '2022-11-26 17:39:17.25387+00', NULL, NULL),
       ('265', 120, 250, '1', 25, 2, 1, 1, '2022-11-26 17:39:17.25387+00', NULL, NULL);

--Stock Details
DROP TABLE IF EXISTS "stock_details";
CREATE TABLE "public"."stock_details"
(
    "sku"           text NOT NULL,
    "language_code" text NOT NULL,
    "name"          text,
    "description"   text,
    "created_at"    timestamptz,
    "updated_at"    timestamptz,
    "deleted_at"    timestamptz,
    CONSTRAINT "stock_details_pkey" PRIMARY KEY ("sku", "language_code")
) WITH (oids = false);

CREATE INDEX "idx_stock_details_deleted_at" ON "public"."stock_details" USING btree ("deleted_at");

INSERT INTO "stock_details" ("sku", "language_code", "name", "description", "created_at", "updated_at", "deleted_at")
VALUES ('122', 'en', 'Back Bag', 'Black Back Bag', '2022-11-26 17:41:21.0649+00', NULL, NULL),
       ('265', 'en', 'Shirt', 'White Shirt', '2022-11-26 17:42:00.124037+00', NULL, NULL),
       ('122', 'ar', 'شنطه ظهر', 'شنطه ظهر سوداء', '2022-11-26 17:41:21.0649+00', NULL, NULL),
       ('265', 'ar', 'قميص', 'قميص ابيض', '2022-11-26 17:42:00.124037+00', NULL, NULL);

ALTER TABLE ONLY "public"."stock_details"
    ADD CONSTRAINT "fk_stock_details_language" FOREIGN KEY (language_code) REFERENCES languages (code) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."stock_details"
    ADD CONSTRAINT "fk_stock_items_details" FOREIGN KEY (sku) REFERENCES stock_items (sku) NOT DEFERRABLE;