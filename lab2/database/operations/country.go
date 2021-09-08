package operations

import (
	"database/sql"
	"github.com/redcuckoo/go-labs-SE/lab2/database"
)

func NewCountryDB(db *sql.DB) database.CountryDB {
	return &countryDB{
		db:  db,
	}
}

type countryDB struct {
	db  *sql.DB
}

func (c *countryDB) New() database.CountryDB {
	return NewCountryDB(c.db)
}

func (c *countryDB) Get(country database.Country) (*database.Country, error) {
	rows, err := c.db.Query("SELECT * FROM country WHERE id=$1", country.Id)

	if err != nil {
		return nil, err
	}

	var result database.Country


	if err == sql.ErrNoRows {
		return nil, nil
	}

	if rows.Next() {
		err = rows.Scan(&result.Id, &result.Name)
		if err != nil {
			return nil, err
		}
	}

	return &result, err
}

func (c *countryDB) Select() ([]database.Country, error) {
	rows, err := c.db.Query("SELECT * FROM country")

	if err != nil {
		return []database.Country{}, err
	}

	var result []database.Country

	for rows.Next() {
		var country database.Country

		err = rows.Scan(&country.Id, &country.Name)
		if err != nil {
			return nil, err
		}

		result = append(result, country)
	}

	return result, err
}

func (c *countryDB) Insert(country database.Country) (database.Country, error) {
	_, err := c.db.Query("INSERT INTO country (name) values ($1) returning *", country.Name)

	if err != nil {
		return database.Country{}, err
	}

	return country, err
}

func (c *countryDB) Update(country database.Country) (database.Country, error) {
	_, err := c.db.Query("UPDATE country SET name=$2 WHERE id=$1", country.Id, country.Name)

	if err != nil {
		return database.Country{}, err
	}

	return country, err
}

func (c *countryDB) Delete(country database.Country) error {
	_, err := c.db.Query("DELETE FROM country WHERE id=$1", country.Id)

	if err != nil {
		return err
	}

	return err
}
