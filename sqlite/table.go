package sqlite

import (
	"errors"

	"github.com/c3sr/database"
	"github.com/c3sr/database/relational"
)

// NewTable ...
func NewTable(db database.Database, tableName string) (database.Table, error) {
	_, ok := db.(*sqliteDatabase)
	if !ok {
		return nil, errors.New("invalid database input. Expecting a sqlite database instance")
	}
	return relational.NewTable(db, tableName)
}
