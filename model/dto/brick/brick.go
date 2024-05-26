package brick

type (
	brickResponse struct {
		Status   int `json:"status"`
		MetaData struct {
			Source string `json:"source"`
			Entity string `json:"entity"`
		} `json:"metaData"`
	}

	BrickCreateTokenResponse struct {
		brickResponse
		Data struct {
			Message     string `json:"message"`
			AccessToken string `json:"accessToken"`
			IssuedAt    string `json:"issuedAt"`
			ExpiresAt   string `json:"expiresAt"`
		} `json:"data"`
	}

	BrickErrorResponse struct {
		brickResponse
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Action  string `json:"action"`
			Reason  string `json:"reason"`
		} `json:"error"`
	}

	BrickCreatePaymentLinkRequest struct {
		Amount       string `json:"amount"`
		ReferenceId  string `json:"referenceId"`
		Description  string `json:"description"`
		EndUserEmail string `json:"endUserEmail"`
		EndUserName  string `json:"endUserName"`
	}

	BrickCreatePaymentLinkResponse struct {
		brickResponse
		Data struct {
			PaymentLinkPath string `json:"paymentLinkPath"`
			ExpiresAt       string `json:"expiresAt"`
			Amount          int    `json:"amount"`
			ReferenceId     string `json:"referenceId"`
			Status          string `json:"status"`
		} `json:"data"`
	}

	BrickPaymentStatusResponse struct {
		Data struct {
			Status string `json:"status"`
		} `json:"data"`
	}
)
