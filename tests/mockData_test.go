package tests

import (
	"go-template/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MockData(t *testing.T) {
	repository.ConnectDatabase()
	err := repository.MockData()
	assert.Nil(t, err)
}
