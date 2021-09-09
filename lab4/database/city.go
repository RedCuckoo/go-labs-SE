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
	Id         uint64 `json:"id,omitempty"`
	Country    uint64 `json:"country"`
	Name       string `json:"name"`
	Population uint64 `json:"population"`
	Capital    bool `json:"capital"`
}
