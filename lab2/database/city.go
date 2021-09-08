package database

type CityDB interface {
	New() CityDB
	Get(city City) (*City, error)
	Select() ([]City, error)
	Insert(city City) (City, error)
	Update(city City) (City, error)
	Delete(city City) error
}

type City struct {
	Id         uint64
	Country    uint64
	Name       string
	Population uint64
	Capital    bool
}
