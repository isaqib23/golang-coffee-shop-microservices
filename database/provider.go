// Package database function that creates a new instance of the database implementation
package database

import "github.com/google/wire"

func NewDatabse() *DBClient {
	return NewDatabaseConnection()
}

var DatabaseSet = wire.NewSet(
	NewDatabse(),
	wire.Build(new(*DBClient)),
)
