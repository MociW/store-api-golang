CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "user_id" VARCHAR(100) UNIQUE NOT NULL,
    "avatar" VARCHAR(255),
    "first_name" VARCHAR(100),
    "last_name" VARCHAR(100),
    "username" VARCHAR(100) UNIQUE NOT NULL,
    "email" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255),
    "phone_number" VARCHAR(20),
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL
);

CREATE TABLE "addresses" (
    "id" SERIAL PRIMARY KEY,
    "user_id" VARCHAR(100) REFERENCES "users" ("user_id"),
    "title" VARCHAR(100),
    "street" VARCHAR(255),
    "country" VARCHAR(100),
    "city" VARCHAR(100),
    "postal_code" VARCHAR(20),
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL
);

CREATE TABLE "products" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(255),
    "description" VARCHAR(255),
    "summary" VARCHAR(255),
    "images" VARCHAR[],
    "user_id" VARCHAR(100) REFERENCES "users"("user_id"),
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL
);

CREATE TABLE "product_skus" (
    "id" SERIAL PRIMARY KEY,
    "product_id" INTEGER REFERENCES "products" ("id"),
    "size" VARCHAR(100),
    "color" VARCHAR(100),
    "sku" VARCHAR(100),
    "price" DECIMAL(10, 2),
    "quantity" INTEGER,
    "user_id" VARCHAR(100) REFERENCES "users" ("user_id"),
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP NULL
);