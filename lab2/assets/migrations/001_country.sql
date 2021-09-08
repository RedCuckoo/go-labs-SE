-- +migrate Up

CREATE TABLE country
(
    "id"   serial       NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    PRIMARY KEY ("id")
);

-- +migrate Down

DROP TABLE country;

