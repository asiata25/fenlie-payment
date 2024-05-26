package brickRepository

import (
	"encoding/base64"
	brickDTO "finpro-fenlie/model/dto/brick"
	"finpro-fenlie/src/brick"
	"fmt"
	"os"

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
	success := brickDTO.BrickCreatePaymentLinkResponse{}
	fail := brickDTO.BrickErrorResponse{}

	resp, err := b.client.R().
		SetHeader("publicAccessToken", "Bearer "+accessToken).
		SetMultipartFormData(map[string]string{
			"referenceId":  payload.ReferenceId,
			"amount":       payload.Amount,
			"endUserEmail": payload.EndUserEmail,
			"endUserName":  payload.EndUserName,
		}).
		SetResult(&success).
		SetError(&fail).
		Post("https://sandbox.onebrick.io/v2/payments/gs/payment-link")
	if err != nil {
		log.Error().Err(err).Msg("error brick: when hit generate payment link endpoint")
		return "", err
	}

	fmt.Println(resp.String())

	if resp.StatusCode() >= 400 {
		respErr := errors.New(fail.Error.Reason)

		log.Error().Err(respErr).Msg("error brick: response error")
		return "", respErr
	} else {
		return success.Data.PaymentLinkPath, nil
	}
}

// GetPaymentStatus implements brick.BrickRepository.
func (b *brickRepository) GetPaymentStatus(accessToken string, transactionId string) (status string, err error) {
	success := brickDTO.BrickPaymentStatusResponse{}
	fail := brickDTO.BrickErrorResponse{}

	resp, err := b.client.R().
		SetHeader("publicAccessToken", "Bearer "+accessToken).
		SetBody(map[string]string{
			"referenceId": transactionId,
		}).
		SetResult(&success).
		SetError(&fail).
		Post("https://sandbox.onebrick.io/v2/payments/gs/payment-link/status")
	if err != nil {
		log.Error().Err(err).Msg("error brick: when hit get payment status endpoint")
		return "", err
	}

	fmt.Println(resp.String())

	if resp.StatusCode() >= 400 {
		respErr := errors.New(fail.Error.Reason)

		log.Error().Err(respErr).Msg("error brick: response error")
		return "", respErr
	} else {
		return success.Data.Status, nil
	}
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
		log.Error().Err(err).Msg("error brick: when hit access token endpoint")
		return "", err
	}

	if resp.StatusCode() >= 400 {
		respErr := errors.New(fail.Error.Reason)

		log.Error().Err(respErr).Msg("error brick: response error")
		return "", respErr
	} else {
		return success.Data.AccessToken, nil
	}
}

func NewBrickRepository(client *resty.Client) brick.BrickRepository {
	return &brickRepository{
		client:      client,
		brickId:     os.Getenv("BRICK_ID"),
		brickSecret: os.Getenv("BRICK_SECRET"),
	}
}
