package helpers

import (
	"context"
	"github.com/redcuckoo/go-labs-SE/lab4/database"
)

type ctxKey int

const (
	countryDBCtxKey ctxKey = iota
	cityDBCtxKey
)

func CtxCountryDB(countryDB database.CountryDB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, countryDBCtxKey, countryDB)
	}
}

func CtxCityDB(cityDB database.CityDB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, cityDBCtxKey, cityDB)
	}
}

func CountryDBFromCtx(ctx context.Context) database.CountryDB {
	return ctx.Value(countryDBCtxKey).(database.CountryDB)
}

func CityDBFromCtx(ctx context.Context) database.CityDB {
	return ctx.Value(cityDBCtxKey).(database.CityDB)
}
