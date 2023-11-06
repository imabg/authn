CREATE TABLE "logins" (
    "id" bigserial PRIMARY KEY,
    "ip" varchar(255) NOT NULL,
    "platform" varchar(255) NOT NULL,
    "user_agent" varchar(255) NOT NULL,
    "access_token" varchar(255) NOT NULL,
    "is_active" boolean NOT NULL,
    "is_blacklisted" boolean NOT NULL,
    "user_id" varchar(50) NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY ("user_id") REFERENCES "users"("id")
);

-- INDEX
CREATE INDEX login_user_id ON logins(user_id, is_active);