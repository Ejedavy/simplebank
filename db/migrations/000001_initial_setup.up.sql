CREATE TABLE  "Account" (
                           "id" bigserial PRIMARY KEY,
                           "owner" varchar NOT NULL,
                           "balance" bigint NOT NULL,
                           "createdAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE  "Entry" (
                         "id" bigserial PRIMARY KEY,
                         "account_id" bigint NOT NULL,
                         "amoount" bigint NOT NULL,
                         "createdAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE  "Transfer" (
                            "id" bigserial PRIMARY KEY,
                            "from_account_id" bigint NOT NULL,
                            "to_account_id" bigint NOT NULL,
                            "amoount" bigint,
                            "createdAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "Account" ("owner");

CREATE INDEX ON "Entry" ("account_id");

CREATE INDEX ON "Transfer" ("from_account_id");

CREATE INDEX ON "Transfer" ("to_account_id");

CREATE INDEX ON "Transfer" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "Entry"."amoount" IS 'Can be positive or negative';

COMMENT ON COLUMN "Transfer"."amoount" IS 'This can only be positive';

ALTER TABLE "Entry" ADD FOREIGN KEY ("account_id") REFERENCES "Account" ("id");

ALTER TABLE "Transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES "Account" ("id");

ALTER TABLE "Transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "Account" ("id");
