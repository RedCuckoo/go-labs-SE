-- +migrate Up

CREATE TABLE city
(
    "id"         serial       NOT NULL,
    "country"    integer      NOT NULL,
    "name"       VARCHAR(255) NOT NULL,
    "population" integer      NOT NULL,
    "capital"    BOOLEAN      NOT NULL,
    PRIMARY KEY ("id"),

    CONSTRAINT fk_country
        FOREIGN KEY (country)
            REFERENCES country (id)
            ON DELETE SET NULL
);

-- +migrate Down

DROP TABLE city;
