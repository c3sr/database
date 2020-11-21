package rethinkdb

import (
	"context"

	"github.com/c3sr/database"
)

const (
	authKeyKey         = "github.com/c3sr/database/rethinkdb/authKey"
	initialCapacityKey = "github.com/c3sr/database/rethinkdb/initialCapacity"
)

// DefaultInitialCapacity ...
var (
	DefaultInitialCapacity = 10
)

// AuthKey ...
func AuthKey(s string) database.Option {
	return func(o *database.Options) {
		o.Context = context.WithValue(o.Context, authKeyKey, s)
	}
}

// InitialCapacity ...
func InitialCapacity(n int) database.Option {
	return func(o *database.Options) {
		o.Context = context.WithValue(o.Context, initialCapacityKey, n)
	}
}
