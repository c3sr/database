package mysql

import (
	"testing"

	"github.com/k0kubun/pp/v3"
	"github.com/c3sr/config"
	"github.com/stretchr/testify/assert"
)

// XXXTestConnection ...
func XXXTestConnection(t *testing.T) {
	config.Init()
	db, err := NewDatabase("abduld")
	if !assert.NoError(t, err) {
		return
	}
	pp.Println(db)
}
