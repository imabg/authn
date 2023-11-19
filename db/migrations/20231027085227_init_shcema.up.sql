-- enum
CREATE TYPE user_type AS ENUM ('USER', 'TENANT_USER');

CREATE TABLE "tenants" (
    "id" varchar(50) PRIMARY KEY,
    "name" varchar(50) NOT NULL,
    "email" varchar(100) NOT NULL UNIQUE,
    "phone" varchar(20),
    "company_size" varchar(20),
    "support_url" varchar(50),
    "support_email" varchar(50),
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "sources" (
  "id" varchar(50) PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "description" varchar(255) NOT NULL,
  "disable_new_user" boolean DEFAULT false,
  "opt_for_passwordless" boolean DEFAULT false,
  "is_active" boolean DEFAULT true,
  "is_system_generated" boolean DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- 1-M relationship between tenants and sources

CREATE TABLE "tenant_sources"(
  "tenant_id" varchar(50) REFERENCES  "tenants" ("id"),
  "source_id" varchar(50) REFERENCES  "sources"("id"),
  PRIMARY KEY (tenant_id, source_id)
);

CREATE TABLE "users" (
  "id" varchar(50) PRIMARY KEY,
  "email" varchar(255),
  "phone" varchar(10),
  "is_email_verified" boolean DEFAULT false,
  "is_phone_verified" boolean DEFAULT false,
  "country_code" varchar(10),
  "is_blacklisted" boolean DEFAULT false,
  "user_type" user_type NOT NULL,
  "source_id" varchar(50),
  FOREIGN KEY ("source_id") REFERENCES "sources" ("id") ON DELETE CASCADE,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT  source_id_check CHECK ( (user_type = 'USER' AND source_id IS NOT NULL ) OR (user_type <> 'USER') )
);


CREATE TABLE "credentials" (
    "id" varchar(50) PRIMARY KEY,
    "password" varchar(255) NOT NULL,
    "user_type" user_type NOT NULL,
    "user_id" varchar(50) NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "sourceConfigs" (
    "id" varchar(50) PRIMARY KEY,
    "password_config" jsonb,
    "is_active" boolean DEFAULT true,
    "max_concurrent_login" int DEFAULT 1,
    "source_id" varchar(50) NOT NULL,
    FOREIGN KEY ("source_id") REFERENCES "sources"("id") ON DELETE CASCADE,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index
CREATE INDEX credentials_user_id ON credentials(user_id);


-- Unique index
CREATE UNIQUE INDEX unique_phone_sourceId ON users (phone, source_id) WHERE phone <> '';
CREATE UNIQUE INDEX unique_email_sourceId ON users (email, source_id) WHERE email <> '';