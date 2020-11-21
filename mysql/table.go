package mysql

import (
	"errors"

	"github.com/c3sr/database"
	"github.com/c3sr/database/relational"
)

// NewTable ...
func NewTable(db database.Database, tableName string) (database.Table, error) {
	_, ok := db.(*mysqlDatabase)
	if !ok {
		return nil, errors.New("invalid database input. Expecting a mysql database instance")
	}
	return relational.NewTable(db, tableName)
}
