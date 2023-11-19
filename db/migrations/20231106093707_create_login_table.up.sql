CREATE TABLE "logins" (
    "id" bigserial PRIMARY KEY,
    "ip" varchar(255) NOT NULL,
    "platform" int NOT NULL,
    "user_agent" varchar(255) NOT NULL,
    "access_token" varchar(1000) NOT NULL,
    "is_active" boolean NOT NULL,
    "logout_at" timestamp,
    "user_id" varchar(50) NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY ("user_id") REFERENCES "users"("id"),
    "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- INDEX
CREATE INDEX login_user_id ON logins(user_id, is_active);