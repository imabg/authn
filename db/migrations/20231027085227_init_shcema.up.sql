CREATE TABLE "sources" (
  "id" varchar(50) PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "description" varchar(255) NOT NULL,
  "disable_user_creation" boolean DEFAULT false,
  "is_active" boolean DEFAULT true,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "users" (
  "id" varchar(50) PRIMARY KEY,
  "email" varchar(50) NOT NULL,
  "phone" varchar(10),
  "is_email_verified" boolean DEFAULT false,
  "is_phone_verified" boolean DEFAULT false,
  "country_code" varchar(10),
  "is_blacklisted" boolean DEFAULT false,
   "source_id" varchar(50) NOT NULL ,
  FOREIGN KEY ("source_id") REFERENCES "sources" ("id") ON DELETE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Unique index
CREATE UNIQUE INDEX unique_phone_sourceId ON users (phone, source_id) WHERE phone <> '';
CREATE UNIQUE INDEX unique_email_sourceId ON users (email, source_id) WHERE email <> '';