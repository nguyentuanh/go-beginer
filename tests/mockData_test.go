package tests

import (
	"go-template/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MockData(t *testing.T) {
	db := repository.ConnectDatabase()
	err := repository.MockData()
	defer db.Close()
	assert.Nil(t, err)
}
