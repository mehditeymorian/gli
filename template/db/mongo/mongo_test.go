package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	cfg := config.Load("config.yaml")

	mdb, err := Connect(cfg.DB)

	assert.NoError(t, err)
	assert.NotNil(t, mdb)

}
