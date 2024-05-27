package brick

import "finpro-fenlie/model/dto/brick"

type BrickRepository interface {
	GenerateAccessToken() (token string, err error)
	GeneratePaymentLink(accessToken string, payload brick.BrickCreatePaymentLinkRequest) (link string, err error)
	GetPaymentStatus(accessToken, transactionId string) (status string, err error)
}
