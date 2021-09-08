package database

type CountryDB interface {
	New() CountryDB
	Get(country Country) (*Country, error)
	Select() ([]Country, error)
	Insert(country Country) (Country, error)
	Update(country Country) (Country, error)
	Delete(country Country) error
}

type Country struct{
	Id   uint64 `db:"id" structs:"-"`
	Name string `db:"name" structs:"name"`
}
