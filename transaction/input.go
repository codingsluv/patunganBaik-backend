package transaction

import "github.com/codingsluv/crowdfounding/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignID int `json:"campaign_id" binding:"required"`
	User       user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transactio_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"paymet_type"`
	FraudStatus       string `json:"fraud_status"`
}
