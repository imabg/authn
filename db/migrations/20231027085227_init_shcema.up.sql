CREATE TABLE "sources" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "description" varchar(255) NOT NULL,
  "disable_user_creation" boolean DEFAULT false,
  "is_active" boolean DEFAULT true,
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "email" varchar(50) NOT NULL,
  "display_id" varchar NOT NULL,
  "phone" int,
  "is_email_verified" boolean DEFAULT false,
  "is_phone_verified" boolean DEFAULT false,
  "country_code" varchar(10),
  "is_blacklisted" boolean DEFAULT false,
  "source_id" bigint,
  FOREIGN KEY ("source_id") REFERENCES "sources" ("id") ON DELETE CASCADE,
  UNIQUE(email, source_id),
  UNIQUE (phone, source_id),
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX users_email ON users(email);
CREATE INDEX users_display_id ON users(display_id);
CREATE INDEX users_phone ON users(phone) WHERE phone IS NOT NULL;


CREATE TABLE "configs" (
  "id" bigserial PRIMARY KEY,
  "password_config" varchar DEFAULT '^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$',
  "source_id" bigint,
  FOREIGN KEY ("source_id") REFERENCES "sources" ("id"),
  "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

