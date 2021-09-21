CREATE TABLE "domains" (
    "name" VARCHAR(255) PRIMARY KEY UNIQUE NOT NULL,
    "deliver" INT,
    "bounce" INT
);