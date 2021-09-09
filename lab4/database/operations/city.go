package operations

import (
	"database/sql"
	"github.com/redcuckoo/go-labs-SE/lab4/database"
)

func NewCityDB(db *sql.DB) database.CityDB {
	return &cityDB{
		db:  db,
	}
}

type cityDB struct {
	db  *sql.DB
}

func (c *cityDB) New() database.CityDB {
	return NewCityDB(c.db)
}

func (c *cityDB) Get(city database.City) (*database.City, error) {
	rows, err := c.db.Query("SELECT * FROM city WHERE id=$1", city.Id)

	if err != nil {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if rows.Next() {
		err = rows.Scan(&city.Id, city.Country, city.Name, city.Population, city.Capital)
		if err != nil {
			return nil, err
		}
	}

	return &city, err
}

func (c *cityDB) Select() ([]database.City, error) {
	rows, err := c.db.Query("SELECT * FROM city")

	if err != nil {
		return []database.City{}, err
	}

	var result []database.City

	for rows.Next() {
		var city database.City

		err = rows.Scan(&city.Id, &city.Country, &city.Name, &city.Population, &city.Capital)
		if err != nil {
			return nil, err
		}

		result = append(result, city)
	}

	return result, err
}

func (c *cityDB) Insert(city database.City) (database.City, error) {
	rows, err := c.db.Query("INSERT INTO city (country, name, population, capital) values ($1, $2, $3, $4) returning *", city.Country, city.Name, city.Population, city.Capital)

	if err != nil {
		return database.City{}, err
	}

	if rows.Next() {
		err = rows.Scan(&city.Id, &city.Country, &city.Name, &city.Population, &city.Capital)
		if err != nil {
			return database.City{}, err
		}
	}

	return city, err
}

func (c *cityDB) Update(city database.City) (database.City, error) {
	rows, err := c.db.Query("UPDATE city SET country=$2, name=$3, population=$4, capital=$5 WHERE id=$1", city.Id, city.Country,city.Name, city.Population, city.Capital)

	if err != nil {
		return database.City{}, err
	}

	if rows.Next() {
		err = rows.Scan(&city.Id, &city.Country, &city.Name, &city.Population, &city.Capital)
		if err != nil {
			return database.City{}, err
		}
	}

	return city, err
}

func (c *cityDB) Delete(city database.City) error {
	_, err := c.db.Query("DELETE FROM city WHERE id=$1", city.Id)

	if err != nil {
		return err
	}

	return err
}
