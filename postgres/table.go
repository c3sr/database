package postgres

import (
	"errors"

	"github.com/c3sr/database"
	"github.com/c3sr/database/relational"
)

// NewTable ...
func NewTable(db database.Database, tableName string) (database.Table, error) {
	_, ok := db.(*postgresDatabase)
	if !ok {
		return nil, errors.New("invalid database input. Expecting a postgres database instance")
	}
	return relational.NewTable(db, tableName)
}
