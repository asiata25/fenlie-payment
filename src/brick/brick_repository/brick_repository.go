package brickRepository

import (
	"encoding/base64"
	brickDTO "finpro-fenlie/model/dto/brick"
	"finpro-fenlie/src/brick"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type brickRepository struct {
	client      *resty.Client
	brickId     string
	brickSecret string
}

// GeneratePaymentLink implements brick.BrickRepository.
func (b *brickRepository) GeneratePaymentLink(accessToken string, payload brickDTO.BrickCreatePaymentLinkRequest) (link string, err error) {
	panic("unimplemented")
}

// GetPaymentStatus implements brick.BrickRepository.
func (b *brickRepository) GetPaymentStatus(accessToken string, transactionId string) {
	panic("unimplemented")
}

// GenerateAccessToken implements brick.BrickRepository.
func (b *brickRepository) GenerateAccessToken() (token string, err error) {
	encodeKey := base64.StdEncoding.EncodeToString([]byte(b.brickId + ":" + b.brickSecret))

	success := brickDTO.BrickCreateTokenResponse{}
	fail := brickDTO.BrickErrorResponse{}

	resp, err := b.client.R().
		SetHeader("Authorization", "Basic "+encodeKey).
		SetResult(&success).
		SetError(&fail).
		Get("https://sandbox.onebrick.io/v2/payments/auth/token")
	if err != nil {
		log.Error().Err(err).Msg("error brick: when generate access token")
	}

	if resp.StatusCode() >= 400 {
		return "", errors.New(fail.Error.Reason)
	} else {
		return success.Data.AccessToken, nil
	}
}

func NewBrickRepository(client *resty.Client, brickId, brickSecret string) brick.BrickRepository {
	return &brickRepository{
		client:      client,
		brickId:     brickId,
		brickSecret: brickSecret,
	}
}
