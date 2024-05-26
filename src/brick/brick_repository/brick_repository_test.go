package brickRepository

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	client := resty.New()
	brick := NewBrickRepository(client)

	token, err := brick.GenerateAccessToken()
	assert.Empty(t, token)
	assert.Error(t, err)
}
