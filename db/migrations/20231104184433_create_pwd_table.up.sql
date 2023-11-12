CREATE TABLE "credentials" (
    "id" varchar(50) PRIMARY KEY,
    "password" varchar(255) NOT NULL,
    "user_id" varchar(50) NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "sourceConfigs" (
    "id" varchar(50) PRIMARY KEY,
    "password_length" int DEFAULT 8,
    "password_lower_allowed" boolean DEFAULT true,
    "password_upper_allowed" boolean DEFAULT true,
    "password_numeric_allowed" boolean DEFAULT true,
    "password_special_allowed" boolean DEFAULT true,
    "is_system_generated" boolean DEFAULT true,
    "is_active" boolean DEFAULT true,
    "max_concurrent_login" int DEFAULT 1,
    "source_id" varchar(50) NOT NULL,
    FOREIGN KEY ("source_id") REFERENCES "sources"("id") ON DELETE CASCADE,
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Index
CREATE INDEX credentials_user_id ON credentials(user_id);


ALTER TABLE "sourceConfigs" ADD CONSTRAINT "source_id_unique" CHECK ( "is_system_generated" = true );