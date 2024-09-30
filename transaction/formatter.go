package transaction

import (
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	campaignTransactionFormatter := CampaignTransactionFormatter{}
	campaignTransactionFormatter.ID = transaction.ID
	campaignTransactionFormatter.Name = transaction.User.Name
	campaignTransactionFormatter.Amount = transaction.Amount
	campaignTransactionFormatter.CreatedAt = transaction.CreatedAt

	return campaignTransactionFormatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	campaignTransactionsFormatter := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		campaignTransactionFormatter := FormatCampaignTransaction(transaction)
		campaignTransactionsFormatter = append(campaignTransactionsFormatter, campaignTransactionFormatter)
	}

	return campaignTransactionsFormatter
}

type UserTransactionFormatter struct {
	ID        int                   `json:"id"`
	Amount    int                   `json:"amount"`
	Status    string                `json:"status"`
	CreatedAt time.Time             `json:"created_at"`
	Campaign  UserCampaignFormatter `json:"campaign"`
}

type UserCampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	userTransactionFormatter := UserTransactionFormatter{}
	userTransactionFormatter.ID = transaction.ID
	userTransactionFormatter.Amount = transaction.Amount
	userTransactionFormatter.Status = transaction.Status
	userTransactionFormatter.CreatedAt = transaction.CreatedAt

	campaign := transaction.Campaign
	userCampaignFormatter := UserCampaignFormatter{}
	userCampaignFormatter.Name = campaign.Name
	userCampaignFormatter.ImageURL = campaign.CampaignImages[0].Filename

	userTransactionFormatter.Campaign = userCampaignFormatter

	return userTransactionFormatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	userTransactionsFormatter := []UserTransactionFormatter{}

	for _, transaction := range transactions {
		userTransactionFormatter := FormatUserTransaction(transaction)
		userTransactionsFormatter = append(userTransactionsFormatter, userTransactionFormatter)
	}

	return userTransactionsFormatter
}
