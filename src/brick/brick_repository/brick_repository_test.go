package brickRepository

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAccessToken(t *testing.T) {
	client := resty.New()
	brick := NewBrickRepository(client, "f41b695e-6e67-4409-96f7-6b2e3b6e07b1", "PRTFiWDFtPUyN9V0FsIoUAqHe0bpvN")

	token, err := brick.GenerateAccessToken()
	assert.Empty(t, token)
	assert.Error(t, err)
}
