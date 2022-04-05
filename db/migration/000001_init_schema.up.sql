CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "balance" float DEFAULT 0
);

INSERT INTO "accounts" VALUES
	(1, 'account_one', 1000),
	(2, 'account_two', 2000.96),
  (3, 'account_three', 7000),
  (4, 'account_four', 5000.86);


CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "from_account" bigint NOT NULL,
  "to_account" bigint NOT NULL,
  "transaction_type" VARCHAR NOT NULL,
  "currency_type" VARCHAR DEFAULT 'USD',
  "amount" float DEFAULT 0
);

ALTER TABLE "transactions" ADD FOREIGN KEY ("from_account") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("to_account") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("name");

CREATE INDEX ON "transactions" ("from_account", "to_account");



