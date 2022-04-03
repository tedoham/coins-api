CREATE TABLE "currency_type" (
  "id" bigserial PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "account" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "account_number" varchar NOT NULL,
  "balance" float DEFAULT 0
);

CREATE TABLE "payments" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "currency_type_id" bigint NOT NULL,
  "direction" varchar NOT NULL,
  "amount" float DEFAULT 0
);

ALTER TABLE "payments" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("currency_type_id") REFERENCES "currency_type" ("id");

CREATE INDEX ON "account" ("name");

CREATE INDEX ON "account" ("account_number");

CREATE INDEX ON "payments" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "account"."name" IS 'account name required';

COMMENT ON COLUMN "account"."account_number" IS 'account number required';
